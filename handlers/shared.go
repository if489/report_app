package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"reports_app/logic"
)

// ErrMissingOrIncorrectReportID will be returned if the uuid is missing or is not the correct format
var ErrMissingOrIncorrectReportID = errors.New("correct report ID is required")

// ErrMissingTicketState is returned when ticket state is missing
var ErrMissingTicketState = errors.New("ticket state required")

// ErrWrongTicketState is returnes when ticket state is incorrect
var ErrWrongTicketState = errors.New("wrong ticket state")

var uuidRegex = regexp.MustCompile(
	"^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$",
)

// ErrorResponse represents the error response
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Response is the response for put report resolver and post report blocker
type Response struct {
	ID string `json:"id"`
}

// ReportsResponse represents the response of the reports loader
type ReportsResponse struct {
	Size    int            `json:"size"`
	Reports []logic.Report `json:"reports"`
}

func errorResponse(w http.ResponseWriter, status int, err ErrorResponse) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
	return
}

// PostReportUpdater handles post/put requests for repport state update
type PostPutReportUpdater struct {
	reportUpdater logic.RUpdater
}

// NewPostPutReportUpdater returns and instance of *PostPutReportUpdater
func NewPostPutReportUpdater(
	reportUpdater logic.RUpdater,
) *PostPutReportUpdater {
	return &PostPutReportUpdater{
		reportUpdater: reportUpdater,
	}
}

type mockReportUpdater struct {
	expectedID    string
	expectedState string
	err           error
}

func (m *mockReportUpdater) Update(id, state string) (string, error) {
	if m.err != nil {
		return "", m.err
	}

	if id != m.expectedID || state != m.expectedState {
		return "", fmt.Errorf(
			"report updater: expected id : '%+v' and expected state : '%+v', received '%+v and '%+v'",
			m.expectedID,
			m.expectedState,
			id,
			state,
		)
	}

	return m.expectedID, nil
}
