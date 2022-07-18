package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/sync_team_groups"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	getTeamGroupsJSON = `
[
  {
    "orgId": 1,
    "teamId": 1,
    "groupId": "test"
  }
]
`
	createdTeamGroupJSON = `
{
  "message":"Group added to Team"
}
`

	deletedTeamGroupJSON = `
{
  "message":"Team Group removed"
}
`
)

func TestTeamGroups(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getTeamGroupsJSON)
	defer mocksrv.Close()

	teamGroups, err := client.SyncTeamGroups.GetTeamGroupsAPI(
		sync_team_groups.NewGetTeamGroupsAPIParams().
			WithTeamID(1),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(teamGroups))

	if len(teamGroups.Payload) != 1 {
		t.Error("Length of returned teamGroups should be 1")
	}
	if teamGroups.Payload[0].TeamID != 1 || teamGroups.Payload[0].OrgID != 1 || teamGroups.Payload[0].GroupID != "test" {
		t.Error("Not correctly parsing returned teamGroups.")
	}
}

func TestNewTeamGroup(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createdTeamGroupJSON)
	defer mocksrv.Close()

	_, err := client.SyncTeamGroups.AddTeamGroupAPI(
		sync_team_groups.NewAddTeamGroupAPIParams().WithTeamID(1).WithBody(&models.TeamGroupMapping{
			GroupID: "test",
		}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTeamGroup(t *testing.T) {
	server, client := gapiTestTools(t, 200, deletedTeamGroupJSON)
	defer server.Close()

	_, err := client.SyncTeamGroups.RemoveTeamGroupAPI(
		sync_team_groups.NewRemoveTeamGroupAPIParams().
			WithTeamID(1).
			WithGroupID("test"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}
