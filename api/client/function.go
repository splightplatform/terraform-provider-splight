package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type FunctionItem struct {
	// ID              string `json:"id"`
	RefID           string `json:"ref_id"`
	Type            string `json:"type"`
	ExpressionPlain string `json:"expression_plain"`
	QueryPlain      string `json:"query_plain"`
}

type FunctionTargetItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FunctionParams struct {
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Type            string             `json:"type"`
	TimeWindow      int                `json:"time_window"`
	TargetAsset     FunctionTargetItem `json:"target_asset"`
	TargetAttribute FunctionTargetItem `json:"target_attribute"`
	TargetVariable  string             `json:"target_variable"`
	RateUnit        string             `json:"rate_unit"`
	RateValue       int                `json:"rate_value"`
	CronMinutes     int                `json:"cron_minutes"`
	CronHours       int                `json:"cron_hours"`
	CronDOM         int                `json:"cron_dom"`
	CronMonth       int                `json:"cron_month"`
	CronDOW         int                `json:"cron_dow"`
	CronYear        int                `json:"cron_year"`
	FunctionItems   []FunctionItem     `json:"function_items"`
}

type Function struct {
	FunctionParams
	ID string `json:"id"`
}

func (c *Client) ListFunctions() (*map[string]Function, error) {
	body, err := c.httpRequest("v2/engine/function/functions/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Function{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateFunction(item *FunctionParams) (*Function, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/function/functions/", "POST", buf)
	if err != nil {
		return nil, err
	}

	function := &Function{}
	err = json.NewDecoder(body).Decode(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (c *Client) UpdateFunction(id string, item *FunctionParams) (*Function, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/function/functions/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	function := &Function{}
	err = json.NewDecoder(body).Decode(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (c *Client) RetrieveFunction(id string) (*Function, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/function/functions/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	function := &Function{}
	err = json.NewDecoder(body).Decode(function)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (c *Client) DeleteFunction(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/function/functions/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
