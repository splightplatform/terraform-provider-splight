package client

import (
	"bytes"
	"encoding/json"
)

type AssetKind struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// API returns the kinds inside the 'results' key
type AssetKinds struct {
	Kinds []AssetKind `json:"results"`
}

func (c *Client) ListAssetKinds() (*[]AssetKind, error) {
	body, err := c.httpRequest("v2/engine/asset/kinds/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := AssetKinds{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items.Kinds, nil
}
