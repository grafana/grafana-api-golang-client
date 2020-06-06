package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type SearchTeam struct {
	TotalCount int64   `json:"totalCount,omitempty"`
	Teams      []*Team `json:"teams,omitempty"`
	Page       int64   `json:"page,omitempty"`
	PerPage    int64   `json:"perPage,omitempty"`
}

// Team consists of a get response
// It's used in  Add and Update API
type Team struct {
	Id          int64  `json:"id,omitempty"`
	OrgId       int64  `json:"orgId,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email,omitempty"`
	AvatarUrl   string `json:"avatarUrl,omitempty"`
	MemberCount int64  `json:"memberCount,omitempty"`
	Permission  int64  `json:"permission,omitempty"`
}

// TeamMember
type TeamMember struct {
	OrgId      int64  `json:"orgId,omitempty"`
	TeamId     int64  `json:"teamId,omitempty"`
	UserId     int64  `json:"userId,omitempty"`
	Email      string `json:"email,omitempty"`
	Login      string `json:"login,omitempty"`
	AvatarUrl  string `json:"avatarUrl,omitempty"`
	Permission int64  `json:"permission,omitempty"`
}

type Preferences struct {
	Theme           string `json:"theme"`
	HomeDashboardId int64  `json:"homeDashboardId"`
	Timezone        string `json:"timezone"`
}

// SearchTeam searches Grafana teams and returns the results.
func (c *Client) SearchTeam(query string) (*SearchTeam, error) {
	var result SearchTeam

	page := "1"
	perPage := "1000"
	path := "/api/teams/search"
	queryValues := url.Values{}
	queryValues.Set("page", page)
	queryValues.Set("perPage", perPage)
	queryValues.Set("query", query)

	resp, err := c.request("GET", path, queryValues, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Team fetches and returns the Grafana team whose ID it's passed.
func (c *Client) Team(id int64) (*Team, error) {
	var team Team

	resp, err := c.request("GET", fmt.Sprintf("/api/teams/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

// AddTeam makes a new team
// email arg is an optional value.
// If you don't want to set email, please set "" (empty string).
func (c *Client) AddTeam(name string, email string) error {
	path := fmt.Sprintf("/api/teams")
	team := Team{
		Name:  name,
		Email: email,
	}
	data, err := json.Marshal(team)
	if err != nil {
		return err
	}

	_, err = c.request("POST", path, nil, bytes.NewBuffer(data))

	return err
}

// UpdateTeam updates a Grafana team.
func (c *Client) UpdateTeam(id int64, name string, email string) error {
	path := fmt.Sprintf("/api/teams/%d", id)
	team := Team{
		Name: name,
	}
	// add param if email exists
	if email != "" {
		team.Email = email
	}
	data, err := json.Marshal(team)
	if err != nil {
		return err
	}

	_, err = c.request("PUT", path, nil, bytes.NewBuffer(data))

	return err
}

// DeleteTeam deletes the Grafana team whose ID it's passed.
func (c *Client) DeleteTeam(id int64) error {
	_, err := c.request("DELETE", fmt.Sprintf("/api/teams/%d", id), nil, nil)

	return err
}

// TeamMembers fetches and returns the team members for the Grafana team whose ID it's passed.
func (c *Client) TeamMembers(id int64) ([]*TeamMember, error) {
	members := make([]*TeamMember, 0)

	resp, err := c.request("GET", fmt.Sprintf("/api/teams/%d/members", id), nil, nil)
	if err != nil {
		return members, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return members, err
	}
	if err := json.Unmarshal(data, &members); err != nil {
		return members, err
	}
	return members, nil
}

// AddTeamMember adds a user to the Grafana team whose ID it's passed.
func (c *Client) AddTeamMember(id int64, userId int64) error {
	path := fmt.Sprintf("/api/teams/%d/members", id)
	member := TeamMember{UserId: userId}
	data, err := json.Marshal(member)
	if err != nil {
		return err
	}
	_, err = c.request("POST", path, nil, bytes.NewBuffer(data))

	return err
}

// RemoveMemberFromTeam removes a user from the Grafana team whose ID it's passed.
func (c *Client) RemoveMemberFromTeam(id int64, userId int64) error {
	path := fmt.Sprintf("/api/teams/%d/members/%d", id, userId)
	_, err := c.request("DELETE", path, nil, nil)

	return err
}

// TeamPreferences fetches and returns preferences for the Grafana team whose ID it's passed.
func (c *Client) TeamPreferences(id int64) (*Preferences, error) {
	var preferences Preferences

	resp, err := c.request("GET", fmt.Sprintf("/api/teams/%d/preferences", id), nil, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &preferences); err != nil {
		return nil, err
	}

	return &preferences, nil
}

// UpdateTeamPreferences updates team preferences for the Grafana team whose ID it's passed.
func (c *Client) UpdateTeamPreferences(id int64, theme string, homeDashboardId int64, timezone string) error {
	path := fmt.Sprintf("/api/teams/%d", id)
	preferences := Preferences{
		Theme:           theme,
		HomeDashboardId: homeDashboardId,
		Timezone:        timezone,
	}
	data, err := json.Marshal(preferences)
	if err != nil {
		return err
	}
	_, err = c.request("PUT", path, nil, bytes.NewBuffer(data))

	return err
}
