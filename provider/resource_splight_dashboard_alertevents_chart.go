package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var alerteventsType string = "alertevents"

func resourceDashboardAlertEventsChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardAlertEventsChart(),
		Create: resourceCreateDashboardAlertEventsChart,
		Read:   resourceReadDashboardAlertEventsChart,
		Update: resourceUpdateDashboardAlertEventsChart,
		Delete: resourceDeleteDashboardAlertEventsChart,
		Exists: resourceExistsDashboardAlertEventsChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardAlertEventsChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	filterOldStatusInterfeace := d.Get("filter_old_status").([]interface{})
	filterOldStatus := make([]string, len(filterOldStatusInterfeace))
	for i, v := range filterOldStatusInterfeace {
		filterOldStatus[i] = v.(string)
	}
	filterNewStatusInterfeace := d.Get("filter_new_status").([]interface{})
	filterNewStatus := make([]string, len(filterNewStatusInterfeace))
	for i, v := range filterNewStatusInterfeace {
		filterNewStatus[i] = v.(string)
	}

	item := client.DashboardAlertEventsChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 alerteventsType,
		FilterName:           d.Get("filter_name").(string),
		FilterOldStatus:      filterOldStatus,
		FilterNewStatus:      filterNewStatus,
	}

	createdDashboardAlertEventsChart, err := apiClient.CreateDashboardAlertEventsChart(&item)
	if err != nil {
		return err
	}
	setDashboardAlertEventsChart(d, createdDashboardAlertEventsChart)

	return nil
}

func resourceUpdateDashboardAlertEventsChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	filterOldStatusInterfeace := d.Get("filter_old_status").([]interface{})
	filterOldStatus := make([]string, len(filterOldStatusInterfeace))
	for i, v := range filterOldStatusInterfeace {
		filterOldStatus[i] = v.(string)
	}
	filterNewStatusInterfeace := d.Get("filter_new_status").([]interface{})
	filterNewStatus := make([]string, len(filterNewStatusInterfeace))
	for i, v := range filterNewStatusInterfeace {
		filterNewStatus[i] = v.(string)
	}
	item := client.DashboardAlertEventsChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 alerteventsType,
		FilterName:           d.Get("filter_name").(string),
		FilterOldStatus:      filterOldStatus,
		FilterNewStatus:      filterNewStatus,
	}

	updatedDashboardAlertEventsChart, err := apiClient.UpdateDashboardAlertEventsChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardAlertEventsChart(d, updatedDashboardAlertEventsChart)

	return nil
}

func resourceReadDashboardAlertEventsChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardAlertEventsChart, err := apiClient.RetrieveDashboardAlertEventsChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardAlertEventsChart(d, retrievedDashboardAlertEventsChart)
	return nil
}

func resourceDeleteDashboardAlertEventsChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardAlertEventsChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardAlertEventsChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardAlertEventsChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardAlertEventsChart(d *schema.ResourceData, DashboardAlertEventsChart *client.DashboardAlertEventsChart) {
	d.SetId(DashboardAlertEventsChart.ID)
	d.Set("name", DashboardAlertEventsChart.Name)
	d.Set("tab", DashboardAlertEventsChart.Tab)
	d.Set("type", DashboardAlertEventsChart.Type)
	d.Set("description", DashboardAlertEventsChart.Description)
	d.Set("position_x", DashboardAlertEventsChart.PositionX)
	d.Set("position_y", DashboardAlertEventsChart.PositionY)
	d.Set("min_height", DashboardAlertEventsChart.MinHeight)
	d.Set("min_width", DashboardAlertEventsChart.MinWidth)
	d.Set("display_time_range", DashboardAlertEventsChart.DisplayTimeRange)
	d.Set("labels_display", DashboardAlertEventsChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardAlertEventsChart.LabelsAggregation)
	d.Set("labels_placement", DashboardAlertEventsChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardAlertEventsChart.RefreshInterval)
	d.Set("relative_window_time", DashboardAlertEventsChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardAlertEventsChart.ShowBeyondData)
	d.Set("timezone", DashboardAlertEventsChart.Timezone)
	d.Set("height", DashboardAlertEventsChart.Height)
	d.Set("width", DashboardAlertEventsChart.Width)
	d.Set("timestamp_gte", DashboardAlertEventsChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardAlertEventsChart.TimestampLTE)
	d.Set("chart_items", DashboardAlertEventsChart.ChartItems)
	d.Set("value_mappings", DashboardAlertEventsChart.ValueMappings)
	d.Set("thresholds", DashboardAlertEventsChart.Thresholds)
	d.Set("filter_name", DashboardAlertEventsChart.FilterName)
	d.Set("filter_old_status", DashboardAlertEventsChart.FilterOldStatus)
	d.Set("filter_new_status", DashboardAlertEventsChart.FilterNewStatus)
}
