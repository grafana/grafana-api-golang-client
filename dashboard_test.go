package gapi

import (
	"testing"

	"github.com/gobs/pretty"
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
	client := gapiTestTools(t, 200, createdAndUpdateDashboardResponse)

	dashboard := Dashboard{
		Model: map[string]interface{}{
			"title": "test",
		},
		FolderID:  0,
		FolderUID: "l3KqBxCMz",
		Overwrite: false,
	}

	resp, err := client.NewDashboard(dashboard)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.UID != "nErXDvCkzz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.UID, "nErXDvCkzz")
	}

	for _, code := range []int{400, 401, 403, 412} {
		client = gapiTestTools(t, code, "error")
		_, err = client.NewDashboard(dashboard)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboardGet(t *testing.T) {
	client := gapiTestTools(t, 200, getDashboardResponse)

	resp, err := client.Dashboard("test")
	if err != nil {
		t.Error(err)
	}
	uid, ok := resp.Model["uid"]
	if !ok || uid != "cIBgcSjkk" {
		t.Errorf("Invalid uid - %s, Expected %s", uid, "cIBgcSjkk")
	}

	client = gapiTestTools(t, 200, getDashboardResponse)

	resp, err = client.DashboardByUID("cIBgcSjkk")
	if err != nil {
		t.Fatal(err)
	}
	uid, ok = resp.Model["uid"]
	if !ok || uid != "cIBgcSjkk" {
		t.Fatalf("Invalid UID - %s, Expected %s", uid, "cIBgcSjkk")
	}

	for _, code := range []int{401, 403, 404} {
		client = gapiTestTools(t, code, "error")
		_, err = client.Dashboard("test")
		if err == nil {
			t.Errorf("%d not detected", code)
		}

		_, err = client.DashboardByUID("cIBgcSjkk")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboardDelete(t *testing.T) {
	client := gapiTestTools(t, 200, "")
	err := client.DeleteDashboard("test")
	if err != nil {
		t.Error(err)
	}

	client = gapiTestTools(t, 200, "")
	err = client.DeleteDashboardByUID("cIBgcSjkk")
	if err != nil {
		t.Fatal(err)
	}

	for _, code := range []int{401, 403, 404, 412} {
		client = gapiTestTools(t, code, "error")

		err = client.DeleteDashboard("test")
		if err == nil {
			t.Errorf("%d not detected", code)
		}

		err = client.DeleteDashboardByUID("cIBgcSjkk")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboards(t *testing.T) {
	client := gapiTestTools(t, 200, getDashboardsJSON)

	dashboards, err := client.Dashboards()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(dashboards))

	if len(dashboards) != 1 {
		t.Error("Length of returned dashboards should be 1")
	}

	if dashboards[0].ID != 1 || dashboards[0].Title != "Grafana Stats" {
		t.Error("Not correctly parsing returned dashboards.")
	}
}
