package logic

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestReportUpdaterUpdate(t *testing.T) {
	cases := []struct {
		title  string
		id     string
		state  string
		expErr error
		expID  string
	}{
		{
			title:  "no rows in result set error",
			id:     "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			expErr: ErrUnknownReportID,
			expID:  "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			db, err := sql.Open("postgres", "host=localhost user=postgres password=password sslmode=disable")
			if err != nil {
				t.Fatalf("Failed to open DB: %v", err)
			}

			updater := NewReportUpdater(db)
			res, err := updater.Update(c.id, c.state)
			if err != nil && !reflect.DeepEqual(err, c.expErr) {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					err, c.expErr)
			}
			if res != c.expID {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					res, c.expID)
			}

		})
	}
}
