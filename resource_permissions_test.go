package gapi

import (
	"net/http"
	"testing"

	"github.com/gobs/pretty"
)

const (
	resourcePermissionsListJSON = `[
		{
			"id": 1,
			"roleName": "basic:admin",
			"isManaged": false,
			"isInherited": false,
			"isServiceAccount": false,
			"builtInRole": "Admin",
			"actions": [
				"datasources:delete",
				"datasources:query",
				"datasources:read",
				"datasources:write",
				"datasources.caching:read",
				"datasources.caching:write",
				"datasources.permissions:read",
				"datasources.permissions:write"
			],
			"permission": "Admin"
		}
	]`
	resourcePermissionsResponseJSON = `{"message":"Permissions updated"}`
)

func TestListResourcePermissions(t *testing.T) {
	client := gapiTestTools(t, http.StatusOK, resourcePermissionsListJSON)
	res, err := client.listResourcePermissions("datasources", ResourceID(1))
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(res))
}

func TestSetResourcePermissions(t *testing.T) {
	client := gapiTestTools(t, http.StatusOK, resourcePermissionsResponseJSON)
	res, err := client.setResourcePermissions("datasources", ResourceID(1), SetResourcePermissionsBody{
		Permissions: []SetResourcePermissionItem{
			{
				UserID:     1,
				Permission: "View",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(res))
}

func TestSetResourcePermissionsByAssignment(t *testing.T) {
	client := gapiTestTools(t, http.StatusOK, resourcePermissionsResponseJSON)
	res, err := client.setResourcePermissionByAssignment("datasources", ResourceID(1), "users", ResourceID(1), SetResourcePermissionBody{
		Permission: SetResourcePermissionItem{
			Permission: "View",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(res))
}
