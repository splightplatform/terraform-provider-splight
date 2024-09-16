package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetParams struct {
	Geometry    json.RawMessage `json:"geometry"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
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
	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.AssetParams = AssetParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Geometry:    json.RawMessage(d.Get("geometry").(string)),
		Tags:        tags,
		Kind:        kind,
	}
	m.Id = d.Id()

	return nil
}

func (m *Asset) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("geometry", string(m.Geometry))
	d.Set("tags", m.Tags)

	if m.Kind != nil {
		d.Set("kind", []map[string]interface{}{
			{
				"id":   m.Kind.Id,
				"name": m.Kind.Name,
			},
		})
	}

	return nil
}
