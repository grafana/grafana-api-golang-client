package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// AlertNotification represents a Grafana alert notification.
type AlertNotification struct {
	Id                    int64       `json:"id,omitempty"`
	Uid                   string      `json:"uid"`
	Name                  string      `json:"name"`
	Type                  string      `json:"type"`
	IsDefault             bool        `json:"isDefault"`
	DisableResolveMessage bool        `json:"disableResolveMessage"`
	SendReminder          bool        `json:"sendReminder"`
	Frequency             string      `json:"frequency"`
	Settings              interface{} `json:"settings"`
}

// AlertNotifications fetches and returns Grafana alert notifications.
func (c *Client) AlertNotifications() ([]AlertNotification, error) {
	alertnotifications := make([]AlertNotification, 0)

	resp, err := c.request("GET", "/api/alert-notifications/", nil, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &alertnotifications)
	return alertnotifications, err
}

// AlertNotification fetches and returns a Grafana alert notification.
func (c *Client) AlertNotification(id int64) (*AlertNotification, error) {
	path := fmt.Sprintf("/api/alert-notifications/%d", id)
	resp, err := c.request("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &AlertNotification{}
	err = json.Unmarshal(data, &result)
	return result, err
}

// NewAlertNotification creates a new Grafana alert notification.
func (c *Client) NewAlertNotification(a *AlertNotification) (int64, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return 0, err
	}
	resp, err := c.request("POST", "/api/alert-notifications", nil, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	result := struct {
		Id int64 `json:"id"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Id, err
}

// UpdateAlertNotification updates a Grafana alert notification.
func (c *Client) UpdateAlertNotification(a *AlertNotification) error {
	path := fmt.Sprintf("/api/alert-notifications/%d", a.Id)
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	_, err = c.request("PUT", path, nil, bytes.NewBuffer(data))

	return err

}

// DeleteAlertNotification deletes a Grafana alert notification.
func (c *Client) DeleteAlertNotification(id int64) error {
	path := fmt.Sprintf("/api/alert-notifications/%d", id)
	_, err := c.request("DELETE", path, nil, nil)

	return err
}
