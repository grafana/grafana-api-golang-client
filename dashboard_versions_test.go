package gapi

import (
	"net/http"
	"strings"
	"testing"
)

const (
	// This must be wraoped in [] to make it a list.
	getDashboardVersionsResponse = `{
		"id": 2,
		"dashboardId": 1,
		"uid": "QA7wKklGz",
		"parentVersion": 1,
		"restoredFrom": 0,
		"version": 2,
		"created": "2017-06-08T17:24:33-04:00",
		"createdBy": "admin",
		"message": "Updated panel title"
	  }`
	getDashboardVersionResponse = `{
		"id": 1,
		"dashboardId": 1,
		"uid": "QA7wKklGz",
		"parentVersion": 0,
		"restoredFrom": 0,
		"version": 1,
		"created": "2017-04-26T17:18:38-04:00",
		"message": "Initial save",
		"data": {
		  "annotations": {"list": []},
		  "editable": true,
		  "gnetId": null,
		  "graphTooltip": 0,
		  "id": 1,
		  "links": [],
		  "rows": [
			{
			  "collapse": false,
			  "height": "250px",
			  "panels": [],
			  "repeat": null,
			  "repeatIteration": null,
			  "repeatRowId": null,
			  "showTitle": false,
			  "title": "Dashboard Row",
			  "titleSize": "h6"
			}
		  ],
		  "schemaVersion": 14,
		  "style": "dark",
		  "tags": [  ],
		  "templating": {"list": []},
		  "time": {"from": "now-6h","to": "now"},
		  "timepicker": {
			"refresh_intervals": ["5s","10s","30s","1m","5m","15m","30m","1h","2h","1d"],
			"time_options": ["5m","15m","1h","6h","12h","24h","2d","7d","30d"]
		  },
		  "timezone": "browser",
		  "title": "test",
		  "version": 1
		},
		"createdBy": "admin"
	  }`
	restoreDashboardResponse = `{
		"id": 70,
		"slug": "my-dashboard",
		"status": "success",
		"uid": "QA7wKklGz",
		"url": "/d/QA7wKklGz/my-dashboard",
		"version": 54
	  }`
	compareDashboardsResponse = `<html>stuff</html>`
)

func TestGetDashboardVersions(t *testing.T) {
	t.Parallel()

	const count = 20

	versions := strings.Repeat(getDashboardVersionsResponse+",", count)
	client := gapiTestToolsFromCalls(t, []mockServerCall{{
		code:      200,
		body:      "[" + strings.TrimSuffix(versions, ",") + "]",
		reqURI:    "/api/dashboards/uid/QA7wKklGz/versions?limit=20&start=0",
		reqMethod: http.MethodGet,
	}})

	dashboardVersions, err := client.GetDashboardVersions("QA7wKklGz", count, 0)
	if err != nil {
		t.Errorf("did not expect an error: %v", err)
	}

	if len(dashboardVersions) != count {
		t.Errorf("wrong dashboard version count returned, expected %d, got %d", count, len(dashboardVersions))
	}
}

func TestGetDashboardVersion(t *testing.T) {
	t.Parallel()
	client := gapiTestToolsFromCalls(t, []mockServerCall{{
		code:      200,
		body:      getDashboardVersionResponse,
		reqURI:    "/api/dashboards/uid/QA7wKklGz/versions/45",
		reqMethod: http.MethodGet,
	}})

	dashboard, err := client.GetDashboardVersion("QA7wKklGz", 45)
	if err != nil {
		t.Errorf("did not expect an error from dashboard version: %v", err)
	}

	if dashboard == nil {
		t.Fatal("dashboard version must not be nil")
	}

	if dashboard.CreatedBy != "admin" {
		t.Error("dashboard version data seems invalid")
	}
}

func TestRestoreDashboardVersion(t *testing.T) {
	t.Parallel()
	client := gapiTestToolsFromCalls(t, []mockServerCall{{
		code:      200,
		body:      restoreDashboardResponse,
		reqURI:    "/api/dashboards/uid/QA7wKklGz/restore",
		reqMethod: http.MethodPost,
		reqBody:   `{"version":54}`,
	}})

	dashboard, err := client.RestoreDashboardVersion("QA7wKklGz", 54)
	if err != nil {
		t.Errorf("did not expect an error from dashboard restore: %v", err)
	}

	if dashboard == nil {
		t.Fatal("dashboard restore data must not be nil")
	}

	if dashboard.Version != 54 {
		t.Error("restored dashboard version data seems invalid")
	}
}

func TestCompareDashboardVersions(t *testing.T) {
	t.Parallel()
	client := gapiTestToolsFromCalls(t, []mockServerCall{{
		code:      200,
		body:      compareDashboardsResponse,
		reqURI:    "/api/dashboards/calculate-diff",
		reqMethod: http.MethodPost,
		reqBody:   `{"base":{"dashboardId":1,"version":2},"new":{"dashboardId":3,"version":4},"diffType":"basic"}`,
	}})

	compare, err := client.CompareDashboardVersions(CompareDashboardsInput{
		BaseDashboardID:      1,
		BaseDashboardVersion: 2,
		NewDashboardID:       3,
		NewDashboardVersion:  4,
	})
	if err != nil {
		t.Errorf("did not expect an error from dashboard compare: %v", err)
	}

	if string(compare) != compareDashboardsResponse {
		t.Error("got wrong responnse from dashboard compare")
	}
}
