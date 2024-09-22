package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardAlertListChartParams struct {
	DashboardChart
	Type          string   `json:"type"`
	FilterName    string   `json:"filter_name"`
	FilterStatus  []string `json:"filter_status"`
	AlertListType string   `json:"alert_list_type,omitempty"`
}

type DashboardAlertListChart struct {
	DashboardAlertListChartParams
	Id string `json:"id"`
}

func (m *DashboardAlertListChart) GetId() string {
	return m.Id
}

func (m *DashboardAlertListChart) GetParams() Params {
	return &m.DashboardAlertListChartParams
}

func (m *DashboardAlertListChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardAlertListChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardAlertListChartParams = DashboardAlertListChartParams{
		DashboardChart: *chartParams,
		Type:           "alertlist",
		FilterName:     d.Get("filter_name").(string),
		FilterStatus:   convertFilterStatus(d, "filter_status"),
		AlertListType:  d.Get("alert_list_type").(string),
	}

	return nil
}

func (m *DashboardAlertListChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("filter_name", m.FilterName)
	d.Set("filter_status", m.FilterStatus)
	d.Set("alert_list_type", m.AlertListType)

	return nil
}
