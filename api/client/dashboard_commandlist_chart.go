package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardCommandListChartParams struct {
	DashboardChartParams
	Type            string `json:"type"`
	CommandListType string `json:"command_list_type"`
	FilterName      string `json:"filter_name"`
}

type DashboardCommandListChart struct {
	DashboardCommandListChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardCommandListCharts() (*map[string]DashboardCommandListChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardCommandListChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardCommandListChart(item *DashboardCommandListChartParams) (*DashboardCommandListChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardCommandListChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardCommandListChart(id string, item *DashboardCommandListChartParams) (*DashboardCommandListChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardCommandListChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardCommandListChart(id string) (*DashboardCommandListChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardCommandListChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardCommandListChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
