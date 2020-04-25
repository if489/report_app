package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"reports_app/logic"
	"testing"
	"time"

	"github.com/go-chi/chi"
)

func TestGetReportLoaderServeHTTP(t *testing.T) {
	dateTimestamp := time.Date(2020, 04, 19, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		title                  string
		loader                 logic.RLoader
		expectedResponseStatus string
		expectedResponseBody   string
	}{
		{
			title: "unknown error",
			loader: &mockReportLoader{
				expectedReports: nil,
				count:           0,
				err:             logic.ErrLoadReports,
			},
			expectedResponseStatus: "500 Internal Server Error",
			expectedResponseBody: `{` +
				`"code":"result_not_loaded",` +
				`"message":"couldn't load reports"` +
				`}` + "\n",
		},
		{
			title: "unmarshal error",
			loader: &mockReportLoader{
				expectedReports: nil,
				count:           0,
				err:             logic.ErrUnmarschal,
			},
			expectedResponseStatus: "500 Internal Server Error",
			expectedResponseBody: `{` +
				`"code":"unmarshalling_error",` +
				`"message":"unmarschal error"` +
				`}` + "\n",
		},
		{
			title: "scan error",
			loader: &mockReportLoader{
				expectedReports: nil,
				count:           0,
				err:             logic.ErrScan,
			},
			expectedResponseStatus: "500 Internal Server Error",
			expectedResponseBody: `{` +
				`"code":"db_error",` +
				`"message":"scan error"` +
				`}` + "\n",
		},
		{
			title:                  "success case",
			expectedResponseStatus: "200 OK",
			loader: &mockReportLoader{
				expectedReports: []logic.Report{
					{
						ID:               "09ecf137-cbda-4d41-a6b2-142d2883da97",
						Source:           "REPORT",
						SourceIdentityID: "6750b4d5-4cb5-45f0-8b60-61be2072cce2",
						Reference: logic.Reference{
							ReferenceID:   "",
							ReferenceType: "",
						},
						State: "OPEN",
						Payload: logic.Payload{
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
				count: 1,
			},
			expectedResponseBody: `{"size":1,"reports":[{` +
				`"id":"09ecf137-cbda-4d41-a6b2-142d2883da97",` +
				`"source":"REPORT",` +
				`"source_identity_id":"6750b4d5-4cb5-45f0-8b60-61be2072cce2",` +
				`"reference":{"reference_id":"","reference_type":""},` +
				`"state":"OPEN",` +
				`"payload":{"source":"","report_type":"","message":"",` +
				`"report_id":"","reference_resource_id":"","reference_resource_type":""},` +
				`"created_at":"2020-04-19T00:00:00Z",` +
				`"updated_at":"2020-04-19T00:00:00Z"` +
				`}]}` + "\n",
		},
		{
			title: "success case for empty result set",
			loader: &mockReportLoader{
				expectedReports: nil,
				count:           0,
			},
			expectedResponseStatus: "200 OK",
			expectedResponseBody:   `{"size":0,"reports":null}` + "\n",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/reports", nil)
			rctx := chi.NewRouteContext()
			request = request.WithContext(context.WithValue(
				request.Context(),
				chi.RouteCtxKey,
				rctx,
			))

			getReportsLoader := NewGetReportsLoader(c.loader)
			handler := http.HandlerFunc(getReportsLoader.ServeHTTP)

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

type mockReportLoader struct {
	expectedReports []logic.Report
	count           int
	err             error
}

func (m *mockReportLoader) Load() ([]logic.Report, int, error) {
	if m.err != nil {
		return nil, 0, m.err
	}

	return m.expectedReports, m.count, nil
}
