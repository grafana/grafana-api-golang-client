package gapi

func (c *Client) ListTeamResourcePermissions(ident ResourceIdent) ([]*ResourcePermission, error) {
	return c.listResourcePermissions("teams", ident)
}

func (c *Client) SetTeamResourcePermissions(ident ResourceIdent, body SetResourcePermissionsBody) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissions("teams", ident, body)
}

func (c *Client) SetUserTeamResourcePermissions(ident ResourceIdent, userID int64, permission string) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissionByAssignment(
		"teams",
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
