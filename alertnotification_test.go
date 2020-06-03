package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getAlertNotificationsJSON = `
[
  {
    "id": 1,
    "uid": "team-a-email-notifier",
    "name": "Team A",
    "type": "email",
    "isDefault": false,
    "sendReminder": false,
    "disableResolveMessage": false,
    "settings": {
      "addresses": "dev@grafana.com"
    },
    "created": "2018-04-23T14:44:09+02:00",
    "updated": "2018-08-20T15:47:49+02:00"
  }
]
	`
	getAlertNotificationJSON = `
{
  "id": 1,
  "uid": "team-a-email-notifier",
  "name": "Team A",
  "type": "email",
  "isDefault": false,
  "sendReminder": false,
  "disableResolveMessage": false,
  "settings": {
    "addresses": "dev@grafana.com"
  },
  "created": "2018-04-23T14:44:09+02:00",
  "updated": "2018-08-20T15:47:49+02:00"
}
`
	createdAlertNotificationJSON = `
{
  "id": 1,
  "uid": "new-alert-notification",
  "name": "Team A",
  "type":  "email",
  "isDefault": false,
  "sendReminder": true,
  "frequency": "15m",
  "settings": {
    "addresses": "dev@grafana.com"
  }
}
`
	updatedAlertNotificationJSON = `
{
  "uid": "new-alert-notification",
  "name": "Team A",
  "type":  "email",
  "isDefault": false,
  "sendReminder": true,
  "frequency": "15m",
  "settings": {
    "addresses": "dev@grafana.com"
  }
}
`
	deletedAlertNotificationJSON = `
{
  "message":"Notification deleted"
}
`
)

func TestAlertNotifications(t *testing.T) {
	server, client := gapiTestTools(200, getAlertNotificationsJSON)
	defer server.Close()

	alertnotifications, err := client.AlertNotifications()
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(alertnotifications))

	if len(alertnotifications) != 1 {
		t.Error("Length of returned alert notifications should be 1")
	}
	if alertnotifications[0].Id != 1 || alertnotifications[0].Name != "Team A" {
		t.Error("Not correctly parsing returned alert notifications.")
	}
}

func TestAlertNotification(t *testing.T) {
	server, client := gapiTestTools(200, getAlertNotificationJSON)
	defer server.Close()

	alertnotification := int64(1)
	resp, err := client.AlertNotification(alertnotification)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Id != alertnotification || resp.Name != "Team A" {
		t.Error("Not correctly parsing returned alert notification.")
	}
}

func TestNewAlertNotification(t *testing.T) {
	server, client := gapiTestTools(200, createdAlertNotificationJSON)
	defer server.Close()

	an := &AlertNotification{
		Name:                  "Team A",
		Type:                  "email",
		IsDefault:             false,
		DisableResolveMessage: true,
		SendReminder:          true,
		Frequency:             "15m",
		Settings: map[string]string{
			"addresses": "dev@grafana.com",
		},
	}
	resp, err := client.NewAlertNotification(an)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp != 1 {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateAlertNotification(t *testing.T) {
	server, client := gapiTestTools(200, updatedAlertNotificationJSON)
	defer server.Close()

	an := &AlertNotification{
		Id:                    1,
		Name:                  "Team A",
		Type:                  "email",
		IsDefault:             false,
		DisableResolveMessage: true,
		SendReminder:          true,
		Frequency:             "15m",
		Settings: map[string]string{
			"addresses": "dev@grafana.com",
		},
	}

	err := client.UpdateAlertNotification(an)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAlertNotification(t *testing.T) {
	server, client := gapiTestTools(200, deletedAlertNotificationJSON)
	defer server.Close()

	err := client.DeleteAlertNotification(1)
	if err != nil {
		t.Error(err)
	}
}
