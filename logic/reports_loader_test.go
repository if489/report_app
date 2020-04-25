package logic

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestReportsLoader(t *testing.T) {
	dateTimestamp := time.Date(2020, 04, 19, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		title      string
		mockRows   *sqlmock.Rows
		mockErr    error
		expReports []Report
		expCount   int
		expErr     error
	}{
		{
			title:      "error while loading the reports",
			mockErr:    ErrLoadReports,
			expReports: nil,
			expCount:   0,
			expErr:     ErrLoadReports,
		},
		{
			title:   "success",
			mockErr: nil,
			mockRows: sqlmock.NewRows([]string{"id", "source", "source_identity_id",
				"reference", "state", "payload", "created_at", "updated_at"}).
				AddRow("09ecf137-cbda-4d41-a6b2-142d2883da97", "REPORT",
					"6750b4d5-4cb5-45f0-8b60-61be2072cce2", `{}`, "OPEN", `{}`,
					dateTimestamp, dateTimestamp),
			expReports: []Report{
				{
					ID:               "09ecf137-cbda-4d41-a6b2-142d2883da97",
					Source:           "REPORT",
					SourceIdentityID: "6750b4d5-4cb5-45f0-8b60-61be2072cce2",
					Reference: Reference{ReferenceID: "",
						ReferenceType: ""},
					State: "OPEN",
					Payload: Payload{
						Source:                "",
						ReportType:            "",
						Message:               "",
						ReportID:              "",
						ReferenceResourceID:   "",
						ReferenceResourceType: "",
					},
					CreatedAt: dateTimestamp,
					UpdatedAt: dateTimestamp,
				},
			},
			expCount: 1,
		},
		{
			title: "unmarshal  error for reference",
			mockRows: sqlmock.NewRows([]string{"id", "source", "source_identity_id",
				"reference", "state", "payload", "created_at", "updated_at"}).
				AddRow("09ecf137-cbda-4d41-a6b2-142d2883da97", "REPORT",
					"6750b4d5-4cb5-45f0-8b60-61be2072cce2", `{not_a_json}`, "OPEN", `{}`,
					dateTimestamp, dateTimestamp),
			expReports: nil,
			expCount:   0,
			expErr:     ErrUnmarschal,
		},
		{
			title: "unmarshal  error for payload",
			mockRows: sqlmock.NewRows([]string{"id", "source", "source_identity_id",
				"reference", "state", "payload", "created_at", "updated_at"}).
				AddRow("09ecf137-cbda-4d41-a6b2-142d2883da97", "REPORT",
					"6750b4d5-4cb5-45f0-8b60-61be2072cce2", `{}`, "OPEN", `{not_a_json}`,
					dateTimestamp, dateTimestamp),
			expReports: nil,
			expCount:   0,
			expErr:     ErrUnmarschal,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf(
					"an error '%s' was not expected when opening a stub database connection",
					err,
				)
			}
			defer db.Close()

			if c.mockErr != nil {
				mock.ExpectExec(LoadReports).WithArgs("OPEN").WillReturnError(c.mockErr)
			}
			mock.ExpectQuery(LoadReports).WithArgs("OPEN").WillReturnRows(c.mockRows)

			loader := NewReportsLoader(db)
			reports, count, err := loader.Load()
			if err != nil && !reflect.DeepEqual(err, c.expErr) {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					err, c.expErr)
			}

			if !reflect.DeepEqual(reports, c.expReports) {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					reports, c.expReports)
			}
			if count != c.expCount {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					err, c.expCount)
			}
		})
	}
}
