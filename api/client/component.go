package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ComponentInputParam struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Multiple    bool            `json:"multiple"`
	Required    bool            `json:"required"`
	Sensitive   bool            `json:"sensitive"`
	Type        string          `json:"type"`
	Value       json.RawMessage `json:"value"`
}

type ComponentParams struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Version     string                `json:"version"`
	Input       []ComponentInputParam `json:"input"`
}

type Component struct {
	ComponentParams
	ID string `json:"id"`
}

func (c *Client) ListComponents() (*map[string]Component, error) {
	body, err := c.httpRequest("v2/engine/component/components/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Component{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateComponent(item *ComponentParams) (*Component, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/component/components/", "POST", buf)
	if err != nil {
		return nil, err
	}

	component := &Component{}
	err = json.NewDecoder(body).Decode(component)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (c *Client) UpdateComponent(id string, item *ComponentParams) (*Component, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/component/components/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	component := &Component{}
	err = json.NewDecoder(body).Decode(component)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (c *Client) RetrieveComponent(id string) (*Component, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/component/components/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	component := &Component{}
	err = json.NewDecoder(body).Decode(component)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (c *Client) DeleteComponent(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/component/components/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
