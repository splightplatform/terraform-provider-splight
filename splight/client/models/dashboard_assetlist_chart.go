package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardAssetListChartParams struct {
	DashboardChart
	Type          string   `json:"type"`
	FilterName    string   `json:"filter_name"`
	FilterStatus  []string `json:"filter_status"`
	AssetListType string   `json:"asset_list_type,omitempty"`
}

type DashboardAssetListChart struct {
	DashboardAssetListChartParams
	Id string `json:"id"`
}

func (m *DashboardAssetListChart) GetId() string {
	return m.Id
}

func (m *DashboardAssetListChart) GetParams() Params {
	return &m.DashboardAssetListChartParams
}

func (m *DashboardAssetListChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardAssetListChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardAssetListChartParams = DashboardAssetListChartParams{
		DashboardChart: *chartParams,
		Type:           "assetlist",
		FilterName:     d.Get("filter_name").(string),
		FilterStatus:   convertFilterStatus(d, "filter_status"),
		AssetListType:  d.Get("asset_list_type").(string),
	}

	return nil
}

func (m *DashboardAssetListChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("filter_name", m.FilterName)
	d.Set("filter_status", m.FilterStatus)
	d.Set("asset_list_type", m.AssetListType)

	return nil
}
