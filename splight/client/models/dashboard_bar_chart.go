package models

type DashboardBarChartParams struct {
	DashboardChartParams
	Type             string `json:"type"`
	YAxisUnit        string `json:"y_axis_unit,omitempty"`
	NumberOfDecimals int    `json:"number_of_decimals,omitempty"`
	Orientation      string `json:"orientation,omitempty"`
}

type DashboardBarChart struct {
	DashboardBarChartParams
	ID string `json:"id"`
}
