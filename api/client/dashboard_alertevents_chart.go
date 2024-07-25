package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DashboardAlertEventsChartParams struct {
	DashboardChartParams
	Type            string   `json:"type"`
	FilterName      string   `json:"filter_name,omitempty"`
	FilterOldStatus []string `json:"filter_old_status"`
	FilterNewStatus []string `json:"filter_new_status"`
}

type DashboardAlertEventsChart struct {
	DashboardAlertEventsChartParams
	ID string `json:"id"`
}

func (c *Client) ListDashboardAlertEventsCharts() (*map[string]DashboardAlertEventsChart, error) {
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]DashboardAlertEventsChart{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateDashboardAlertEventsChart(item *DashboardAlertEventsChartParams) (*DashboardAlertEventsChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/dashboard/charts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	asset := &DashboardAlertEventsChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) UpdateDashboardAlertEventsChart(id string, item *DashboardAlertEventsChartParams) (*DashboardAlertEventsChart, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	asset := &DashboardAlertEventsChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) RetrieveDashboardAlertEventsChart(id string) (*DashboardAlertEventsChart, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	asset := &DashboardAlertEventsChart{}
	err = json.NewDecoder(body).Decode(asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *Client) DeleteDashboardAlertEventsChart(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/dashboard/charts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
