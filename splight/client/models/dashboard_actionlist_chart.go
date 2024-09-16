package models

type DashboardActionListChartParams struct {
	DashboardChartParams
	Type            string `json:"type"`
	ActionListType  string `json:"action_list_type"`
	FilterName      string `json:"filter_name"`
	FilterAssetName string `json:"filter_asset_name,omitempty"`
}

type DashboardActionListChart struct {
	DashboardActionListChartParams
	Id string `json:"id"`
}
