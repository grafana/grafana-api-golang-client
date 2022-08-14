// +build acceptance

package gapi

import (
	"testing"
)

func TestTeamAcceptance(t *testing.T) {
	teamID, err := client.AddTeam("foo", "foo@bar.com")
	if err != nil {
		t.Fatalf("failed to add team: %v", err)
	}

	var hDBID int64 = 123
	err = client.UpdateTeamPreferences(teamID, Preferences{
		HomeDashboardID: hDBID,
	})
	if err != nil {
		t.Fatalf("failed to update team preferences: %v", err)
	}

	prefs, err := client.TeamPreferences(teamID)
	if err != nil {
		t.Fatalf("failed to get team preferences: %v", err)
	}

	if prefs.HomeDashboardID != hDBID {
		t.Fatalf("expected home dashboard ID to be '%d'; got '%d'", hDBID, prefs.HomeDashboardID)
	}

	err = client.DeleteTeam(teamID)
	if err != nil {
		t.Fatalf("failed to delete team: %v", err)
	}
}
