package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var assetlistType string = "assetlist"

func resourceDashboardAssetListChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardAssetListChart(),
		Create: resourceCreateDashboardAssetListChart,
		Read:   resourceReadDashboardAssetListChart,
		Update: resourceUpdateDashboardAssetListChart,
		Delete: resourceDeleteDashboardAssetListChart,
		Exists: resourceExistsDashboardAssetListChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardAssetListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)

	filterStatusInterfeace := d.Get("filter_status").([]interface{})
	filterStatus := make([]string, len(filterStatusInterfeace))
	for i, v := range filterStatusInterfeace {
		filterStatus[i] = v.(string)
	}

	item := client.DashboardAssetListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 assetlistType,
		FilterName:           d.Get("filter_name").(string),
		FilterStatus:         filterStatus,
		AssetListType:        d.Get("asset_list_type").(string),
	}

	createdDashboardAssetListChart, err := apiClient.CreateDashboardAssetListChart(&item)
	if err != nil {
		return err
	}
	setDashboardAssetListChart(d, createdDashboardAssetListChart)

	return nil
}

func resourceUpdateDashboardAssetListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)

	filterStatusInterfeace := d.Get("filter_status").([]interface{})
	filterStatus := make([]string, len(filterStatusInterfeace))
	for i, v := range filterStatusInterfeace {
		filterStatus[i] = v.(string)
	}

	item := client.DashboardAssetListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 assetlistType,
		FilterName:           d.Get("filter_name").(string),
		FilterStatus:         filterStatus,
		AssetListType:        d.Get("asset_list_type").(string),
	}

	updatedDashboardAssetListChart, err := apiClient.UpdateDashboardAssetListChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardAssetListChart(d, updatedDashboardAssetListChart)

	return nil
}

func resourceReadDashboardAssetListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardAssetListChart, err := apiClient.RetrieveDashboardAssetListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardAssetListChart(d, retrievedDashboardAssetListChart)
	return nil
}

func resourceDeleteDashboardAssetListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardAssetListChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardAssetListChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardAssetListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardAssetListChart(d *schema.ResourceData, DashboardAssetListChart *client.DashboardAssetListChart) {
	d.SetId(DashboardAssetListChart.ID)
	d.Set("name", DashboardAssetListChart.Name)
	d.Set("tab", DashboardAssetListChart.Tab)
	d.Set("type", DashboardAssetListChart.Type)
	d.Set("description", DashboardAssetListChart.Description)
	d.Set("position_x", DashboardAssetListChart.PositionX)
	d.Set("position_y", DashboardAssetListChart.PositionY)
	d.Set("min_height", DashboardAssetListChart.MinHeight)
	d.Set("min_width", DashboardAssetListChart.MinWidth)
	d.Set("display_time_range", DashboardAssetListChart.DisplayTimeRange)
	d.Set("labels_display", DashboardAssetListChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardAssetListChart.LabelsAggregation)
	d.Set("labels_placement", DashboardAssetListChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardAssetListChart.RefreshInterval)
	d.Set("relative_window_time", DashboardAssetListChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardAssetListChart.ShowBeyondData)
	d.Set("timezone", DashboardAssetListChart.Timezone)
	d.Set("height", DashboardAssetListChart.Height)
	d.Set("width", DashboardAssetListChart.Width)
	d.Set("timestamp_gte", DashboardAssetListChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardAssetListChart.TimestampLTE)
	d.Set("chart_items", DashboardAssetListChart.ChartItems)
	d.Set("value_mappings", DashboardAssetListChart.ValueMappings)
	d.Set("thresholds", DashboardAssetListChart.Thresholds)
	d.Set("filter_name", DashboardAssetListChart.FilterName)
	d.Set("filter_status", DashboardAssetListChart.FilterStatus)
	d.Set("asset_list_type", DashboardAssetListChart.AssetListType)
}
