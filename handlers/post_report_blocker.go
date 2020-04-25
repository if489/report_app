package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reports_app/logic"

	"github.com/go-chi/chi"
)

// PostServeHTTP reads and validates request and writes response
func (p *PostPutReportUpdater) PostServeHTTP(w http.ResponseWriter, r *http.Request) {

	reportID := chi.URLParam(r, "reportId")
	match := uuidRegex.MatchString(reportID)
	if len(reportID) == 0 || !match {
		err := ErrorResponse{"path_not_found", ErrMissingOrIncorrectReportID.Error()}
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	// call logic
	responseID, err := p.reportUpdater.Update(reportID, "BLOCKED")
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
	w.WriteHeader(http.StatusOK)

	jsonEnc := json.NewEncoder(w)
	err = jsonEnc.Encode(&response)
	if err != nil {
		log.Fatal(err)
	}
}
