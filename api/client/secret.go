package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type SecretParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Secret struct {
	SecretParams
	ID string `json:"id"`
}

func (c *Client) ListSecrets() (*map[string]Secret, error) {
	body, err := c.httpRequest("v2/engine/secret/secrets/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Secret{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateSecret(item *SecretParams) (*Secret, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/secret/secrets/", "POST", buf)
	if err != nil {
		return nil, err
	}

	secret := &Secret{}
	err = json.NewDecoder(body).Decode(secret)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (c *Client) UpdateSecret(id string, item *SecretParams) (*Secret, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/secret/secrets/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	secret := &Secret{}
	err = json.NewDecoder(body).Decode(secret)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (c *Client) RetrieveSecret(id string) (*Secret, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/secret/secrets/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	secret := &Secret{}
	err = json.NewDecoder(body).Decode(secret)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (c *Client) DeleteSecret(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/secret/secrets/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
