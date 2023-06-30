package folders

import (
	"testing"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	testUID        = "nErXDvCkzz"
	testTitle      = "test-folder"
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

func TestGetFolder(t *testing.T) {
	client := GetClient(t, 200, getFolderJSON)

	folder := int64(1)
	params := NewGetFolderByIDParams().WithFolderID(folder)
	resp, err := client.GetFolderByID(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Payload.ID != folder || resp.Payload.Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folder.")
	}
}

func TestGetFolderByUid(t *testing.T) {
	client := GetClient(t, 200, getFolderJSON)

	params := NewGetFolderByUIDParams().WithFolderUID(testUID)
	resp, err := client.GetFolderByUID(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Payload.UID != testUID || resp.Payload.Title != "Departmenet ABC" {
		t.Error("Not correctly parsing returned folder.")
	}
}

func TestCreateFolder(t *testing.T) {
	client := GetClient(t, 200, createdFolderJSON)

	params := NewCreateFolderParams().
		WithBody(&models.CreateFolderCommand{UID: testUID})

	ok, err := client.CreateFolder(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	if ok.Payload.UID != testUID {
		t.Errorf("UID does not match expected value; expected: %s, got: %s", testUID, ok.Payload.UID)
	}
}

func TestUpdateFolder(t *testing.T) {
	client := GetClient(t, 200, updatedFolderJSON)

	params := NewUpdateFolderParams().
		WithBody(&models.UpdateFolderCommand{
			UID:   testUID,
			Title: testTitle,
		})
	ok, err := client.UpdateFolder(params, nil)
	if err != nil {
		t.Fatal(err)
	}
	if ok.Payload.Title != "Departmenet DEF" || ok.Payload.UID != testUID {
		t.Errorf("expected Title %s and UID %s, got Title %s and UID %s", "Departmenet DEF", testUID, ok.Payload.Title, ok.Payload.UID)
	}
}

func TestDeleteFolder(t *testing.T) {
	client := GetClient(t, 200, deletedFolderJSON)

	params := NewDeleteFolderParams().WithFolderUID(testUID)
	ok, err := client.DeleteFolder(params, nil)
	if err != nil {
		t.Fatal(err)
	}
	msg := "Folder deleted"
	if *ok.Payload.Message != msg {
		t.Errorf("expected message '%s', got: %s", msg, *ok.Payload.Message)
	}
}
