package gapi

import (
	"github.com/gobs/pretty"
	"testing"
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
	updateTeamJSON     = `{ "message":"Team updated"}`
	deleteTeamJSON     = `{"message":"Team deleted"}`
	getTeamMembersJSON = `
[
  {
    "orgId": 1,
    "teamId": 1,
    "userId": 3,
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
    "userId": 2,
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
  "userId": 2
}
`
	removeMemberFromTeamJSON = `{"message":"Team Member removed"}`
	getTeamPreferencesJSON   = `
{
  "theme": "",
  "homeDashboardId": 0,
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
	server, client := gapiTestTools(200, searchTeamJSON)
	defer server.Close()

	query := "myteam"
	resp, err := client.SearchTeam(query)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expect := &SearchTeam{
		TotalCount: 1,
		Teams: []Team{
			{
				Id:          1,
				OrgId:       1,
				Name:        "MyTestTeam",
				Email:       "",
				AvatarUrl:   "avatar/3f49c15916554246daa714b9bd0ee398",
				MemberCount: 1,
				Permission:  0,
			},
		},
		Page:    1,
		PerPage: 1000,
	}
	t.Run("check data", func(t *testing.T) {
		if expect.TotalCount != resp.TotalCount || expect.Teams[0].Name != resp.Teams[0].Name {
			t.Error("Not correctly data")
		}
	})
}

func TestTeam(t *testing.T) {
	server, client := gapiTestTools(200, getTeamJSON)
	defer server.Close()

	id := int64(1)
	resp, err := client.Team(id)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expect := &Team{
		Id:          1,
		OrgId:       1,
		Name:        "MyTestTeam",
		Email:       "",
		AvatarUrl:   "avatar/abcdef",
		MemberCount: 1,
		Permission:  0,
	}
	t.Run("check data", func(t *testing.T) {
		if expect.Id != resp.Id || expect.Name != expect.Name {
			t.Error("Not correctly data")
		}
	})
}

func TestAddTeam(t *testing.T) {
	server, client := gapiTestTools(200, addTeamsJSON)
	defer server.Close()

	name := "TestTeam"
	email := ""

	err := client.AddTeam(name, email)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTeam(t *testing.T) {
	server, client := gapiTestTools(200, updateTeamJSON)
	defer server.Close()

	id := int64(1)
	name := "TestTeam"
	email := ""

	err := client.UpdateTeam(id, name, email)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTeam(t *testing.T) {
	server, client := gapiTestTools(200, deleteTeamJSON)
	defer server.Close()

	id := int64(1)

	err := client.DeleteTeam(id)
	if err != nil {
		t.Error(err)
	}
}

func TestTeamMembers(t *testing.T) {
	server, client := gapiTestTools(200, getTeamMembersJSON)
	defer server.Close()

	id := int64(1)

	resp, err := client.TeamMembers(id)
	if err != nil {
		t.Error(err)
	}
	expects := []TeamMember{
		{
			OrgId:      1,
			TeamId:     1,
			UserId:     3,
			Email:      "user1@email.com",
			Login:      "user1",
			AvatarUrl:  "/avatar/1b3c32f6386b0185c40d359cdc733a79",
			Permission: 0,
		},
		{
			OrgId:      1,
			TeamId:     1,
			UserId:     2,
			Email:      "user2@email.com",
			Login:      "user2",
			AvatarUrl:  "/avatar/cad3c68da76e45d10269e8ef02f8e73e",
			Permission: 0,
		},
	}

	for i, expect := range expects {
		t.Run("check data", func(t *testing.T) {
			if expect.Email != resp[i].Email || expect.AvatarUrl != resp[i].AvatarUrl {
				t.Error("Not correctly data")
			}
		})
	}
}

func TestAddTeamMember(t *testing.T) {
	server, client := gapiTestTools(200, addTeamMemberJSON)
	defer server.Close()

	id := int64(1)
	userId := int64(2)

	if err := client.AddTeamMember(id, userId); err != nil {
		t.Error(err)
	}
}

func TestRemoveMemberFromTeam(t *testing.T) {
	server, client := gapiTestTools(200, removeMemberFromTeamJSON)
	defer server.Close()

	id := int64(1)
	userId := int64(2)

	if err := client.RemoveMemberFromTeam(id, userId); err != nil {
		t.Error(err)
	}
}

func TestTeamPreferences(t *testing.T) {
	server, client := gapiTestTools(200, getTeamPreferencesJSON)
	defer server.Close()

	id := int64(1)

	resp, err := client.TeamPreferences(id)
	if err != nil {
		t.Error(err)
	}
	expect := &Preferences{
		Theme:           "",
		HomeDashboardId: 0,
		Timezone:        "",
	}

	t.Run("check data", func(t *testing.T) {
		if expect.Theme != resp.Theme || expect.HomeDashboardId != resp.HomeDashboardId {
			t.Error("Not correctly data")
		}
	})
}

func TestUpdateTeamPreferences(t *testing.T) {
	server, client := gapiTestTools(200, updateTeamPreferencesJSON)
	defer server.Close()

	id := int64(1)
	theme := ""
	homeDashboardId := int64(0)
	timezone := ""

	if err := client.UpdateTeamPreferences(id, theme, homeDashboardId, timezone); err != nil {
		t.Error(err)
	}
}
