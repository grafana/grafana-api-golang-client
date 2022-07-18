package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/dashboards"
	"github.com/grafana/grafana-api-golang-client/goclient/client/search"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	createdAndUpdateDashboardResponse = `{
		"slug": "test",
		"id": 1,
		"uid": "nErXDvCkzz",
		"status": "success",
		"version": 1
	}`

	getDashboardResponse = `{
		"dashboard": {
			"id": 1,
			"uid": "cIBgcSjkk",
			"title": "Production Overview",
			"version": 0
		},
		"meta": {
			"isStarred": false,
			"url": "/d/cIBgcSjkk/production-overview",
			"slug": "production-overview"
		}
	}`

	getDashboardsJSON = `[
    {
      "id": 1,
      "uid": "RGAPB1cZz",
      "title": "Grafana Stats",
      "uri": "db/grafana-stats",
      "url": "/dashboards/d/RGAPB1cZz/grafana-stat",
      "slug": "",
      "type": "dash-db",
      "tags": [],
      "isStarred": false
    }
  ]`
)

func TestDashboardCreateAndUpdate(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdAndUpdateDashboardResponse)
	defer server.Close()

	dashboard := models.SaveDashboardCommand{
		Dashboard: map[string]interface{}{
			"title": "test",
		},
		FolderID:  0,
		Overwrite: false,
	}

	resp, err := client.Dashboards.PostDashboard(
		dashboards.NewPostDashboardParams().WithBody(&dashboard),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if *resp.Payload.UID != "nErXDvCkzz" {
		t.Errorf("Invalid uid - %s, Expected %s", *resp.Payload.UID, "nErXDvCkzz")
	}

	for _, code := range []int{400, 401, 403, 412} {
		server.code = code
		_, err := client.Dashboards.PostDashboard(
			dashboards.NewPostDashboardParams().WithBody(&dashboard),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboardGet(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDashboardResponse)
	defer server.Close()

	resp, err := client.Dashboards.GetDashboardByUID(
		dashboards.NewGetDashboardByUIDParams().WithUID("cIBgcSjkk"),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	dashboardData, ok := resp.Payload.Dashboard.(map[string]interface{})

	uid, ok := dashboardData["uid"]
	if !ok || uid != "cIBgcSjkk" {
		t.Errorf("Invalid uid - %s, Expected %s", uid, "cIBgcSjkk")
	}

	for _, code := range []int{401, 403, 404} {
		server.code = code

		_, err = client.Dashboards.GetDashboardByUID(
			dashboards.NewGetDashboardByUIDParams().WithUID("cIBgcSjkk"),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboardDelete(t *testing.T) {
	server, client := gapiTestTools(t, 200, "")
	defer server.Close()

	_, err := client.Dashboards.DeleteDashboardByUID(
		dashboards.NewDeleteDashboardByUIDParams().
			WithUID("cIBgcSjkk"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, code := range []int{401, 403, 404, 412} {
		server.code = code

		_, err := client.Dashboards.DeleteDashboardByUID(
			dashboards.NewDeleteDashboardByUIDParams().
				WithUID("cIBgcSjkk"),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboards(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDashboardsJSON)
	defer server.Close()

	typ := "dash-db"
	resp, err := client.Search.Search(
		search.NewSearchParams().WithType(&typ),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if len(resp.Payload) != 1 {
		t.Error("Length of returned dashboards should be 1")
	}

	if resp.Payload[0].ID != 1 || resp.Payload[0].Title != "Grafana Stats" {
		t.Error("Not correctly parsing returned dashboards.")
	}
}
