package gapi

func (c *Client) ListServiceAccountResourcePermissions(uid string) ([]*ResourcePermission, error) {
	return c.listResourcePermissions(ServiceAccountsResource, ResourceUID(uid))
}

func (c *Client) SetServiceAccountResourcePermissions(uid string, body SetResourcePermissionsBody) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissions(ServiceAccountsResource, ResourceUID(uid), body)
}

func (c *Client) SetUserServiceAccountResourcePermissions(serviceAccountUID string, userID int64, permission string) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissionByAssignment(
		ServiceAccountsResource,
		ResourceUID(serviceAccountUID),
		UsersResource,
		ResourceID(userID),
		SetResourcePermissionBody{
			Permission: SetResourcePermissionItem{
				UserID:     userID,
				Permission: permission,
			},
		},
	)
}

func (c *Client) SetTeamServiceAccountResourcePermissions(serviceAccountUID string, teamID int64, permission string) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissionByAssignment(
		ServiceAccountsResource,
		ResourceUID(serviceAccountUID),
		TeamsResource,
		ResourceID(teamID),
		SetResourcePermissionBody{
			Permission: SetResourcePermissionItem{
				TeamID:     teamID,
				Permission: permission,
			},
		},
	)
}
