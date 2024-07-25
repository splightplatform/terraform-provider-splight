package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var barType string = "bar"

func resourceDashboardBarChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardBarChart(),
		Create: resourceCreateDashboardBarChart,
		Read:   resourceReadDashboardBarChart,
		Update: resourceUpdateDashboardBarChart,
		Delete: resourceDeleteDashboardBarChart,
		Exists: resourceExistsDashboardBarChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardBarChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardBarChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 barType,
		YAxisUnit:            d.Get("y_axis_unit").(int),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
		Orientation:          d.Get("orientation").(string),
	}

	createdDashboardBarChart, err := apiClient.CreateDashboardBarChart(&item)
	if err != nil {
		return err
	}
	setDashboardBarChart(d, createdDashboardBarChart)

	return nil
}

func resourceUpdateDashboardBarChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardBarChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 barType,
		YAxisUnit:            d.Get("y_axis_unit").(int),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
		Orientation:          d.Get("orientation").(string),
	}

	updatedDashboardBarChart, err := apiClient.UpdateDashboardBarChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardBarChart(d, updatedDashboardBarChart)

	return nil
}

func resourceReadDashboardBarChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardBarChart, err := apiClient.RetrieveDashboardBarChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardBarChart(d, retrievedDashboardBarChart)
	return nil
}

func resourceDeleteDashboardBarChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardBarChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardBarChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardBarChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardBarChart(d *schema.ResourceData, DashboardBarChart *client.DashboardBarChart) {
	d.SetId(DashboardBarChart.ID)
	d.Set("name", DashboardBarChart.Name)
	d.Set("tab", DashboardBarChart.Tab)
	d.Set("type", DashboardBarChart.Type)
	d.Set("description", DashboardBarChart.Description)
	d.Set("position_x", DashboardBarChart.PositionX)
	d.Set("position_y", DashboardBarChart.PositionY)
	d.Set("min_height", DashboardBarChart.MinHeight)
	d.Set("min_width", DashboardBarChart.MinWidth)
	d.Set("display_time_range", DashboardBarChart.DisplayTimeRange)
	d.Set("labels_display", DashboardBarChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardBarChart.LabelsAggregation)
	d.Set("labels_placement", DashboardBarChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardBarChart.RefreshInterval)
	d.Set("relative_window_time", DashboardBarChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardBarChart.ShowBeyondData)
	d.Set("timezone", DashboardBarChart.Timezone)
	d.Set("height", DashboardBarChart.Height)
	d.Set("width", DashboardBarChart.Width)
	d.Set("timestamp_gte", DashboardBarChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardBarChart.TimestampLTE)
	d.Set("chart_items", DashboardBarChart.ChartItems)
	d.Set("value_mappings", DashboardBarChart.ValueMappings)
	d.Set("thresholds", DashboardBarChart.Thresholds)
	d.Set("y_axis_unit", DashboardBarChart.YAxisUnit)
	d.Set("number_of_decimals", DashboardBarChart.NumberOfDecimals)
	d.Set("orientation", DashboardBarChart.Orientation)
}
