package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AlertItem struct {
	ID                   string          `json:"id,omitempty"`
	RefID                string          `json:"ref_id"`
	Type                 string          `json:"type"`
	Expression           string          `json:"expression"`
	ExpressionPlain      string          `json:"expression_plain"`
	QueryPlain           string          `json:"query_plain"`
	QueryFilterAsset     AlertTargetItem `json:"query_filter_asset"`
	QueryFilterAttribute AlertTargetItem `json:"query_filter_attribute"`
	QueryGroupFunction   string          `json:"query_group_function"`
	QueryGroupUnit       string          `json:"query_group_unit"`
}

type AlertTargetItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Implement custom JSON marshalling to omit the struct if both fields are empty
func (ati AlertTargetItem) MarshalJSON() ([]byte, error) {
	if ati.ID == "" && ati.Name == "" {
		return []byte("null"), nil
	}
	type Alias AlertTargetItem
	return json.Marshal((Alias)(ati))
}

// Implement custom JSON unmarshalling to initialize the struct if the field is null
func (ati *AlertTargetItem) UnmarshalJSON(data []byte) error {
	type Alias AlertTargetItem
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(ati),
	}

	if string(data) == "null" {
		*ati = AlertTargetItem{}
		return nil
	}

	return json.Unmarshal(data, &aux)
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
	RelatedAssets  []RelatedAsset   `json:"assets"`
}

type Alert struct {
	AlertParams
	ID string `json:"id"`
}

func (m *Alert) GetID() string {
	return m.ID
}

func (m *Alert) GetParams() Params {
	return &m.AlertParams
}

func (m *Alert) ResourcePath() string {
	return "v2/engine/alert/alerts/"
}

func (m *Alert) FromSchema(d *schema.ResourceData) error {
	// Convert alert items
	alertItems := convertAlertItems(d.Get("alert_items").([]interface{}))

	// Convert alert thresholds
	alertThresholds := convertAlertThresholds(d.Get("thresholds").([]interface{}))

	// Convert related assets
	alertRelatedAssets := convertRelatedAssets(d.Get("related_assets").(*schema.Set).List())

	// Create the AlertParams object
	m.AlertParams = AlertParams{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Type:           d.Get("type").(string),
		TimeWindow:     d.Get("time_window").(int),
		RateUnit:       d.Get("rate_unit").(string),
		RateValue:      d.Get("rate_value").(int),
		Severity:       d.Get("severity").(string),
		Operator:       d.Get("operator").(string),
		Aggregation:    d.Get("aggregation").(string),
		Thresholds:     alertThresholds,
		TargetVariable: d.Get("target_variable").(string),
		AlertItems:     alertItems,
		RelatedAssets:  alertRelatedAssets,
	}

	return nil
}

func convertAlertItems(alertItemsInterface []interface{}) []AlertItem {
	alertItems := make([]AlertItem, len(alertItemsInterface))
	for i, item := range alertItemsInterface {
		alertItem := item.(map[string]interface{})
		queryFilterAsset := alertItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
		queryFilterAttribute := alertItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
		queryGroupFunction := alertItem["query_group_function"].(string)
		queryGroupUnit := alertItem["query_group_unit"].(string)
		alertItems[i] = AlertItem{
			RefID:           alertItem["ref_id"].(string),
			Type:            alertItem["type"].(string),
			Expression:      alertItem["expression"].(string),
			ExpressionPlain: alertItem["expression_plain"].(string),
			QueryPlain:      alertItem["query_plain"].(string),
			QueryFilterAsset: AlertTargetItem{
				Name: queryFilterAsset["name"].(string),
				ID:   queryFilterAsset["id"].(string),
			},
			QueryFilterAttribute: AlertTargetItem{
				Name: queryFilterAttribute["name"].(string),
				ID:   queryFilterAttribute["id"].(string),
			},
			QueryGroupFunction: queryGroupFunction,
			QueryGroupUnit:     queryGroupUnit,
		}
	}
	return alertItems
}

func convertAlertThresholds(thresholdsInterface []interface{}) []AlertThreshold {
	alertThresholds := make([]AlertThreshold, len(thresholdsInterface))
	for i, item := range thresholdsInterface {
		threshold := item.(map[string]interface{})
		alertThresholds[i] = AlertThreshold{
			Value:      threshold["value"].(float64),
			Status:     threshold["status"].(string),
			StatusText: threshold["status_text"].(string),
		}
	}
	return alertThresholds
}

func convertRelatedAssets(relatedAssetsInterface []interface{}) []RelatedAsset {
	alertRelatedAssets := make([]RelatedAsset, len(relatedAssetsInterface))
	for i, item := range relatedAssetsInterface {
		alertRelatedAssets[i] = RelatedAsset{
			Id: item.(string),
		}
	}
	return alertRelatedAssets
}

func (m *Alert) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("type", m.Type)
	d.Set("time_window", m.TimeWindow)
	d.Set("operator", m.Operator)
	d.Set("aggregation", m.Aggregation)
	d.Set("rate_unit", m.RateUnit)
	d.Set("rate_value", m.RateValue)
	d.Set("cron_minutes", m.CronMinutes)
	d.Set("cron_hours", m.CronHours)
	d.Set("cron_dom", m.CronDOM)
	d.Set("cron_month", m.CronMonth)
	d.Set("cron_dow", m.CronDOW)
	d.Set("cron_year", m.CronYear)
	d.Set("severity", m.Severity)
	d.Set("related_assets", m.RelatedAssets)

	thresholds := make([]map[string]interface{}, len(m.Thresholds))
	for i, m := range m.Thresholds {
		thresholds[i] = map[string]interface{}{
			"value":       m.Value,
			"status":      m.Status,
			"status_text": m.StatusText,
		}
	}
	d.Set("thresholds", thresholds)

	alertItems := make([]map[string]interface{}, len(m.AlertItems))
	for i, alert := range m.AlertItems {
		alertItems[i] = map[string]interface{}{
			"id":               alert.ID,
			"ref_id":           alert.RefID,
			"type":             alert.Type,
			"expression":       alert.Expression,
			"expression_plain": alert.ExpressionPlain,
			"query_plain":      alert.QueryPlain,
			"query_filter_asset": []map[string]interface{}{
				{
					"id":   alert.QueryFilterAsset.ID,
					"name": alert.QueryFilterAsset.Name,
				},
			},
			"query_filter_attribute": []map[string]interface{}{
				{
					"id":   alert.QueryFilterAttribute.ID,
					"name": alert.QueryFilterAttribute.Name,
				},
			},
			"query_group_function": alert.QueryGroupFunction,
			"query_group_unit":     alert.QueryGroupUnit,
		}
	}
	d.Set("alert_items", alertItems)

	return nil
}
