package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Setpoint struct {
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name"`
	Value     json.RawMessage `json:"value"`
	Attribute QueryFilter     `json:"attribute"`
}

type ActionParams struct {
	Asset     QueryFilter `json:"asset"`
	Name      string      `json:"name"`
	Setpoints []Setpoint  `json:"setpoints"`
}

type Action struct {
	ActionParams
	ID string `json:"id"`
}

func (c *Client) ListActions() (*map[string]Action, error) {
	body, err := c.httpRequest("v2/engine/asset/actions/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Action{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateAction(item *ActionParams) (*Action, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/asset/actions/", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &Action{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateAction(id string, item *ActionParams) (*Action, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/actions/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &Action{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveAction(id string) (*Action, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/actions/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &Action{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteAction(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/actions/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
