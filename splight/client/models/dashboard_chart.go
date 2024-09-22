package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
	Color                string       `json:"color"`
	RefId                string       `json:"ref_id"`
	Type                 string       `json:"type"`
	Label                string       `json:"label"`
	Hidden               bool         `json:"hidden"`
	ExpressionPlain      string       `json:"expression_plain"`
	QueryFilterAsset     *QueryFilter `json:"query_filter_asset"`
	QueryFilterAttribute *QueryFilter `json:"query_filter_attribute"`
	QueryPlain           string       `json:"query_plain"`
	QueryGroupUnit       string       `json:"query_group_unit"`
	QueryGroupFunction   string       `json:"query_group_function"`
	QuerySortDirection   int          `json:"query_sort_direction"`
	QueryLimit           int          `json:"query_limit"`
}

type DashboardChart struct {
	Name               string                  `json:"name"`
	Tab                string                  `json:"tab"`
	Description        string                  `json:"description,omitempty"`
	PositionX          int                     `json:"position_x,omitempty"`
	PositionY          int                     `json:"position_y,omitempty"`
	MinHeight          int                     `json:"min_height"`
	MinWidth           int                     `json:"min_width"`
	DisplayTimeRange   bool                    `json:"display_time_range"`
	LabelsDisplay      bool                    `json:"labels_display"`
	LabelsAggregation  string                  `json:"labels_aggregation"`
	LabelsPlacement    string                  `json:"labels_placement"`
	RefreshInterval    string                  `json:"refresh_interval,omitempty"`
	RelativeWindowTime string                  `json:"relative_window_time,omitempty"`
	ShowBeyondData     bool                    `json:"show_beyond_data"`
	Timezone           string                  `json:"timezone,omitempty"`
	TimestampGTE       string                  `json:"timestamp_gte"`
	TimestampLTE       string                  `json:"timestamp_lte"`
	Height             int                     `json:"height"`
	Width              int                     `json:"width"`
	Collection         string                  `json:"collection"`
	ChartItems         []DashboardChartItem    `json:"chart_items"`
	Thresholds         []DashboardThreshold    `json:"thresholds"`
	ValueMappings      []DashboardValueMapping `json:"value_mappings"`
}

// Helper to extract slice of map[string]any{}
func convertList[T any](list []any, extractFunc func(map[string]any) T) []T {
	result := make([]T, len(list))
	for i, item := range list {
		result[i] = extractFunc(item.(map[string]any))
	}
	return result
}

func convertChartItem(item map[string]any) DashboardChartItem {
	queryFilterAsset := convertSingleQueryFilter(item["query_filter_asset"].(*schema.Set).List())
	queryFilterAttribute := convertSingleQueryFilter(item["query_filter_attribute"].(*schema.Set).List())

	if queryFilterAsset.isEmpty() {
		queryFilterAsset = nil
	}
	if queryFilterAttribute.isEmpty() {
		queryFilterAttribute = nil
	}

	return DashboardChartItem{
		Color:                item["color"].(string),
		RefId:                item["ref_id"].(string),
		Type:                 item["type"].(string),
		Label:                item["label"].(string),
		Hidden:               item["hidden"].(bool),
		QueryGroupUnit:       item["query_group_unit"].(string),
		QueryGroupFunction:   item["query_group_function"].(string),
		ExpressionPlain:      item["expression_plain"].(string),
		QueryFilterAsset:     queryFilterAsset,
		QueryFilterAttribute: queryFilterAttribute,
		QueryPlain:           item["query_plain"].(string),
		QuerySortDirection:   item["query_sort_direction"].(int),
		QueryLimit:           item["query_limit"].(int),
	}
}

func convertValueMapping(item map[string]any) DashboardValueMapping {
	return DashboardValueMapping{
		Type:        item["type"].(string),
		Order:       item["order"].(int),
		DisplayText: item["display_text"].(string),
		MatchValue:  item["match_value"].(string),
	}
}

func convertThreshold(item map[string]any) DashboardThreshold {
	return DashboardThreshold{
		Value:       item["value"].(float64),
		Color:       item["color"].(string),
		DisplayText: item["display_text"].(string),
	}
}

func convertDashboardChartParams(d *schema.ResourceData) *DashboardChart {
	chartItems := convertList(d.Get("chart_items").(*schema.Set).List(), convertChartItem)
	valueMappings := convertList(d.Get("value_mappings").(*schema.Set).List(), convertValueMapping)
	thresholds := convertList(d.Get("thresholds").(*schema.Set).List(), convertThreshold)

	return &DashboardChart{
		Name:               d.Get("name").(string),
		Tab:                d.Get("tab").(string),
		Description:        d.Get("description").(string),
		PositionX:          d.Get("position_x").(int),
		PositionY:          d.Get("position_y").(int),
		MinHeight:          d.Get("min_height").(int),
		MinWidth:           d.Get("min_width").(int),
		DisplayTimeRange:   d.Get("display_time_range").(bool),
		LabelsDisplay:      d.Get("labels_display").(bool),
		LabelsAggregation:  d.Get("labels_aggregation").(string),
		LabelsPlacement:    d.Get("labels_placement").(string),
		RefreshInterval:    d.Get("refresh_interval").(string),
		RelativeWindowTime: d.Get("relative_window_time").(string),
		ShowBeyondData:     d.Get("show_beyond_data").(bool),
		Timezone:           d.Get("timezone").(string),
		TimestampGTE:       d.Get("timestamp_gte").(string),
		TimestampLTE:       d.Get("timestamp_lte").(string),
		Height:             d.Get("height").(int),
		Width:              d.Get("width").(int),
		Collection:         d.Get("collection").(string),
		ChartItems:         chartItems,
		Thresholds:         thresholds,
		ValueMappings:      valueMappings,
	}
}

func saveDashboardChartToSchema(d *schema.ResourceData, m *DashboardChart) error {
	d.Set("name", m.Name)
	d.Set("tab", m.Tab)
	d.Set("description", m.Description)
	d.Set("position_x", m.PositionX)
	d.Set("position_y", m.PositionY)
	d.Set("min_height", m.MinHeight)
	d.Set("min_width", m.MinWidth)
	d.Set("display_time_range", m.DisplayTimeRange)
	d.Set("labels_display", m.LabelsDisplay)
	d.Set("labels_aggregation", m.LabelsAggregation)
	d.Set("labels_placement", m.LabelsPlacement)
	d.Set("refresh_interval", m.RefreshInterval)
	d.Set("relative_window_time", m.RelativeWindowTime)
	d.Set("show_beyond_data", m.ShowBeyondData)
	d.Set("timezone", m.Timezone)
	d.Set("height", m.Height)
	d.Set("width", m.Width)
	d.Set("timestamp_gte", m.TimestampGTE)
	d.Set("timestamp_lte", m.TimestampLTE)

	chartItems := make([]map[string]any, len(m.ChartItems))
	for i, chartItem := range m.ChartItems {
		var queryFilterAsset []map[string]string
		var queryFilterAttribute []map[string]string

		// Set to empty map in case of nil, since thats how
		// we allow it in the schema
		if chartItem.QueryFilterAsset != nil {
			queryFilterAsset = chartItem.QueryFilterAsset.toMap()
		} else {
			queryFilterAsset = (&QueryFilter{}).toMap()
		}

		if chartItem.QueryFilterAttribute != nil {
			queryFilterAttribute = chartItem.QueryFilterAttribute.toMap()
		} else {
			queryFilterAttribute = (&QueryFilter{}).toMap()
		}

		chartItems[i] = map[string]any{
			"color":                  chartItem.Color,
			"ref_id":                 chartItem.RefId,
			"type":                   chartItem.Type,
			"label":                  chartItem.Label,
			"hidden":                 chartItem.Hidden,
			"expression_plain":       chartItem.ExpressionPlain,
			"query_filter_asset":     queryFilterAsset,
			"query_filter_attribute": queryFilterAttribute,
			"query_plain":            chartItem.QueryPlain,
			"query_group_unit":       chartItem.QueryGroupUnit,
			"query_group_function":   chartItem.QueryGroupFunction,
			"query_sort_direction":   chartItem.QuerySortDirection,
			"query_limit":            chartItem.QueryLimit,
		}
	}

	d.Set("chart_items", chartItems)
	d.Set("value_mappings", m.ValueMappings)
	d.Set("thresholds", m.Thresholds)

	return nil
}
