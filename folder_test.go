package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/folders"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	getFoldersJSON = `
[
  {
    "id":1,
    "uid": "nErXDvCkzz",
    "title": "Departmenet ABC",
    "url": "/dashboards/f/nErXDvCkzz/department-abc",
    "hasAcl": false,
    "canSave": true,
    "canEdit": true,
    "canAdmin": true,
    "createdBy": "admin",
    "created": "2018-01-31T17:43:12+01:00",
    "updatedBy": "admin",
    "updated": "2018-01-31T17:43:12+01:00",
    "version": 1
  }
]
	`
	getFolderJSON = `
{
  "id":1,
  "uid": "nErXDvCkzz",
  "title": "Departmenet ABC",
  "url": "/dashboards/f/nErXDvCkzz/department-abc",
  "hasAcl": false,
  "canSave": true,
  "canEdit": true,
  "canAdmin": true,
  "createdBy": "admin",
  "created": "2018-01-31T17:43:12+01:00",
  "updatedBy": "admin",
  "updated": "2018-01-31T17:43:12+01:00",
  "version": 1
}
`
	createdFolderJSON = `
{
  "id":1,
  "uid": "nErXDvCkzz",
  "title": "Departmenet ABC",
  "url": "/dashboards/f/nErXDvCkzz/department-abc",
  "hasAcl": false,
  "canSave": true,
  "canEdit": true,
  "canAdmin": true,
  "createdBy": "admin",
  "created": "2018-01-31T17:43:12+01:00",
  "updatedBy": "admin",
  "updated": "2018-01-31T17:43:12+01:00",
  "version": 1
}
`
	updatedFolderJSON = `
{
  "id":1,
  "uid": "nErXDvCkzz",
  "title": "Departmenet DEF",
  "url": "/dashboards/f/nErXDvCkzz/department-def",
  "hasAcl": false,
  "canSave": true,
  "canEdit": true,
  "canAdmin": true,
  "createdBy": "admin",
  "created": "2018-01-31T17:43:12+01:00",
  "updatedBy": "admin",
  "updated": "2018-01-31T17:43:12+01:00",
  "version": 1
}
`
	deletedFolderJSON = `
{
  "message":"Folder deleted"
}
`
)

func TestFolders(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getFoldersJSON)
	defer mocksrv.Close()

	folders, err := client.Folders.GetFolders(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(folders))

	if len(folders.Payload) != 1 {
		t.Error("Length of returned folders should be 1")
	}
	if folders.Payload[0].ID != 1 || folders.Payload[0].Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folders.")
	}
}

func TestFolder(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getFolderJSON)
	defer mocksrv.Close()

	folder := int64(1)
	respByID, err := client.Folders.GetFolderByID(
		folders.NewGetFolderByIDParams().WithFolderID(folder),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(respByID))

	if respByID.Payload.ID != folder || respByID.Payload.Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folder.")
	}
}

func TestFolderByUid(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getFolderJSON)
	defer mocksrv.Close()

	folder := "nErXDvCkzz"
	resp, err := client.Folders.GetFolderByUID(
		folders.NewGetFolderByUIDParams().WithFolderUID(folder),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.UID != folder || resp.Payload.Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folder.")
	}
}

func TestNewFolder(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createdFolderJSON)
	defer mocksrv.Close()

	resp, err := client.Folders.CreateFolder(
		folders.NewCreateFolderParams().
			WithBody(&models.CreateFolderCommand{
				Title: "test-folder",
			}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.UID != "nErXDvCkzz" {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateFolder(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updatedFolderJSON)
	defer mocksrv.Close()

	_, err := client.Folders.UpdateFolder(
		folders.NewUpdateFolderParams().
			WithFolderUID("nErXDvCkzz").
			WithBody(&models.UpdateFolderCommand{
				Title: "test-folder",
			}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteFolder(t *testing.T) {
	server, client := gapiTestTools(t, 200, deletedFolderJSON)
	defer server.Close()

	_, err := client.Folders.DeleteFolder(
		folders.NewDeleteFolderParams().WithFolderUID("nErXDvCkzz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}
