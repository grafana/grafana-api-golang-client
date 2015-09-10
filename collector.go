package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	m "github.com/grafana/grafana/pkg/models"
)

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

func (c *Client) Collectors(settings m.GetCollectorsQuery) ([]Collector, error) {
	data, _ := json.Marshal(settings)
	collectors := make([]Collector, 0)

	req, err := c.newRequest("GET", "/api/collectors/", bytes.NewBuffer(data))
	if err != nil {
		return collectors, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return collectors, err
	}
	if resp.StatusCode != 200 {
		return collectors, errors.New(resp.Status)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return collectors, err
	}
	err = json.Unmarshal(data, &collectors)
	return collectors, err
}
