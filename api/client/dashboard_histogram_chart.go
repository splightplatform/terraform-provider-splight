package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardHistogramChartParams struct {
	DashboardChartParams
	Type                  string `json:"type"`
	NumberOfDecimals      int    `json:"number_of_decimals,omitempty"`
	BucketCount           int    `json:"bucket_count"`
	BucketSize            int    `json:"bucket_size,omitempty"`
	HistogramType         string `json:"histogram_type"`
	Sorting               string `json:"sorting"`
	Stacked               bool   `json:"stacked"`
	CategoriesTopMaxLimit int    `json:"categories_top_max_limit,omitempty"`
}

type DashboardHistogramChart struct {
	DashboardHistogramChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardHistogramCharts() (*map[string]DashboardHistogramChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardHistogramChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardHistogramChart(item *DashboardHistogramChartParams) (*DashboardHistogramChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardHistogramChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardHistogramChart(id string, item *DashboardHistogramChartParams) (*DashboardHistogramChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardHistogramChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardHistogramChart(id string) (*DashboardHistogramChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardHistogramChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardHistogramChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
