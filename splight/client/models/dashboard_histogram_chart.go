package models

type DashboardHistogramChartParams struct {
	DashboardChartParams
	Type                  string `json:"type"`
	NumberOfDecimals      int    `json:"number_of_decimals,omitempty"`
	BucketCount           int    `json:"bucket_count"`
	BucketSize            int    `json:"bucket_size,omitempty"`
	HistogramType         string `json:"histogram_type"`
	Sorting               string `json:"sorting"`
	Stacked               bool   `json:"stacked"`
	CategoriesTopMaxLimit int    `json:"categories_top_max_limit,omitempty"`
}

type DashboardHistogramChart struct {
	DashboardHistogramChartParams
	Id string `json:"id"`
}
