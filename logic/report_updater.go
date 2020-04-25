package logic

import (
	"database/sql"
	"errors"
)

// ErrUnknownReportID returned when report id not found
var ErrUnknownReportID = errors.New("unknown report id")

// RUpdater manages updating report state
type RUpdater interface {
	Update(string, string) (string, error)
}

// ReportUpdater holds the logic for updating the report state
type ReportUpdater struct {
	db *sql.DB
}

// NewReportUpdater return a new instance of *ReportUpdater
func NewReportUpdater(db *sql.DB) *ReportUpdater {
	return &ReportUpdater{
		db: db,
	}
}

// ReportUpdaterUpdate will update the state of the report
func (ru *ReportUpdater) Update(id, state string) (string, error) {
	var returnedID string

	row := ru.db.QueryRow(UpdateReport, state, id)
	err := row.Scan(&returnedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUnknownReportID
		}
		return "", err
	}
	return returnedID, nil
}

// UpdateReport updates the state of the report
const UpdateReport string = `
UPDATE reports
SET state=$1, updated_at=CURRENT_TIMESTAMP
WHERE id = $2
RETURNING id
`
