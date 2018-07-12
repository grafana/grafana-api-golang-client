package gapi

import (
	"testing"
)

const (
	createUserJSON = `{"id":1,"message":"User created"}`
	deleteUserJSON = `{"message":"User deleted"}`
)

func TestCreateUser(t *testing.T) {
	server, client := gapiTestTools(200, createUserJSON)
	defer server.Close()
	user := User{
		Email:    "admin@localhost",
		Login:    "admin",
		Name:     "Administrator",
		Password: "password",
	}
	resp, err := client.CreateUser(user)
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
