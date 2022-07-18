package gapi

import (
	"testing"
)

const (
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
		"id": 34,
		"name": "Some Plugin",
		"slug": "some-plugin",
		"version": "1.2.3",
		"description": "Some Plugin for adding functionality"
	}`
)

func TestInstallCloudPlugin(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, installPluginJSON)
	defer mocksrv.Close()

	client := getCloudClient(t, mocksrv.server.URL)

	installation, err := client.InstallCloudPlugin("some-stack", "some-plugin", "1.2.3")
	if err != nil {
		t.Error(err)
	}

	expectedInstallation := CloudPluginInstallation{}
	err = UnmarshalJSONToStruct(installPluginJSON, &expectedInstallation)
	if err != nil {
		t.Fatal(err)
	}

	if *installation != expectedInstallation {
		t.Errorf("Unexpected installation - Actual: %v, Expected: %v", installation, expectedInstallation)
	}

	for _, code := range []int{401, 403, 404, 412} {
		mocksrv.code = code

		installation, err = client.InstallCloudPlugin("some-stack", "some-plugin", "1.2.3")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
		if installation != nil {
			t.Errorf("Expected empty installation, got %v", installation)
		}
	}
}

func TestUninstallCloudPlugin(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, uninstallPluginJSON)
	defer mocksrv.Close()

	client := getCloudClient(t, mocksrv.server.URL)

	err := client.UninstallCloudPlugin("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	for _, code := range []int{401, 403, 404, 412} {
		mocksrv.code = code

		err = client.UninstallCloudPlugin("some-stack", "some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestIsCloudPluginInstalled(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, getPluginJSON)

	client := getCloudClient(t, mocksrv.server.URL)

	ok, err := client.IsCloudPluginInstalled("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Errorf("Expected plugin installation - Expected true, got false")
	}

	mocksrv.code = 404
	ok, err = client.IsCloudPluginInstalled("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	if ok {
		t.Errorf("Unexpected plugin installation - Expected false, got true")
	}

	for _, code := range []int{401, 403, 412} {
		mocksrv.code = code

		_, err := client.IsCloudPluginInstalled("some-stack", "some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestGetCloudPluginInstallation(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, installPluginJSON)
	defer mocksrv.Close()

	client := getCloudClient(t, mocksrv.server.URL)

	installation, err := client.GetCloudPluginInstallation("some-stack", "some-plugin")
	if err != nil {
		t.Error(err)
	}

	expectedInstallation := CloudPluginInstallation{}
	err = UnmarshalJSONToStruct(installPluginJSON, &expectedInstallation)
	if err != nil {
		t.Fatal(err)
	}

	if *installation != expectedInstallation {
		t.Errorf("Unexpected installation - Actual: %v, Expected: %v", installation, expectedInstallation)
	}

	for _, code := range []int{401, 403, 404, 412} {
		mocksrv.code = code

		installation, err = client.GetCloudPluginInstallation("some-stack", "some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
		if installation != nil {
			t.Errorf("Expected empty installation, got %v", installation)
		}
	}
}

func TestPlugin(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, getPluginJSON)
	defer mocksrv.Close()

	client := getCloudClient(t, mocksrv.server.URL)

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
		mocksrv.code = code

		_, err = client.PluginBySlug("some-plugin")
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}

func TestPluginByID(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, getPluginJSON)
	defer mocksrv.Close()

	client := getCloudClient(t, mocksrv.server.URL)

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
		mocksrv.code = code

		_, err = client.PluginByID(123)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}
