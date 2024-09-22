package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardStatChartParams struct {
	DashboardChart
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit"`
	Border           bool   `json:"border"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
}

type DashboardStatChart struct {
	DashboardStatChartParams
	Id string `json:"id"`
}

func (m *DashboardStatChart) GetId() string {
	return m.Id
}

func (m *DashboardStatChart) GetParams() Params {
	return &m.DashboardStatChartParams
}

func (m *DashboardStatChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardStatChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardStatChartParams = DashboardStatChartParams{
		DashboardChart:   *chartParams,
		Type:             "stat",
		YAxisUnit:        d.Get("y_axis_unit").(string),
		Border:           d.Get("border").(bool),
		NumberOfDecimals: d.Get("number_of_decimals").(int),
	}

	return nil
}

func (m *DashboardStatChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("y_axis_unit", m.YAxisUnit)
	d.Set("border", m.Border)
	d.Set("number_of_decimals", m.NumberOfDecimals)

	return nil
}
