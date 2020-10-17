package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getOrgUsersJSON   = `[{"orgID":1,"userID":1,"email":"admin@localhost","avatarUrl":"/avatar/46d229b033af06a191ff2267bca9ae56","login":"admin","role":"Admin","lastSeenAt":"2018-06-28T14:16:11Z","lastSeenAtAge":"\u003c 1m"}]`
	addOrgUserJSON    = `{"message":"User added to organization"}`
	updateOrgUserJSON = `{"message":"Organization user updated"}`
	removeOrgUserJSON = `{"message":"User removed from organization"}`
)

func TestOrgUsers(t *testing.T) {
	server, client := gapiTestTools(t, 200, getOrgUsersJSON)
	defer server.Close()

	org := int64(1)
	resp, err := client.OrgUsers(org)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	user := OrgUser{
		OrgID:  1,
		UserID: 1,
		Email:  "admin@localhost",
		Login:  "admin",
		Role:   "Admin",
	}

	if resp[0] != user {
		t.Error("Not correctly parsing returned organization users.")
	}
}

func TestAddOrgUser(t *testing.T) {
	server, client := gapiTestTools(t, 200, addOrgUserJSON)
	defer server.Close()

	orgID, user, role := int64(1), "admin@localhost", "Admin"

	err := client.AddOrgUser(orgID, user, role)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateOrgUser(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateOrgUserJSON)
	defer server.Close()

	orgID, userID, role := int64(1), int64(1), "Editor"

	err := client.UpdateOrgUser(orgID, userID, role)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveOrgUser(t *testing.T) {
	server, client := gapiTestTools(t, 200, removeOrgUserJSON)
	defer server.Close()

	orgID, userID := int64(1), int64(1)

	err := client.RemoveOrgUser(orgID, userID)
	if err != nil {
		t.Error(err)
	}
}
