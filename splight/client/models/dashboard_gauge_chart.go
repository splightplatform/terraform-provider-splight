package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardGaugeChartParams struct {
	DashboardChart
	Type             string `json:"type"`
	MaxLimit         int    `json:"max_limit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
}

type DashboardGaugeChart struct {
	DashboardGaugeChartParams
	Id string `json:"id"`
}

func (m *DashboardGaugeChart) GetId() string {
	return m.Id
}

func (m *DashboardGaugeChart) GetParams() Params {
	return &m.DashboardGaugeChartParams
}

func (m *DashboardGaugeChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardGaugeChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardGaugeChartParams = DashboardGaugeChartParams{
		DashboardChart:   *chartParams,
		Type:             "gauge",
		MaxLimit:         d.Get("max_limit").(int),
		NumberOfDecimals: d.Get("number_of_decimals").(int),
	}

	return nil
}

func (m *DashboardGaugeChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("max_limit", m.MaxLimit)
	d.Set("number_of_decimals", m.NumberOfDecimals)

	return nil
}
