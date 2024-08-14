package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) RetrieveOrgId() (string, error) {
	body, err := c.httpRequest("v2/account/user/organizations/", "GET", bytes.Buffer{})
	if err != nil {
		return "", err
	}
	orgs := map[string]interface{}{}
	err = json.NewDecoder(body).Decode(&orgs)
	if err != nil {
		return "", err
	}
	if len(orgs) == 0 {
		return "", fmt.Errorf("No organizations found")
	}
	orgId := orgs["id"].(string)
	return orgId, nil
}

func (c *Client) RetrieveEmail() (string, error) {
	body, err := c.httpRequest("v2/account/user/profile/", "GET", bytes.Buffer{})
	if err != nil {
		return "", err
	}
	profile := map[string]interface{}{}
	err = json.NewDecoder(body).Decode(&profile)
	if err != nil {
		return "", err
	}
	if email, ok := profile["email"].(string); ok {
		return email, nil
	}
	return "", fmt.Errorf("User email not found")
}
