package models

type DashboardAlertListChartParams struct {
	DashboardChartParams
	Type          string   `json:"type"`
	FilterName    string   `json:"filter_name"`
	FilterStatus  []string `json:"filter_status"`
	AlertListType string   `json:"alert_list_type,omitempty"`
}

type DashboardAlertListChart struct {
	DashboardAlertListChartParams
	ID string `json:"id"`
}
