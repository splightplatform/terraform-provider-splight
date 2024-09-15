package models

import "encoding/json"

type QueryFilter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DashboardValueMapping struct {
	Type        string `json:"type"`
	Order       int    `json:"order"`
	DisplayText string `json:"display_text"`
	MatchValue  string `json:"match_value"`
}

type DashboardThreshold struct {
	Value       float64 `json:"value"`
	Color       string  `json:"color"`
	DisplayText string  `json:"display_text"`
}

type DashboardChartItem struct {
	Color                string      `json:"color"`
	RefID                string      `json:"ref_id"`
	Type                 string      `json:"type"`
	Label                string      `json:"label"`
	Hidden               bool        `json:"hidden"`
	QueryGroupUnit       string      `json:"query_group_unit"`
	QueryGroupFunction   string      `json:"query_group_function"`
	ExpressionPlain      string      `json:"expression_plain"`
	QueryFilterAsset     QueryFilter `json:"query_filter_asset"`
	QueryFilterAttribute QueryFilter `json:"query_filter_attribute"`
	QueryPlain           string      `json:"query_plain"`
	QuerySortDirection   int         `json:"query_sort_direction"`
	QueryLimit           int         `json:"query_limit"`
}

type DashboardChartParams struct {
	Name               string                  `json:"name"`
	Tab                string                  `json:"tab"`
	Description        string                  `json:"description,omitempty"`
	PositionX          int                     `json:"position_x,omitempty"`
	PositionY          int                     `json:"position_y,omitempty"`
	MinHeight          int                     `json:"min_height"`
	MinWidth           int                     `json:"min_width"`
	DisplayTimeRange   bool                    `json:"display_time_range"`
	LabelsDisplay      bool                    `json:"labels_display"`
	LabelsAggregation  string                  `json:"labels_aggregation"`
	LabelsPlacement    string                  `json:"labels_placement"`
	RefreshInterval    string                  `json:"refresh_interval,omitempty"`
	RelativeWindowTime string                  `json:"relative_window_time,omitempty"`
	ShowBeyondData     bool                    `json:"show_beyond_data"`
	Timezone           string                  `json:"timezone,omitempty"`
	TimestampGTE       string                  `json:"timestamp_gte"`
	TimestampLTE       string                  `json:"timestamp_lte"`
	Height             int                     `json:"height"`
	Width              int                     `json:"width"`
	Collection         string                  `json:"collection"`
	ChartItems         []DashboardChartItem    `json:"chart_items"`
	Thresholds         []DashboardThreshold    `json:"thresholds"`
	ValueMappings      []DashboardValueMapping `json:"value_mappings"`
}

type DashboardChart struct {
	DashboardChartParams
	ID string `json:"id"`
}

// Implement custom JSON marshalling to omit the struct if both fields are empty
func (fti QueryFilter) MarshalJSON() ([]byte, error) {
	if fti.Id == "" && fti.Name == "" {
		return []byte("null"), nil
	}
	type Alias QueryFilter
	return json.Marshal((Alias)(fti))
}

// Implement custom JSON unmarshalling to initialize the struct if the field is null
func (fti *QueryFilter) UnmarshalJSON(data []byte) error {
	type Alias QueryFilter
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(fti),
	}

	if string(data) == "null" {
		*fti = QueryFilter{}
		return nil
	}

	return json.Unmarshal(data, &aux)
}
