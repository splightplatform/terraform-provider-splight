package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

var histogramType string = "histogram"

func resourceDashboardHistogramChart() *schema.Resource {
	return &schema.Resource{
		Schema: schemaDashboardHistogramChart(),
		Create: resourceCreateDashboardHistogramChart,
		Read:   resourceReadDashboardHistogramChart,
		Update: resourceUpdateDashboardHistogramChart,
		Delete: resourceDeleteDashboardHistogramChart,
		Exists: resourceExistsDashboardHistogramChart,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateDashboardHistogramChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardHistogramChartParams{
		DashboardChartParams:  DashboardChartParams,
		Type:                  histogramType,
		NumberOfDecimals:      d.Get("number_of_decimals").(int),
		BucketCount:           d.Get("bucket_count").(int),
		BucketSize:            d.Get("bucket_size").(int),
		HistogramType:         d.Get("histogram_type").(string),
		Sorting:               d.Get("sorting").(string),
		Stacked:               d.Get("stacked").(bool),
		CategoriesTopMaxLimit: d.Get("categories_top_max_limit").(int),
	}

	createdDashboardHistogramChart, err := apiClient.CreateDashboardHistogramChart(&item)
	if err != nil {
		return err
	}
	setDashboardHistogramChart(d, createdDashboardHistogramChart)

	return nil
}

func resourceUpdateDashboardHistogramChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	DashboardChartParams := getDashboardChartParams(d)
	item := client.DashboardHistogramChartParams{
		DashboardChartParams:  DashboardChartParams,
		Type:                  histogramType,
		NumberOfDecimals:      d.Get("number_of_decimals").(int),
		BucketCount:           d.Get("bucket_count").(int),
		BucketSize:            d.Get("bucket_size").(int),
		HistogramType:         d.Get("histogram_type").(string),
		Sorting:               d.Get("sorting").(string),
		Stacked:               d.Get("stacked").(bool),
		CategoriesTopMaxLimit: d.Get("categories_top_max_limit").(int),
	}

	updatedDashboardHistogramChart, err := apiClient.UpdateDashboardHistogramChart(itemId, &item)
	if err != nil {
		return err
	}

	setDashboardHistogramChart(d, updatedDashboardHistogramChart)

	return nil
}

func resourceReadDashboardHistogramChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedDashboardHistogramChart, err := apiClient.RetrieveDashboardHistogramChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding DashboardChart with ID %s", itemId)
		}
	}
	setDashboardHistogramChart(d, retrievedDashboardHistogramChart)
	return nil
}

func resourceDeleteDashboardHistogramChart(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteDashboardHistogramChart(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsDashboardHistogramChart(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveDashboardHistogramChart(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func setDashboardHistogramChart(d *schema.ResourceData, DashboardHistogramChart *client.DashboardHistogramChart) {
	d.SetId(DashboardHistogramChart.ID)
	d.Set("name", DashboardHistogramChart.Name)
	d.Set("tab", DashboardHistogramChart.Tab)
	d.Set("type", DashboardHistogramChart.Type)
	d.Set("description", DashboardHistogramChart.Description)
	d.Set("position_x", DashboardHistogramChart.PositionX)
	d.Set("position_y", DashboardHistogramChart.PositionY)
	d.Set("min_height", DashboardHistogramChart.MinHeight)
	d.Set("min_width", DashboardHistogramChart.MinWidth)
	d.Set("display_time_range", DashboardHistogramChart.DisplayTimeRange)
	d.Set("labels_display", DashboardHistogramChart.LabelsDisplay)
	d.Set("labels_aggregation", DashboardHistogramChart.LabelsAggregation)
	d.Set("labels_placement", DashboardHistogramChart.LabelsPlacement)
	d.Set("refresh_interval", DashboardHistogramChart.RefreshInterval)
	d.Set("relative_window_time", DashboardHistogramChart.RelativeWindowTime)
	d.Set("show_beyond_data", DashboardHistogramChart.ShowBeyondData)
	d.Set("timezone", DashboardHistogramChart.Timezone)
	d.Set("height", DashboardHistogramChart.Height)
	d.Set("width", DashboardHistogramChart.Width)
	d.Set("timestamp_gte", DashboardHistogramChart.TimestampGTE)
	d.Set("timestamp_lte", DashboardHistogramChart.TimestampLTE)
	d.Set("chart_items", DashboardHistogramChart.ChartItems)
	d.Set("value_mappings", DashboardHistogramChart.ValueMappings)
	d.Set("thresholds", DashboardHistogramChart.Thresholds)
	d.Set("number_of_decimals", DashboardHistogramChart.NumberOfDecimals)
	d.Set("bucket_count", DashboardHistogramChart.BucketCount)
	d.Set("bucket_size", DashboardHistogramChart.BucketSize)
	d.Set("histogram_type", DashboardHistogramChart.HistogramType)
	d.Set("sorting", DashboardHistogramChart.Sorting)
	d.Set("stacked", DashboardHistogramChart.Stacked)
	d.Set("categories_top_max_limit", DashboardHistogramChart.CategoriesTopMaxLimit)
}
