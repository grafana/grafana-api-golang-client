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
)

func TestDashboardCreateAndUpdate(t *testing.T) {
	server, client := gapiTestTools(200, createdAndUpdateDashboardResponse)
	defer server.Close()

	dashboard := Dashboard{
		Model: map[string]interface{}{
			"title": "test",
		},
		Folder:    0,
		Overwrite: false,
	}

	resp, err := client.NewDashboard(dashboard)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Uid != "nErXDvCkzz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.Uid, "nErXDvCkzz")
	}

	for _, code := range []int{400, 401, 403, 412} {
		server.code = code
		_, err = client.NewDashboard(dashboard)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestDashboardGet(t *testing.T) {
	server, client := gapiTestTools(200, getDashboardResponse)
	defer server.Close()

	resp, err := client.Dashboard("test")
	if err != nil {
		t.Error(err)
	}
	uid, ok := resp.Model["uid"]
	if !ok || uid != "cIBgcSjkk" {
		t.Errorf("Invalid uid - %s, Expected %s", uid, "cIBgcSjkk")
	}

	resp, err = client.DashboardByUID("cIBgcSjkk")
	if err != nil {
		t.Error(err)
	}
	uid, ok = resp.Model["uid"]
	if !ok || uid != "cIBgcSjkk" {
		t.Errorf("Invalid uid - %s, Expected %s", uid, "cIBgcSjkk")
	}

	for _, code := range []int{401, 403, 404} {
		server.code = code
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
	server, client := gapiTestTools(200, "")
	defer server.Close()

	err := client.DeleteDashboard("test")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteDashboardByUID("cIBgcSjkk")
	if err != nil {
		t.Error(err)
	}

	for _, code := range []int{401, 403, 404, 412} {
		server.code = code

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
