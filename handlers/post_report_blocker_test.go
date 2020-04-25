package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"reports_app/logic"
	"testing"

	"github.com/go-chi/chi"
)

func TestPostReportBlockerServeHTTP(t *testing.T) {

	cases := []struct {
		title                  string
		reportID               string
		updater                logic.RUpdater
		expectedResponseStatus string
		expectedResponseBody   string
	}{
		{
			title:                  "incorrect report ID",
			reportID:               "test",
			expectedResponseStatus: "400 Bad Request",
			expectedResponseBody: `{` +
				`"code":"path_not_found",` +
				`"message":"correct report ID is required"` +
				`}` + "\n",
		},
		{
			title:    "no rows in result set error",
			reportID: "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			updater: &mockReportUpdater{
				expectedID:    "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
				expectedState: "BLOCKED",
				err:           logic.ErrUnknownReportID,
			},
			expectedResponseStatus: "404 Not Found",
			expectedResponseBody: `{` +
				`"code":"entity_not_found",` +
				`"message":"unknown report id"` +
				`}` + "\n",
		},
		{
			title:    "success case",
			reportID: "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			updater: &mockReportUpdater{
				expectedID:    "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
				expectedState: "BLOCKED",
				err:           nil,
			},
			expectedResponseStatus: "200 OK",
			expectedResponseBody: `{` +
				`"id":"01322891-c5cb-4ac5-90d4-3c4224f40ba2"` +
				`}` + "\n",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", fmt.Sprintf("/reports/block/%s", c.reportID), nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("reportId", c.reportID)

			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

			postReportUpdater := NewPostPutReportUpdater(c.updater)
			handler := http.HandlerFunc(postReportUpdater.PostServeHTTP)

			handler.ServeHTTP(recorder, request)

			result := recorder.Result()

			if !reflect.DeepEqual(result.Status, c.expectedResponseStatus) {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					result.Status, c.expectedResponseStatus)
			}

			if result.Header.Get("Content-Type") != "application/json" {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n application/json\n",
					result.Header.Get("Content-Type"))
			}

			body, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Errorf("Couldn't open result body")
			}

			if !reflect.DeepEqual(string(body), c.expectedResponseBody) {
				t.Errorf(
					"\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
					string(body), c.expectedResponseBody)
			}
		})
	}
}
