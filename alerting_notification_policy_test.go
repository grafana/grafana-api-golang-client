package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestNotificationPolicies(t *testing.T) {
	t.Run("get policy tree succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, notificationPolicyJSON)
		defer server.Close()
		/*cfg := Config{
			BasicAuth: url.UserPassword("admin", "admin"),
		}
		client, _ := New("http://localhost:3000", cfg)*/

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
