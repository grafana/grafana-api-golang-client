//+build !integration

package gapi

import (
	"testing"
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
	server, client := gapiTestTools(t, 200, newBuiltInRoleAssignmentResponse)
	t.Cleanup(func() {
		server.Close()
	})

	br := BuiltInRoleAssignment{
		Global:      false,
		RoleUID:     "test:policy",
		BuiltinRole: "Viewer",
	}

	_, err := client.NewBuiltInRoleAssignment(br)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBuiltInRoleAssignments(t *testing.T) {
	server, client := gapiTestTools(t, 200, getBuiltInRoleAssignmentsResponse)
	t.Cleanup(func() {
		server.Close()
	})

	resp, err := client.GetBuiltInRoleAssignments()

	if err != nil {
		t.Error(err)
	}

	expected := map[string][]*Role{
		"Grafana Admin": {
			{
				Version:     1,
				Global:      true,
				Name:        "grafana:roles:users:admin:read",
				UID:         "tJTyTNqMk",
				Description: "",
			},
		},
		"Viewer": {
			{
				Version:     2,
				Global:      false,
				Name:        "custom:reports:editor",
				UID:         "tJTyTNqMk1",
				Description: "Role to allow users to create/read reports",
			},
		},
	}

	if len(expected["Viewer"]) != len(resp["Viewer"]) || len(expected["Grafana Admin"]) != len(resp["Grafana Admin"]) {
		t.Error("Unexpected built-in role assignments.")
	}
}

func TestDeleteBuiltInRoleAssignment(t *testing.T) {
	server, client := gapiTestTools(t, 200, removeBuiltInRoleAssignmentResponse)
	t.Cleanup(func() {
		server.Close()
	})

	br := BuiltInRoleAssignment{
		Global:      false,
		RoleUID:     "test:policy",
		BuiltinRole: "Viewer",
	}
	err := client.DeleteBuiltInRoleAssignment(br)
	if err != nil {
		t.Error(err)
	}
}
