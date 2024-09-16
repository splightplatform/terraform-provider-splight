package models

type DashboardTextChartParams struct {
	DashboardChartParams
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type DashboardTextChart struct {
	DashboardTextChartParams
	Id string `json:"id"`
}
