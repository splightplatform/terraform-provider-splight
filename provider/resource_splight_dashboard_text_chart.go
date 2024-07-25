package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var textType string = "text"

func resourceDashboardTextChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardTextChart(),
		Create: resourceCreateDashboardTextChart,
		Read:   resourceReadDashboardTextChart,
		Update: resourceUpdateDashboardTextChart,
		Delete: resourceDeleteDashboardTextChart,
		Exists: resourceExistsDashboardTextChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardTextChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardTextChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 textType,
		Text:                 d.Get("text").(string),
	}

	createdDashboardTextChart, err := apiClient.CreateDashboardTextChart(&item)
	if err != nil {
		return err
	}
	setDashboardTextChart(d, createdDashboardTextChart)

	return nil
}

func resourceUpdateDashboardTextChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardTextChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 textType,
		Text:                 d.Get("text").(string),
	}

	updatedDashboardTextChart, err := apiClient.UpdateDashboardTextChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardTextChart(d, updatedDashboardTextChart)

	return nil
}

func resourceReadDashboardTextChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardTextChart, err := apiClient.RetrieveDashboardTextChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardTextChart(d, retrievedDashboardTextChart)
	return nil
}

func resourceDeleteDashboardTextChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardTextChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardTextChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardTextChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardTextChart(d *schema.ResourceData, DashboardTextChart *client.DashboardTextChart) {
	d.SetId(DashboardTextChart.ID)
	d.Set("name", DashboardTextChart.Name)
	d.Set("tab", DashboardTextChart.Tab)
	d.Set("type", DashboardTextChart.Type)
	d.Set("description", DashboardTextChart.Description)
	d.Set("position_x", DashboardTextChart.PositionX)
	d.Set("position_y", DashboardTextChart.PositionY)
	d.Set("min_height", DashboardTextChart.MinHeight)
	d.Set("min_width", DashboardTextChart.MinWidth)
	d.Set("display_time_range", DashboardTextChart.DisplayTimeRange)
	d.Set("labels_display", DashboardTextChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardTextChart.LabelsAggregation)
	d.Set("labels_placement", DashboardTextChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardTextChart.RefreshInterval)
	d.Set("relative_window_time", DashboardTextChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardTextChart.ShowBeyondData)
	d.Set("timezone", DashboardTextChart.Timezone)
	d.Set("height", DashboardTextChart.Height)
	d.Set("width", DashboardTextChart.Width)
	d.Set("timestamp_gte", DashboardTextChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardTextChart.TimestampLTE)
	d.Set("chart_items", DashboardTextChart.ChartItems)
	d.Set("value_mappings", DashboardTextChart.ValueMappings)
	d.Set("thresholds", DashboardTextChart.Thresholds)
	d.Set("text", DashboardTextChart.Text)
}
