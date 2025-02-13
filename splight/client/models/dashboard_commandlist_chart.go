package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardCommandListChartParams struct {
	DashboardChart
	Type            string `json:"type"`
	CommandListType string `json:"command_list_type"`
	FilterName      string `json:"filter_name"`
}

type DashboardCommandListChart struct {
	DashboardCommandListChartParams
	Id string `json:"id"`
}

func (m *DashboardCommandListChart) GetId() string {
	return m.Id
}

func (m *DashboardCommandListChart) GetParams() Params {
	return &m.DashboardCommandListChartParams
}

func (m *DashboardCommandListChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardCommandListChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardCommandListChartParams = DashboardCommandListChartParams{
		DashboardChart:  *chartParams,
		Type:            "commandlist",
		CommandListType: d.Get("command_list_type").(string),
		FilterName:      d.Get("filter_name").(string),
	}

	return nil
}

func (m *DashboardCommandListChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("command_list_type", m.CommandListType)
	d.Set("filter_name", m.FilterName)

	return nil
}
