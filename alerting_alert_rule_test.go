//go:build !integration
// +build !integration

package gapi

import (
	"encoding/json"
	"testing"

	"github.com/gobs/pretty"
)

func TestAlertRules(t *testing.T) {
	t.Run("get alert rule succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getAlertRuleJSON)
		defer server.Close()

		alertRule, err := client.AlertRule("123abcd")

		if err != nil {
			t.Error(err)
		}
		if alertRule.UID != "123abcd" {
			t.Errorf("incorrect UID - expected %s got %s", "123abcd", alertRule.UID)
		}
	})

	t.Run("get alert rule group succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, getAlertRuleGroupJSON)
		defer server.Close()

		group, err := client.AlertRuleGroup("d8-gk06nz", "test")

		if err != nil {
			t.Error(err)
		}
		t.Log(pretty.PrettyFormat(group))
		if group.Title != "test" {
			t.Errorf("incorrect title - expected %s got %s", "test", group.Title)
		}
		if group.FolderUID != "d8-gk06nz" {
			t.Errorf("incorrect folderUID - expected %s got %s", "d8-gk06nz", group.FolderUID)
		}
		if len(group.Rules) != 1 {
			t.Errorf("wrong number of rules, got %d", len(group.Rules))
		}
	})

	t.Run("get non-existent alert rule fails", func(t *testing.T) {
		server, client := gapiTestTools(t, 404, "")
		defer server.Close()

		alertRule, err := client.AlertRule("does not exist")

		if err == nil {
			t.Errorf("expected error but got nil")
			t.Log(pretty.PrettyFormat(alertRule))
		}
	})

	t.Run("create alert rule succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 201, writeAlertRuleJSON)
		defer server.Close()
		alertRule := createAlertRule()
		uid, err := client.NewAlertRule(&alertRule)

		if err != nil {
			t.Error(err)
		}
		if uid != "123abcd" {
			t.Errorf("unexpected UID returned, got %s", uid)
		}
	})

	t.Run("update alert rule succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 200, writeAlertRuleJSON)
		defer server.Close()
		alertRule := createAlertRule()
		alertRule.UID = "foobar"

		err := client.UpdateAlertRule(&alertRule)

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("delete alert rule succeeds", func(t *testing.T) {
		server, client := gapiTestTools(t, 204, "")
		defer server.Close()

		err := client.DeleteAlertRule("123abcd")

		if err != nil {
			t.Error(err)
		}
	})
}

func createAlertRule() AlertRule {
	return AlertRule{
		Condition:    "A",
		Data:         createAlertQueries(),
		ExecErrState: ErrOK,
		FolderUID:    "project_test",
		NoDataState:  NoDataOk,
		OrgID:        1,
		RuleGroup:    "eval_group_1",
		Title:        "Always in alarm",
		ForDuration:  0,
	}
}

func createAlertQueries() []*AlertQuery {
	alertQueries := make([]*AlertQuery, 1)
	alertQueries[0] = &AlertQuery{
		DatasourceUID:     "-100",
		Model:             json.RawMessage(`{"datasourceUid":"-100","model":{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"},"queryType":"","refId":"A","relativeTimeRange":{"from":0,"to":0}}`),
		QueryType:         "",
		RefID:             "A",
		RelativeTimeRange: RelativeTimeRange{From: 0, To: 0},
	}
	return alertQueries
}

const writeAlertRuleJSON = `
	{
	"conditions": "A",
	"data": [{"datasourceUid":"-100","model":{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"},"queryType":"","refId":"A","relativeTimeRange":{"from":0,"to":0}}],
	"uid": "123abcd",
	"execErrState": "OK",
	"folderUID": "project_test",
	"noDataState": "OK",
	"orgId": 1,
	"ruleGroup": "eval_group_1",
	"title": "Always in alarm",
	"for": 0
}
`

const getAlertRuleJSON = `
	{
	"conditions": "A",
	"data": [{"datasourceUid":"-100","model":{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"},"queryType":"","refId":"A","relativeTimeRange":{"from":0,"to":0}}],
	"execErrState": "OK",
	"folderUID": "project_test",
	"noDataState": "OK",
	"orgId": 1,
	"uid": "123abcd",
	"ruleGroup": "eval_group_1",
	"title": "Always in alarm",
	"for": 0
}
`

const getAlertRuleGroupJSON = `
{
	"title": "test",
	"folderUid": "d8-gk06nz",
	"interval": 60,
	"rules": [
		{
			"ID": 1,
			"OrgID": 1,
			"Title": "abc",
			"Condition": "B",
			"Data": [
				{
					"refId": "A",
					"queryType": "",
					"relativeTimeRange": {
						"from": 600,
						"to": 0
					},
					"datasourceUid": "PD8C576611E62080A",
					"model": {
						"hide": false,
						"intervalMs": 1000,
						"maxDataPoints": 43200,
						"refId": "A"
					}
				},
				{
					"refId": "B",
					"queryType": "",
					"relativeTimeRange": {
						"from": 0,
						"to": 0
					},
					"datasourceUid": "-100",
					"model": {
						"conditions": [
							{
								"evaluator": {
									"params": [
										3
									],
									"type": "gt"
								},
								"operator": {
									"type": "and"
								},
								"query": {
									"params": [
										"A"
									]
								},
								"reducer": {
									"params": [],
									"type": "last"
								},
								"type": "query"
							}
						],
						"datasource": {
							"type": "__expr__",
							"uid": "-100"
						},
						"hide": false,
						"intervalMs": 1000,
						"maxDataPoints": 43200,
						"refId": "B",
						"type": "classic_conditions"
					}
				}
			],
			"Updated": "2022-07-07T16:23:56-05:00",
			"IntervalSeconds": 60,
			"Version": 1,
			"UID": "hsXgz0enz",
			"NamespaceUID": "d8-gk06nz",
			"DashboardUID": null,
			"PanelID": null,
			"RuleGroup": "test",
			"RuleGroupIndex": 1,
			"NoDataState": "NoData",
			"ExecErrState": "Alerting",
			"For": 300000000000,
			"Annotations": {},
			"Labels": {}
		}
	]
}`
