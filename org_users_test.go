package gapi

import (
	"testing"
	"github.com/gobs/pretty"
)

func TestOrgUsers(t *testing.T) {
	server, client := gapiTestTools(200, getOrgUsersJSON)
	defer server.Close()

	org := int64(1)
	resp, err := client.OrgUsers(org)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	user := OrgUser{
		OrgId: 1,
		UserId: 1,
		Email: "admin@localhost",
		Login: "admin",
		Role: "Admin",
	}

	if resp[0] != user {
		t.Error("Not correctly parsing returned organization users.")
	}
}

func TestAddOrgUser(t *testing.T) {
	server, client := gapiTestTools(200, addOrgUserJSON)
	defer server.Close()

	orgId, user, role := int64(1), "admin@localhost", "Admin"

	err := client.AddOrgUser(orgId, user, role)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateOrgUser(t *testing.T) {
	server, client := gapiTestTools(200, updateOrgUserJSON)
	defer server.Close()

	orgId, userId, role := int64(1), int64(1), "Editor"

	err := client.UpdateOrgUser(orgId, userId, role)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveOrgUser(t *testing.T) {
	server, client := gapiTestTools(200, removeOrgUserJSON)
	defer server.Close()

	orgId, userId := int64(1), int64(1)

	err := client.RemoveOrgUser(orgId, userId)
	if err != nil {
		t.Error(err)
	}
}
