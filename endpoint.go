package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	m "github.com/grafana/grafana/pkg/models"
)

func (c *Client) NewEndpoint(settings m.AddEndpointCommand) error {
	data, err := json.Marshal(settings)
	req, err := c.newRequest("PUT", "/api/endpoints", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("response BODY", string(data))
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}
