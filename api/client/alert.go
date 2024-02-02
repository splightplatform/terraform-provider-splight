package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AlertItem struct {
	// ID              string `json:"id"`
	RefID           string `json:"ref_id"`
	Type            string `json:"type"`
	ExpressionPlain string `json:"expression_plain"`
	QueryPlain      string `json:"query_plain"`
}

type AlertTargetItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AlertThreshold struct {
	Value      float64 `json:"value"`
	Status     string  `json:"status"`
	StatusText string  `json:"status_text"`
}

type AlertParams struct {
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Type           string           `json:"type"`
	Severity       string           `json:"severity"`
	TimeWindow     int              `json:"stmt_time_window"`
	Operator       string           `json:"stmt_operator"`
	Aggregation    string           `json:"stmt_aggregation"`
	Thresholds     []AlertThreshold `json:"stmt_thresholds"`
	TargetVariable string           `json:"stmt_target_variable"`
	RateUnit       string           `json:"rate_unit"`
	RateValue      int              `json:"rate_value"`
	CronMinutes    int              `json:"cron_minutes"`
	CronHours      int              `json:"cron_hours"`
	CronDOM        int              `json:"cron_dom"`
	CronMonth      int              `json:"cron_month"`
	CronDOW        int              `json:"cron_dow"`
	CronYear       int              `json:"cron_year"`
	AlertItems     []AlertItem      `json:"alert_items"`
}

type Alert struct {
	AlertParams
	ID string `json:"id"`
}

func (c *Client) ListAlerts() (*map[string]Alert, error) {
	body, err := c.httpRequest("v2/engine/alert/alerts/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Alert{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateAlert(item *AlertParams) (*Alert, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/alert/alerts/", "POST", buf)
	if err != nil {
		return nil, err
	}

	function := &Alert{}
	err = json.NewDecoder(body).Decode(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (c *Client) UpdateAlert(id string, item *AlertParams) (*Alert, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/alert/alerts/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	function := &Alert{}
	err = json.NewDecoder(body).Decode(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (c *Client) RetrieveAlert(id string) (*Alert, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/alert/alerts/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	function := &Alert{}
	err = json.NewDecoder(body).Decode(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (c *Client) DeleteAlert(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/alert/alerts/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
