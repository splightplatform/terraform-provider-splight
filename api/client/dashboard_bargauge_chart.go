package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardBarGaugeChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	MaxLimit         int    `json:"max_limit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
	Orientation      string `json:"orientation,omitempty"`
}

type DashboardBarGaugeChart struct {
	DashboardBarGaugeChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardBarGaugeCharts() (*map[string]DashboardBarGaugeChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardBarGaugeChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardBarGaugeChart(item *DashboardBarGaugeChartParams) (*DashboardBarGaugeChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardBarGaugeChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardBarGaugeChart(id string, item *DashboardBarGaugeChartParams) (*DashboardBarGaugeChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardBarGaugeChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardBarGaugeChart(id string) (*DashboardBarGaugeChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardBarGaugeChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardBarGaugeChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
