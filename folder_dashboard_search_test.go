package gapi

import (
	"net/url"
	"testing"
)

const (

	// This response is copied from the examples in the API docs:
	// https://grafana.com/docs/grafana/latest/http_api/folder_dashboard_search/
	getFolderDashboardSearchResponse = `[
		{
			"id": 163,
			"uid": "000000163",
			"title": "Folder",
			"url": "/dashboards/f/000000163/folder",
			"type": "dash-folder",
			"tags": [],
			"isStarred": false,
			"uri":"db/folder"
		},
		{
			"id":1,
			"uid": "cIBgcSjkk",
			"title":"Production Overview",
			"url": "/d/cIBgcSjkk/production-overview",
			"type":"dash-db",
			"tags":["prod"],
			"isStarred":true,
			"uri":"db/production-overview"
		},
		{
			"id":1,
			"uid": "cIBgcSjkk",
			"title":"Production Overview",
			"url": "/d/cIBgcSjkk/production-overview",
			"type":"dash-db",
			"tags":["prod"],
			"isStarred":true,
			"folderId": 2,
			"folderUid": "000000163",
			"folderTitle": "Folder",
			"folderUrl": "/dashboards/f/000000163/folder",
			"uri":"db/production-overview"
		}
	]`
)

func TestFolderDashboardSearch(t *testing.T) {
	client := gapiTestTools(t, 200, getFolderDashboardSearchResponse)
	resp, err := client.FolderDashboardSearch(url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp) != 3 {
		t.Errorf("Expected 3 objects in response, got %d", len(resp))
	}
	if resp[0].ID != 163 || resp[0].Title != "Folder" {
		t.Error("Not correctly parsing response.")
	}
}
