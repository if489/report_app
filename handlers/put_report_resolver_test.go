package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"reports_app/logic"
	"testing"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func TestPutReportResolverServeHTTP(t *testing.T) {

	cases := []struct {
		title                  string
		reportID               string
		reqBody                string
		updater                logic.RUpdater
		expectedResponseStatus string
		expectedResponseBody   string
	}{
		{
			title:                  "empty report ID",
			reportID:               "",
			reqBody:                `{}`,
			expectedResponseStatus: "400 Bad Request",
			expectedResponseBody: `{` +
				`"code":"path_not_found",` +
				`"message":"correct report ID is required"` +
				`}` + "\n",
		},
		{
			title:    "unmarshalling error",
			reportID: "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			reqBody:  "a_string",
			updater: &mockReportUpdater{
				expectedID:    "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
				expectedState: "CLOSED",
			},
			expectedResponseStatus: "400 Bad Request",
			expectedResponseBody: `{` +
				`"code":"unmarshalling_error",` +
				`"message":"invalid character 'a' looking for beginning of value"` +
				`}` + "\n",
		},
		{
			title:                  "empty reqeust body",
			reportID:               "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			reqBody:                `{}`,
			expectedResponseStatus: "400 Bad Request",
			expectedResponseBody: `{` +
				`"code":"validation_error",` +
				`"message":"ticket state required"` +
				`}` + "\n",
		},
		{
			title:    "wrong reqeust body",
			reportID: "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			reqBody: `{
				"ticketState":"bla"
				}`,
			expectedResponseStatus: "400 Bad Request",
			expectedResponseBody: `{` +
				`"code":"validation_error",` +
				`"message":"wrong ticket state"` +
				`}` + "\n",
		},
		{
			title:    "no rows in result set error",
			reportID: "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
			reqBody: `{
				"ticketState":"CLOSED"
				}`,
			updater: &mockReportUpdater{
				expectedID:    "01322891-c5cb-4ac5-90d4-3c4224f40ba2",
				expectedState: "CLOSED",
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
			reportID: "06c6055f-5cf1-4153-9e8e-a9966deaae68",
			reqBody: `{
				"ticketState":"CLOSED"
				}`,
			updater: &mockReportUpdater{
				expectedID:    "06c6055f-5cf1-4153-9e8e-a9966deaae68",
				expectedState: "CLOSED",
			},
			expectedResponseStatus: "200 OK",
			expectedResponseBody: `{` +
				`"id":"06c6055f-5cf1-4153-9e8e-a9966deaae68"` +
				`}` + "\n",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			bodyBuf := bytes.NewBuffer([]byte(c.reqBody))
			request := httptest.NewRequest("PUT", fmt.Sprintf("/reports/%s", c.reportID), bodyBuf)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("reportId", c.reportID)

			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))

			putReportUpdater := NewPostPutReportUpdater(c.updater)
			handler := http.HandlerFunc(putReportUpdater.PutServeHTTP)

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

func TestTicketStateRequestValidate(t *testing.T) {

	cases := []struct {
		title       string
		req         *TicketStateRequest
		expectedErr error
	}{
		{
			title:       "empty ticket state",
			req:         &TicketStateRequest{},
			expectedErr: ErrMissingTicketState,
		},
		{
			title: "wrong ticket state",
			req: &TicketStateRequest{
				TicketState: "bla",
			},
			expectedErr: ErrWrongTicketState,
		},
	}

	for _, c := range cases {
		err := validate(c.req)
		if err != nil && !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("\nGot unexpected output: \n%#v\nExpected output:\n%#v\n",
				err,
				c.expectedErr,
			)
		}
	}
}
