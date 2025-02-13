package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardHistogramChartParams struct {
	DashboardChart
	Type                  string `json:"type"`
	NumberOfDecimals      int    `json:"number_of_decimals,omitempty"`
	BucketCount           int    `json:"bucket_count"`
	BucketSize            int    `json:"bucket_size,omitempty"`
	HistogramType         string `json:"histogram_type"`
	Sorting               string `json:"sorting"`
	Stacked               bool   `json:"stacked"`
	CategoriesTopMaxLimit int    `json:"categories_top_max_limit,omitempty"`
}

type DashboardHistogramChart struct {
	DashboardHistogramChartParams
	Id string `json:"id"`
}

func (m *DashboardHistogramChart) GetId() string {
	return m.Id
}

func (m *DashboardHistogramChart) GetParams() Params {
	return &m.DashboardHistogramChartParams
}

func (m *DashboardHistogramChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardHistogramChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardHistogramChartParams = DashboardHistogramChartParams{
		DashboardChart:        *chartParams,
		Type:                  "histogram",
		NumberOfDecimals:      d.Get("number_of_decimals").(int),
		BucketCount:           d.Get("bucket_count").(int),
		BucketSize:            d.Get("bucket_size").(int),
		HistogramType:         d.Get("histogram_type").(string),
		Sorting:               d.Get("sorting").(string),
		Stacked:               d.Get("stacked").(bool),
		CategoriesTopMaxLimit: d.Get("categories_top_max_limit").(int),
	}

	return nil
}

func (m *DashboardHistogramChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("number_of_decimals", m.NumberOfDecimals)
	d.Set("bucket_count", m.BucketCount)
	d.Set("bucket_size", m.BucketSize)
	d.Set("histogram_type", m.HistogramType)
	d.Set("sorting", m.Sorting)
	d.Set("stacked", m.Stacked)
	d.Set("categories_top_max_limit", m.CategoriesTopMaxLimit)

	return nil
}
