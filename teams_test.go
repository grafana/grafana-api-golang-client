package gapi

const (
	getTeamJSON = `
{
  "id": 1,
  "orgId": 1,
  "name": "MyTestTeam",
  "email": "",
  "created": "2017-12-15T10:40:45+01:00",
  "updated": "2017-12-15T10:40:45+01:00"
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
