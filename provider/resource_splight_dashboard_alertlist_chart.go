package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var alertlistType string = "alertlist"

func resourceDashboardAlertListChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardAlertListChart(),
		Create: resourceCreateDashboardAlertListChart,
		Read:   resourceReadDashboardAlertListChart,
		Update: resourceUpdateDashboardAlertListChart,
		Delete: resourceDeleteDashboardAlertListChart,
		Exists: resourceExistsDashboardAlertListChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardAlertListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)

	filterStatusInterfeace := d.Get("filter_status").([]interface{})
	filterStatus := make([]string, len(filterStatusInterfeace))
	for i, v := range filterStatusInterfeace {
		filterStatus[i] = v.(string)
	}

	item := client.DashboardAlertListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 alertlistType,
		FilterName:           d.Get("filter_name").(string),
		FilterStatus:         filterStatus,
		AlertListType:        d.Get("alert_list_type").(string),
	}

	createdDashboardAlertListChart, err := apiClient.CreateDashboardAlertListChart(&item)
	if err != nil {
		return err
	}
	setDashboardAlertListChart(d, createdDashboardAlertListChart)

	return nil
}

func resourceUpdateDashboardAlertListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)

	filterStatusInterfeace := d.Get("filter_status").([]interface{})
	filterStatus := make([]string, len(filterStatusInterfeace))
	for i, v := range filterStatusInterfeace {
		filterStatus[i] = v.(string)
	}

	item := client.DashboardAlertListChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 alertlistType,
		FilterName:           d.Get("filter_name").(string),
		FilterStatus:         filterStatus,
		AlertListType:        d.Get("alert_list_type").(string),
	}

	updatedDashboardAlertListChart, err := apiClient.UpdateDashboardAlertListChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardAlertListChart(d, updatedDashboardAlertListChart)

	return nil
}

func resourceReadDashboardAlertListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardAlertListChart, err := apiClient.RetrieveDashboardAlertListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardAlertListChart(d, retrievedDashboardAlertListChart)
	return nil
}

func resourceDeleteDashboardAlertListChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardAlertListChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardAlertListChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardAlertListChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardAlertListChart(d *schema.ResourceData, DashboardAlertListChart *client.DashboardAlertListChart) {
	d.SetId(DashboardAlertListChart.ID)
	d.Set("name", DashboardAlertListChart.Name)
	d.Set("tab", DashboardAlertListChart.Tab)
	d.Set("type", DashboardAlertListChart.Type)
	d.Set("description", DashboardAlertListChart.Description)
	d.Set("position_x", DashboardAlertListChart.PositionX)
	d.Set("position_y", DashboardAlertListChart.PositionY)
	d.Set("min_height", DashboardAlertListChart.MinHeight)
	d.Set("min_width", DashboardAlertListChart.MinWidth)
	d.Set("display_time_range", DashboardAlertListChart.DisplayTimeRange)
	d.Set("labels_display", DashboardAlertListChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardAlertListChart.LabelsAggregation)
	d.Set("labels_placement", DashboardAlertListChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardAlertListChart.RefreshInterval)
	d.Set("relative_window_time", DashboardAlertListChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardAlertListChart.ShowBeyondData)
	d.Set("timezone", DashboardAlertListChart.Timezone)
	d.Set("height", DashboardAlertListChart.Height)
	d.Set("width", DashboardAlertListChart.Width)
	d.Set("timestamp_gte", DashboardAlertListChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardAlertListChart.TimestampLTE)
	d.Set("chart_items", DashboardAlertListChart.ChartItems)
	d.Set("value_mappings", DashboardAlertListChart.ValueMappings)
	d.Set("thresholds", DashboardAlertListChart.Thresholds)
	d.Set("filter_name", DashboardAlertListChart.FilterName)
	d.Set("filter_status", DashboardAlertListChart.FilterStatus)
	d.Set("alert_list_type", DashboardAlertListChart.AlertListType)
}
