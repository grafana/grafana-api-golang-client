package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getDashboardsJSON = `
[
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
]
	`
)

func TestDashboards(t *testing.T) {
	server, client := gapiTestTools(200, getDashboardsJSON)
	defer server.Close()

	dashboards, err := client.Dashboards()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(dashboards))

	if len(dashboards) != 1 {
		t.Error("Length of returned dashboards should be 1")
	}

	if dashboards[0].Id != 1 || dashboards[0].Title != "Grafana Stats" {
		t.Error("Not correctly parsing returned dashboards.")
	}
}
