package gapi

import (
	"testing"

	"github.com/gobs/pretty"
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
	client := gapiTestTools(t, 201, newRoleResponse)

	roleReq := Role{
		Global:      false,
		Name:        "test:policy",
		Description: "test:policy",
		Permissions: []Permission{
			{
				Action: "test:self",
				Scope:  "test:self",
			},
		},
	}

	resp, err := client.NewRole(roleReq)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.UID != "vc3SCSsGz" {
		t.Error("Not correctly parsing returned role uid.")
	}
}

func TestGetRole(t *testing.T) {
	client := gapiTestTools(t, 200, getRoleResponse)

	uid := "vc3SCSsGz"

	resp, err := client.GetRole(uid)

	if err != nil {
		t.Error(err)
	}

	expected := Role{
		Global:      false,
		Version:     1,
		UID:         "vc3SCSsGz",
		Name:        "test:policy",
		Description: "Test policy description",
		Group:       "test group",
		DisplayName: "test display",
		Hidden:      false,
		Permissions: []Permission{
			{
				Action: "test:self",
				Scope:  "test:self",
			},
		},
	}

	t.Run("check response data", func(t *testing.T) {
		if expected.UID != resp.UID || expected.Name != resp.Name {
			t.Error("Not correctly parsing returned role.")
		}
	})
}

func TestUpdateRole(t *testing.T) {
	client := gapiTestTools(t, 200, updatedRoleResponse)

	roleReq := Role{
		Global:      false,
		Name:        "test:policy",
		Description: "test:policy",
		Permissions: []Permission{
			{
				Action: "test:self1",
				Scope:  "test:self1",
			},
		},
	}

	err := client.UpdateRole(roleReq)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteRole(t *testing.T) {
	client := gapiTestTools(t, 200, deleteRoleResponse)

	err := client.DeleteRole("vc3SCSsGz", false)
	if err != nil {
		t.Error(err)
	}
}
