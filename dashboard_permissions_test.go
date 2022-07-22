package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/dashboard_permissions"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
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
	server, client := gapiTestTools(t, 200, getDashboardPermissionsJSON)
	defer server.Close()

	resp, err := client.DashboardPermissions.GetDashboardPermissionsListByUID(
		dashboard_permissions.NewGetDashboardPermissionsListByUIDParams().WithUID("nErXDvCkzz"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*models.DashboardACLInfoDTO{
		{
			DashboardID: 1,
			//DashboardUID:   "nErXDvCkzz",
			Role:           "Viewer",
			UserID:         0,
			TeamID:         0,
			IsFolder:       false,
			Inherited:      false,
			Permission:     1,
			PermissionName: "View",
		},
		{
			DashboardID: 2,
			//DashboardUID:   "nErXDvCkzz",
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
			if resp.Payload[i].DashboardID != expect.DashboardID ||
				// resp.Payload[i].DashboardUID != expect.DashboardUID ||
				resp.Payload[i].Role != expect.Role ||
				resp.Payload[i].UserID != expect.UserID ||
				resp.Payload[i].TeamID != expect.TeamID ||
				resp.Payload[i].IsFolder != expect.IsFolder ||
				resp.Payload[i].Inherited != expect.Inherited ||
				resp.Payload[i].Permission != expect.Permission ||
				resp.Payload[i].PermissionName != expect.PermissionName {
				t.Error("Not correctly parsing returned dashboard permission")
			}
		})
	}
}

func TestUpdateDashboardPermissions(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateDashboardPermissionsJSON)
	defer server.Close()

	items := models.UpdateDashboardACLCommand{
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
	_, err := client.DashboardPermissions.UpdateDashboardPermissionsByUID(
		dashboard_permissions.NewUpdateDashboardPermissionsByUIDParams().
			WithUID("uid").
			WithBody(&items),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
