package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestNotificationPolicies(t *testing.T) {
	t.Run("get policy tree succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, notificationPolicyJSON)
		defer server.Close()

		np, err := client.NotificationPolicy()

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(np))
		if np.Receiver != "grafana-default-email" {
			t.Errorf("wrong receiver, got %#v", np.Receiver)
		}
		if len(np.Routes) != 1 {
			t.Errorf("wrong number of routes returned, got %#v", np)
		}
	})

	t.Run("set policy tree succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 202, `{"message":"created"}`)
		defer server.Close()
		np := createNotificationPolicy()

		err := client.SetNotificationPolicy(&np)

		if err != nil {
			t.Error(err)
		}
	})
}

func createNotificationPolicy() NotificationPolicy {
	return NotificationPolicy{
		Receiver: "grafana-default-email",
		GroupBy:  []string{"asdfasdf", "alertname"},
		Routes: []SpecificPolicy{
			{
				Receiver: "grafana-default-email",
				ObjectMatchers: Matchers{
					{
						Type:  MatchNotEqual,
						Name:  "abc",
						Value: "def",
					},
				},
				Continue: true,
			},
			{
				Receiver: "grafana-default-email",
				ObjectMatchers: Matchers{
					{
						Type:  MatchRegexp,
						Name:  "jkl",
						Value: "something.*",
					},
				},
				Continue: false,
			},
		},
		GroupWait:      "10s",
		GroupInterval:  "5m",
		RepeatInterval: "4h",
	}
}

const notificationPolicyJSON = `
{
	"receiver": "grafana-default-email",
	"group_by": [
		"..."
	],
	"routes": [
		{
			"receiver": "grafana-default-email",
			"object_matchers": [
				[
					"a",
					"=",
					"b"
				],
				[
					"asdf",
					"!=",
					"jk"
				]
			]
		}
	],
	"group_wait": "5s",
	"group_interval": "1m",
	"repeat_interval": "1h"
}`
