package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AlertItem struct {
	Id                   string       `json:"id,omitempty"`
	RefId                string       `json:"ref_id"`
	Type                 string       `json:"type"`
	Expression           string       `json:"expression"`
	ExpressionPlain      string       `json:"expression_plain"`
	QueryPlain           string       `json:"query_plain"`
	QueryFilterAsset     *QueryFilter `json:"query_filter_asset"`
	QueryFilterAttribute *QueryFilter `json:"query_filter_attribute"`
	QueryGroupFunction   string       `json:"query_group_function"`
	QueryGroupUnit       string       `json:"query_group_unit"`
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
	Tags           []QueryFilter    `json:"tags"`
	AlertItems     []AlertItem      `json:"alert_items"`
	RelatedAssets  []QueryFilter    `json:"assets"`
}

type Alert struct {
	AlertParams
	Id string `json:"id"`
}

func (m *Alert) GetId() string {
	return m.Id
}

func (m *Alert) GetParams() Params {
	return &m.AlertParams
}

func (m *Alert) ResourcePath() string {
	return "v2/engine/alert/alerts/"
}

func (m *Alert) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	// Convert alert items
	alertItems := convertAlertItems(d.Get("alert_items").([]interface{}))

	// Convert alert thresholds
	alertThresholds := convertAlertThresholds(d.Get("thresholds").([]interface{}))

	// Convert tags
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Convert related assets
	assets := convertQueryFilters(d.Get("related_assets").(*schema.Set).List())

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
		Tags:           tags,
		AlertItems:     alertItems,
		RelatedAssets:  assets,
	}

	return nil
}

func convertAlertItems(alertItemsInterface []interface{}) []AlertItem {
	alertItems := make([]AlertItem, len(alertItemsInterface))
	for i, item := range alertItemsInterface {
		alertItem := item.(map[string]interface{})
		queryFilterAsset := convertSingleQueryFilter(alertItem["query_filter_asset"].(*schema.Set).List())
		queryFilterAttribute := convertSingleQueryFilter(alertItem["query_filter_attribute"].(*schema.Set).List())

		if queryFilterAttribute.isEmpty() {
			queryFilterAsset = nil
		}
		if queryFilterAttribute.isEmpty() {
			queryFilterAttribute = nil
		}
		queryGroupFunction := alertItem["query_group_function"].(string)
		queryGroupUnit := alertItem["query_group_unit"].(string)
		alertItems[i] = AlertItem{
			RefId:                alertItem["ref_id"].(string),
			Type:                 alertItem["type"].(string),
			Expression:           alertItem["expression"].(string),
			ExpressionPlain:      alertItem["expression_plain"].(string),
			QueryPlain:           alertItem["query_plain"].(string),
			QueryFilterAsset:     queryFilterAsset,
			QueryFilterAttribute: queryFilterAttribute,
			QueryGroupFunction:   queryGroupFunction,
			QueryGroupUnit:       queryGroupUnit,
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

func (m *Alert) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

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

	thresholds := make([]map[string]interface{}, len(m.Thresholds))
	for i, m := range m.Thresholds {
		thresholds[i] = map[string]interface{}{
			"value":       m.Value,
			"status":      m.Status,
			"status_text": m.StatusText,
		}
	}
	d.Set("thresholds", thresholds)

	var tags []map[string]any
	for _, tag := range m.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	// Query filters are always set
	alertItems := make([]map[string]interface{}, len(m.AlertItems))
	for i, alert := range m.AlertItems {
		var queryFilterAsset []map[string]string
		var queryFilterAttribute []map[string]string

		// Set to empty map in case of nil, since thats how
		// we allow it in the schema
		if alert.QueryFilterAsset != nil {
			queryFilterAsset = alert.QueryFilterAsset.toMap()
		} else {
			queryFilterAsset = (&QueryFilter{}).toMap()
		}

		if alert.QueryFilterAttribute != nil {
			queryFilterAttribute = alert.QueryFilterAttribute.toMap()
		} else {
			queryFilterAttribute = (&QueryFilter{}).toMap()
		}
		alertItems[i] = map[string]interface{}{
			"id":                     alert.Id,
			"ref_id":                 alert.RefId,
			"type":                   alert.Type,
			"expression":             alert.Expression,
			"expression_plain":       alert.ExpressionPlain,
			"query_plain":            alert.QueryPlain,
			"query_filter_asset":     queryFilterAsset,
			"query_filter_attribute": queryFilterAttribute,
			"query_group_function":   alert.QueryGroupFunction,
			"query_group_unit":       alert.QueryGroupUnit,
		}
	}
	d.Set("alert_items", alertItems)

	var relatedasets []map[string]any
	for _, relatedAsset := range m.RelatedAssets {
		relatedasets = append(relatedasets, map[string]any{
			"id":   relatedAsset.Id,
			"name": relatedAsset.Name,
		})
	}
	d.Set("related_assets", relatedasets)

	return nil
}
