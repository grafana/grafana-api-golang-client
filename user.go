package gapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type User struct {
	Id      int64 	`json:"id"`
	Email   string	`json:"email"`
	Name    string	`json:"name"`
	Login   string	`json:"login"`
	IsAdmin bool	`json:"isAdmin"`
}

func (c *Client) Users() ([]User, error) {
	users := make([]User, 0)
	req, err := c.newRequest("GET", "/api/users", nil)
	if err != nil {
		return users, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return users, err
	}
	if resp.StatusCode != 200 {
		return users, errors.New(resp.Status)
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

func (c *Client) UserByEmail(email string) (User, error) {
	user := User{}
	req, err := c.newQueryRequest("GET", "/api/users/lookup", fmt.Sprintf("loginOrEmail=%s", email))
	if err != nil {
		return user, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return user, err
	}
	if resp.StatusCode != 200 {
		return user, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	tmp := struct {
		Id		int64	`json:"id"`
		Email	string	`json:"email"`
		Name	string	`json:"name"`
		Login	string	`json:"login"`
		IsAdmin	bool 	`json:"isGrafanaAdmin"`
	}{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return user, err
	}
	user = User(tmp)
	return user, err
}
