package gapi

import (
	"testing"
	"github.com/gobs/pretty"
)

func TestUsers(t *testing.T) {
	server, client := gapiTestTools(200, getUsersJSON)
	defer server.Close()

	resp, err := client.Users()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	user := User{
		Id: 1,
		Email: "admin@localhost",
        Name: "",
		Login: "admin",
		IsAdmin: true,
	}

	if len(resp) != 1 || resp[0] != user {
		t.Error("Not correctly parsing returned users.")
	}
}

func TestUserByEmail(t *testing.T) {
    server, client := gapiTestTools(200, getUserByEmailJSON)
    defer server.Close()

    resp, err := client.UserByEmail("admin@localhost")
    if err != nil {
        t.Error(err)
    }

    t.Log(pretty.PrettyFormat(resp))

    user := User{
		Id: 1,
		Email: "admin@localhost",
        Name: "",
		Login: "admin",
		IsAdmin: true,
	}
    if resp != user {
        t.Error("Not correctly parsing returned user.")
    }
}
