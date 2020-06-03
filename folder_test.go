package gapi

import (
	"testing"

	"github.com/gobs/pretty"
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
	server, client := gapiTestTools(200, getFoldersJSON)
	defer server.Close()

	folders, err := client.Folders()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(folders))

	if len(folders) != 1 {
		t.Error("Length of returned folders should be 1")
	}
	if folders[0].Id != 1 || folders[0].Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folders.")
	}
}

func TestFolder(t *testing.T) {
	server, client := gapiTestTools(200, getFolderJSON)
	defer server.Close()

	folder := int64(1)
	resp, err := client.Folder(folder)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Id != folder || resp.Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folder.")
	}
}

func TestNewFolder(t *testing.T) {
	server, client := gapiTestTools(200, createdFolderJSON)
	defer server.Close()

	resp, err := client.NewFolder("test-folder")
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Uid != "nErXDvCkzz" {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateFolder(t *testing.T) {
	server, client := gapiTestTools(200, updatedFolderJSON)
	defer server.Close()

	err := client.UpdateFolder("nErXDvCkzz", "test-folder")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFolder(t *testing.T) {
	server, client := gapiTestTools(200, deletedFolderJSON)
	defer server.Close()

	err := client.DeleteFolder("nErXDvCkzz")
	if err != nil {
		t.Error(err)
	}
}
