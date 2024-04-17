package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AssetMetadataParams struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Unit  string      `json:"unit"`
	Asset string      `json:"asset"`
	Value interface{} `json:"value"`
}

type AssetMetadata struct {
	Id string `json:"id"`
	AssetMetadataParams
}

func (c *Client) ListAssetMetadatas() (*map[string]AssetMetadata, error) {
	body, err := c.httpRequest("v2/engine/asset/metadata/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]AssetMetadata{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateAssetMetadata(item *AssetMetadataParams) (*AssetMetadata, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/asset/metadata/", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &AssetMetadata{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateAssetMetadata(id string, item *AssetMetadataParams) (*AssetMetadata, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/metadata/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &AssetMetadata{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveAssetMetadata(id string) (*AssetMetadata, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/metadata/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &AssetMetadata{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteAssetMetadata(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/metadata/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
