package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/legacy_alerts_notification_channels"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
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
	mocksrv, client := gapiTestTools(t, 200, getAlertNotificationsJSON)
	defer mocksrv.Close()

	alertnotifications, err := client.LegacyAlertsNotificationChannels.GetAlertNotificationChannels(
		legacy_alerts_notification_channels.NewGetAlertNotificationChannelsParams(),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(alertnotifications))

	if len(alertnotifications.Payload) != 1 {
		t.Error("Length of returned alert notifications should be 1")
	}
	if alertnotifications.Payload[0].ID != 1 || alertnotifications.Payload[0].Name != "Team A" {
		t.Error("Not correctly parsing returned alert notifications.")
	}
}

func TestAlertNotification(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getAlertNotificationJSON)
	defer mocksrv.Close()

	alertnotification := int64(1)
	resp, err := client.LegacyAlertsNotificationChannels.GetAlertNotificationChannelByID(
		legacy_alerts_notification_channels.NewGetAlertNotificationChannelByIDParams().
			WithNotificationChannelID(alertnotification),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.ID != alertnotification || resp.Payload.Name != "Team A" {
		t.Error("Not correctly parsing returned alert notification.")
	}
}

func TestNewAlertNotification(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createdAlertNotificationJSON)
	defer mocksrv.Close()

	an := &models.CreateAlertNotificationCommand{
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
	resp, err := client.LegacyAlertsNotificationChannels.CreateAlertNotificationChannel(
		legacy_alerts_notification_channels.NewCreateAlertNotificationChannelParams().
			WithBody(an),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.ID != 1 {
		t.Error("Not correctly parsing returned creation message.")
	}
}

func TestUpdateAlertNotification(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updatedAlertNotificationJSON)
	defer mocksrv.Close()

	an := &models.UpdateAlertNotificationCommand{
		ID:                    1,
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

	_, err := client.LegacyAlertsNotificationChannels.UpdateAlertNotificationChannel(
		legacy_alerts_notification_channels.NewUpdateAlertNotificationChannelParams().
			WithNotificationChannelID(an.ID).
			WithBody(an),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAlertNotification(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deletedAlertNotificationJSON)
	defer mocksrv.Close()

	_, err := client.LegacyAlertsNotificationChannels.DeleteAlertNotificationChannel(
		legacy_alerts_notification_channels.NewDeleteAlertNotificationChannelParams().
			WithNotificationChannelID(1),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
