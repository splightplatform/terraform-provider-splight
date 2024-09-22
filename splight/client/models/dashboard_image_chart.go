package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardImageChartParams struct {
	DashboardChart
	Type      string `json:"type"`
	ImageURL  string `json:"image_url,omitempty"`
	ImageFile string `json:"image_file,omitempty"`
}

type DashboardImageChart struct {
	DashboardImageChartParams
	Id string `json:"id"`
}

func (m *DashboardImageChart) GetId() string {
	return m.Id
}

func (m *DashboardImageChart) GetParams() Params {
	return &m.DashboardImageChartParams
}

func (m *DashboardImageChart) ResourcePath() string {
	return "v2/engine/dashboard/charts/"
}

func (m *DashboardImageChart) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	chartParams := convertDashboardChartParams(d)

	m.DashboardImageChartParams = DashboardImageChartParams{
		DashboardChart: *chartParams,
		Type:           "image",
		ImageURL:       d.Get("image_url").(string),
		ImageFile:      d.Get("image_file").(string),
	}

	return nil
}

func (m *DashboardImageChart) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	saveDashboardChartToSchema(d, &m.DashboardChart)

	d.Set("type", m.Type)
	d.Set("image_url", m.ImageURL)
	d.Set("image_file", m.ImageFile)

	return nil
}
