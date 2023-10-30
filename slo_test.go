package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestSLOs(t *testing.T) {
	t.Run("list all SLOs succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 200, getSlosJSON)

		resp, err := client.ListSlos()

		slos := resp.Slos

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(slos))
		if len(slos) != 1 {
			t.Errorf("wrong number of contact points returned, got %d", len(slos))
		}
		if slos[0].Name != "list-slos" {
			t.Errorf("incorrect name - expected Name-Test, got %s", slos[0].Name)
		}
	})

	t.Run("get individual SLO succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 200, getSloJSON)

		slo, err := client.GetSlo("qkkrknp12w6tmsdcrfkdf")

		t.Log(pretty.PrettyFormat(slo))
		if err != nil {
			t.Error(err)
		}
		if slo.Uuid != "qkkrknp12w6tmsdcrfkdf" {
			t.Errorf("incorrect UID - expected qkkrknp12w6tmsdcrfkdf, got %s", slo.Uuid)
		}
	})

	t.Run("get non-existent SLOs fails", func(t *testing.T) {
		client := gapiTestTools(t, 404, getSlosJSON)

		slo, err := client.GetSlo("qkkrknp12w6tmsdcrfkdf")

		if err == nil {
			t.Log(pretty.PrettyFormat(slo))
			t.Error("expected error but got nil")
		}
	})

	t.Run("create SLO succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 201, createSloJSON)
		slo := generateSlo()

		resp, err := client.CreateSlo(slo)

		if err != nil {
			t.Error(err)
		}
		if resp.UUID != "sjnp8wobcbs3eit28n8yb" {
			t.Errorf("unexpected UID returned, got %s", resp.UUID)
		}
	})

	t.Run("update SLO succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 200, createSloJSON)
		slo := generateSlo()
		slo.Description = "Updated Description"

		err := client.UpdateSlo(slo.Uuid, slo)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("delete SLO succeeds", func(t *testing.T) {
		client := gapiTestTools(t, 204, "")

		err := client.DeleteSlo("qkkrknp12w6tmsdcrfkdf")

		if err != nil {
			t.Log(err)
			t.Error(err)
		}
	})
}

const getSlosJSON = `
{
    "slos": [
        {
            "uuid": "qkkrknp12w6tmsdcrfkdf",
            "name": "list-slos",
            "description": "list-slos-description",
			"query": {
                "freeform": {
                    "query": "sum(rate(apiserver_request_total{code!=\"500\"}[$__rate_interval])) / sum(rate(apiserver_request_total[$__rate_interval]))"
                },
                "type": "freeform"
            },
            "objectives": [
                {
                    "value": 0.995,
                    "window": "28d"
                }
            ],
            "drillDownDashboardRef": {
                "uid": "5IkqX6P4k"
            }
        }
    ]
}`

const getSloJSON = `
{
    "uuid": "qkkrknp12w6tmsdcrfkdf",
    "name": "Name-Test",
    "description": "Description-Test",
	"query": {
		"freeform": {
			"query": "sum(rate(apiserver_request_total{code!=\"500\"}[$__rate_interval])) / sum(rate(apiserver_request_total[$__rate_interval]))"
		},
		"type": "freeform"
	},
    "objectives": [
        {
            "value": 0.995,
            "window": "28d"
        }
    ],
    "drillDownDashboardRef": {
        "uid": "5IkqX6P4k"
    }
}`

const createSloJSON = `
{
    "uuid": "sjnp8wobcbs3eit28n8yb",
    "name": "test-name",
    "description": "test-description",
	"query": {
		"freeform": {
			"query": "sum(rate(apiserver_request_total{code!=\"500\"}[$__rate_interval])) / sum(rate(apiserver_request_total[$__rate_interval]))"
		},
		"type": "freeform"
	},
    "objectives": [
        {
            "value": 0.995,
            "window": "30d"
        }
    ],
    "drillDownDashboardRef": {
        "uid": "zz5giRyVk"
    }
}
`

func generateSlo() Slo {
	objective := []Objective{{Value: 0.995, Window: "30d"}}
	query := Query{
		Freeform: &FreeformQuery{
			Query: "sum(rate(apiserver_request_total{code!=\"500\"}[$__rate_interval])) / sum(rate(apiserver_request_total[$__rate_interval]))",
		},
		Type: QueryTypeFreeform,
	}

	slo := Slo{
		Name:        "test-name",
		Description: "test-description",
		Objectives:  objective,
		Query:       query,
	}

	return slo
}
