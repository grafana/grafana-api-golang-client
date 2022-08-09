package gapi

// NavLink represents a Grafana nav link.
type NavLink struct {
	ID     string `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	URL    string `json:"url,omitempty"`
	Target string `json:"target,omitempty"`
}

// NavbarPreference represents a Grafana navbar preference.
type NavbarPreference struct {
	SavedItems []NavLink `json:"savedItems"`
}

// QueryHistoryPreference represents a Grafana query history preference.
type QueryHistoryPreference struct {
	HomeTab string `json:"homeTab"`
}

// Preferences represents Grafana preferences.
type Preferences struct {
	Theme string `json:"theme"`
	// TODO: research whether this should be homeDashboardId or homeDashboardID
	// https://github.com/grafana/grafana/blob/8185b6fdf7b2f9bcef1f62dade4bb4087e23eba0/pkg/api/dtos/prefs.go#L7
	HomeDashboardID  int64                  `json:"homeDashboardId"`
	HomeDashboardUID string                 `json:"homeDashboardUID,omitempty"`
	Timezone         string                 `json:"timezone,omitempty"`
	WeekStart        string                 `json:"weekStart,omitempty"`
	Locale           string                 `json:"locale,omitempty"`
	Navbar           NavbarPreference       `json:"navbar,omitempty"`
	QueryHistory     QueryHistoryPreference `json:"queryHistory,omitempty"`
}
