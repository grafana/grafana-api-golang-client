package gapi

import (
	"testing"
)

const (
	installPluginID   = 123
	installPluginJSON = `
	{
		"id": 123,
		"instanceId": 2,
		"instanceUrl": "mystack.grafana.net",
		"instanceSlug": "mystack",
		"pluginId": 3,
		"pluginSlug": "some-plugin",
		"pluginName": "Some Plugin",
		"version": "1.2.3",
		"latestVersion": "1.2.3",
		"createdAt": "2021-12-22T14:02:46.000Z",
		"updatedAt": null
	}`
	uninstallPluginJSON = `
	{
		"id": 123,
		"instanceId": 2,
		"instanceUrl": "mystack.grafana.net",
		"instanceSlug": "mystack",
		"pluginId": 3,
		"pluginSlug": "some-plugin",
		"pluginName": "Some Plugin",
		"version": "1.2.3",
		"latestVersion": "1.2.3",
		"createdAt": "2021-12-22T14:02:46.000Z",
		"updatedAt": null
	}`
	getPluginJSON = `
	{
		"name": "Some Plugin",
		"slug": "some-plugin",
		"version": "1.2.3",
		"description": "Some Plugin for adding functionality"
	}`
)

func TestInstallCloudPlugin(t *testing.T) {
	server, client := gapiTestTools(t, 200, installPluginJSON)
	defer server.Close()

	id, err := client.InstallCloudPlugin("some-stack", "some-plugin", "1.2.3")
	if err != nil {
		t.Error(err)
	}

	if id != installPluginID {
		t.Errorf("unexpected ID received: %d", id)
	}

	for _, code := range []int{401, 403, 404, 412} {
		server.code = code

		id, err = client.InstallCloudPlugin("some-stack", "some-plugin", "1.2.3")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
		if id != 0 {
			t.Errorf("unexpected ID received: %d", id)
		}
	}
}

func TestUninstallCloudPlugin(t *testing.T) {
	server, client := gapiTestTools(t, 200, uninstallPluginJSON)
	defer server.Close()

	err := client.UninstallCloudPlugin("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	for _, code := range []int{401, 403, 404, 412} {
		server.code = code

		err = client.UninstallCloudPlugin("some-stack", "some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestIsCloudPluginInstalled(t *testing.T) {
	server, client := gapiTestTools(t, 200, getPluginJSON)

	ok, err := client.IsCloudPluginInstalled("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Errorf("Expected plugin installation - Expected true, got false")
	}

	server.code = 404
	ok, err = client.IsCloudPluginInstalled("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	if ok {
		t.Errorf("Unexpected plugin installation - Expected false, got true")
	}

	for _, code := range []int{401, 403, 412} {
		server.code = code

		_, err := client.IsCloudPluginInstalled("some-stack", "some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestPlugin(t *testing.T) {
	server, client := gapiTestTools(t, 200, getPluginJSON)
	defer server.Close()

	plugin, err := client.PluginBySlug("some-plugin")
	if err != nil {
		t.Error(err)
	}

	expectedPlugin := Plugin{}
	err = UnmarshalJSONToStruct(getPluginJSON, &expectedPlugin)
	if err != nil {
		t.Fatal(err)
	}

	if *plugin != expectedPlugin {
		t.Errorf("Unexpected plugin - Actual: %v, Expected: %v", plugin, expectedPlugin)
	}

	for _, code := range []int{404} {
		server.code = code

		_, err = client.PluginBySlug("some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestPluginByID(t *testing.T) {
	server, client := gapiTestTools(t, 200, getPluginJSON)
	defer server.Close()

	plugin, err := client.PluginBySlug("some-plugin")
	if err != nil {
		t.Error(err)
	}

	expectedPlugin := Plugin{}
	err = UnmarshalJSONToStruct(getPluginJSON, &expectedPlugin)
	if err != nil {
		t.Fatal(err)
	}

	if *plugin != expectedPlugin {
		t.Errorf("Unexpected plugin - Actual: %v, Expected: %v", plugin, expectedPlugin)
	}

	for _, code := range []int{404} {
		server.code = code

		_, err = client.PluginByID(123)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}
