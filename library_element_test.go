package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/library_elements"
	"github.com/grafana/grafana-api-golang-client/goclient/client/search"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
	"github.com/stretchr/testify/require"
)

const (
	getLibraryElementNameResponse = `{
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

	getLibraryElementUIDResponse = `{
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

	patchLibraryElementResponse = `{
		"result": {
			"id": 25,
			"orgId": 1,
			"folderId": 0,
			"uid": "V--OrYHnz",
			"name": "Updated library element name",
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

	deleteLibraryElementResponse = `{
		"message": "Library element deleted",
		"id": 28
	}`

	getLibraryElementConnectionsResponse = `{
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

	getLibraryElementConnectedDashboardsResponse = `[
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

func TestLibraryElementCreate(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getLibraryElementUIDResponse)
	defer mocksrv.Close()

	element := models.CreateLibraryElementCommand{
		FolderID: 0,
		Name:     "API docs Example",
		Model:    map[string]interface{}{"description": "", "type": ""},
	}

	resp, err := client.LibraryElements.CreateLibraryElement(
		library_elements.NewCreateLibraryElementParams().
			WithBody(&element),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.Result.UID != "V--OrYHnz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.Payload.Result.UID, "V--OrYHnz")
	}

	for _, code := range []int{400, 401, 403} {
		mocksrv.code = code
		_, err := client.LibraryElements.CreateLibraryElement(
			library_elements.NewCreateLibraryElementParams().
				WithBody(&element),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryElementGetByName(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getLibraryElementNameResponse)
	defer mocksrv.Close()

	resp, err := client.LibraryElements.GetLibraryElementByName(
		library_elements.NewGetLibraryElementByNameParams().
			WithLibraryElementName("API docs Example"),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	if resp.Payload.Result.UID != "V--OrYHnz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.Payload.Result.UID, "V--OrYHnz")
	}

	for _, code := range []int{401, 403, 404} {
		mocksrv.code = code
		_, err := client.LibraryElements.GetLibraryElementByName(
			library_elements.NewGetLibraryElementByNameParams().
				WithLibraryElementName("API docs Example"),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryElementGetByUID(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getLibraryElementUIDResponse)
	defer mocksrv.Close()

	resp, err := client.LibraryElements.GetLibraryElementByUID(
		library_elements.NewGetLibraryElementByUIDParams().
			WithLibraryElementUID("V--OrYHnz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Payload.Result.Name != "API docs Example" {
		t.Fatalf("Invalid Name - %s, Expected %s", resp.Payload.Result.Name, "API docs Example")
	}

	for _, code := range []int{401, 403, 404} {
		mocksrv.code = code
		_, err := client.LibraryElements.GetLibraryElementByUID(
			library_elements.NewGetLibraryElementByUIDParams().
				WithLibraryElementUID("V--OrYHnz"),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestPatchLibraryElement(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, patchLibraryElementResponse)
	defer mocksrv.Close()

	element := models.PatchLibraryElementCommand{
		FolderID: 1,
		Name:     "Updated library element name",
		Model:    map[string]interface{}{"description": "new description", "type": ""},
	}
	resp, err := client.LibraryElements.UpdateLibraryElement(
		library_elements.NewUpdateLibraryElementParams().
			WithLibraryElementUID("V--OrYHnz").
			WithBody(&element),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Payload.Result.Name != "Updated library element name" {
		t.Fatalf("Invalid Name - %s, Expected %s", resp.Payload.Result.Name, "Updated library element name")
	}

	m, ok := resp.Payload.Result.Model.(map[string]interface{})
	require.True(t, ok)
	if m["description"] != "new description" {
		t.Fatalf("Invalid element JSON description - %s, Expected %s", m["description"], "Updated library element name")
	}

	for _, code := range []int{401, 403, 404} {
		mocksrv.code = code

		_, err := client.LibraryElements.UpdateLibraryElement(
			library_elements.NewUpdateLibraryElementParams().
				WithLibraryElementUID("V--OrYHnz").
				WithBody(&element),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryElementGetConnections(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getLibraryElementConnectionsResponse)
	defer mocksrv.Close()

	resp, err := client.LibraryElements.GetLibraryElementConnections(
		library_elements.NewGetLibraryElementConnectionsParams().
			WithLibraryElementUID("V--OrYHnz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Payload.Result[0].ID != int64(148) {
		t.Fatalf("Invalid connection id - %d, Expected %d", resp.Payload.Result[0].ID, 148)
	}
}

func TestLibraryElementConnectedDashboards(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getLibraryElementConnectionsResponse)
	defer mocksrv.Close()

	resp, err := client.LibraryElements.GetLibraryElementConnections(
		library_elements.NewGetLibraryElementConnectionsParams().
			WithLibraryElementUID("V--OrYHnz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	var dashboardIds []int64
	for _, connection := range resp.Payload.Result {
		dashboardIds = append(dashboardIds, connection.ConnectionID)
	}

	typ := "dash-db"
	_, client = gapiTestTools(t, 200, getLibraryElementConnectedDashboardsResponse)
	dashboardsResp, err := client.Search.Search(
		search.NewSearchParams().
			WithType(&typ).
			WithDashboardIds(dashboardIds),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if dashboardsResp.Payload[0].Title != "Production Overview" {
		t.Fatalf("Invalid title from connected dashboard 0 - %s, Expected %s", dashboardsResp.Payload[0].Title, "Production Overview")
	}
}

func TestLibraryElementDelete(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deleteLibraryElementResponse)
	defer mocksrv.Close()

	resp, err := client.LibraryElements.DeleteLibraryElementByUID(
		library_elements.NewDeleteLibraryElementByUIDParams().WithLibraryElementUID("V--OrYHnz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Payload.Message != "Library element deleted" {
		t.Error("Failed to delete. Response should contain the correct response message")
	}
}
