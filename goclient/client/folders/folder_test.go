package folders

import (
	"testing"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	getFoldersJSON = `{
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
	}`
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

// func TestListFolders(t *testing.T) {
// 	mockData := strings.Repeat(getFoldersJSON+",", 1000) // make 1000 folders.
// 	mockData = "[" + mockData[:len(mockData)-1] + "]"    // remove trailing comma; make a json list.

// 	// This creates 1000 + 1000 + 1 (2001, 3 calls) worth of folders.
// 	client := client.GetClient(t, []mockServerCall{
// 		{200, mockData},
// 		{200, mockData},
// 		{200, "[" + getFolderJSON + "]"},
// 	})

// 	const dashCount = 2001

// 	folders, err := client.Folders()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(pretty.PrettyFormat(folders))

// 	if len(folders) != dashCount {
// 		t.Errorf("Length of returned folders should be %d", dashCount)
// 	}
// 	if folders[0].ID != 1 || folders[0].Title != "Departmenet ABC" {
// 		t.Error("Not correctly parsing returned folders.")
// 	}
// 	if folders[dashCount-1].ID != 1 || folders[dashCount-1].Title != "Departmenet ABC" {
// 		t.Error("Not correctly parsing returned folders.")
// 	}
// }

// func TestGetFolder(t *testing.T) {
// 	client := client.GetClient(t, 200, getFolderJSON)

// 	folder := int64(1)
// 	resp, err := client.Folders.GetFolderByID(folder)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(pretty.PrettyFormat(resp))

// 	if resp.ID != folder || resp.Title != "Departmenet ABC" {
// 		t.Error("Not correctly parsing returned folder.")
// 	}
// }

// func TestGetFolderByUid(t *testing.T) {
// 	client := client.GetClient(t, 200, getFolderJSON)

// 	folder := "nErXDvCkzz"
// 	resp, err := client.Folders.GetFolderByUID(folder)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(pretty.PrettyFormat(resp))

// 	if resp.UID != folder || resp.Title != "Departmenet ABC" {
// 		t.Error("Not correctly parsing returned folder.")
// 	}
// }

func TestCreateFolder(t *testing.T) {
	client := GetClient(t, 200, createdFolderJSON, nil)

	const UID = "nErXDvCkzz"
	params := NewCreateFolderParams().
		WithDefaults().
		WithBody(&models.CreateFolderCommand{UID: UID})

	ok, err := client.CreateFolder(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	if ok.Payload.UID != UID {
		t.Errorf("UID does not match expected value; expected: %s, got: %s", UID, ok.Payload.UID)
	}
}

// func TestUpdateFolder(t *testing.T) {
// 	client := client.GetClient(t, 200, updatedFolderJSON)

// 	err := client.UpdateFolder("nErXDvCkzz", "test-folder")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestDeleteFolder(t *testing.T) {
// 	client := client.GetClient(t, 200, deletedFolderJSON)

// 	err := client.DeleteFolder("nErXDvCkzz")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
