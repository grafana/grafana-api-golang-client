package gapi

import (
	"testing"
	"time"

	"github.com/gobs/pretty"
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
	now        = time.Now()
	testReport = Report{
		DashboardID: 33,
		Name:        "My Report",
		Recipients:  "test@test.com",
		Schedule: ReportSchedule{
			StartDate:         &now,
			EndDate:           nil,
			Frequency:         "custom",
			IntervalFrequency: "weeks",
			IntervalAmount:    2,
			WorkdaysOnly:      true,
			TimeZone:          "GMT",
		},
		Options: ReportOptions{
			Orientation: "landscape",
			Layout:      "grid",
			TimeRange: ReportTimeRange{
				From: "now-1h",
				To:   "now",
			},
		},
		EnableDashboardURL: true,
		EnableCSV:          true,
	}
)

func TestReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, getReportJSON)
	defer server.Close()

	report := int64(4)
	resp, err := client.Report(report)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.ID != report || resp.Name != "My Report" {
		t.Error("Not correctly parsing returned report.")
	}
}

func TestNewReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, createReportJSON)
	defer server.Close()

	resp, err := client.NewReport(testReport)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp != 4 {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, "")
	defer server.Close()

	err := client.UpdateReport(testReport)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteReport(t *testing.T) {
	server, client := gapiTestTools(t, 200, "")
	defer server.Close()

	err := client.DeleteReport(4)
	if err != nil {
		t.Fatal(err)
	}
}
