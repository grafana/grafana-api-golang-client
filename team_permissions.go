package gapi

func (c *Client) ListTeamResourcePermissions(ident ResourceIdent) ([]*ResourcePermission, error) {
	return c.listResourcePermissions(TeamsResource, ident)
}

func (c *Client) SetTeamResourcePermissions(ident ResourceIdent, body SetResourcePermissionsBody) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissions(TeamsResource, ident, body)
}

func (c *Client) SetUserTeamResourcePermissions(ident ResourceIdent, userID int64, permission string) (*SetResourcePermissionsResponse, error) {
	return c.setResourcePermissionByAssignment(
		TeamsResource,
		ident,
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
