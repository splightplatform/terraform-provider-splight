package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardTimeseriesChartParams struct {
	DashboardChart
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

func (m *DashboardTimeseriesChart) GetId() string {
	return m.Id
}

func (m *DashboardTimeseriesChart) GetParams() Params {
	return &m.DashboardTimeseriesChartParams
}

func (m *DashboardTimeseriesChart) ResourcePath() string {
	return "v3/engine/dashboard/charts/"
}

func (m *DashboardTimeseriesChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardTimeseriesChartParams = DashboardTimeseriesChartParams{
		DashboardChart:         *chartParams,
		Type:                   "timeseries",
		YAxisMaxLimit:          d.Get("y_axis_max_limit").(int),
		YAxisMinLimit:          d.Get("y_axis_min_limit").(int),
		YAxisUnit:              d.Get("y_axis_unit").(string),
		NumberOfDecimals:       d.Get("number_of_decimals").(int),
		XAxisFormat:            d.Get("x_axis_format").(string),
		XAxisAutoSkip:          d.Get("x_axis_auto_skip").(bool),
		XAxisMaxTicksLimit:     d.Get("x_axis_max_ticks_limit").(int),
		LineInterpolationStyle: d.Get("line_interpolation_style").(string),
		TimeseriesType:         d.Get("timeseries_type").(string),
		Fill:                   d.Get("fill").(bool),
		ShowLine:               d.Get("show_line").(bool),
	}

	return nil
}

func (m *DashboardTimeseriesChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("y_axis_max_limit", m.YAxisMaxLimit)
	d.Set("y_axis_min_limit", m.YAxisMinLimit)
	d.Set("y_axis_unit", m.YAxisUnit)
	d.Set("number_of_decimals", m.NumberOfDecimals)
	d.Set("x_axis_format", m.XAxisFormat)
	d.Set("x_axis_auto_skip", m.XAxisAutoSkip)
	d.Set("x_axis_max_ticks_limit", m.XAxisMaxTicksLimit)
	d.Set("line_interpolation_style", m.LineInterpolationStyle)
	d.Set("timeseries_type", m.TimeseriesType)
	d.Set("fill", m.Fill)
	d.Set("show_line", m.ShowLine)

	return nil
}
