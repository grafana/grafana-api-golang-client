package gapi

import (
	"testing"
	// "github.com/gobs/pretty"
    "github.com/grafana/grafana/pkg/api/dtos"
)

const (
    createUserJSON = `{"id":1,"message":"User created"}`
	deleteUserJSON = `{"message":"User deleted"}`
)

func TestCreateUserForm(t *testing.T) {
	server, client := gapiTestTools(200, createUserJSON)
	defer server.Close()

    user := dtos.AdminCreateUserForm{
		Email: "admin@localhost",
		Login: "admin",
        Name: "Administrator",
		Password: "password",
	}
	resp, err := client.CreateUserForm(user)
	if err != nil {
		t.Error(err)
	}

	if resp != 1 {
		t.Error("Not correctly parsing returned user message.")
	}
}

func TestCreateUser(t *testing.T) {
	server, client := gapiTestTools(200, createUserJSON)
	defer server.Close()

    user := dtos.AdminCreateUserForm{
		Email: "admin@localhost",
		Login: "admin",
        Name: "Administrator",
		Password: "password",
	}
	resp, err := client.CreateUser(user.Email, user.Login, user.Email, user.Password)
	if err != nil {
		t.Error(err)
	}

	if resp != 1 {
		t.Error("Not correctly parsing returned user message.")
	}
}

func TestDeleteUser(t *testing.T) {
    server, client := gapiTestTools(200, deleteUserJSON)
    defer server.Close()

    err := client.DeleteUser(int64(1))
    if err != nil {
        t.Error(err)
    }
}
