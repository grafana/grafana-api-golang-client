package gapi

import (
	"testing"

	"github.com/gobs/pretty"
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
	server, client := gapiTestTools(t, 200, getTeamGroupsJSON)
	defer server.Close()

	teamID := int64(1)
	teamGroups, err := client.TeamGroups(teamID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(teamGroups))

	if len(teamGroups) != 1 {
		t.Error("Length of returned teamGroups should be 1")
	}
	if teamGroups[0].TeamID != 1 || teamGroups[0].OrgID != 1 || teamGroups[0].GroupID != "test" {
		t.Error("Not correctly parsing returned teamGroups.")
	}
}

func TestNewTeamGroup(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdTeamGroupJSON)
	defer server.Close()

	err := client.NewTeamGroup(int64(1), "test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTeamGroup(t *testing.T) {
	server, client := gapiTestTools(t, 200, deletedTeamGroupJSON)
	defer server.Close()

	err := client.DeleteTeamGroup(int64(1), "test")
	if err != nil {
		t.Fatal(err)
	}
}
