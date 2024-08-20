package client

import (
	"bytes"
	"encoding/json"
	"runtime"
)

type AssetKind struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (c *Client) ListAssetKinds() ([]AssetKind, error) {
	body, err := c.httpRequest("v2/engine/asset/kinds/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := []AssetKind{}
	// TODO: revert back to the original code that read from "results"
	runtime.Breakpoint()
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
