package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetParams struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Geometry    json.RawMessage `json:"geometry"`
	Tags        []QueryFilter   `json:"tags"`
	Kind        *QueryFilter    `json:"kind"`
}

type Asset struct {
	AssetParams
	Id string `json:"id"`
}

func (m *Asset) GetId() string {
	return m.Id
}

func (m *Asset) GetParams() Params {
	return &m.AssetParams
}

func (m *Asset) ResourcePath() string {
	return "v2/engine/asset/assets/"
}

func (m *Asset) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.AssetParams = AssetParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Geometry:    json.RawMessage(d.Get("geometry").(string)),
		Tags:        tags,
		Kind:        kind,
	}

	return nil
}

func (m *Asset) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("geometry", string(m.Geometry))

	var tags []map[string]any
	for _, tag := range m.AssetParams.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	if m.Kind != nil {
		d.Set("kind", []map[string]any{
			{
				"id":   m.Kind.Id,
				"name": m.Kind.Name,
			},
		})
	}

	return nil
}
