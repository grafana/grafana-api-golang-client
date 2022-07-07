package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestContactPoints(t *testing.T) {
	t.Run("get contact points succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getContactPointsJSON)
		defer server.Close()

		ps, err := client.ContactPoints()

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(ps))
		if len(ps) != 2 {
			t.Errorf("wrong number of contact points returned, got %#v", ps)
		}
		if ps[0].UID != "" {
			t.Errorf("incorrect UID - expected %s on element %d, got %#v", "", 0, ps)
		}
		if ps[1].UID != "rc5r0bjnz" {
			t.Errorf("incorrect UID - expected %s on element %d, got %#v", "rc5r0bjnz", 0, ps)
		}
	})

	t.Run("get contact points by name succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getContactPointsQueryJSON)
		defer server.Close()

		ps, err := client.ContactPointsByName("slack-receiver-1")

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(ps))
		if len(ps) != 1 {
			t.Errorf("wrong number of contact points returned, got %#v", ps)
		}
		if ps[0].UID != "rc5r0bjnz" {
			t.Errorf("incorrect UID - expected %s on element %d, got %#v", "rc5r0bjnz", 0, ps)
		}
	})

	t.Run("get contact point succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getContactPointsJSON)
		defer server.Close()

		p, err := client.ContactPoint("rc5r0bjnz")

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(p))
		if p.UID != "rc5r0bjnz" {
			t.Errorf("incorrect UID - expected %s got %#v", "rc5r0bjnz", p)
		}
	})

	t.Run("get non-existent contact point fails", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getContactPointsJSON)
		defer server.Close()

		p, err := client.ContactPoint("does not exist")

		if err == nil {
			t.Errorf("expected error but got nil")
			t.Log(pretty.PrettyFormat(p))
		}
	})

	t.Run("create contact point succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 201, writeContactPointJSON)
		defer server.Close()
		p := createContactPoint()

		uid, err := client.NewContactPoint(&p)

		if err != nil {
			t.Error(err)
		}
		if uid != "rc5r0bjnz" {
			t.Errorf("unexpected UID returned, got %s", uid)
		}
	})

	t.Run("update contact point succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, writeContactPointJSON)
		defer server.Close()
		p := createContactPoint()
		p.UID = "on7otbj7k"

		err := client.UpdateContactPoint(&p)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("delete contact point succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 204, "")
		defer server.Close()

		err := client.DeleteContactPoint("rc5r0bjnz")

		if err != nil {
			t.Error(err)
		}
	})
}

func createContactPoint() ContactPoint {
	return ContactPoint{
		Name:                  "slack-receiver-123",
		Type:                  "slack",
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

const getContactPointsQueryJSON = `
[
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
