package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type TagParams struct {
	Name string `json:"name"`
}

type Tag struct {
	TagParams
	ID string `json:"id"`
}

func (c *Client) ListTags() ([]Tag, error) {
	body, err := c.httpRequest("v2/account/tags", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := []Tag{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) CreateTag(item *TagParams) (*Tag, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/account/tags", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &Tag{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateTag(id string, item *TagParams) (*Tag, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/account/tags/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &Tag{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveTag(id string) (*Tag, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/account/tags/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &Tag{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteTag(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/account/tag/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
