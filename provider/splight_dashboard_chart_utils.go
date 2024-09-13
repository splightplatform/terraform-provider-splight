package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func getDashboardChartParams(d *schema.ResourceData) client.DashboardChartParams {
	chartItemInterface := d.Get("chart_items").(*schema.Set).List()
	chartItemInterfaceList := make([]map[string]interface{}, len(chartItemInterface))
	for i, chartItemInterfaceItem := range chartItemInterface {
		chartItemInterfaceList[i] = chartItemInterfaceItem.(map[string]interface{})
	}
	chartItems := make([]client.DashboardChartItem, len(chartItemInterfaceList))
	for i, chartItemItem := range chartItemInterfaceList {

		chartItems[i] = client.DashboardChartItem{
			Color:              chartItemItem["color"].(string),
			RefID:              chartItemItem["ref_id"].(string),
			Type:               chartItemItem["type"].(string),
			Label:              chartItemItem["label"].(string),
			Hidden:             chartItemItem["hidden"].(bool),
			QueryGroupUnit:     chartItemItem["query_group_unit"].(string),
			QueryGroupFunction: chartItemItem["query_group_function"].(string),
			ExpressionPlain:    chartItemItem["expression_plain"].(string),
			QueryPlain:         chartItemItem["query_plain"].(string),
			QuerySortDirection: chartItemItem["query_sort_direction"].(int),
			QueryLimit:         chartItemItem["query_limit"].(int),
		}

		if chartItemItem["query_filter_asset"].(*schema.Set).Len() > 0 {
			filter_asset_item := chartItemItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
			chartItems[i].QueryFilterAsset = &client.QueryFilter{
				Id:   filter_asset_item["id"].(string),
				Name: filter_asset_item["name"].(string),
			}
		}
		if chartItemItem["query_filter_attribute"].(*schema.Set).Len() > 0 {
			filter_attr_item := chartItemItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
			chartItems[i].QueryFilterAttribute = &client.QueryFilter{
				Id:   filter_attr_item["id"].(string),
				Name: filter_attr_item["name"].(string),
			}
		}

	}
	valueMappingInterface := d.Get("value_mappings").(*schema.Set).List()
	valueMappingInterfaceList := make([]map[string]interface{}, len(valueMappingInterface))
	for i, valueMappingInterfaceItem := range valueMappingInterface {
		valueMappingInterfaceList[i] = valueMappingInterfaceItem.(map[string]interface{})
	}
	valueMappings := make([]client.DashboardValueMapping, len(valueMappingInterfaceList))
	for i, valueMappingItem := range valueMappingInterfaceList {
		valueMappings[i] = client.DashboardValueMapping{
			Type:        valueMappingItem["type"].(string),
			Order:       valueMappingItem["order"].(int),
			DisplayText: valueMappingItem["display_text"].(string),
			MatchValue:  valueMappingItem["match_value"].(string),
		}
	}
	thresholdInterface := d.Get("thresholds").(*schema.Set).List()
	thresholdInterfaceList := make([]map[string]interface{}, len(thresholdInterface))
	for i, thresholdInterfaceItem := range thresholdInterface {
		thresholdInterfaceList[i] = thresholdInterfaceItem.(map[string]interface{})
	}
	thresholds := make([]client.DashboardThreshold, len(thresholdInterfaceList))
	for i, thresholdItem := range thresholdInterfaceList {
		thresholds[i] = client.DashboardThreshold{
			Value:       thresholdItem["value"].(float64),
			Color:       thresholdItem["color"].(string),
			DisplayText: thresholdItem["display_text"].(string),
		}
	}

	return client.DashboardChartParams{
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
		ValueMappings:      valueMappings,
		Thresholds:         thresholds,
	}
}
