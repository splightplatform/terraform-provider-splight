package models

type DashboardTimeseriesChartParams struct {
	DashboardChartParams
	Type                   string `json:"type"`
	YAxisMaxLimit          int    `json:"y_axis_max_limit,omitempty"`
	YAxisMinLimit          int    `json:"y_axis_min_limit,omitempty"`
	YAxisUnit              string `json:"y_axis_unit,omitempty"`
	NumberOfDecimals       int    `json:"number_of_decimals,omitempty"`
	XAxisFormat            string `json:"x_axis_format"`
	XAxisAutoSkip          bool   `json:"x_axis_auto_skip"`
	XAxisMaxTicksLimit     int    `json:"x_axis_max_ticks_limit,omitempty"`
	LineInterpolationStyle string `json:"line_interpolation_style,omitempty"`
	TimeseriesType         string `json:"timeseries_type,omitempty"`
	Fill                   bool   `json:"fill"`
	ShowLine               bool   `json:"show_line"`
}

type DashboardTimeseriesChart struct {
	DashboardTimeseriesChartParams
	Id string `json:"id"`
}
