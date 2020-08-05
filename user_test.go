package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getUsersJSON       = `[{"id":1,"email":"users@localhost","isAdmin":true}]`
	getUserJSON        = `{"id":2,"email":"user@localhost","isGrafanaAdmin":false}`
	getUserByEmailJSON = `{"id":3,"email":"userByEmail@localhost","isGrafanaAdmin":true}`
	getUserUpdateJSON  = `{"id":4,"email":"userUpdate@localhost","isGrafanaAdmin":false}`
)

func TestUsers(t *testing.T) {
	server, client := gapiTestTools(200, getUsersJSON)
	defer server.Close()

	resp, err := client.Users()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if len(resp) != 1 {
		t.Fatal("No users were returned.")
	}

	user := resp[0]

	if user.Email != "users@localhost" ||
		user.Id != 1 ||
		user.IsAdmin != true {
		t.Error("Not correctly parsing returned users.")
	}
}

func TestUser(t *testing.T) {
	server, client := gapiTestTools(200, getUserJSON)
	defer server.Close()

	user, err := client.User(1)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(user))

	if user.Email != "user@localhost" ||
		user.Id != 2 ||
		user.IsAdmin != false {
		t.Error("Not correctly parsing returned user.")
	}
}

func TestUserByEmail(t *testing.T) {
	server, client := gapiTestTools(200, getUserByEmailJSON)
	defer server.Close()

	user, err := client.UserByEmail("admin@localhost")
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(user))

	if user.Email != "userByEmail@localhost" ||
		user.Id != 3 ||
		user.IsAdmin != true {
		t.Error("Not correctly parsing returned user.")
	}
}

func TestUserUpdate(t *testing.T) {
	server, client := gapiTestTools(200, getUserUpdateJSON)
	defer server.Close()

	user, err := client.User(4)
	if err != nil {
		t.Error(err)
	}
	user.IsAdmin = true
	err = client.UserUpdate(user)
	if err != nil {
		t.Error(err)
	}
}
