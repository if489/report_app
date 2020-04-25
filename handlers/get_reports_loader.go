package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reports_app/logic"
)

// GetReportsLoader handles get request for loading of open reports
type GetReportsLoader struct {
	reportsLoader logic.RLoader
}

// NewGetReportsLoader returns a new instance of *GetReportsLoader
func NewGetReportsLoader(
	reportsLoader logic.RLoader,
) *GetReportsLoader {
	return &GetReportsLoader{
		reportsLoader: reportsLoader,
	}
}

// ServeHTTP reads and validates request and writes response
func (g *GetReportsLoader) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// call logic
	Reports, count, err := g.reportsLoader.Load()
	if err != nil {
		switch err {
		case logic.ErrLoadReports:
			err := ErrorResponse{"result_not_loaded", err.Error()}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		case logic.ErrUnmarschal:
			err := ErrorResponse{"unmarshalling_error", err.Error()}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		case logic.ErrScan:
			err := ErrorResponse{"db_error", err.Error()}
			errorResponse(w, http.StatusInternalServerError, err)
			return
		default:
			log.Fatal(err)
		}
	}

	ReportsResponse := ReportsResponse{count, Reports}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonEnc := json.NewEncoder(w)
	err = jsonEnc.Encode(&ReportsResponse)
	if err != nil {
		log.Fatal(err)
	}
}
