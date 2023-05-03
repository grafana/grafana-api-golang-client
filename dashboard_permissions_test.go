package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getDashboardPermissionsJSON = `
[
  {
    "dashboardId": 1,
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
    "url": "",
    "inherited": false
  },
  {
    "dashboardId": 2,
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
    "uid": "nErXDvCkzz",
    "title": "",
    "slug": "",
    "isFolder": false,
    "url": "",
    "inherited": true
  }
]
`
	updateDashboardPermissionsJSON = `
{
	"message": "Dashboard permissions updated"
}
`
)

func TestDashboardPermissions(t *testing.T) {
	client := gapiTestTools(t, 200, getDashboardPermissionsJSON)

	resp, err := client.DashboardPermissions(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*DashboardPermission{
		{
			DashboardID:    1,
			DashboardUID:   "nErXDvCkzz",
			Role:           "Viewer",
			UserID:         0,
			TeamID:         0,
			IsFolder:       false,
			Inherited:      false,
			Permission:     1,
			PermissionName: "View",
		},
		{
			DashboardID:    2,
			DashboardUID:   "nErXDvCkzz",
			Role:           "Editor",
			UserID:         0,
			TeamID:         0,
			IsFolder:       false,
			Inherited:      true,
			Permission:     2,
			PermissionName: "Edit",
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if resp[i].DashboardID != expect.DashboardID ||
				resp[i].DashboardUID != expect.DashboardUID ||
				resp[i].Role != expect.Role ||
				resp[i].UserID != expect.UserID ||
				resp[i].TeamID != expect.TeamID ||
				resp[i].IsFolder != expect.IsFolder ||
				resp[i].Inherited != expect.Inherited ||
				resp[i].Permission != expect.Permission ||
				resp[i].PermissionName != expect.PermissionName {
				t.Error("Not correctly parsing returned dashboard permission")
			}
		})
	}
}

func TestUpdateDashboardPermissions(t *testing.T) {
	client := gapiTestTools(t, 200, updateDashboardPermissionsJSON)

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
				TeamID:     1,
				Permission: 1,
			},
			{
				UserID:     11,
				Permission: 4,
			},
		},
	}
	err := client.UpdateDashboardPermissions(1, items)
	if err != nil {
		t.Error(err)
	}
}

func TestDashboardPermissionsByUID(t *testing.T) {
	client := gapiTestTools(t, 200, getDashboardPermissionsJSON)

	resp, err := client.DashboardPermissionsByUID("nErXDvCkzz")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*DashboardPermission{
		{
			DashboardID:    1,
			DashboardUID:   "nErXDvCkzz",
			Role:           "Viewer",
			UserID:         0,
			TeamID:         0,
			IsFolder:       false,
			Inherited:      false,
			Permission:     1,
			PermissionName: "View",
		},
		{
			DashboardID:    2,
			DashboardUID:   "nErXDvCkzz",
			Role:           "Editor",
			UserID:         0,
			TeamID:         0,
			IsFolder:       false,
			Inherited:      true,
			Permission:     2,
			PermissionName: "Edit",
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if resp[i].DashboardID != expect.DashboardID ||
				resp[i].DashboardUID != expect.DashboardUID ||
				resp[i].Role != expect.Role ||
				resp[i].UserID != expect.UserID ||
				resp[i].TeamID != expect.TeamID ||
				resp[i].IsFolder != expect.IsFolder ||
				resp[i].Inherited != expect.Inherited ||
				resp[i].Permission != expect.Permission ||
				resp[i].PermissionName != expect.PermissionName {
				t.Error("Not correctly parsing returned dashboard permission")
			}
		})
	}
}

func TestUpdateDashboardPermissionsByUID(t *testing.T) {
	client := gapiTestTools(t, 200, updateDashboardPermissionsJSON)

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
				TeamID:     1,
				Permission: 1,
			},
			{
				UserID:     11,
				Permission: 4,
			},
		},
	}
	err := client.UpdateDashboardPermissionsByUID("nErXDvCkzz", items)
	if err != nil {
		t.Error(err)
	}
}
