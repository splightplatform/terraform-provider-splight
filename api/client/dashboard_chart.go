package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardValueMapping struct {
	Type        string `json:"type"`
	Order       int    `json:"order"`
	DisplayText string `json:"display_text"`
	MatchValue  string `json:"match_value"`
}

type DashboardThreshold struct {
	Value       float64 `json:"value"`
	Color       string  `json:"color"`
	DisplayText string  `json:"display_text"`
}

type DashboardChartItem struct {
	Color              string `json:"color"`
	RefID              string `json:"ref_id"`
	Type               string `json:"type"`
	Label              string `json:"label"`
	QueryGroupUnit     string `json:"query_group_unit"`
	QueryGroupFunction string `json:"query_group_function"`
	ExpressionPlain    string `json:"expression_plain"`
	QueryPlain         string `json:"query_plain"`
	QuerySortDirection int    `json:"query_sort_direction"`
	QueryLimit         int    `json:"query_limit"`
}

type DashboardChartParams struct {
	Name          string                  `json:"name"`
	Tab           string                  `json:"tab"`
	Type          string                  `json:"type"`
	TimestampGTE  string                  `json:"timestamp_gte"`
	TimestampLTE  string                  `json:"timestamp_lte"`
	Height        int                     `json:"height"`
	Width         int                     `json:"width"`
	ChartItems    []DashboardChartItem    `json:"chart_items"`
	Thresholds    []DashboardThreshold    `json:"thresholds"`
	ValueMappings []DashboardValueMapping `json:"value_mappings"`
}

type DashboardChart struct {
	DashboardChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardCharts() (*map[string]DashboardChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardChart(item *DashboardChartParams) (*DashboardChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardChart(id string, item *DashboardChartParams) (*DashboardChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardChart(id string) (*DashboardChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
