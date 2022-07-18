package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/folder_permissions"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
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
	server, client := gapiTestTools(t, 200, getFolderPermissionsJSON)
	defer server.Close()

	resp, err := client.FolderPermissions.GetFolderPermissions(
		folder_permissions.NewGetFolderPermissionsParams().
			WithFolderUID("nErXDvCkzz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*models.DashboardACLInfoDTO{
		{
			UID:            "perm-1",
			UserID:         0,
			TeamID:         0,
			Role:           "Viewer",
			IsFolder:       false,
			Permission:     1,
			PermissionName: "View",
			FolderID:       -1,
			DashboardID:    0,
		},
		{
			UID:            "perm-2",
			UserID:         0,
			TeamID:         0,
			Role:           "Editor",
			IsFolder:       false,
			Permission:     2,
			PermissionName: "Edit",
			FolderID:       0,
			DashboardID:    -1,
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if resp.Payload[i].UID != expect.UID || resp.Payload[i].Role != expect.Role {
				t.Error("Not correctly parsing returned folder permission")
			}
		})
	}
}

func TestUpdateFolderPermissions(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateFolderPermissionsJSON)
	defer server.Close()

	cmd := models.UpdateDashboardACLCommand{
		Items: []*models.DashboardACLUpdateItem{
			{
				Role:       "viewer",
				Permission: 1,
			},
			{
				Role:       "Editor",
				Permission: 2,
			},
			{
				TeamID:     1,
				Permission: 1,
			},
			{
				UserID:     11,
				Permission: 4,
			},
		},
	}
	_, err := client.FolderPermissions.UpdateFolderPermissions(
		folder_permissions.NewUpdateFolderPermissionsParams().
			WithFolderUID("nErXDvCkzz").
			WithBody(&cmd),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
