package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Plugin struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// InstallCloudPlugin installs the specified plugin to the given stack.
func (c *Client) InstallCloudPlugin(stackID int64, pluginSlug string, pluginVersion string) error {
	installPluginRequest := struct {
		Plugin  string `json:"plugin"`
		Version string `json:"version"`
	}{
		Plugin:  pluginSlug,
		Version: pluginVersion,
	}

	data, err := json.Marshal(installPluginRequest)
	if err != nil {
		return err
	}

	err = c.request("POST", fmt.Sprintf("/api/instances/%d/plugins", stackID), nil, bytes.NewBuffer(data), nil)
	if err != nil {
		return err
	}

	return nil
}

// UninstallCloudPlugin uninstalls the specified plugin to the given stack.
func (c *Client) UninstallCloudPlugin(stackID int64, pluginSlug string) error {
	return c.request("DElETE", fmt.Sprintf("/api/instances/%d/plugins/%s", stackID, pluginSlug), nil, nil, nil)
}

func (c *Client) IsCloudPluginInstalled(stackID int64, pluginSlug string) (bool, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/instances/%d/plugins/%s", stackID, pluginSlug), nil, nil)
	if err != nil {
		return false, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return false, nil
		}
		bodyContents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		return false, fmt.Errorf("status: %d, body: %v", resp.StatusCode, string(bodyContents))
	}

	return true, nil
}

// PluginBySlug returns the plugin with the given slug.
// An error will be returned given an unknown slug.
func (c *Client) PluginBySlug(slug string) (*Plugin, error) {
	p := Plugin{}

	err := c.request("GET", fmt.Sprintf("/api/plugins/%s", slug), nil, nil, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// PluginByID returns the plugin with the given id.
// An error will be returned given an unknown ID.
func (c *Client) PluginByID(pluginID int64) (*Plugin, error) {
	p := Plugin{}

	err := c.request("GET", fmt.Sprintf("/api/plugins/%d", pluginID), nil, nil, p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
