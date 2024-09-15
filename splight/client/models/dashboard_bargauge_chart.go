package models

type DashboardBarGaugeChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	MaxLimit         int    `json:"max_limit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
	Orientation      string `json:"orientation,omitempty"`
}

type DashboardBarGaugeChart struct {
	DashboardBarGaugeChartParams
	ID string `json:"id"`
}
