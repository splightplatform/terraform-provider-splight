package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AssetParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Asset struct {
	AssetParams
	ID string `json:"id"`
}

func (c *Client) ListAssets() (*map[string]Asset, error) {
	body, err := c.httpRequest("v2/engine/assets", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Asset{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateAsset(item *AssetParams) (*Asset, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/assets/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &Asset{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateAsset(item *AssetParams) (*Asset, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/assets/%s", item.Name), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &Asset{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveAsset(id string) (*Asset, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/assets/%v", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &Asset{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteAsset(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/assets/%s", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
