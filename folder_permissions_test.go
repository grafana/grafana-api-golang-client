package gapi

import (
	"github.com/gobs/pretty"
	"testing"
)

const (
	getFolderPermissionsJSON = `
[
  {
    "id": 1,
    "folderId": -1,
    "created": "2017-06-20T02:00:00+02:00",
    "updated": "2017-06-20T02:00:00+02:00",
    "userId": 0,
    "userLogin": "",
    "userEmail": "",
    "teamId": 0,
    "team": "",
    "role": "Viewer",
    "permission": 1,
    "permissionName": "View",
    "uid": "nErXDvCkzz",
    "title": "",
    "slug": "",
    "isFolder": false,
    "url": ""
  },
  {
    "id": 2,
    "dashboardId": -1,
    "created": "2017-06-20T02:00:00+02:00",
    "updated": "2017-06-20T02:00:00+02:00",
    "userId": 0,
    "userLogin": "",
    "userEmail": "",
    "teamId": 0,
    "team": "",
    "role": "Editor",
    "permission": 2,
    "permissionName": "Edit",
    "uid": "",
    "title": "",
    "slug": "",
    "isFolder": false,
    "url": ""
  }
]
`
	updateFolderPermissionsJSON = `
{
	"message": "Folder permissions updated"
}
`
)

func TestFolderPermissions(t *testing.T) {
	server, client := gapiTestTools(200, getFolderPermissionsJSON)
	defer server.Close()

	fid := "nErXDvCkzz"
	resp, err := client.FolderPermissions(fid)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*FolderPermission{
		{
			Id:             1,
			FolderUid:      "nErXDvCkzz",
			UserId:         0,
			TeamId:         0,
			Role:           "Viewer",
			IsFolder:       false,
			Permission:     1,
			PermissionName: "View",
			FolderId:       -1,
			DashboardId:    0,
		},
		{
			Id:             2,
			FolderUid:      "",
			UserId:         0,
			TeamId:         0,
			Role:           "Editor",
			IsFolder:       false,
			Permission:     2,
			PermissionName: "Edit",
			FolderId:       0,
			DashboardId:    -1,
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if resp[i].Id != expect.Id || resp[i].Role != expect.Role {
				t.Error("Not correctly data")
			}
		})
	}
}

func TestUpdateFolderPermissions(t *testing.T) {
	server, client := gapiTestTools(200, updateFolderPermissionsJSON)
	defer server.Close()

	items := &PermissionItems{
		Items: []*PermissionItem{
			{
				Role:       "viewer",
				Permission: 1,
			},
			{
				Role:       "Editor",
				Permission: 2,
			},
			{
				TeamId:     1,
				Permission: 1,
			},
			{
				UserId:     11,
				Permission: 4,
			},
		},
	}
	err := client.UpdateFolderPermissions("nErXDvCkzz", items)
	if err != nil {
		t.Error(err)
	}
}
