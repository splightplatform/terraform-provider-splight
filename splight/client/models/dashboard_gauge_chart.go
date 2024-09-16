package models

type DashboardGaugeChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	MaxLimit         int    `json:"max_limit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
}

type DashboardGaugeChart struct {
	DashboardGaugeChartParams
	Id string `json:"id"`
}
