package scrive

import "fmt"

func (c *Client) GetUserAccessRoles(userID string) (*[]*AccessRole, *ScriveError) {
	resp := &[]*AccessRole{}
	_, se := c.gw(
		fmt.Sprintf("getuserroles/%s", userID),
		nil,
		resp,
	)
	return resp, se
}

func (c *Client) GetAccessRole(roleID string) (*AccessRole, *ScriveError) {
	resp := &AccessRole{}
	_, se := c.gw(
		fmt.Sprintf("accesscontrol/roles/%s", roleID),
		nil,
		resp,
	)
	return resp, se
}
