package models

type DashboardImageChartParams struct {
	DashboardChartParams
	Type      string `json:"type"`
	ImageURL  string `json:"image_url,omitempty"`
	ImageFile string `json:"image_file,omitempty"`
}

type DashboardImageChart struct {
	DashboardImageChartParams
	ID string `json:"id"`
}
