package gapi

import (
	"strings"
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/legacy_alerts"
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
	mocksrv, client := gapiTestTools(t, 200, alertsJSON)
	defer mocksrv.Close()

	as, err := client.LegacyAlerts.GetAlerts(
		legacy_alerts.NewGetAlertsParams().WithDashboardID([]string{"123"}),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(as))

	if as.Payload[0].ID != 1 {
		t.Error("alerts response should contain alerts with an ID")
	}
}

func TestAlerts_500(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 500, alertsJSON)
	defer mocksrv.Close()

	_, err := client.LegacyAlerts.GetAlerts(
		legacy_alerts.NewGetAlertsParams().WithDashboardID([]string{"123"}),
		nil,
	)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}

func TestAlert(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, alertJSON)
	defer mocksrv.Close()

	res, err := client.LegacyAlerts.GetAlertByID(
		legacy_alerts.NewGetAlertByIDParams().WithAlertID("1"),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.Payload.ID != 1 {
		t.Error("alert response should contain the ID of the queried alert")
	}
}

func TestAlert_500(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 500, alertJSON)
	defer mocksrv.Close()

	_, err := client.LegacyAlerts.GetAlertByID(
		legacy_alerts.NewGetAlertByIDParams().WithAlertID("1"),
		nil,
	)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}

func TestPauseAlert(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, pauseAlertJSON)
	defer mocksrv.Close()

	res, err := client.LegacyAlerts.PauseAlert(
		legacy_alerts.NewPauseAlertParams().WithAlertID("1"),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.Payload.State != "Paused" {
		t.Error("pause alert response should contain the correct response message")
	}
}

func TestPauseAlert_500(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 500, pauseAlertJSON)
	defer mocksrv.Close()

	_, err := client.LegacyAlerts.PauseAlert(
		legacy_alerts.NewPauseAlertParams().WithAlertID("1"),
		nil,
	)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}
