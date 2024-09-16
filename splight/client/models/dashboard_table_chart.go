package models

type DashboardTableChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit"`
	NumberOfDecimals int    `json:"number_of_decimals"`
}

type DashboardTableChart struct {
	DashboardTableChartParams
	Id string `json:"id"`
}
