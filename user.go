package scrive

import "fmt"

func (c *Client) LoginAndGetSession(p OAuthAuthorization) (*Session, *ScriveError) {
	sess := &Session{}
	_, se := c.pw(
		"loginandgetsession",
		func(req *request) {
			req.writeMJSON("personal_token", p)
		},
		sess,
	)
	return sess, se
}

func (c *Client) GetTokenForPersonalCredentials(userID string, minutes *uint64) (*PersonalCredentialsToken, *ScriveError) {
	tok := &PersonalCredentialsToken{}
	_, se := c.pw(
		fmt.Sprintf("gettokenforpersonalcredentials/%s", userID),
		func(req *request) {
			req.writeMUInt("minutes", minutes)
		},
		tok,
	)
	return tok, se
}

func (c *Client) Setup2FA() (*Setup2FAResponse, *ScriveError) {
	resp := &Setup2FAResponse{}
	_, se := c.pw(
		"2fa/setup",
		nil,
		resp,
	)
	return resp, se
}

func (c *Client) Confirm2FA(totp string) (*Confirm2FAResp, *ScriveError) {
	resp := &Confirm2FAResp{}
	_, se := c.pw(
		"2fa/confirm",
		func(req *request) {
			req.writeMString("totp", &totp)
		},
		resp,
	)
	return resp, se
}

func (c *Client) Disable2FA() (*Disable2FAResp, *ScriveError) {
	resp := &Disable2FAResp{}
	_, se := c.pw(
		"2fa/disable",
		nil,
		resp,
	)
	return resp, se
}

func (c *Client) IsUserDeletable() (*IsUserDeletableResp, *ScriveError) {
	resp := &IsUserDeletableResp{}
	_, se := c.pw(
		"isuserdeletable",
		nil,
		resp,
	)
	return resp, se
}

func (c *Client) DeleteUser(email string) *ScriveError {
	_, se := c.pwb(
		"deleteuser",
		func(req *request) {
			req.writeMString("email", &email)
		},
	)
	return se
}

func (c *Client) GetDataRetentionPolicy() (*DataRetentionPolicy, *ScriveError) {
	resp := &DataRetentionPolicy{}
	_, se := c.gw(
		"dataretentionpolicy",
		nil,
		resp,
	)
	return resp, se
}

func (c *Client) SetDataRetentionPolicy(p DataRetentionPolicy) *ScriveError {
	_, se := c.pwb(
		"dataretentionpolicy/set",
		func(req *request) {
			req.writeMJSON("data_retention_policy", p)
		},
	)
	return se
}
