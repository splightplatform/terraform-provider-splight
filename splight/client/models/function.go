package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FunctionItem struct {
	Id                   string            `json:"id,omitempty"`
	RefId                string            `json:"ref_id"`
	Type                 string            `json:"type"`
	Expression           string            `json:"expression"`
	ExpressionPlain      string            `json:"expression_plain"`
	QueryPlain           string            `json:"query_plain"`
	QueryFilterAsset     *QueryFilter      `json:"query_filter_asset"`
	QueryFilterAttribute *TypedQueryFilter `json:"query_filter_attribute"`
	QueryGroupFunction   string            `json:"query_group_function"`
	QueryGroupUnit       string            `json:"query_group_unit"`
}

type FunctionParams struct {
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Type            string           `json:"type"`
	TimeWindow      int              `json:"time_window"`
	TargetAsset     QueryFilter      `json:"target_asset"`
	TargetAttribute TypedQueryFilter `json:"target_attribute"`
	TargetVariable  string           `json:"target_variable"`
	RateUnit        string           `json:"rate_unit"`
	RateValue       int              `json:"rate_value"`
	CronMinutes     int              `json:"cron_minutes"`
	CronHours       int              `json:"cron_hours"`
	CronDOM         int              `json:"cron_dom"`
	CronMonth       int              `json:"cron_month"`
	CronDOW         int              `json:"cron_dow"`
	CronYear        int              `json:"cron_year"`
	FunctionItems   []FunctionItem   `json:"function_items"`
	Tags            []QueryFilter    `json:"tags"`
}

type Function struct {
	FunctionParams
	Id string `json:"id"`
}

func (m *Function) GetId() string {
	return m.Id
}

func (m *Function) GetParams() Params {
	return &m.FunctionParams
}

func (m *Function) ResourcePath() string {
	return "v2/engine/function/functions/"
}

func (m *Function) FromSchema(d *schema.ResourceData) error {
	targetAsset := convertSingleQueryFilter(d.Get("target_asset").(*schema.Set).List())
	targetAttribute := convertSingleTypedQueryFilter(d.Get("target_attribute").(*schema.Set).List())

	// Convert function items
	functionItems := convertFunctionItems(d.Get("function_items").([]any))

	// Convert tags
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Create the FunctionParams object
	m.FunctionParams = FunctionParams{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Type:            d.Get("type").(string),
		TimeWindow:      d.Get("time_window").(int),
		RateUnit:        d.Get("rate_unit").(string),
		RateValue:       d.Get("rate_value").(int),
		TargetVariable:  d.Get("target_variable").(string),
		TargetAsset:     *targetAsset,
		TargetAttribute: *targetAttribute,
		Tags:            tags,
		FunctionItems:   functionItems,
	}

	return nil
}

func convertFunctionItems(functionItemsInterface []any) []FunctionItem {
	functionItems := make([]FunctionItem, len(functionItemsInterface))
	for i, item := range functionItemsInterface {
		functionItem := item.(map[string]any)

		queryFilterAsset := convertSingleQueryFilter(functionItem["query_filter_asset"].(*schema.Set).List())
		queryFilterAttribute := convertSingleTypedQueryFilter(functionItem["query_filter_attribute"].(*schema.Set).List())

		if queryFilterAsset.isEmpty() {
			queryFilterAsset = nil
		}
		if queryFilterAttribute.isEmpty() {
			queryFilterAttribute = nil
		}

		queryGroupFunction := functionItem["query_group_function"].(string)
		queryGroupUnit := functionItem["query_group_unit"].(string)
		functionItems[i] = FunctionItem{
			RefId:                functionItem["ref_id"].(string),
			Type:                 functionItem["type"].(string),
			Expression:           functionItem["expression"].(string),
			ExpressionPlain:      functionItem["expression_plain"].(string),
			QueryPlain:           functionItem["query_plain"].(string),
			QueryFilterAsset:     queryFilterAsset,
			QueryFilterAttribute: queryFilterAttribute,
			QueryGroupFunction:   queryGroupFunction,
			QueryGroupUnit:       queryGroupUnit,
		}
	}
	return functionItems
}

func (m *Function) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("type", m.Type)
	d.Set("time_window", m.TimeWindow)

	// Since the schemas for these params are 'TypeSet' we must convert
	// our structs to a slice of 'FunctionTargetItem'.
	// Otherwise the SDK will raise a type error.
	d.Set("target_asset", []map[string]any{
		{
			"id":   m.TargetAsset.Id,
			"name": m.TargetAsset.Name,
		},
	})
	d.Set("target_attribute", []map[string]any{
		{
			"id":   m.TargetAttribute.Id,
			"name": m.TargetAttribute.Name,
			"type": m.TargetAttribute.Type,
		},
	})

	d.Set("rate_unit", m.RateUnit)
	d.Set("rate_value", m.RateValue)
	d.Set("cron_minutes", m.CronMinutes)
	d.Set("cron_hours", m.CronHours)
	d.Set("cron_dom", m.CronDOM)
	d.Set("cron_month", m.CronMonth)
	d.Set("cron_dow", m.CronDOW)
	d.Set("cron_year", m.CronYear)
	d.Set("tags", m.Tags)

	functionItems := make([]map[string]any, len(m.FunctionItems))
	for i, function := range m.FunctionItems {
		var queryFilterAsset []map[string]string
		var queryFilterAttribute []map[string]string

		// Set to empty map in case of nil, since thats how
		// we allow it in the schema
		if function.QueryFilterAsset != nil {
			queryFilterAsset = function.QueryFilterAsset.toMap()
		} else {
			queryFilterAsset = (&QueryFilter{}).toMap()
		}

		if function.QueryFilterAttribute != nil {
			queryFilterAttribute = function.QueryFilterAttribute.toMap()
		} else {
			queryFilterAttribute = (&QueryFilter{}).toMap()
		}

		functionItems[i] = map[string]any{
			"id":                     function.Id,
			"ref_id":                 function.RefId,
			"type":                   function.Type,
			"expression":             function.Expression,
			"expression_plain":       function.ExpressionPlain,
			"query_plain":            function.QueryPlain,
			"query_filter_asset":     queryFilterAsset,
			"query_filter_attribute": queryFilterAttribute,
			"query_group_function":   function.QueryGroupFunction,
			"query_group_unit":       function.QueryGroupUnit,
		}
	}

	d.Set("function_items", functionItems)

	return nil
}
