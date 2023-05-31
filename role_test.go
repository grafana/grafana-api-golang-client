package gapi

import (
	"strings"
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

	roleUID = "vc3SCSsGz"
)

func TestRoles(t *testing.T) {
	mockData := strings.Repeat(getRoleResponse+",", 1000) // make 1000 roles.
	mockData = "[" + mockData[:len(mockData)-1] + "]"     // remove trailing comma; make a json list.

	// This creates 1000 + 1000 + 1 (2001, 3 calls) worth of roles.
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{200, mockData},
		{200, mockData},
		{200, "[" + getRoleResponse + "]"},
	})

	const dashCount = 2001

	roles, err := client.GetRoles()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(roles))

	if len(roles) != dashCount {
		t.Fatalf("Length of returned roles should be %d", dashCount)
	}
	if roles[0].UID != roleUID || roles[0].Name != "test:policy" {
		t.Error("Not correctly parsing returned roles.")
	}
	if roles[dashCount-1].UID != roleUID || roles[dashCount-1].Name != "test:policy" {
		t.Error("Not correctly parsing returned roles.")
	}
}

func TestRolesZeroResults(t *testing.T) {
	// This return zero roles.
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{200, "[]"},
	})

	roles, err := client.GetRoles()
	if err != nil {
		t.Fatal(err)
	}

	if len(roles) != 0 {
		t.Errorf("Length of returned roles should be zero")
	}
}

func TestRolesSinglePage(t *testing.T) {
	mockData := strings.Repeat(getRoleResponse+",", 999) // make 999 roles.
	mockData = "[" + mockData[:len(mockData)-1] + "]"    // remove trailing comma; make a json list.

	// This creates 999 worth of roles.
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{200, mockData},
	})

	const dashCount = 999

	roles, err := client.GetRoles()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(roles))

	if len(roles) != dashCount {
		t.Fatalf("Length of returned roles should be %d", dashCount)
	}
	if roles[0].UID != roleUID || roles[0].Name != "test:policy" {
		t.Error("Not correctly parsing returned roles.")
	}
	if roles[dashCount-1].UID != roleUID || roles[dashCount-1].Name != "test:policy" {
		t.Error("Not correctly parsing returned roles.")
	}
}

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

	if resp.UID != roleUID {
		t.Error("Not correctly parsing returned role uid.")
	}
}

func TestGetRole(t *testing.T) {
	client := gapiTestTools(t, 200, getRoleResponse)

	resp, err := client.GetRole(roleUID)
	if err != nil {
		t.Error(err)
	}

	expected := Role{
		Global:      false,
		Version:     1,
		UID:         roleUID,
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

	err := client.DeleteRole(roleUID, false)
	if err != nil {
		t.Error(err)
	}
}
