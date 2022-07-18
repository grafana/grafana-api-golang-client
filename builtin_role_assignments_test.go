package gapi

import (
	"testing"

	"github.com/grafana/grafana-api-golang-client/goclient/client/access_control"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	newBuiltInRoleAssignmentResponse = `
{
    "message": "Built-in role grant added"
}
`
	getBuiltInRoleAssignmentsResponse = `
{
    "Grafana Admin": [
        {
            "version": 1,
            "uid": "tJTyTNqMk",
            "name": "grafana:roles:users:admin:read",
            "description": "",
            "global": true
        }
    ],
    "Viewer": [
        {
            "version": 2,
            "uid": "tJTyTNqMk1",
            "name": "custom:reports:editor",
            "description": "Role to allow users to create/read reports",
            "global": false
        }
    ]
}
`

	removeBuiltInRoleAssignmentResponse = `
{
    "message": "Built-in role grant removed"
}
`
)

func TestNewBuiltInRoleAssignment(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, newBuiltInRoleAssignmentResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	br := models.AddBuiltInRoleCommand{
		Global:      false,
		RoleUID:     "test:policy",
		BuiltInRole: "Viewer",
	}

	_, err := client.AccessControl.AddBuiltinRole(
		access_control.NewAddBuiltinRoleParams().WithBody(&br),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBuiltInRoleAssignments(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getBuiltInRoleAssignmentsResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	resp, err := client.AccessControl.ListBuiltinRoles(access_control.NewListBuiltinRolesParams(), nil)

	if err != nil {
		t.Error(err)
	}

	expected := map[string][]*models.RoleDTO{
		"Grafana Admin": {
			{
				Version:     1,
				Name:        "grafana:roles:users:admin:read",
				UID:         "tJTyTNqMk",
				Description: "",
			},
		},
		"Viewer": {
			{
				Version:     2,
				Name:        "custom:reports:editor",
				UID:         "tJTyTNqMk1",
				Description: "Role to allow users to create/read reports",
			},
		},
	}

	if len(expected["Viewer"]) != len(resp.Payload["Viewer"]) || len(expected["Grafana Admin"]) != len(resp.Payload["Grafana Admin"]) {
		t.Error("Unexpected built-in role assignments.")
	}
}

func TestDeleteBuiltInRoleAssignment(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, removeBuiltInRoleAssignmentResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	global := false

	_, err := client.AccessControl.RemoveBuiltinRole(
		access_control.NewRemoveBuiltinRoleParams().
			WithRoleUID("test:policy").
			WithBuiltinRole("Viewer").
			WithGlobal(&global),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
