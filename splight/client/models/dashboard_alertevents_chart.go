package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardAlertEventsChartParams struct {
	DashboardChart
	Type            string   `json:"type"`
	FilterName      string   `json:"filter_name"`
	FilterOldStatus []string `json:"filter_old_status"`
	FilterNewStatus []string `json:"filter_new_status"`
}

type DashboardAlertEventsChart struct {
	DashboardAlertEventsChartParams
	Id string `json:"id"`
}

func (m *DashboardAlertEventsChart) GetId() string {
	return m.Id
}

func (m *DashboardAlertEventsChart) GetParams() Params {
	return &m.DashboardAlertEventsChartParams
}

func (m *DashboardAlertEventsChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func convertFilterStatus(d *schema.ResourceData, key string) []string {
	raw := d.Get(key).([]interface{})
	status := make([]string, len(raw))
	for i, v := range raw {
		status[i] = v.(string)
	}
	return status
}

func (m *DashboardAlertEventsChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardAlertEventsChartParams = DashboardAlertEventsChartParams{
		DashboardChart:  *chartParams,
		Type:            "alertevents",
		FilterName:      d.Get("filter_name").(string),
		FilterOldStatus: convertFilterStatus(d, "filter_old_status"),
		FilterNewStatus: convertFilterStatus(d, "filter_new_status"),
	}

	return nil
}

func (m *DashboardAlertEventsChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("filter_name", m.FilterName)
	d.Set("filter_old_status", m.FilterOldStatus)
	d.Set("filter_new_status", m.FilterNewStatus)

	return nil
}
