package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExternalGridParams struct {
	AssetParams
}

type ExternalGrid struct {
	ExternalGridParams
	Id string `json:"id"`
}

func (m *ExternalGrid) GetId() string {
	return m.Id
}

func (m *ExternalGrid) GetParams() Params {
	return &m.ExternalGridParams
}

func (m *ExternalGrid) ResourcePath() string {
	return "v2/engine/asset/external-grids/"
}

func (m *ExternalGrid) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.ExternalGridParams = ExternalGridParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       json.RawMessage(d.Get("geometry").(string)),
			CustomTimezone: d.Get("custom_timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

	return nil
}

func (m *ExternalGrid) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)
	d.Set("geometry", string(m.AssetParams.Geometry))
	d.Set("custom_timezone", m.AssetParams.CustomTimezone)

	var tags []map[string]any
	for _, tag := range m.AssetParams.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	d.Set("kind", []map[string]any{
		{
			"id":   m.AssetParams.Kind.Id,
			"name": m.AssetParams.Kind.Name,
		},
	})

	return nil
}
