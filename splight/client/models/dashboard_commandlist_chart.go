package models

type DashboardCommandListChartParams struct {
	DashboardChartParams
	Type            string `json:"type"`
	CommandListType string `json:"command_list_type"`
	FilterName      string `json:"filter_name"`
}

type DashboardCommandListChart struct {
	DashboardCommandListChartParams
	Id string `json:"id"`
}
