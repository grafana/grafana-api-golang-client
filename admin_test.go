package gapi

import (
	"strings"
	"testing"

	"github.com/gobs/pretty"
)

const (
	createUserJSON = `{"id":1,"message":"User created"}`
	deleteUserJSON = `{"message":"User deleted"}`

	pauseAllAlertsJSON = `{
		"alertsAffected": 1,
		"state": "Paused",
		"message": "alert paused"
	}`
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

func TestPauseAllAlerts(t *testing.T) {
	server, client := gapiTestTools(200, pauseAllAlertsJSON)
	defer server.Close()

	res, err := client.PauseAllAlerts()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.State != "Paused" {
		t.Error("pause all alerts response should contain the correct response message")
	}
}

func TestPauseAllAlerts_500(t *testing.T) {
	server, client := gapiTestTools(500, pauseAllAlertsJSON)
	defer server.Close()

	_, err := client.PauseAllAlerts()
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}
