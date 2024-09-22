package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type DashboardParams struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	RelatedAssets []QueryFilter `json:"assets"`
	Tags          []QueryFilter `json:"tags"`
}

type Dashboard struct {
	DashboardParams
	Id string `json:"id"`
}

func (m *Dashboard) GetId() string {
	return m.Id
}

func (m *Dashboard) GetParams() Params {
	return &m.DashboardParams
}

func (m *Dashboard) ResourcePath() string {
	return "v2/engine/dashboard/dashboards/"
}

func (m *Dashboard) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	assets := convertQueryFilters(d.Get("related_assets").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.DashboardParams = DashboardParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		RelatedAssets: assets,
		Tags:          tags,
	}

	return nil
}

func (m *Dashboard) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("related_assets", m.RelatedAssets)
	d.Set("description", m.Description)

	return nil
}
