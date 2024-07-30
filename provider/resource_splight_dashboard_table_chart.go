package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var tableType string = "table"

func resourceDashboardTableChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardTableChart(),
		Create: resourceCreateDashboardTableChart,
		Read:   resourceReadDashboardTableChart,
		Update: resourceUpdateDashboardTableChart,
		Delete: resourceDeleteDashboardTableChart,
		Exists: resourceExistsDashboardTableChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardTableChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardTableChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 tableType,
		YAxisUnit:            d.Get("y_axis_unit").(string),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
	}

	createdDashboardTableChart, err := apiClient.CreateDashboardTableChart(&item)
	if err != nil {
		return err
	}
	setDashboardTableChart(d, createdDashboardTableChart)

	return nil
}

func resourceUpdateDashboardTableChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardTableChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 tableType,
		YAxisUnit:            d.Get("y_axis_unit").(string),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
	}

	updatedDashboardTableChart, err := apiClient.UpdateDashboardTableChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardTableChart(d, updatedDashboardTableChart)

	return nil
}

func resourceReadDashboardTableChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardTableChart, err := apiClient.RetrieveDashboardTableChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardTableChart(d, retrievedDashboardTableChart)
	return nil
}

func resourceDeleteDashboardTableChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardTableChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardTableChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardTableChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardTableChart(d *schema.ResourceData, DashboardTableChart *client.DashboardTableChart) {
	d.SetId(DashboardTableChart.ID)
	d.Set("name", DashboardTableChart.Name)
	d.Set("tab", DashboardTableChart.Tab)
	d.Set("type", DashboardTableChart.Type)
	d.Set("description", DashboardTableChart.Description)
	d.Set("position_x", DashboardTableChart.PositionX)
	d.Set("position_y", DashboardTableChart.PositionY)
	d.Set("min_height", DashboardTableChart.MinHeight)
	d.Set("min_width", DashboardTableChart.MinWidth)
	d.Set("display_time_range", DashboardTableChart.DisplayTimeRange)
	d.Set("labels_display", DashboardTableChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardTableChart.LabelsAggregation)
	d.Set("labels_placement", DashboardTableChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardTableChart.RefreshInterval)
	d.Set("relative_window_time", DashboardTableChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardTableChart.ShowBeyondData)
	d.Set("timezone", DashboardTableChart.Timezone)
	d.Set("height", DashboardTableChart.Height)
	d.Set("width", DashboardTableChart.Width)
	d.Set("timestamp_gte", DashboardTableChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardTableChart.TimestampLTE)
	d.Set("chart_items", DashboardTableChart.ChartItems)
	d.Set("value_mappings", DashboardTableChart.ValueMappings)
	d.Set("thresholds", DashboardTableChart.Thresholds)
	d.Set("y_axis_unit", DashboardTableChart.YAxisUnit)
	d.Set("number_of_decimals", DashboardTableChart.NumberOfDecimals)
}
