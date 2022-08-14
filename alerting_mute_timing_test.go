//+build !integration

package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestMuteTimings(t *testing.T) {
	t.Run("get mute timings succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getMuteTimingsJSON)
		defer server.Close()

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

	t.Run("get mute timing succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, muteTimingJSON)
		defer server.Close()

		mt, err := client.MuteTiming("timing one")

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(mt))
		if mt.Name != "timing one" {
			t.Errorf("incorrect name - expected %s, got %#v", "timing one", mt)
		}
	})

	t.Run("get non-existent mute timing fails", func(t *testing.T) {
		server, client := gapiTestTools(t, 404, muteTimingJSON)
		defer server.Close()

		mt, err := client.MuteTiming("does not exist")

		if err == nil {
			t.Errorf("expected error but got nil")
			t.Log(pretty.PrettyFormat(mt))
		}
	})

	t.Run("create mute timing succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 201, muteTimingJSON)
		defer server.Close()
		mt := createMuteTiming()

		err := client.NewMuteTiming(&mt)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("update mute timing succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, muteTimingJSON)
		defer server.Close()
		mt := createMuteTiming()
		mt.TimeIntervals[0].Weekdays = []WeekdayRange{"tuesday", "thursday"}

		err := client.UpdateMuteTiming(&mt)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("delete mute timing succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 204, muteTimingJSON)
		defer server.Close()

		err := client.DeleteMuteTiming("timing two")

		if err != nil {
			t.Error(err)
		}
	})
}

func createMuteTiming() MuteTiming {
	return MuteTiming{
		Name: "timing two",
		TimeIntervals: []TimeInterval{
			{
				Times: []TimeRange{
					{
						StartMinute: "13:13",
						EndMinute:   "15:15",
					},
				},
				Weekdays: []WeekdayRange{"monday", "wednesday"},
				Months:   []MonthRange{"1:3", "4"},
				Years:    []YearRange{"2022", "2023"},
			},
		},
	}
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

const muteTimingJSON = `
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
}`
