package gapi

import (
	"testing"
)

const (
	getOrgPreferencesJSON    = `{"theme": "foo","homeDashboardId": 0,"timezone": "","weekStart": "","navbar": {"savedItems": null},"queryHistory": {"homeTab": ""}}`
	updateOrgPreferencesJSON = `{"message":"Preferences updated"}`
)

func TestOrgPreferences(t *testing.T) {
	server, client := gapiTestTools(t, 200, getOrgPreferencesJSON)
	defer server.Close()

	resp, err := client.OrgPreferences()
	if err != nil {
		t.Fatal(err)
	}

	expected := "foo"
	if resp.Theme != expected {
		t.Errorf("Expected org preferences theme '%s'; got '%s'", expected, resp.Theme)
	}
}

func TestUpdateOrgPreferences(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateOrgPreferencesJSON)
	defer server.Close()

	resp, err := client.UpdateOrgPreferences(Preferences{
		Theme: "foo",
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := "Preferences updated"
	if resp.Message != expected {
		t.Errorf("Expected org preferences message '%s'; got '%s'", expected, resp.Message)
	}
}

func TestUpdateAllOrgPreference(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateOrgPreferencesJSON)
	defer server.Close()

	resp, err := client.UpdateAllOrgPreferences(Preferences{
		Theme: "foo",
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := "Preferences updated"
	if resp.Message != expected {
		t.Errorf("Expected org preferences message '%s'; got '%s'", expected, resp.Message)
	}
}
