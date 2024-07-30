package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var gaugeType string = "gauge"

func resourceDashboardGaugeChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardGaugeChart(),
		Create: resourceCreateDashboardGaugeChart,
		Read:   resourceReadDashboardGaugeChart,
		Update: resourceUpdateDashboardGaugeChart,
		Delete: resourceDeleteDashboardGaugeChart,
		Exists: resourceExistsDashboardGaugeChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardGaugeChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 gaugeType,
		MaxLimit:             d.Get("max_limit").(int),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
	}

	createdDashboardGaugeChart, err := apiClient.CreateDashboardGaugeChart(&item)
	if err != nil {
		return err
	}
	setDashboardGaugeChart(d, createdDashboardGaugeChart)

	return nil
}

func resourceUpdateDashboardGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardGaugeChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 gaugeType,
		MaxLimit:             d.Get("max_limit").(int),
		NumberOfDecimals:     d.Get("number_of_decimals").(int),
	}

	updatedDashboardGaugeChart, err := apiClient.UpdateDashboardGaugeChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardGaugeChart(d, updatedDashboardGaugeChart)

	return nil
}

func resourceReadDashboardGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardGaugeChart, err := apiClient.RetrieveDashboardGaugeChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardGaugeChart(d, retrievedDashboardGaugeChart)
	return nil
}

func resourceDeleteDashboardGaugeChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardGaugeChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardGaugeChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardGaugeChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardGaugeChart(d *schema.ResourceData, DashboardGaugeChart *client.DashboardGaugeChart) {
	d.SetId(DashboardGaugeChart.ID)
	d.Set("name", DashboardGaugeChart.Name)
	d.Set("tab", DashboardGaugeChart.Tab)
	d.Set("type", DashboardGaugeChart.Type)
	d.Set("description", DashboardGaugeChart.Description)
	d.Set("position_x", DashboardGaugeChart.PositionX)
	d.Set("position_y", DashboardGaugeChart.PositionY)
	d.Set("min_height", DashboardGaugeChart.MinHeight)
	d.Set("min_width", DashboardGaugeChart.MinWidth)
	d.Set("display_time_range", DashboardGaugeChart.DisplayTimeRange)
	d.Set("labels_display", DashboardGaugeChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardGaugeChart.LabelsAggregation)
	d.Set("labels_placement", DashboardGaugeChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardGaugeChart.RefreshInterval)
	d.Set("relative_window_time", DashboardGaugeChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardGaugeChart.ShowBeyondData)
	d.Set("timezone", DashboardGaugeChart.Timezone)
	d.Set("height", DashboardGaugeChart.Height)
	d.Set("width", DashboardGaugeChart.Width)
	d.Set("timestamp_gte", DashboardGaugeChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardGaugeChart.TimestampLTE)
	d.Set("chart_items", DashboardGaugeChart.ChartItems)
	d.Set("value_mappings", DashboardGaugeChart.ValueMappings)
	d.Set("thresholds", DashboardGaugeChart.Thresholds)
	d.Set("max_limit", DashboardGaugeChart.MaxLimit)
	d.Set("number_of_decimals", DashboardGaugeChart.NumberOfDecimals)
}
