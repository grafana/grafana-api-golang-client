package gapi

import (
	"github.com/gobs/pretty"
	"testing"
)

const (
	getOrgsJSON    = `[{"id":1,"name":"Main Org."},{"id":2,"name":"Test Org."}]`
	getOrgJSON     = `{"id":1,"name":"Main Org.","address":{"address1":"","address2":"","city":"","zipCode":"","state":"","country":""}}`
	createdOrgJSON = `{"message":"Organization created","orgId":1}`
	updatedOrgJSON = `{"message":"Organization updated"}`
	deletedOrgJSON = `{"message":"Organization deleted"}`
)

func TestOrgs(t *testing.T) {
	server, client := gapiTestTools(200, getOrgsJSON)
	defer server.Close()

	orgs, err := client.Orgs()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(orgs))

	if len(orgs) != 2 {
		t.Error("Length of returned orgs should be 2")
	}
	if orgs[0].Id != 1 || orgs[0].Name != "Main Org." {
		t.Error("Not correctly parsing returned organizations.")
	}
}

func TestOrgByName(t *testing.T) {
	server, client := gapiTestTools(200, getOrgJSON)
	defer server.Close()

	org := "Main Org."
	resp, err := client.OrgByName(org)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Id != 1 || resp.Name != org {
		t.Error("Not correctly parsing returned organization.")
	}
}

func TestOrg(t *testing.T) {
	server, client := gapiTestTools(200, getOrgJSON)
	defer server.Close()

	org := int64(1)
	resp, err := client.Org(org)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Id != org || resp.Name != "Main Org." {
		t.Error("Not correctly parsing returned organization.")
	}
}

func TestNewOrg(t *testing.T) {
	server, client := gapiTestTools(200, createdOrgJSON)
	defer server.Close()

	resp, err := client.NewOrg("test-org")
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp != 1 {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateOrg(t *testing.T) {
	server, client := gapiTestTools(200, updatedOrgJSON)
	defer server.Close()

	err := client.UpdateOrg(int64(1), "test-org")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOrg(t *testing.T) {
	server, client := gapiTestTools(200, deletedOrgJSON)
	defer server.Close()

	err := client.DeleteOrg(int64(1))
	if err != nil {
		t.Error(err)
	}
}
