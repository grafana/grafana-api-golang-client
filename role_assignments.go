package gapi

type RoleAssignments struct {
	UID            string               `json:"uid"`
	Users          []UserRoleAssignment `json:"users,omitempty"`
	Teams          []int                `json:"teams,omitempty"`
	ServiceAccount []int                `json:"serviceAccounts,omitempty"`
}

type UserRoleAssignment struct {
	ID     int  `json:"id,omitempty"`
	Global bool `json:"global"`
}

func (c *Client) GetRoleAssignments(uid string) (*RoleAssignments, error) {
	// implement me
	return &RoleAssignments{}, nil
}

func (c *Client) UpdateRoleAssignments(ra RoleAssignments) error {
	// implement me
	return nil
}
