package gapi

import (
	"strings"
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/admin"
	"github.com/grafana/grafana-api-golang-client/goclient/client/admin_users"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

const (
	createUserJSON            = `{"id":1,"message":"User created"}`
	deleteUserJSON            = `{"message":"User deleted"}`
	updateUserPasswordJSON    = `{"message":"User password updated"}`
	updateUserPermissionsJSON = `{"message":"User permissions updated"}`

	pauseAllAlertsJSON = `{
		"alertsAffected": 1,
		"state": "Paused",
		"message": "alert paused"
	}`
)

func TestCreateUser(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createUserJSON)
	defer mocksrv.Close()

	resp, err := client.AdminUsers.CreateUser(admin_users.NewCreateUserParams().WithBody(
		&models.AdminCreateUserForm{
			Email:    "admin@localhost",
			Login:    "admin",
			Name:     "Administrator",
			Password: "password",
		},
	), nil)
	if err != nil {
		t.Error(err)
	}

	if resp.Payload.ID != 1 {
		t.Error("Not correctly parsing returned user message.")
	}
}

func TestDeleteUser(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deleteUserJSON)
	defer mocksrv.Close()

	_, err := client.AdminUsers.DeleteUser(admin_users.NewDeleteUserParams().WithUserID(1), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updateUserPasswordJSON)
	defer mocksrv.Close()

	_, err := client.AdminUsers.SetPassword(admin_users.NewSetPasswordParams().
		WithUserID(1).
		WithBody(&models.AdminUpdateUserPasswordForm{
			Password: "new-password",
		},
		), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUserPermissions(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updateUserPermissionsJSON)
	defer mocksrv.Close()

	_, err := client.AdminUsers.SetPermissions(admin_users.NewSetPermissionsParams().
		WithUserID(1).
		WithBody(&models.AdminUpdateUserPermissionsForm{
			IsGrafanaAdmin: false,
		},
		), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestPauseAllAlerts(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, pauseAllAlertsJSON)
	defer mocksrv.Close()

	res, err := client.Admin.PauseAllAlerts(admin.NewPauseAllAlertsParams().WithBody(
		&models.PauseAllAlertsCommand{
			Paused: true,
		},
	), nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.Payload.State != "Paused" {
		t.Error("pause all alerts response should contain the correct response message")
	}
}

func TestPauseAllAlerts_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, pauseAllAlertsJSON)
	defer server.Close()

	_, err := client.Admin.PauseAllAlerts(admin.NewPauseAllAlertsParams(), nil)
	if !strings.Contains(err.Error(), "status: 500") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}
}
