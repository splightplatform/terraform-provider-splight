package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardBarGaugeChartParams struct {
	DashboardChart
	Type             string `json:"type"`
	MaxLimit         int    `json:"max_limit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
	Orientation      string `json:"orientation,omitempty"`
}

type DashboardBarGaugeChart struct {
	DashboardBarGaugeChartParams
	Id string `json:"id"`
}

func (m *DashboardBarGaugeChart) GetId() string {
	return m.Id
}

func (m *DashboardBarGaugeChart) GetParams() Params {
	return &m.DashboardBarGaugeChartParams
}

func (m *DashboardBarGaugeChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardBarGaugeChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardBarGaugeChartParams = DashboardBarGaugeChartParams{
		DashboardChart:   *chartParams,
		Type:             "bargauge",
		MaxLimit:         d.Get("max_limit").(int),
		NumberOfDecimals: d.Get("number_of_decimals").(int),
		Orientation:      d.Get("orientation").(string),
	}

	return nil
}

func (m *DashboardBarGaugeChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("max_limit", m.MaxLimit)
	d.Set("number_of_decimals", m.NumberOfDecimals)
	d.Set("orientation", m.Orientation)

	return nil
}
