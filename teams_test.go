package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/google/go-cmp/cmp"
	"github.com/grafana/grafana-api-golang-client/goclient/client/teams"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	searchTeamJSON = `
{
  "totalCount": 1,
  "teams": [
    {
      "id": 1,
      "orgId": 1,
      "name": "MyTestTeam",
      "email": "",
      "avatarUrl": "/avatar/3f49c15916554246daa714b9bd0ee398",
      "memberCount": 1
    }
  ],
  "page": 1,
  "perPage": 1000
}
`
	getTeamJSON = `
{
  "id": 1,
  "orgId": 1,
  "name": "MyTestTeam",
  "email": "",
  "avatarUrl": "avatar/abcdef",
  "memberCount": 1,
  "permission": 0
}
`
	addTeamsJSON = `
{
  "message":"Team created",
  "teamId":2
}
`
	updateTeamJSON     = `{"message":"Team updated"}`
	deleteTeamJSON     = `{"message":"Team deleted"}`
	getTeamMembersJSON = `
[
  {
    "orgId": 1,
    "teamId": 1,
    "userID": 3,
    "auth_module": "oauth_github",
    "email": "user1@email.com",
    "login": "user1",
    "avatarUrl": "/avatar/1b3c32f6386b0185c40d359cdc733a79",
    "labels": [],
    "permission": 0
  },
  {
    "orgId": 1,
    "teamId": 1,
    "userID": 2,
    "auth_module": "oauth_github",
    "email": "user2@email.com",
    "login": "user2",
    "avatarUrl": "/avatar/cad3c68da76e45d10269e8ef02f8e73e",
    "labels": [],
    "permission": 0
  }
]
`
	addTeamMemberJSON = `
{
  "userID": 2
}
`
	removeMemberFromTeamJSON = `{"message":"Team Member removed"}`
	getTeamPreferencesJSON   = `
{
  "theme": "",
  "homeDashboardID": 0,
  "timezone": ""
}
`
	updateTeamPreferencesJSON = `
{
  "message":"Preferences updated"
}
`
)

func TestSearchTeam(t *testing.T) {
	mocksv, client := gapiTestTools(t, 200, searchTeamJSON)
	defer mocksv.Close()

	query := "myteam"
	resp, err := client.Teams.SearchTeams(
		teams.NewSearchTeamsParams().WithQuery(&query),
		nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expect := &models.SearchTeamQueryResult{
		TotalCount: 1,
		Teams: []*models.TeamDTO{
			{
				ID:          1,
				OrgID:       1,
				Name:        "MyTestTeam",
				Email:       "",
				AvatarURL:   "avatar/3f49c15916554246daa714b9bd0ee398",
				MemberCount: 1,
				Permission:  0,
			},
		},
		Page:    1,
		PerPage: 1000,
	}
	t.Run("check data", func(t *testing.T) {
		if cmp.Diff(resp.Payload, expect) != "" {
			t.Error("Not correctly parsing returned team search.")
		}
	})
}

func TestTeam(t *testing.T) {
	server, client := gapiTestTools(t, 200, getTeamJSON)
	defer server.Close()

	resp, err := client.Teams.GetTeamByID(
		teams.NewGetTeamByIDParams().WithTeamID("1"),
		nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expect := &models.TeamDTO{
		ID:          1,
		OrgID:       1,
		Name:        "MyTestTeam",
		Email:       "",
		AvatarURL:   "avatar/abcdef",
		MemberCount: 1,
		Permission:  0,
	}
	t.Run("check data", func(t *testing.T) {
		if resp.Payload.ID != expect.ID || resp.Payload.Name != expect.Name {
			t.Error("Not correctly parsing returned team.")
		}
	})
}

func TestAddTeam(t *testing.T) {
	server, client := gapiTestTools(t, 200, addTeamsJSON)
	defer server.Close()

	resp, err := client.Teams.CreateTeam(
		teams.NewCreateTeamParams().
			WithBody(&models.CreateTeamCommand{
				Name: "TestTeam",
			}),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	if resp.Payload.TeamID == 0 {
		t.Error("AddTeam returned an invalid ID")
	}
}

func TestUpdateTeam(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateTeamJSON)
	defer server.Close()

	_, err := client.Teams.UpdateTeam(
		teams.NewUpdateTeamParams().
			WithTeamID("1").
			WithBody(&models.UpdateTeamCommand{
				Name: "TestTeam",
			}),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTeam(t *testing.T) {
	server, client := gapiTestTools(t, 200, deleteTeamJSON)
	defer server.Close()

	_, err := client.Teams.DeleteTeamByID(
		teams.NewDeleteTeamByIDParams().WithTeamID("1"),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestTeamMembers(t *testing.T) {
	server, client := gapiTestTools(t, 200, getTeamMembersJSON)
	defer server.Close()

	resp, err := client.Teams.GetTeamMembers(
		teams.NewGetTeamMembersParams().WithTeamID("1"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	expects := []*models.TeamMemberDTO{
		{
			OrgID:      1,
			TeamID:     1,
			UserID:     3,
			Email:      "user1@email.com",
			Login:      "user1",
			AvatarURL:  "/avatar/1b3c32f6386b0185c40d359cdc733a79",
			Permission: 0,
		},
		{
			OrgID:      1,
			TeamID:     1,
			UserID:     2,
			Email:      "user2@email.com",
			Login:      "user2",
			AvatarURL:  "/avatar/cad3c68da76e45d10269e8ef02f8e73e",
			Permission: 0,
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if expect.Email != resp.Payload[i].Email || expect.AvatarURL != resp.Payload[i].AvatarURL {
				t.Error("Not correctly parsing returned team members.")
			}
		})
	}
}

func TestAddTeamMember(t *testing.T) {
	server, client := gapiTestTools(t, 200, addTeamMemberJSON)
	defer server.Close()

	if _, err := client.Teams.AddTeamMember(
		teams.NewAddTeamMemberParams().
			WithTeamID("1").
			WithBody(&models.AddTeamMemberCommand{
				UserID: int64(2),
			}),
		nil,
	); err != nil {
		t.Error(err)
	}
}

func TestRemoveMemberFromTeam(t *testing.T) {
	server, client := gapiTestTools(t, 200, removeMemberFromTeamJSON)
	defer server.Close()

	if _, err := client.Teams.RemoveTeamMember(
		teams.NewRemoveTeamMemberParams().
			WithTeamID("1").
			WithUserID(int64(2)),
		nil,
	); err != nil {
		t.Error(err)
	}
}

func TestTeamPreferences(t *testing.T) {
	server, client := gapiTestTools(t, 200, getTeamPreferencesJSON)
	defer server.Close()

	resp, err := client.Teams.GetTeamPreferences(
		teams.NewGetTeamPreferencesParams().
			WithTeamID("1"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	expect := &models.Prefs{
		Theme:           "",
		HomeDashboardID: 0,
		Timezone:        "",
	}

	t.Run("check data", func(t *testing.T) {
		if expect.Theme != resp.Payload.Theme || expect.HomeDashboardID != resp.Payload.HomeDashboardID {
			t.Error("Not correctly parsing returned team preferences.")
		}
	})
}

func TestUpdateTeamPreferences(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateTeamPreferencesJSON)
	defer server.Close()

	preferences := models.UpdatePrefsCmd{
		Theme:           "",
		HomeDashboardID: int64(0),
		Timezone:        "",
	}

	if _, err := client.Teams.UpdateTeamPreferences(
		teams.NewUpdateTeamPreferencesParams().WithTeamID("1").WithBody(&preferences),
		nil,
	); err != nil {
		t.Error(err)
	}
}
