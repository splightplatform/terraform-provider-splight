package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
)

type RelatedAsset struct {
	Id string `json:"id"`
}

type Asset struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	RelatedAssets []RelatedAsset `json:"assets"`
	Geometry      string         `json:"geometry"`
}

func (a *Asset) UnmarshalJSON(data []byte) error {
	// Create a type to avoid infinite recursion
	type Alias Asset
	aux := &struct {
		Geometry json.RawMessage `json:"geometry"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	// Unmarshal everything except for Geometry
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Store Geometry as a string
	a.Geometry = string(aux.Geometry)

	return nil
}

func (a Asset) MarshalJSON() ([]byte, error) {

	// Create an alias to avoid infinite recursion
	type Alias Asset

	return json.Marshal(&struct {
		Geometry string `json:"geometry"`
		Alias
	}{
		Geometry: a.Geometry,
		Alias:    (Alias)(a),
	})
}

func (c *Client) ListAssets() (*map[string]Asset, error) {
	body, err := c.httpRequest("v2/engine/asset/assets/", "GET", bytes.Buffer{})
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

func (c *Client) CreateAsset(item *Asset) (*Asset, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	runtime.Breakpoint()
	body, err := c.httpRequest("v2/engine/asset/assets/", "POST", buf)
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

func (c *Client) UpdateAsset(id string, item *Asset) (*Asset, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/assets/%s/", id), "PUT", buf)
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

func (c *Client) RetrieveAsset(id string) (*Asset, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/assets/%s/", id), "GET", bytes.Buffer{})
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
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/assets/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
