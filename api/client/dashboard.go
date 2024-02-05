package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Dashboard struct {
	DashboardParams
	ID string `json:"id"`
}

func (c *Client) ListDashboards() (*map[string]Dashboard, error) {
	body, err := c.httpRequest("v2/engine/dashboard/dashboards/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Dashboard{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboard(item *DashboardParams) (*Dashboard, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/dashboards/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &Dashboard{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboard(id string, item *DashboardParams) (*Dashboard, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/dashboards/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &Dashboard{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboard(id string) (*Dashboard, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/dashboards/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &Dashboard{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboard(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/dashboards/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
