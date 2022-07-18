package gapi

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/reports"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
	"github.com/stretchr/testify/require"
)

var (
	getReportJSON = `
	{
		"id": 4,
		"userId": 0,
		"orgId": 1,
		"dashboardId": 33,
		"dashboardName": "Terraform Acceptance Test",
		"dashboardUid": "",
		"name": "My Report",
		"recipients": "test@test.com",
		"replyTo": "",
		"message": "",
		"schedule": {
			"startDate": "2020-01-01T00:00:00Z",
			"endDate": null,
			"frequency": "custom",
			"intervalFrequency": "weeks",
			"intervalAmount": 2,
			"workdaysOnly": true,
			"dayOfMonth": "1",
			"day": "wednesday",
			"hour": 0,
			"minute": 0,
			"timeZone": "GMT"
		},
		"options": {
			"orientation": "landscape",
			"layout": "grid",
			"timeRange": {
				"from": "now-1h",
				"to": "now"
			}
		},
		"templateVars": {},
		"enableDashboardUrl": true,
		"enableCsv": true,
		"state": "",
		"created": "2022-01-11T15:09:13Z",
		"updated": "2022-01-11T16:18:34Z"
	}
`
	createReportJSON = `
	{
		"id": 4
	}
`
)

func testReport(t *testing.T) *models.CreateOrUpdateConfigCmd {
	t.Helper()

	startDate, err := strfmt.ParseDateTime("2020-01-01T00:00:00Z")
	require.NoError(t, err)

	return &models.CreateOrUpdateConfigCmd{
		DashboardID: 33,
		Name:        "My Report",
		Recipients:  "test@test.com",
		Schedule: &models.ScheduleDTO{
			StartDate:         startDate,
			EndDate:           strfmt.NewDateTime(),
			Frequency:         "custom",
			IntervalFrequency: "weeks",
			IntervalAmount:    2,
			WorkdaysOnly:      true,
			TimeZone:          "GMT",
		},
		Options: &models.ReportOptionsDTO{
			Orientation: "landscape",
			Layout:      "grid",
			TimeRange: &models.TimeRangeDTO{
				From: "now-1h",
				To:   "now",
			},
		},
		EnableDashboardURL: true,
		EnableCsv:          true,
	}
}

func TestReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, getReportJSON)
	defer server.Close()

	reportID := int64(4)
	resp, err := client.Reports.GetReport(
		reports.NewGetReportParams().
			WithReportID(reportID),
		nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.ID != reportID || resp.Payload.Name != "My Report" {
		t.Error("Not correctly parsing returned report.")
	}
}

func TestNewReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, createReportJSON)
	defer server.Close()

	resp, err := client.Reports.CreateReport(
		reports.NewCreateReportParams().
			WithBody(testReport(t)),
		nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.ID != 4 {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, "")
	defer server.Close()

	_, err := client.Reports.UpdateReport(
		reports.NewUpdateReportParams().
			WithReportID(4).
			WithBody(testReport(t)),
		nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, "")
	defer server.Close()

	_, err := client.Reports.DeleteReport(
		reports.NewDeleteReportParams().
			WithReportID(int64(4)),
		nil)
	if err != nil {
		t.Fatal(err)
	}
}
