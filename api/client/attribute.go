package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AttributeParams struct {
	Asset string  `json:"asset"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Unit  *string `json:"unit"`
}

type Attribute struct {
	AttributeParams
	ID string `json:"id"`
}

func (c *Client) ListAttributes() (*map[string]Attribute, error) {
	body, err := c.httpRequest("v2/engine/asset/attributes/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Attribute{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateAttribute(item *AttributeParams) (*Attribute, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/asset/attributes/", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &Attribute{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateAttribute(id string, item *AttributeParams) (*Attribute, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/attributes/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &Attribute{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveAttribute(id string) (*Attribute, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/attributes/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &Attribute{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteAttribute(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/attributes/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
