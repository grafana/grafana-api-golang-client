package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getLibraryPanelNameResponse = `{
		"result": [
	    {
				"id": 25,
				"orgId": 1,
				"folderId": 0,
				"uid": "V--OrYHnz",
				"name": "API docs Example",
				"kind": 1,
				"model": {
					"description": "",
					"type": ""
				},
				"version": 1,
				"meta": {
					"folderName": "General",
					"folderUid": "",
					"connectedDashboards": 1,
					"created": "2021-09-27T09:56:17+02:00",
					"updated": "2021-09-27T09:56:17+02:00",
					"createdBy": {
						"id": 1,
						"name": "admin",
						"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
					},
					"updatedBy": {
						"id": 1,
						"name": "admin",
						"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
					}
				}
	 		}
		]
	}`

	getLibraryPanelUIDResponse = `{
		"result": {
			"id": 25,
			"orgId": 1,
			"folderId": 0,
			"uid": "V--OrYHnz",
			"name": "API docs Example",
			"kind": 1,
			"model": {
				"description": "",
				"type": ""
			},
			"version": 1,
			"meta": {
				"folderName": "General",
				"folderUid": "",
				"connectedDashboards": 1,
				"created": "2021-09-27T09:56:17+02:00",
				"updated": "2021-09-27T09:56:17+02:00",
				"createdBy": {
					"id": 1,
					"name": "admin",
					"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
				},
				"updatedBy": {
					"id": 1,
					"name": "admin",
					"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
				}
			}
	 	}
	}`

	patchLibraryPanelResponse = `{
		"result": {
			"id": 25,
			"orgId": 1,
			"folderId": 0,
			"uid": "V--OrYHnz",
			"name": "Updated library panel name",
			"kind": 1,
			"model": {
				"description": "new description",
				"type": ""
			},
			"version": 1,
			"meta": {
				"folderName": "General",
				"folderUid": "",
				"connectedDashboards": 1,
				"created": "2021-09-27T09:56:17+02:00",
				"updated": "2021-09-27T09:56:17+02:00",
				"createdBy": {
					"id": 1,
					"name": "admin",
					"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
				},
				"updatedBy": {
					"id": 1,
					"name": "admin",
					"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
				}
			}
	 	}
	}`

	deleteLibraryPanelResponse = `{
		"message": "Library element deleted",
		"id": 28
	}`

	getLibraryPanelConnectionsResponse = `{
    "result": [
      {
        "id": 148,
        "kind": 1,
        "elementId": 25,
        "connectionId": 527,
        "created": "2021-09-27T10:00:07+02:00",
        "createdBy": {
            "id": 1,
            "name": "admin",
            "avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56"
        }
      }
    ]
	}`

	getLibraryPanelConnectedDashboardsResponse = `[
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
			"id":2,
			"uid": "SjkkcIBgc",
			"title":"Production Overview 2",
			"url": "/d/SjkkcIBgc/production-overview-2",
			"type":"dash-db",
			"tags":["prod"],
			"isStarred":true,
			"uri":"db/production-overview"
		}
	]`
)

func TestLibraryPanelCreate(t *testing.T) {
	client := gapiTestTools(t, 200, getLibraryPanelUIDResponse)

	panel := LibraryPanel{
		Folder: 0,
		Name:   "API docs Example",
		Model:  map[string]interface{}{"description": "", "type": ""},
	}

	resp, err := client.NewLibraryPanel(panel)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.UID != "V--OrYHnz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.UID, "V--OrYHnz")
	}

	for _, code := range []int{400, 401, 403} {
		client = gapiTestTools(t, code, "error")
		_, err = client.NewLibraryPanel(panel)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryPanelGetByName(t *testing.T) {
	client := gapiTestTools(t, 200, getLibraryPanelNameResponse)

	resp, err := client.LibraryPanelByName("API docs Example")
	if err != nil {
		t.Error(err)
	}
	if resp.UID != "V--OrYHnz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.UID, "V--OrYHnz")
	}

	for _, code := range []int{401, 403, 404} {
		client = gapiTestTools(t, code, "error")
		_, err = client.LibraryPanelByName("test")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryPanelGetByUID(t *testing.T) {
	client := gapiTestTools(t, 200, getLibraryPanelUIDResponse)

	resp, err := client.LibraryPanelByUID("V--OrYHnz")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Name != "API docs Example" {
		t.Fatalf("Invalid Name - %s, Expected %s", resp.Name, "API docs Example")
	}

	for _, code := range []int{401, 403, 404} {
		client = gapiTestTools(t, code, "error")
		_, err = client.LibraryPanelByUID("V--OrYHnz")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestPatchLibraryPanel(t *testing.T) {
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{code: 200, body: getLibraryPanelUIDResponse},
		{code: 200, body: patchLibraryPanelResponse},
	})

	panel := LibraryPanel{
		Folder: 1,
		Name:   "Updated library panel name",
		Model:  map[string]interface{}{"description": "new description", "type": ""},
	}
	resp, err := client.PatchLibraryPanel("V--OrYHnz", panel)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Name != "Updated library panel name" {
		t.Fatalf("Invalid Name - %s, Expected %s", resp.Name, "Updated library panel name")
	}
	if resp.Model["description"] != "new description" {
		t.Fatalf("Invalid panel JSON description - %s, Expected %s", resp.Name, "Updated library panel name")
	}

	for _, code := range []int{401, 403, 404} {
		client = gapiTestTools(t, code, "error")

		_, err := client.LibraryPanelByUID("V--OrYHnz")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryPanelGetConnections(t *testing.T) {
	client := gapiTestTools(t, 200, getLibraryPanelConnectionsResponse)

	resp, err := client.LibraryPanelConnections("V--OrYHnz")
	if err != nil {
		t.Fatal(err)
	}

	if (*resp)[0].ID != int64(148) {
		t.Fatalf("Invalid connection id - %d, Expected %d", (*resp)[0].ID, 148)
	}
}

func TestLibraryPanelConnectedDashboards(t *testing.T) {
	client := gapiTestTools(t, 200, getLibraryPanelConnectionsResponse)

	connections, err := client.LibraryPanelConnections("V--OrYHnz")
	if err != nil {
		t.Fatal(err)
	}

	var dashboardIds []int64
	for _, connection := range *connections {
		dashboardIds = append(dashboardIds, connection.DashboardID)
	}

	client = gapiTestTools(t, 200, getLibraryPanelConnectedDashboardsResponse)
	dashboards, err := client.DashboardsByIDs(dashboardIds)
	if err != nil {
		t.Fatal(err)
	}

	if dashboards[0].Title != "Production Overview" {
		t.Fatalf("Invalid title from connected dashboard 0 - %s, Expected %s", dashboards[0].Title, "Production Overview")
	}
}

func TestLibraryPanelDelete(t *testing.T) {
	client := gapiTestTools(t, 200, deleteLibraryPanelResponse)

	resp, err := client.DeleteLibraryPanel("V--OrYHnz")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message != "Library element deleted" {
		t.Error("Failed to delete. Response should contain the correct response message")
	}
}
