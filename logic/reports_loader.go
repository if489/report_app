package logic

import (
	"database/sql"
	"encoding/json"
	"errors"
)

// ErrLoadReports will be returned when reports couldn't be loaded from the db
var ErrLoadReports = errors.New("couldn't load reports")

// ErrUnmarschal is returnes when unmarshaling returns an error
var ErrUnmarschal = errors.New("unmarschal error")

// ErrScan is returned when a db returns an error from scan
var ErrScan = errors.New("scan error")

// RLoader manages loads all open reports
type RLoader interface {
	Load() ([]Report, int, error)
}

// ReportsLoader holds the logic for loading open reports
type ReportsLoader struct {
	db *sql.DB
}

// NewReportsLoader return a new instance of *ReportsLoader
func NewReportsLoader(db *sql.DB) *ReportsLoader {
	return &ReportsLoader{
		db: db,
	}
}

// ReportsLoaderLoad will load all open  reports
func (rl *ReportsLoader) Load() ([]Report, int, error) {
	rows, err := rl.db.Query(LoadReports, "OPEN")
	if err != nil {
		return nil, 0, ErrLoadReports
	}
	defer rows.Close()

	Reports := []Report{}
	counter := 0

	for rows.Next() {
		Report := Report{}
		Reference := []byte{}
		Payload := []byte{}

		if err := rows.Scan(
			&Report.ID,
			&Report.Source,
			&Report.SourceIdentityID,
			&Reference,
			&Report.State,
			&Payload,
			&Report.CreatedAt,
			&Report.UpdatedAt); err != nil {
			return nil, 0, ErrScan
		}

		err = json.Unmarshal(Reference, &Report.Reference)
		if err != nil {
			return nil, 0, ErrUnmarschal
		}

		err = json.Unmarshal(Payload, &Report.Payload)
		if err != nil {
			return nil, 0, ErrUnmarschal
		}

		counter++
		Reports = append(Reports, Report)
	}

	return Reports, counter, nil
}

// LoadReports queries all the open reports
const LoadReports string = `
SELECT
	id,
    source,
    source_identity_id,
    reference,
    state,
    payload,
    created_at,
    updated_at
FROM reports
WHERE state = $1
`
