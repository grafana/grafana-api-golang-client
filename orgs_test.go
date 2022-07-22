package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/orgs"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	getOrgsJSON    = `[{"id":1,"name":"Main Org."},{"id":2,"name":"Test Org."}]`
	getOrgJSON     = `{"id":1,"name":"Main Org.","address":{"address1":"","address2":"","city":"","zipCode":"","state":"","country":""}}`
	createdOrgJSON = `{"message":"Organization created","orgId":1}`
	updatedOrgJSON = `{"message":"Organization updated"}`
	deletedOrgJSON = `{"message":"Organization deleted"}`
)

func TestOrgs(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getOrgsJSON)
	defer mocksrv.Close()

	orgs, err := client.Orgs.SearchOrgs(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(orgs))

	if len(orgs.Payload) != 2 {
		t.Error("Length of returned orgs should be 2")
	}
	if orgs.Payload[0].ID != 1 || orgs.Payload[0].Name != "Main Org." {
		t.Error("Not correctly parsing returned organizations.")
	}
}

func TestOrgByName(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getOrgJSON)
	defer mocksrv.Close()

	org := "Main Org."
	resp, err := client.Orgs.GetOrgByName(
		orgs.NewGetOrgByNameParams().WithOrgName(org),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.ID != 1 || resp.Payload.Name != org {
		t.Error("Not correctly parsing returned organization.")
	}
}

func TestOrg(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getOrgJSON)
	defer mocksrv.Close()

	org := int64(1)
	resp, err := client.Orgs.GetOrgByID(
		orgs.NewGetOrgByIDParams().WithOrgID(org),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.ID != org || resp.Payload.Name != "Main Org." {
		t.Error("Not correctly parsing returned organization.")
	}
}

func TestNewOrg(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createdOrgJSON)
	defer mocksrv.Close()

	resp, err := client.Orgs.CreateOrg(
		orgs.NewCreateOrgParams().
			WithBody(&models.CreateOrgCommand{
				Name: "test-org",
			}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if *resp.Payload.OrgID != 1 {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateOrg(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updatedOrgJSON)
	defer mocksrv.Close()

	_, err := client.Orgs.UpdateOrg(
		orgs.NewUpdateOrgParams().
			WithOrgID(1).
			WithBody(&models.UpdateOrgForm{
				Name: "test-org",
			}),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOrg(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deletedOrgJSON)
	defer mocksrv.Close()

	_, err := client.Orgs.DeleteOrgByID(
		orgs.NewDeleteOrgByIDParams().WithOrgID(1),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
