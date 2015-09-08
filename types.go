package gapi

import "time"

type Collector struct {
	Id             int
	Org_id         int
	Slug           string
	Name           string
	Tags           []string
	Public         bool
	Latitude       int
	Longitude      int
	Enabled        bool
	Online         bool
	Enabled_change time.Time
	Online_change  time.Time
}

type Org struct {
	Id   int
	Name string
}
