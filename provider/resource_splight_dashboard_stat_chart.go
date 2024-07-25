package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var statType string = "stat"

func resourceDashboardStatChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardStatChart(),
		Create: resourceCreateDashboardStatChart,
		Read:   resourceReadDashboardStatChart,
		Update: resourceUpdateDashboardStatChart,
		Delete: resourceDeleteDashboardStatChart,
		Exists: resourceExistsDashboardStatChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardStatChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardStatChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 statType,
		YAxisUnit:            d.Get("y_axis_unit").(string),
		Border:               d.Get("border").(bool),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
	}

	createdDashboardStatChart, err := apiClient.CreateDashboardStatChart(&item)
	if err != nil {
		return err
	}
	setDashboardStatChart(d, createdDashboardStatChart)

	return nil
}

func resourceUpdateDashboardStatChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardStatChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 statType,
		YAxisUnit:            d.Get("y_axis_unit").(string),
		Border:               d.Get("border").(bool),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
	}

	updatedDashboardStatChart, err := apiClient.UpdateDashboardStatChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardStatChart(d, updatedDashboardStatChart)

	return nil
}

func resourceReadDashboardStatChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardStatChart, err := apiClient.RetrieveDashboardStatChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardStatChart(d, retrievedDashboardStatChart)
	return nil
}

func resourceDeleteDashboardStatChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardStatChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardStatChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardStatChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardStatChart(d *schema.ResourceData, DashboardStatChart *client.DashboardStatChart) {
	d.SetId(DashboardStatChart.ID)
	d.Set("name", DashboardStatChart.Name)
	d.Set("tab", DashboardStatChart.Tab)
	d.Set("type", DashboardStatChart.Type)
	d.Set("description", DashboardStatChart.Description)
	d.Set("position_x", DashboardStatChart.PositionX)
	d.Set("position_y", DashboardStatChart.PositionY)
	d.Set("min_height", DashboardStatChart.MinHeight)
	d.Set("min_width", DashboardStatChart.MinWidth)
	d.Set("display_time_range", DashboardStatChart.DisplayTimeRange)
	d.Set("labels_display", DashboardStatChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardStatChart.LabelsAggregation)
	d.Set("labels_placement", DashboardStatChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardStatChart.RefreshInterval)
	d.Set("relative_window_time", DashboardStatChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardStatChart.ShowBeyondData)
	d.Set("timezone", DashboardStatChart.Timezone)
	d.Set("height", DashboardStatChart.Height)
	d.Set("width", DashboardStatChart.Width)
	d.Set("timestamp_gte", DashboardStatChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardStatChart.TimestampLTE)
	d.Set("chart_items", DashboardStatChart.ChartItems)
	d.Set("value_mappings", DashboardStatChart.ValueMappings)
	d.Set("thresholds", DashboardStatChart.Thresholds)
	d.Set("y_axis_unit", DashboardStatChart.YAxisUnit)
	d.Set("border", DashboardStatChart.Border)
	d.Set("number_of_decimals", DashboardStatChart.NumberOfDecimals)
}
