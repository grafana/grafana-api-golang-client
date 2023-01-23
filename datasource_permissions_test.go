//go:build !integration
// +build !integration

package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getDatasourcePermissionsJSON = `{
	"datasourceId": 1,
	"enabled": true,
	"permissions": [
		{
			"datasourceId": 1,
			"userId": 1,
			"userLogin": "user",
			"userEmail": "user@test.com",
			"userAvatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56",
			"permission": 1,
			"permissionName": "Query",
			"created": "2017-06-20T02:00:00+02:00",
			"updated": "2017-06-20T02:00:00+02:00"
		},
		{
			"datasourceId": 2,
			"teamId": 1,
			"team": "A Team",
			"teamAvatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56",
			"permission": 1,
			"permissionName": "Query",
			"created": "2017-06-20T02:00:00+02:00",
			"updated": "2017-06-20T02:00:00+02:00"
		},
		{
			"datasourceId": 1,
			"permission": 2,
			"permissionName": "Edit",
			"builtInRole": "Viewer",
			"created": "2017-06-20T02:00:00+02:00",
			"updated": "2017-06-20T02:00:00+02:00"
		}
	]
}`
	addDatasourcePermissionsJSON = `{
	"message": "Datasource permission added"
}`
)

func TestDatasourcePermissions(t *testing.T) {
	client := gapiTestTools(t, 200, getDatasourcePermissionsJSON)

	resp, err := client.DatasourcePermissions(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*DatasourcePermission{
		{
			DatasourceID:   1,
			UserID:         1,
			TeamID:         0,
			Permission:     1,
			PermissionName: "Query",
		},
		{
			DatasourceID:   2,
			UserID:         0,
			TeamID:         1,
			Permission:     1,
			PermissionName: "Query",
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if resp.Permissions[i].DatasourceID != expect.DatasourceID ||
				resp.Permissions[i].UserID != expect.UserID ||
				resp.Permissions[i].TeamID != expect.TeamID ||
				resp.Permissions[i].Permission != expect.Permission ||
				resp.Permissions[i].PermissionName != expect.PermissionName {
				t.Error("Not correctly parsing returned datasource permission")
			}
		})
	}
}

func TestAddDatasourcePermissions(t *testing.T) {
	for _, item := range []*DatasourcePermissionAddPayload{
		{
			TeamID:     1,
			Permission: 1,
		},
		{
			UserID:     11,
			Permission: 1,
		},
		{
			BuiltInRole: "Viewer",
			Permission:  2,
		},
	} {
		client := gapiTestTools(t, 200, addDatasourcePermissionsJSON)
		err := client.AddDatasourcePermission(1, item)
		if err != nil {
			t.Error(err)
		}
	}
}
