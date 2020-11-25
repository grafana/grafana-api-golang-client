package gapi

import (
	"net/url"
	"strings"
	"testing"

	"github.com/gobs/pretty"
)

const (
	alertsJSON = `[{
		"id": 1,
		"dashboardId": 1,
		"dashboardUId": "ABcdEFghij",
		"dashboardSlug": "sensors",
		"panelId": 1,
		"name": "fire place sensor",
		"state": "alerting",
		"newStateDate": "2018-05-14T05:55:20+02:00",
		"evalDate": "0001-01-01T00:00:00Z",
		"evalData": null,
		"executionError": "",
		"url": "http://grafana.com/dashboard/db/sensors"
	}]`

	alertJSON = `{
		"id": 1,
		"dashboardId": 1,
		"dashboardUId": "ABcdEFghij",
		"dashboardSlug": "sensors",
		"panelId": 1,
		"name": "fire place sensor",
		"state": "alerting",
		"message": "Someone is trying to break in through the fire place",
		"newStateDate": "2018-05-14T05:55:20+02:00",
		"evalDate": "0001-01-01T00:00:00Z",
		"executionError": "",
		"url": "http://grafana.com/dashboard/db/sensors"
	}`

	pauseAlertJSON = `{
		"alertId": 1,
		"state": "Paused",
		"message": "alert paused"
	}`
)

func TestAlerts(t *testing.T) {
	server, client := gapiTestTools(t, 200, alertsJSON)
	defer server.Close()

	params := url.Values{}
	params.Add("dashboardId", "123")

	as, err := client.Alerts(params)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(as))

	if as[0].ID != 1 {
		t.Error("alerts response should contain alerts with an ID")
	}
}

func TestAlerts_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, alertsJSON)
	defer server.Close()

	params := url.Values{}
	params.Add("dashboardId", "123")

	_, err := client.Alerts(params)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}

func TestAlert(t *testing.T) {
	server, client := gapiTestTools(t, 200, alertJSON)
	defer server.Close()

	res, err := client.Alert(1)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.ID != 1 {
		t.Error("alert response should contain the ID of the queried alert")
	}
}

func TestAlert_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, alertJSON)
	defer server.Close()

	_, err := client.Alert(1)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}

func TestPauseAlert(t *testing.T) {
	server, client := gapiTestTools(t, 200, pauseAlertJSON)
	defer server.Close()

	res, err := client.PauseAlert(1)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.State != "Paused" {
		t.Error("pause alert response should contain the correct response message")
	}
}

func TestPauseAlert_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, pauseAlertJSON)
	defer server.Close()

	_, err := client.PauseAlert(1)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}
