package gapi

func (c *Client) ListServiceAccountResourcePermissions(ident ResourceIdent) ([]*ResourcePermission, error) {
	return c.listResourcePermissions("serviceaccounts", ident)
}

func (c *Client) SetServiceAccountResourcePermissions(ident ResourceIdent, body SetResourcePermissionsBody) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissions("serviceaccounts", ident, body)
}

func (c *Client) SetUserServiceAccountResourcePermissions(ident ResourceIdent, userID int64, permission string) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissionByAssignment(
		"serviceaccounts",
		ident,
		"users",
		ResourceID(userID),
		SetResourcePermissionBody{
			Permission: SetResourcePermissionItem{
				UserID:     userID,
				Permission: permission,
			},
		},
	)
}

func (c *Client) SetTeamServiceAccountResourcePermissions(ident ResourceIdent, teamID int64, permission string) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissionByAssignment(
		"serviceaccounts",
		ident,
		"teams",
		ResourceID(teamID),
		SetResourcePermissionBody{
			Permission: SetResourcePermissionItem{
				TeamID:     teamID,
				Permission: permission,
			},
		},
	)
}
