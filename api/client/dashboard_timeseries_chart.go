package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardTimeseriesChartParams struct {
	DashboardChartParams
	Type                   string `json:"type"`
	YAxisMaxLimit          int    `json:"y_axis_max_limit,omitempty"`
	YAxisMinLimit          int    `json:"y_axis_min_limit,omitempty"`
	YAxisUnit              string `json:"y_axis_unit,omitempty"`
	NumberOfDecimals       int    `json:"number_of_decimals,omitempty"`
	XAxisFormat            string `json:"x_axis_format"`
	XAxisAutoSkip          bool   `json:"x_axis_auto_skip"`
	XAxisMaxTicksLimit     int    `json:"x_axis_max_ticks_limit,omitempty"`
	LineInterpolationStyle string `json:"line_interpolation_style,omitempty"`
	TimeseriesType         string `json:"timeseries_type,omitempty"`
	Fill                   bool   `json:"fill"`
	ShowLine               bool   `json:"show_line"`
}

type DashboardTimeseriesChart struct {
	DashboardTimeseriesChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardTimeseriesCharts() (*map[string]DashboardTimeseriesChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardTimeseriesChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardTimeseriesChart(item *DashboardTimeseriesChartParams) (*DashboardTimeseriesChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardTimeseriesChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardTimeseriesChart(id string, item *DashboardTimeseriesChartParams) (*DashboardTimeseriesChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardTimeseriesChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardTimeseriesChart(id string) (*DashboardTimeseriesChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardTimeseriesChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardTimeseriesChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
