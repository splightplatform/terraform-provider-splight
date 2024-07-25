package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardStatChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit,omitempty"`
	Border           bool   `json:"border"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
}

type DashboardStatChart struct {
	DashboardStatChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardStatCharts() (*map[string]DashboardStatChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardStatChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardStatChart(item *DashboardStatChartParams) (*DashboardStatChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardStatChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardStatChart(id string, item *DashboardStatChartParams) (*DashboardStatChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardStatChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardStatChart(id string) (*DashboardStatChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardStatChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardStatChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
