package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceDashboardChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardChart(),
		Create: resourceCreateDashboardChart,
		Read:   resourceReadDashboardChart,
		Update: resourceUpdateDashboardChart,
		Delete: resourceDeleteDashboardChart,
		Exists: resourceExistsDashboardChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	chartItemInterface := d.Get("chart_items").(*schema.Set).List()
	chartItemInterfaceList := make([]map[string]interface{}, len(chartItemInterface))
	for i, chartItemInterfaceItem := range chartItemInterface {
		chartItemInterfaceList[i] = chartItemInterfaceItem.(map[string]interface{})
	}
	chartItems := make([]client.DashboardChartItem, len(chartItemInterfaceList))
	for i, chartItemItem := range chartItemInterfaceList {

		var filter_asset_value *client.QueryFilter
		var filter_attr_value *client.QueryFilter

		if chartItemItem["query_filter_asset"].(*schema.Set).Len() > 0 {
			filter_asset_item := chartItemItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
			filter_asset_value = &client.QueryFilter{
				Id:   filter_asset_item["id"].(string),
				Name: filter_asset_item["name"].(string),
			}
		}
		if chartItemItem["query_filter_attribute"].(*schema.Set).Len() > 0 {
			filter_attr_item := chartItemItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
			filter_attr_value = &client.QueryFilter{
				Id:   filter_attr_item["id"].(string),
				Name: filter_attr_item["name"].(string),
			}
		}

		chartItems[i] = client.DashboardChartItem{
			Color:                chartItemItem["color"].(string),
			RefID:                chartItemItem["ref_id"].(string),
			Type:                 chartItemItem["type"].(string),
			Label:                chartItemItem["label"].(string),
			Hidden:               chartItemItem["hidden"].(bool),
			QueryGroupUnit:       chartItemItem["query_group_unit"].(string),
			QueryGroupFunction:   chartItemItem["query_group_function"].(string),
			ExpressionPlain:      chartItemItem["expression_plain"].(string),
			QueryFilterAsset:     filter_asset_value,
			QueryFilterAttribute: filter_attr_value,
			QueryPlain:           chartItemItem["query_plain"].(string),
			QuerySortDirection:   chartItemItem["query_sort_direction"].(int),
			QueryLimit:           chartItemItem["query_limit"].(int),
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
	typeInterface := d.Get("type").(string)
	timeseriesTypeInteraface := d.Get("timeseries_type").(string)
	var timeseriesType *string
	if typeInterface == "timeseries" {
		timeseriesType = &timeseriesTypeInteraface
	} else {
		timeseriesType = nil
	}

	item := client.DashboardChartParams{
		Name:           d.Get("name").(string),
		Tab:            d.Get("tab").(string),
		Type:           d.Get("type").(string),
		TimeseriesType: timeseriesType,
		TimestampGTE:   d.Get("timestamp_gte").(string),
		TimestampLTE:   d.Get("timestamp_lte").(string),
		Height:         d.Get("height").(int),
		Width:          d.Get("width").(int),
		Collection:     d.Get("collection").(string),
		ChartItems:     chartItems,
		ValueMappings:  valueMappings,
		Thresholds:     thresholds,
	}

	createdDashboardChart, err := apiClient.CreateDashboardChart(&item)
	if err != nil {
		return err
	}

	d.SetId(createdDashboardChart.ID)
	d.Set("name", createdDashboardChart.Name)
	d.Set("tab", createdDashboardChart.Tab)
	d.Set("type", createdDashboardChart.Type)
	if createdDashboardChart.Type == "timeseries" {
		d.Set("timeseries_type", createdDashboardChart.TimeseriesType)
	}
	d.Set("height", createdDashboardChart.Height)
	d.Set("width", createdDashboardChart.Width)
	d.Set("timestamp_gte", createdDashboardChart.TimestampGTE)
	d.Set("timestamp_lte", createdDashboardChart.TimestampLTE)
	d.Set("chart_items", createdDashboardChart.ChartItems)
	d.Set("value_mappings", createdDashboardChart.ValueMappings)
	d.Set("thresholds", createdDashboardChart.Thresholds)
	return nil
}

func resourceUpdateDashboardChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	chartItemInterface := d.Get("chart_items").(*schema.Set).List()
	chartItemInterfaceList := make([]map[string]interface{}, len(chartItemInterface))
	for i, chartItemInterfaceItem := range chartItemInterface {
		chartItemInterfaceList[i] = chartItemInterfaceItem.(map[string]interface{})
	}
	chartItems := make([]client.DashboardChartItem, len(chartItemInterfaceList))
	for i, chartItemItem := range chartItemInterfaceList {

		var filter_asset_value *client.QueryFilter
		var filter_attr_value *client.QueryFilter

		if chartItemItem["query_filter_asset"].(*schema.Set).Len() > 0 {
			filter_asset_item := chartItemItem["query_filter_asset"].(*schema.Set).List()[0].(map[string]interface{})
			filter_asset_value = &client.QueryFilter{
				Id:   filter_asset_item["id"].(string),
				Name: filter_asset_item["name"].(string),
			}
		}
		if chartItemItem["query_filter_attribute"].(*schema.Set).Len() > 0 {
			filter_attr_item := chartItemItem["query_filter_attribute"].(*schema.Set).List()[0].(map[string]interface{})
			filter_attr_value = &client.QueryFilter{
				Id:   filter_attr_item["id"].(string),
				Name: filter_attr_item["name"].(string),
			}
		}

		chartItems[i] = client.DashboardChartItem{
			Color:                chartItemItem["color"].(string),
			RefID:                chartItemItem["ref_id"].(string),
			Type:                 chartItemItem["type"].(string),
			Label:                chartItemItem["label"].(string),
			Hidden:               chartItemItem["hidden"].(bool),
			QueryGroupUnit:       chartItemItem["query_group_unit"].(string),
			QueryGroupFunction:   chartItemItem["query_group_function"].(string),
			ExpressionPlain:      chartItemItem["expression_plain"].(string),
			QueryFilterAsset:     filter_asset_value,
			QueryFilterAttribute: filter_attr_value,
			QueryPlain:           chartItemItem["query_plain"].(string),
			QuerySortDirection:   chartItemItem["query_sort_direction"].(int),
			QueryLimit:           chartItemItem["query_limit"].(int),
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

	typeInterface := d.Get("type").(string)
	timeseriesTypeInteraface := d.Get("timeseries_type").(string)
	var timeseriesType *string
	if typeInterface == "timeseries" {
		timeseriesType = &timeseriesTypeInteraface
	} else {
		timeseriesType = nil
	}

	item := client.DashboardChartParams{
		Name:           d.Get("name").(string),
		Tab:            d.Get("tab").(string),
		Type:           d.Get("type").(string),
		TimeseriesType: timeseriesType,
		TimestampGTE:   d.Get("timestamp_gte").(string),
		TimestampLTE:   d.Get("timestamp_lte").(string),
		Height:         d.Get("height").(int),
		Width:          d.Get("width").(int),
		Collection:     d.Get("collection").(string),
		ChartItems:     chartItems,
		ValueMappings:  valueMappings,
		Thresholds:     thresholds,
	}
	createdDashboardChart, err := apiClient.UpdateDashboardChart(itemId, &item)
	if err != nil {
		return err
	}

	d.SetId(createdDashboardChart.ID)
	d.Set("name", createdDashboardChart.Name)
	d.Set("tab", createdDashboardChart.Tab)
	d.Set("type", createdDashboardChart.Type)
	if createdDashboardChart.Type == "timeseries" {
		d.Set("timeseries_type", createdDashboardChart.TimeseriesType)
	}
	d.Set("height", createdDashboardChart.Height)
	d.Set("width", createdDashboardChart.Width)
	d.Set("timestamp_gte", createdDashboardChart.TimestampGTE)
	d.Set("timestamp_lte", createdDashboardChart.TimestampLTE)
	d.Set("chart_items", createdDashboardChart.ChartItems)
	d.Set("value_mappings", createdDashboardChart.ValueMappings)
	d.Set("thresholds", createdDashboardChart.Thresholds)
	return nil
}

func resourceReadDashboardChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardChart, err := apiClient.RetrieveDashboardChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}

	chartItemsDict := make([]map[interface{}]interface{}, len(retrievedDashboardChart.ChartItems))
	for i, item := range retrievedDashboardChart.ChartItems {
		chartItem := map[interface{}]interface{}{
			"color":                item.Color,
			"ref_id":               item.RefID,
			"type":                 item.Type,
			"label":                item.Label,
			"hidden":               item.Hidden,
			"expression_plain":     item.ExpressionPlain,
			"query_group_unit":     item.QueryGroupUnit,
			"query_group_function": item.QueryGroupFunction,
			"query_plain":          item.QueryPlain,
			"query_sort_direction": item.QuerySortDirection,
			"query_limit":          item.QueryLimit,
		}
		if item.QueryFilterAsset != nil {
			chartItem["query_filter_asset"] = map[string]interface{}{
				"id":   item.QueryFilterAsset.Id,
				"name": item.QueryFilterAsset.Name,
			}
		} else {
			chartItem["query_filter_asset"] = nil
		}
		if item.QueryFilterAttribute != nil {
			chartItem["query_filter_attribute"] = map[string]interface{}{
				"id":   item.QueryFilterAttribute.Id,
				"name": item.QueryFilterAttribute.Name,
			}
		} else {
			chartItem["query_filter_attribute"] = nil
		}
		chartItemsDict[i] = chartItem
	}

	thresholdsDict := make([]map[interface{}]interface{}, len(retrievedDashboardChart.Thresholds))
	for i, item := range retrievedDashboardChart.Thresholds {
		thresholdsDict[i] = map[interface{}]interface{}{
			"color":        item.Color,
			"display_text": item.DisplayText,
			"value":        item.Value,
		}
	}
	valueMappingsDict := make([]map[interface{}]interface{}, len(retrievedDashboardChart.ValueMappings))
	for i, item := range retrievedDashboardChart.ValueMappings {
		valueMappingsDict[i] = map[interface{}]interface{}{
			"order":        item.Order,
			"type":         item.Type,
			"display_text": item.DisplayText,
			"match_value":  item.MatchValue,
		}
	}
	d.SetId(retrievedDashboardChart.ID)
	d.Set("name", retrievedDashboardChart.Name)
	d.Set("tab", retrievedDashboardChart.Tab)
	d.Set("type", retrievedDashboardChart.Type)
	if retrievedDashboardChart.Type == "timeseries" {
		d.Set("timeseries_type", retrievedDashboardChart.TimeseriesType)
	}
	d.Set("timestamp_gte", retrievedDashboardChart.TimestampGTE)
	d.Set("timestamp_lte", retrievedDashboardChart.TimestampLTE)
	d.Set("height", retrievedDashboardChart.Height)
	d.Set("width", retrievedDashboardChart.Width)
	d.Set("chart_items", chartItemsDict)
	d.Set("value_mappings", valueMappingsDict)
	d.Set("thresholds", thresholdsDict)
	return nil
}

func resourceDeleteDashboardChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
