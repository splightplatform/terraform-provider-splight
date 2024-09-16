package models

type DashboardAssetListChartParams struct {
	DashboardChartParams
	Type          string   `json:"type"`
	FilterName    string   `json:"filter_name"`
	FilterStatus  []string `json:"filter_status"`
	AssetListType string   `json:"asset_list_type,omitempty"`
}

type DashboardAssetListChart struct {
	DashboardAssetListChartParams
	Id string `json:"id"`
}
