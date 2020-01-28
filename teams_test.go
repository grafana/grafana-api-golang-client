package gapi

import (
	"github.com/gobs/pretty"
	"testing"
)

const (
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
    "email": "user1@email.com",
    "login": "user1",
    "avatarUrl": "\/avatar\/1b3c32f6386b0185c40d359cdc733a79"
  },
  {
    "orgId": 1,
    "teamId": 1,
    "userId": 2,
    "email": "user2@email.com",
    "login": "user2",
    "avatarUrl": "\/avatar\/cad3c68da76e45d10269e8ef02f8e73e"
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
	server, client := gapiTestTools(200, addTeamsJSON)
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
	server, client := gapiTestTools(200, addTeamsJSON)
	defer server.Close()

	id := int64(1)

	err := client.DeleteTeam(id)
	if err != nil {
		t.Error(err)
	}
}
