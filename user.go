package gapi

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type User struct {
	Id       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"isAdmin,omitempty"`
}

// Users fetches and returns Grafana users.
func (c *Client) Users() ([]User, error) {
	users := make([]User, 0)
	resp, err := c.request("GET", "/api/users", nil, nil)
	if err != nil {
		return users, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return users, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return users, err
	}
	return users, err
}

// UserByEmail fetches and returns the user whose email matches that passed.
func (c *Client) UserByEmail(email string) (User, error) {
	user := User{}
	query := url.Values{}
	query.Add("loginOrEmail", email)
	resp, err := c.request("GET", "/api/users/lookup", query, nil)
	if err != nil {
		return user, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	tmp := struct {
		Id       int64  `json:"id,omitempty"`
		Email    string `json:"email,omitempty"`
		Name     string `json:"name,omitempty"`
		Login    string `json:"login,omitempty"`
		Password string `json:"password,omitempty"`
		IsAdmin  bool   `json:"isGrafanaAdmin,omitempty"`
	}{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return user, err
	}
	user = User(tmp)
	return user, err
}
