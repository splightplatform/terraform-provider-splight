package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardBarChartParams struct {
	DashboardChart
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
	Orientation      string `json:"orientation,omitempty"`
}

type DashboardBarChart struct {
	DashboardBarChartParams
	Id string `json:"id"`
}

func (m *DashboardBarChart) GetId() string {
	return m.Id
}

func (m *DashboardBarChart) GetParams() Params {
	return &m.DashboardBarChartParams
}

func (m *DashboardBarChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardBarChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardBarChartParams = DashboardBarChartParams{
		DashboardChart:   *chartParams,
		Type:             "bar",
		YAxisUnit:        d.Get("y_axis_unit").(string),
		NumberOfDecimals: d.Get("number_of_decimals").(int),
		Orientation:      d.Get("orientation").(string),
	}

	return nil
}

func (m *DashboardBarChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("y_axis_unit", m.YAxisUnit)
	d.Set("number_of_decimals", m.NumberOfDecimals)
	d.Set("orientation", m.Orientation)

	return nil
}
