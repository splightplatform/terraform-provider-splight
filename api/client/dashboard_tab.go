package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardTabParams struct {
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Dashboard string `json:"dashboard"`
}

type DashboardTab struct {
	DashboardTabParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardTabs() (*map[string]DashboardTab, error) {
	body, err := c.httpRequest("v2/engine/dashboard/tabs/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardTab{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardTab(item *DashboardTabParams) (*DashboardTab, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/tabs/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardTab{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardTab(id string, item *DashboardTabParams) (*DashboardTab, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/tabs/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardTab{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardTab(id string) (*DashboardTab, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/tabs/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardTab{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardTab(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/tabs/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
