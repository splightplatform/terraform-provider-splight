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
	d.Set("description", m.Description)

	var relatedasets []map[string]any
	for _, relatedAsset := range m.RelatedAssets {
		relatedasets = append(relatedasets, map[string]any{
			"id":   relatedAsset.Id,
			"name": relatedAsset.Name,
		})
	}
	d.Set("related_assets", relatedasets)

	d.Set("description", m.Description)

	return nil
}
