package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getLibraryPanelResponse = `{
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
)

func TestLibraryPanelCreate(t *testing.T) {
	server, client := gapiTestTools(t, 200, getLibraryPanelResponse)
	defer server.Close()

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
		server.code = code
		_, err = client.NewLibraryPanel(panel)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryPanelGet(t *testing.T) {
	server, client := gapiTestTools(t, 200, getLibraryPanelResponse)
	defer server.Close()

	resp, err := client.LibraryPanelByName("API docs Example")
	if err != nil {
		t.Error(err)
	}
	if resp.UID != "V--OrYHnz" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.UID, "V--OrYHnz")
	}

	resp, err = client.LibraryPanelByUID("V--OrYHnz")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Name != "API docs Example" {
		t.Fatalf("Invalid Name - %s, Expected %s", resp.Name, "API docs Example")
	}

	for _, code := range []int{401, 403, 404} {
		server.code = code
		_, err = client.LibraryPanelByName("test")
		if err == nil {
			t.Errorf("%d not detected", code)
		}

		_, err = client.LibraryPanelByUID("V--OrYHnz")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestPatchLibraryPanel(t *testing.T) {
	server, client := gapiTestTools(t, 200, patchLibraryPanelResponse)
	defer server.Close()

	panel := &LibraryPanel{
		Folder: 1,
		Name:   "Updated library panel name",
		Model:  map[string]interface{}{"description": "new description", "type": ""},
	}
	resp, err := client.PatchLibraryPanelByUID("V--OrYHnz", panel)
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
		server.code = code

		_, err := client.LibraryPanelByUID("V--OrYHnz")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestLibraryPanelDelete(t *testing.T) {
	server, client := gapiTestTools(t, 200, deleteLibraryPanelResponse)
	defer server.Close()

	resp, err := client.DeleteLibraryPanel("V--OrYHnz")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Message != "Library element deleted" {
		t.Error("Failed to delete. Response should contain the correct response message")
	}
}
