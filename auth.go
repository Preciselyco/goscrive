package scrive

import "fmt"

//PAC represents Personal Access Credentials.
//Found in Scrive dashboard under "Integration Settings"
type PAC struct {
	ClientCredentialsIdentifier string
	ClientCredentialsSecret     string
	TokenCredentialsIdentifier  string
	TokenCredentialsSecret      string
}

func (c *Client) constructAuthHeaderPAC() string {
	pac := c.config.PAC
	return fmt.Sprintf("oauth_signature_method=\"PLAINTEXT\", oauth_consumer_key=\"%s\", oauth_token=\"%s\", oauth_signature=\"%s&%s\"",
		pac.ClientCredentialsIdentifier,
		pac.TokenCredentialsIdentifier,
		pac.ClientCredentialsSecret,
		pac.TokenCredentialsSecret)
}

type LoginToken struct {
	LoginToken     string `json:"login_token"`
	QRCode         string `json:"qr_code"`
	ExpirationTime string `json:"expiration_time"`
}
