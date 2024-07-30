package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var actionlistType string = "actionlist"

func resourceDashboardActionListChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardActionListChart(),
		Create: resourceCreateDashboardActionListChart,
		Read:   resourceReadDashboardActionListChart,
		Update: resourceUpdateDashboardActionListChart,
		Delete: resourceDeleteDashboardActionListChart,
		Exists: resourceExistsDashboardActionListChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardActionListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)

	item := client.DashboardActionListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 actionlistType,
		FilterName:           d.Get("filter_name").(string),
		ActionListType:       d.Get("action_list_type").(string),
		FilterAssetName:      d.Get("filter_asset_name").(string),
	}

	createdDashboardActionListChart, err := apiClient.CreateDashboardActionListChart(&item)
	if err != nil {
		return err
	}
	setDashboardActionListChart(d, createdDashboardActionListChart)

	return nil
}

func resourceUpdateDashboardActionListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)

	item := client.DashboardActionListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 actionlistType,
		FilterName:           d.Get("filter_name").(string),
		ActionListType:       d.Get("action_list_type").(string),
		FilterAssetName:      d.Get("filter_asset_name").(string),
	}

	updatedDashboardActionListChart, err := apiClient.UpdateDashboardActionListChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardActionListChart(d, updatedDashboardActionListChart)

	return nil
}

func resourceReadDashboardActionListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardActionListChart, err := apiClient.RetrieveDashboardActionListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardActionListChart(d, retrievedDashboardActionListChart)
	return nil
}

func resourceDeleteDashboardActionListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardActionListChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardActionListChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardActionListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardActionListChart(d *schema.ResourceData, DashboardActionListChart *client.DashboardActionListChart) {
	d.SetId(DashboardActionListChart.ID)
	d.Set("name", DashboardActionListChart.Name)
	d.Set("tab", DashboardActionListChart.Tab)
	d.Set("type", DashboardActionListChart.Type)
	d.Set("description", DashboardActionListChart.Description)
	d.Set("position_x", DashboardActionListChart.PositionX)
	d.Set("position_y", DashboardActionListChart.PositionY)
	d.Set("min_height", DashboardActionListChart.MinHeight)
	d.Set("min_width", DashboardActionListChart.MinWidth)
	d.Set("display_time_range", DashboardActionListChart.DisplayTimeRange)
	d.Set("labels_display", DashboardActionListChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardActionListChart.LabelsAggregation)
	d.Set("labels_placement", DashboardActionListChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardActionListChart.RefreshInterval)
	d.Set("relative_window_time", DashboardActionListChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardActionListChart.ShowBeyondData)
	d.Set("timezone", DashboardActionListChart.Timezone)
	d.Set("height", DashboardActionListChart.Height)
	d.Set("width", DashboardActionListChart.Width)
	d.Set("timestamp_gte", DashboardActionListChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardActionListChart.TimestampLTE)
	d.Set("chart_items", DashboardActionListChart.ChartItems)
	d.Set("value_mappings", DashboardActionListChart.ValueMappings)
	d.Set("thresholds", DashboardActionListChart.Thresholds)
	d.Set("filter_name", DashboardActionListChart.FilterName)
	d.Set("action_list_type", DashboardActionListChart.ActionListType)
	d.Set("filter_asset_name", DashboardActionListChart.FilterAssetName)
}
