package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var commandlistType string = "commandlist"

func resourceDashboardCommandListChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardCommandListChart(),
		Create: resourceCreateDashboardCommandListChart,
		Read:   resourceReadDashboardCommandListChart,
		Update: resourceUpdateDashboardCommandListChart,
		Delete: resourceDeleteDashboardCommandListChart,
		Exists: resourceExistsDashboardCommandListChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardCommandListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardCommandListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 commandlistType,
		CommandListType:      d.Get("command_list_type").(string),
		FilterName:           d.Get("filter_name").(string),
	}

	createdDashboardCommandListChart, err := apiClient.CreateDashboardCommandListChart(&item)
	if err != nil {
		return err
	}
	setDashboardCommandListChart(d, createdDashboardCommandListChart)

	return nil
}

func resourceUpdateDashboardCommandListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardCommandListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 commandlistType,
		CommandListType:      d.Get("command_list_type").(string),
		FilterName:           d.Get("filter_name").(string),
	}

	updatedDashboardCommandListChart, err := apiClient.UpdateDashboardCommandListChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardCommandListChart(d, updatedDashboardCommandListChart)

	return nil
}

func resourceReadDashboardCommandListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardCommandListChart, err := apiClient.RetrieveDashboardCommandListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardCommandListChart(d, retrievedDashboardCommandListChart)
	return nil
}

func resourceDeleteDashboardCommandListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardCommandListChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardCommandListChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardCommandListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardCommandListChart(d *schema.ResourceData, DashboardCommandListChart *client.DashboardCommandListChart) {
	d.SetId(DashboardCommandListChart.ID)
	d.Set("name", DashboardCommandListChart.Name)
	d.Set("tab", DashboardCommandListChart.Tab)
	d.Set("type", DashboardCommandListChart.Type)
	d.Set("description", DashboardCommandListChart.Description)
	d.Set("position_x", DashboardCommandListChart.PositionX)
	d.Set("position_y", DashboardCommandListChart.PositionY)
	d.Set("min_height", DashboardCommandListChart.MinHeight)
	d.Set("min_width", DashboardCommandListChart.MinWidth)
	d.Set("display_time_range", DashboardCommandListChart.DisplayTimeRange)
	d.Set("labels_display", DashboardCommandListChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardCommandListChart.LabelsAggregation)
	d.Set("labels_placement", DashboardCommandListChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardCommandListChart.RefreshInterval)
	d.Set("relative_window_time", DashboardCommandListChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardCommandListChart.ShowBeyondData)
	d.Set("timezone", DashboardCommandListChart.Timezone)
	d.Set("height", DashboardCommandListChart.Height)
	d.Set("width", DashboardCommandListChart.Width)
	d.Set("timestamp_gte", DashboardCommandListChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardCommandListChart.TimestampLTE)
	d.Set("chart_items", DashboardCommandListChart.ChartItems)
	d.Set("value_mappings", DashboardCommandListChart.ValueMappings)
	d.Set("thresholds", DashboardCommandListChart.Thresholds)
	d.Set("command_list_type", DashboardCommandListChart.CommandListType)
	d.Set("filter_name", DashboardCommandListChart.FilterName)
}
