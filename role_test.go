package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/access_control"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	newRoleResponse = `
{
    "global": false,
    "uid": "vc3SCSsGz",
    "name": "test:policy",
	"version": 1,
    "description": "Test policy description",
	"displayName": "test display",
	"group": "test group",
    "hidden": false,
    "permissions": [
        {
            "id": 6,
            "permission": "test:self",
            "scope": "test:self",
            "updated": "2021-02-22T16:16:05.646913+01:00",
            "created": "2021-02-22T16:16:05.646912+01:00"
        }
    ],
    "updated": "2021-02-22T16:16:05.644216+01:00",
    "created": "2021-02-22T16:16:05.644216+01:00"
}
`
	getRoleResponse = `
{
    "global": false,
    "uid": "vc3SCSsGz",
    "name": "test:policy",
	"version": 1,
    "description": "Test policy description",
	"displayName": "test display",
	"group": "test group",
    "hidden": false,
    "permissions": [
        {
            "permission": "test:self",
            "scope": "test:self",
            "updated": "2021-02-22T16:16:05.646913+01:00",
            "created": "2021-02-22T16:16:05.646912+01:00"
        }
    ],
    "updated": "2021-02-22T16:16:05.644216+01:00",
    "created": "2021-02-22T16:16:05.644216+01:00"
}
`

	updatedRoleResponse = `{"message":"Role updated"}`
	deleteRoleResponse  = `{"message":"Role deleted"}`
)

func TestNewRole(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 201, newRoleResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	roleReq := models.CreateRoleForm{
		Global:      false,
		Name:        "test:policy",
		Description: "test:policy",
		Permissions: []*models.Permission{
			{
				Action: "test:self",
				Scope:  "test:self",
			},
		},
	}

	resp, err := client.AccessControl.CreateRole(
		access_control.NewCreateRoleParams().
			WithBody(&roleReq),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.UID != "vc3SCSsGz" {
		t.Error("Not correctly parsing returned role uid.")
	}
}

func TestGetRole(t *testing.T) {
	server, client := gapiTestTools(t, 200, getRoleResponse)
	t.Cleanup(func() {
		server.Close()
	})

	uid := "vc3SCSsGz"

	resp, err := client.AccessControl.GetRole(
		access_control.NewGetRoleParams().WithRoleUID(uid),
		nil,
	)

	if err != nil {
		t.Error(err)
	}

	expected := models.RoleDTO{
		//Global:      false,
		Version:     1,
		UID:         "vc3SCSsGz",
		Name:        "test:policy",
		Description: "Test policy description",
		Group:       "test group",
		DisplayName: "test display",
		Hidden:      false,
		Permissions: []*models.Permission{
			{
				Action: "test:self",
				Scope:  "test:self",
			},
		},
	}

	t.Run("check response data", func(t *testing.T) {
		if expected.UID != resp.Payload.UID || expected.Name != resp.Payload.Name {
			t.Error("Not correctly parsing returned role.")
		}
	})
}

func TestUpdateRole(t *testing.T) {
	server, client := gapiTestTools(t, 200, updatedRoleResponse)
	t.Cleanup(func() {
		server.Close()
	})

	roleReq := models.UpdateRoleCommand{
		Global:      false,
		Name:        "test:policy",
		Description: "test:policy",
		Permissions: []*models.Permission{
			{
				Action: "test:self1",
				Scope:  "test:self1",
			},
		},
	}

	_, err := client.AccessControl.UpdateRoleWithPermissions(
		access_control.NewUpdateRoleWithPermissionsParams().WithBody(&roleReq),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteRole(t *testing.T) {
	server, client := gapiTestTools(t, 200, deleteRoleResponse)
	t.Cleanup(func() {
		server.Close()
	})

	global := false
	_, err := client.AccessControl.RemoveUserRole(
		access_control.NewRemoveUserRoleParams().
			WithRoleUID("vc3SCSsGz").
			WithGlobal(&global),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
