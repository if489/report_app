package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reports_app/logic"

	"github.com/go-chi/chi"
)

// TicketStateRequest is the request body of put resolver endpoint
type TicketStateRequest struct {
	TicketState string `json:"ticketState"`
}

// PutServeHTTP reads and validates request and writes response
func (p *PostPutReportUpdater) PutServeHTTP(w http.ResponseWriter, r *http.Request) {

	reportID := chi.URLParam(r, "reportId")
	match := uuidRegex.MatchString(reportID)
	if len(reportID) == 0 || !match {
		err := ErrorResponse{"path_not_found", ErrMissingOrIncorrectReportID.Error()}
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	TicketStateRequest := &TicketStateRequest{}
	if err = json.Unmarshal(reqBody, &TicketStateRequest); err != nil {
		err := ErrorResponse{"unmarshalling_error", err.Error()}
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = validate(TicketStateRequest)
	if err != nil {
		err := ErrorResponse{"validation_error", err.Error()}
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	// call logic
	responseID, err := p.reportUpdater.Update(reportID, TicketStateRequest.TicketState)
	if err != nil {
		if err == logic.ErrUnknownReportID {
			err := ErrorResponse{"entity_not_found", err.Error()}
			errorResponse(w, http.StatusNotFound, err)
			return
		}
		err := ErrorResponse{"entity_not_modified", err.Error()}
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// build response
	response := Response{responseID}
	w.Header().Set("Content-Type", "application/json")

	jsonEnc := json.NewEncoder(w)
	err = jsonEnc.Encode(&response)
	if err != nil {
		log.Fatal(err)
	}
}

func validate(ticketStateRequest *TicketStateRequest) error {

	if ticketStateRequest.TicketState == "" {
		return ErrMissingTicketState
	}

	if ticketStateRequest.TicketState != "CLOSED" {
		return ErrWrongTicketState
	}

	return nil
}
