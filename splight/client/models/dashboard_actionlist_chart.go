package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardActionListChartParams struct {
	DashboardChart
	Type            string `json:"type"`
	ActionListType  string `json:"action_list_type"`
	FilterName      string `json:"filter_name"`
	FilterAssetName string `json:"filter_asset_name,omitempty"`
}

type DashboardActionListChart struct {
	DashboardActionListChartParams
	Id string `json:"id"`
}

func (m *DashboardActionListChart) GetId() string {
	return m.Id
}

func (m *DashboardActionListChart) GetParams() Params {
	return &m.DashboardActionListChartParams
}

func (m *DashboardActionListChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardActionListChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardActionListChartParams = DashboardActionListChartParams{
		DashboardChart:  *chartParams,
		Type:            "actionlist",
		FilterName:      d.Get("filter_name").(string),
		ActionListType:  d.Get("action_list_type").(string),
		FilterAssetName: d.Get("filter_asset_name").(string),
	}

	return nil
}

func (m *DashboardActionListChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("filter_name", m.ActionListType)
	d.Set("action_list_type", m.FilterName)
	d.Set("filter_asset_name", m.FilterAssetName)

	return nil
}
