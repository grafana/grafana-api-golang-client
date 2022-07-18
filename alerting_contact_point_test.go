package gapi

import (
	"fmt"
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client"
	"github.com/grafana/grafana-api-golang-client/goclient/client/provisioning"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
	"github.com/stretchr/testify/require"
)

func TestContactPoints(t *testing.T) {
	t.Run("get contact points succeeds", func(t *testing.T) {
		mocksrv, client := gapiTestTools(t, 200, getContactPointsJSON)
		defer mocksrv.Close()

		ps, err := client.Provisioning.RouteGetContactpoints(
			provisioning.NewRouteGetContactpointsParams(),
			nil,
		)

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(ps))
		if len(ps.Payload) != 2 {
			t.Errorf("wrong number of contact points returned, got %#v", ps)
		}
		if ps.Payload[0].UID != "" {
			t.Errorf("incorrect UID - expected %s on element %d, got %#v", "", 0, ps)
		}
		if ps.Payload[1].UID != "rc5r0bjnz" {
			t.Errorf("incorrect UID - expected %s on element %d, got %#v", "rc5r0bjnz", 0, ps)
		}
	})

	t.Run("get contact point succeeds", func(t *testing.T) {
		mocksrv, client := gapiTestTools(t, 200, getContactPointsJSON)
		defer mocksrv.Close()

		cp, err := getContactPointByUID(t, client, "rc5r0bjnz")
		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(cp))
		if cp.UID != "rc5r0bjnz" {
			t.Errorf("incorrect UID - expected %s got %#v", "rc5r0bjnz", cp)
		}
	})

	t.Run("get non-existent contact point fails", func(t *testing.T) {
		mocksrv, client := gapiTestTools(t, 200, getContactPointsJSON)
		defer mocksrv.Close()

		cp, err := getContactPointByUID(t, client, "does not exist")

		if err == nil {
			t.Errorf("expected error but got nil")
			t.Log(pretty.PrettyFormat(cp))
		}
	})

	t.Run("create contact point succeeds", func(t *testing.T) {
		mocksrv, client := gapiTestTools(t, 201, writeContactPointJSON)
		defer mocksrv.Close()
		p := createContactPoint()

		res, err := client.Provisioning.RoutePostContactpoints(
			provisioning.NewRoutePostContactpointsParams().
				WithBody(&p),
			nil,
		)

		if err != nil {
			t.Error(err)
		}
		if res.Payload.UID != "rc5r0bjnz" {
			t.Errorf("unexpected UID returned, got %s", res.Payload.UID)
		}
	})

	t.Run("update contact point succeeds", func(t *testing.T) {
		mocksrv, client := gapiTestTools(t, 200, writeContactPointJSON)
		defer mocksrv.Close()
		p := createContactPoint()
		existingUID := p.UID
		newUID := "on7otbj7k"
		p.UID = newUID

		_, err := client.Provisioning.RoutePutContactpoint(
			provisioning.NewRoutePutContactpointParams().
				WithUID(existingUID).
				WithBody(&p),
			nil,
		)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("delete contact point succeeds", func(t *testing.T) {
		mocksrv, client := gapiTestTools(t, 204, "")
		defer mocksrv.Close()

		_, err := client.Provisioning.RouteDeleteContactpoints(
			provisioning.NewRouteDeleteContactpointsParams().
				WithUID("rc5r0bjnz"),
			nil,
		)

		if err != nil {
			t.Error(err)
		}
	})
}

func createContactPoint() models.EmbeddedContactPoint {
	t := "slack"
	return models.EmbeddedContactPoint{
		Name:                  "slack-receiver-123",
		Type:                  &t,
		DisableResolveMessage: false,
		Settings: map[string]interface{}{
			"recipient": "@zxcv",
			"token":     "test-token",
			"url":       "https://test-url",
		},
	}
}

const getContactPointsJSON = `
[
	{
		"uid": "",
		"name": "default-email-receiver",
		"type": "email",
		"disableResolveMessage": false,
		"settings": {
			"addresses": "<example@email.com>"
		}
	},
	{
		"uid": "rc5r0bjnz",
		"name": "slack-receiver-1",
		"type": "slack",
		"disableResolveMessage": false,
		"settings": {
			"recipient": "@foo",
			"token": "[REDACTED]",
			"url": "[REDACTED]"
		}
	}
]`

const writeContactPointJSON = `
{
	"uid": "rc5r0bjnz",
	"name": "slack-receiver-1",
	"type": "slack",
	"disableResolveMessage": false,
	"settings": {
		"recipient": "@foo",
		"token": "[REDACTED]",
		"url": "[REDACTED]"
	}
}
`

func getContactPointByUID(t *testing.T, client *client.GrafanaHTTPAPI, uid string) (*models.EmbeddedContactPoint, error) {
	t.Helper()

	ps, err := client.Provisioning.RouteGetContactpoints(
		provisioning.NewRouteGetContactpointsParams(),
		nil,
	)
	require.NoError(t, err)
	for _, p := range ps.Payload {
		if p.UID == uid {
			return p, nil
		}
	}
	return nil, fmt.Errorf("contact point with uid %s not found", uid)
}
