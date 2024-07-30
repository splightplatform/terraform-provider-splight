package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var imageType string = "image"

func resourceDashboardImageChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardImageChart(),
		Create: resourceCreateDashboardImageChart,
		Read:   resourceReadDashboardImageChart,
		Update: resourceUpdateDashboardImageChart,
		Delete: resourceDeleteDashboardImageChart,
		Exists: resourceExistsDashboardImageChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardImageChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardImageChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 imageType,
		ImageURL:             d.Get("image_url").(string),
		ImageFile:            d.Get("image_file").(string),
	}

	createdDashboardImageChart, err := apiClient.CreateDashboardImageChart(&item)
	if err != nil {
		return err
	}
	setDashboardImageChart(d, createdDashboardImageChart)

	return nil
}

func resourceUpdateDashboardImageChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardImageChartParams{
		DashboardChartParams: DashboardChartParams,
		Type:                 imageType,
		ImageURL:             d.Get("image_url").(string),
		ImageFile:            d.Get("image_file").(string),
	}

	updatedDashboardImageChart, err := apiClient.UpdateDashboardImageChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardImageChart(d, updatedDashboardImageChart)

	return nil
}

func resourceReadDashboardImageChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardImageChart, err := apiClient.RetrieveDashboardImageChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardImageChart(d, retrievedDashboardImageChart)
	return nil
}

func resourceDeleteDashboardImageChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardImageChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardImageChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardImageChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardImageChart(d *schema.ResourceData, DashboardImageChart *client.DashboardImageChart) {
	d.SetId(DashboardImageChart.ID)
	d.Set("name", DashboardImageChart.Name)
	d.Set("tab", DashboardImageChart.Tab)
	d.Set("type", DashboardImageChart.Type)
	d.Set("description", DashboardImageChart.Description)
	d.Set("position_x", DashboardImageChart.PositionX)
	d.Set("position_y", DashboardImageChart.PositionY)
	d.Set("min_height", DashboardImageChart.MinHeight)
	d.Set("min_width", DashboardImageChart.MinWidth)
	d.Set("display_time_range", DashboardImageChart.DisplayTimeRange)
	d.Set("labels_display", DashboardImageChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardImageChart.LabelsAggregation)
	d.Set("labels_placement", DashboardImageChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardImageChart.RefreshInterval)
	d.Set("relative_window_time", DashboardImageChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardImageChart.ShowBeyondData)
	d.Set("timezone", DashboardImageChart.Timezone)
	d.Set("height", DashboardImageChart.Height)
	d.Set("width", DashboardImageChart.Width)
	d.Set("timestamp_gte", DashboardImageChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardImageChart.TimestampLTE)
	d.Set("chart_items", DashboardImageChart.ChartItems)
	d.Set("value_mappings", DashboardImageChart.ValueMappings)
	d.Set("thresholds", DashboardImageChart.Thresholds)
	d.Set("image_url", DashboardImageChart.ImageURL)
	d.Set("image_file", DashboardImageChart.ImageFile)
}
