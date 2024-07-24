package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceDashboardTimeseriesChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardTimeseriesChart(),
		Create: resourceCreateDashboardTimeseriesChart,
		Read:   resourceReadDashboardTimeseriesChart,
		Update: resourceUpdateDashboardTimeseriesChart,
		Delete: resourceDeleteDashboardTimeseriesChart,
		Exists: resourceExistsDashboardTimeseriesChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardTimeseriesChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardTimeseriesChartParams{
		DashboardChartParams:   DashboardChartParams,
		YAxisMaxLimit:          d.Get("y_axis_max_limit").(int),
		YAxisMinLimit:          d.Get("y_axis_min_limit").(int),
		YAxisUnit:              d.Get("y_axis_unit").(string),
		NumberOfDecimals:       d.Get("number_of_decimals").(int),
		XAxisFormat:            d.Get("x_axis_format").(string),
		XAxisAutoSkip:          d.Get("x_axis_auto_skip").(bool),
		XAxisMaxTicksLimit:     d.Get("x_axis_max_ticks_limit").(int),
		LineInterpolationStyle: d.Get("line_interpolation_style").(string),
		TimeseriesType:         d.Get("timeseries_type").(string),
		Fill:                   d.Get("fill").(bool),
		ShowLine:               d.Get("show_line").(bool),
	}

	createdDashboardTimeseriesChart, err := apiClient.CreateDashboardTimeseriesChart(&item)
	if err != nil {
		return err
	}
	setDashboardTimeseriesChart(d, createdDashboardTimeseriesChart)

	return nil
}

func resourceUpdateDashboardTimeseriesChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardTimeseriesChartParams{
		DashboardChartParams:   DashboardChartParams,
		YAxisMaxLimit:          d.Get("y_axis_max_limit").(int),
		YAxisMinLimit:          d.Get("y_axis_min_limit").(int),
		YAxisUnit:              d.Get("y_axis_unit").(string),
		NumberOfDecimals:       d.Get("number_of_decimals").(int),
		XAxisFormat:            d.Get("x_axis_format").(string),
		XAxisAutoSkip:          d.Get("x_axis_auto_skip").(bool),
		XAxisMaxTicksLimit:     d.Get("x_axis_max_ticks_limit").(int),
		LineInterpolationStyle: d.Get("line_interpolation_style").(string),
		TimeseriesType:         d.Get("timeseries_type").(string),
		Fill:                   d.Get("fill").(bool),
		ShowLine:               d.Get("show_line").(bool),
	}

	updatedDashboardTimeseriesChart, err := apiClient.UpdateDashboardTimeseriesChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardTimeseriesChart(d, updatedDashboardTimeseriesChart)

	return nil
}

func resourceReadDashboardTimeseriesChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardTimeseriesChart, err := apiClient.RetrieveDashboardTimeseriesChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardTimeseriesChart(d, retrievedDashboardTimeseriesChart)
	return nil
}

func resourceDeleteDashboardTimeseriesChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardTimeseriesChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardTimeseriesChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardTimeseriesChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardTimeseriesChart(d *schema.ResourceData, DashboardTimeseriesChart *client.DashboardTimeseriesChart) {
	d.SetId(DashboardTimeseriesChart.ID)
	d.Set("name", DashboardTimeseriesChart.Name)
	d.Set("tab", DashboardTimeseriesChart.Tab)
	d.Set("type", DashboardTimeseriesChart.Type)
	d.Set("description", DashboardTimeseriesChart.Description)
	d.Set("position_x", DashboardTimeseriesChart.PositionX)
	d.Set("position_y", DashboardTimeseriesChart.PositionY)
	d.Set("min_height", DashboardTimeseriesChart.MinHeight)
	d.Set("min_width", DashboardTimeseriesChart.MinWidth)
	d.Set("display_time_range", DashboardTimeseriesChart.DisplayTimeRange)
	d.Set("labels_display", DashboardTimeseriesChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardTimeseriesChart.LabelsAggregation)
	d.Set("labels_placement", DashboardTimeseriesChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardTimeseriesChart.RefreshInterval)
	d.Set("relative_window_time", DashboardTimeseriesChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardTimeseriesChart.ShowBeyondData)
	d.Set("timezone", DashboardTimeseriesChart.Timezone)
	d.Set("height", DashboardTimeseriesChart.Height)
	d.Set("width", DashboardTimeseriesChart.Width)
	d.Set("timestamp_gte", DashboardTimeseriesChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardTimeseriesChart.TimestampLTE)
	d.Set("chart_items", DashboardTimeseriesChart.ChartItems)
	d.Set("value_mappings", DashboardTimeseriesChart.ValueMappings)
	d.Set("thresholds", DashboardTimeseriesChart.Thresholds)
	d.Set("y_axis_max_limit", DashboardTimeseriesChart.YAxisMaxLimit)
	d.Set("y_axis_min_limit", DashboardTimeseriesChart.YAxisMinLimit)
	d.Set("y_axis_unit", DashboardTimeseriesChart.YAxisUnit)
	d.Set("number_of_decimals", DashboardTimeseriesChart.NumberOfDecimals)
	d.Set("x_axis_format", DashboardTimeseriesChart.XAxisFormat)
	d.Set("x_axis_auto_skip", DashboardTimeseriesChart.XAxisAutoSkip)
	d.Set("x_axis_max_ticks_limit", DashboardTimeseriesChart.XAxisMaxTicksLimit)
	d.Set("line_interpolation_style", DashboardTimeseriesChart.LineInterpolationStyle)
	d.Set("timeseries_type", DashboardTimeseriesChart.TimeseriesType)
	d.Set("fill", DashboardTimeseriesChart.Fill)
	d.Set("show_line", DashboardTimeseriesChart.ShowLine)
}
