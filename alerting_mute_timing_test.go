package gapi

import (
	"net/url"
	"testing"

	"github.com/gobs/pretty"
)

func TestMuteTimings(t *testing.T) {
	t.Run("get mute timings succeeds", func(t *testing.T) {
		/*server, client := gapiTestTools(t, 200, getMessageTemplatesJSON)
		defer server.Close()*/
		cfg := Config{
			BasicAuth: url.UserPassword("admin", "admin"),
		}
		client, _ := New("http://localhost:3000", cfg)

		mts, err := client.MuteTimings()

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(mts))
		if len(mts) != 2 {
			t.Errorf("wrong number of mute timings returned, got %#v", mts)
		}
		if mts[0].Name != "timing one" {
			t.Errorf("incorrect name - expected %s on element %d, got %#v", "timing one", 0, mts)
		}
		if mts[1].Name != "another timing" {
			t.Errorf("incorrect name - expected %s on element %d, got %#v", "another timing", 1, mts)
		}
	})
}

const getMuteTimingsJSON = `
[
	{
		"name": "timing one",
		"time_intervals": [
			{
				"times": [
					{
						"start_time": "13:13",
						"end_time": "15:15"
					}
				],
				"weekdays": [
					"monday:wednesday"
				],
				"months": [
					"1"
				]
			}
		]
	},
	{
		"name": "another timing",
		"time_intervals": [
			{
				"days_of_month": [
					"1"
				],
				"years": [
					"2030"
				]
			}
		]
	}
]`
