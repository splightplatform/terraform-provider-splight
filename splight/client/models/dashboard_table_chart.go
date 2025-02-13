package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardTableChartParams struct {
	DashboardChart
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit"`
	NumberOfDecimals int    `json:"number_of_decimals"`
}

type DashboardTableChart struct {
	DashboardTableChartParams
	Id string `json:"id"`
}

func (m *DashboardTableChart) GetId() string {
	return m.Id
}

func (m *DashboardTableChart) GetParams() Params {
	return &m.DashboardTableChartParams
}

func (m *DashboardTableChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardTableChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardTableChartParams = DashboardTableChartParams{
		DashboardChart:   *chartParams,
		Type:             "table",
		YAxisUnit:        d.Get("y_axis_unit").(string),
		NumberOfDecimals: d.Get("number_of_decimals").(int),
	}

	return nil
}

func (m *DashboardTableChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("y_axis_unit", m.YAxisUnit)
	d.Set("number_of_decimals", m.NumberOfDecimals)

	return nil
}
