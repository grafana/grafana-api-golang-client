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
		"dashboards": [
			{
				"dashboard": {
					"id": 33,
					"uid": "nErXDvCkzz",
					"name": "Terraform Acceptance Test"
				},
				"timeRange": {
					"from": "now-1h",
					"to": "now"
				}
			}
		],
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
			"layout": "grid"
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
		Name:       "My Report",
		Recipients: "test@test.com",
		Schedule: ReportSchedule{
			StartDate:         &now,
			EndDate:           nil,
			Frequency:         "custom",
			IntervalFrequency: "weeks",
			IntervalAmount:    2,
			WorkdaysOnly:      true,
			TimeZone:          "GMT",
		},
		Dashboards: []ReportDashboard{
			{
				Dashboard: ReportDashboardIdentifier{
					ID:  33,
					UID: "nErXDvCkzz",
				},
				TimeRange: ReportDashboardTimeRange{
					From: "now-1h",
					To:   "now",
				},
			},
		},
		Options: ReportOptions{
			Orientation: "landscape",
			Layout:      "grid",
		},
		EnableDashboardURL: true,
		EnableCSV:          true,
	}
)

func TestReport(t *testing.T) {
	client := gapiTestTools(t, 200, getReportJSON)

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
	client := gapiTestTools(t, 200, createReportJSON)

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
	client := gapiTestTools(t, 200, "")

	err := client.UpdateReport(testReport)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteReport(t *testing.T) {
	client := gapiTestTools(t, 200, "")

	err := client.DeleteReport(4)
	if err != nil {
		t.Fatal(err)
	}
}
