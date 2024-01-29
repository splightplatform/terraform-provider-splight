package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MetadataParams struct {
	Asset string  `json:"asset"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Value string  `json:"value"`
	Unit  *string `json:"unit"`
}

type Metadata struct {
	MetadataParams
	ID string `json:"id"`
}

func (c *Client) ListMetadatas() (*map[string]Metadata, error) {
	body, err := c.httpRequest("v2/engine/asset/metadata/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Metadata{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateMetadata(item *MetadataParams) (*Metadata, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/asset/metadata/", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &Metadata{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateMetadata(id string, item *MetadataParams) (*Metadata, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/metadata/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &Metadata{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveMetadata(id string) (*Metadata, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/metadata/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &Metadata{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteMetadata(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/asset/metadata/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
