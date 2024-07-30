package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var bargaugeType string = "bargauge"

func resourceDashboardBarGaugeChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardBarGaugeChart(),
		Create: resourceCreateDashboardBarGaugeChart,
		Read:   resourceReadDashboardBarGaugeChart,
		Update: resourceUpdateDashboardBarGaugeChart,
		Delete: resourceDeleteDashboardBarGaugeChart,
		Exists: resourceExistsDashboardBarGaugeChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardBarGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardBarGaugeChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 bargaugeType,
		MaxLimit:             d.Get("max_limit").(int),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
		Orientation:          d.Get("orientation").(string),
	}

	createdDashboardBarGaugeChart, err := apiClient.CreateDashboardBarGaugeChart(&item)
	if err != nil {
		return err
	}
	setDashboardBarGaugeChart(d, createdDashboardBarGaugeChart)

	return nil
}

func resourceUpdateDashboardBarGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardBarGaugeChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 bargaugeType,
		MaxLimit:             d.Get("max_limit").(int),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
		Orientation:          d.Get("orientation").(string),
	}

	updatedDashboardBarGaugeChart, err := apiClient.UpdateDashboardBarGaugeChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardBarGaugeChart(d, updatedDashboardBarGaugeChart)

	return nil
}

func resourceReadDashboardBarGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardBarGaugeChart, err := apiClient.RetrieveDashboardBarGaugeChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardBarGaugeChart(d, retrievedDashboardBarGaugeChart)
	return nil
}

func resourceDeleteDashboardBarGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardBarGaugeChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardBarGaugeChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardBarGaugeChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardBarGaugeChart(d *schema.ResourceData, DashboardBarGaugeChart *client.DashboardBarGaugeChart) {
	d.SetId(DashboardBarGaugeChart.ID)
	d.Set("name", DashboardBarGaugeChart.Name)
	d.Set("tab", DashboardBarGaugeChart.Tab)
	d.Set("type", DashboardBarGaugeChart.Type)
	d.Set("description", DashboardBarGaugeChart.Description)
	d.Set("position_x", DashboardBarGaugeChart.PositionX)
	d.Set("position_y", DashboardBarGaugeChart.PositionY)
	d.Set("min_height", DashboardBarGaugeChart.MinHeight)
	d.Set("min_width", DashboardBarGaugeChart.MinWidth)
	d.Set("display_time_range", DashboardBarGaugeChart.DisplayTimeRange)
	d.Set("labels_display", DashboardBarGaugeChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardBarGaugeChart.LabelsAggregation)
	d.Set("labels_placement", DashboardBarGaugeChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardBarGaugeChart.RefreshInterval)
	d.Set("relative_window_time", DashboardBarGaugeChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardBarGaugeChart.ShowBeyondData)
	d.Set("timezone", DashboardBarGaugeChart.Timezone)
	d.Set("height", DashboardBarGaugeChart.Height)
	d.Set("width", DashboardBarGaugeChart.Width)
	d.Set("timestamp_gte", DashboardBarGaugeChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardBarGaugeChart.TimestampLTE)
	d.Set("chart_items", DashboardBarGaugeChart.ChartItems)
	d.Set("value_mappings", DashboardBarGaugeChart.ValueMappings)
	d.Set("thresholds", DashboardBarGaugeChart.Thresholds)
	d.Set("max_limit", DashboardBarGaugeChart.MaxLimit)
	d.Set("number_of_decimals", DashboardBarGaugeChart.NumberOfDecimals)
	d.Set("orientation", DashboardBarGaugeChart.Orientation)
}
