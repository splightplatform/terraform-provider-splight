package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardTextChartParams struct {
	DashboardChart
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type DashboardTextChart struct {
	DashboardTextChartParams
	Id string `json:"id"`
}

func (m *DashboardTextChart) GetId() string {
	return m.Id
}

func (m *DashboardTextChart) GetParams() Params {
	return &m.DashboardTextChartParams
}

func (m *DashboardTextChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardTextChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardTextChartParams = DashboardTextChartParams{
		DashboardChart: *chartParams,
		Type:           "text",
		Text:           d.Get("text").(string),
	}

	return nil
}

func (m *DashboardTextChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("text", m.Text)

	return nil
}
