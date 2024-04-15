package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AssetAttributeParams struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
	Asset string `json:"asset"`
}

type AssetAttribute struct {
	AssetAttributeParams
	Id string `json:"id"`
}

func (c *Client) ListAssetAttributes() (*map[string]AssetAttribute, error) {
	body, err := c.httpRequest("v2/engine/asset/attributes/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]AssetAttribute{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateAssetAttribute(item *AssetAttributeParams) (*AssetAttribute, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/asset/attributes/", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &AssetAttribute{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateAssetAttribute(id string, item *AssetAttributeParams) (*AssetAttribute, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/attributes/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &AssetAttribute{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveAssetAttribute(id string) (*AssetAttribute, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/attributes/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &AssetAttribute{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteAssetAttribute(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/attributes/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
