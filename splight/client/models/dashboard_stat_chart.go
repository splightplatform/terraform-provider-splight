package models

type DashboardStatChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit"`
	Border           bool   `json:"border"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
}

type DashboardStatChart struct {
	DashboardStatChartParams
	Id string `json:"id"`
}
