package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BusParams struct {
	AssetParams
	NominalVoltage AssetMetadata `json:"nominal_voltage"`
}

type Bus struct {
	BusParams
	Id string `json:"id"`
}

func (m *Bus) GetId() string {
	return m.Id
}

func (m *Bus) GetParams() Params {
	return &m.BusParams
}

func (m *Bus) ResourcePath() string {
	return "v2/engine/asset/buses/"
}

func (m *Bus) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.BusParams = BusParams{
		AssetParams: AssetParams{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Geometry:    json.RawMessage(d.Get("geometry").(string)),
			Tags:        tags,
			Kind:        kind,
		},
	}

	nominalVoltage := convertAssetMetadata(d.Get("nominal_voltage").(*schema.Set).List())
	if nominalVoltage.Type == "" {
		nominalVoltage.Type = "Number"
	}
	if nominalVoltage.Name == "" {
		nominalVoltage.Name = "nominal_voltage"
	}
	m.BusParams.NominalVoltage = *nominalVoltage

	return nil
}

func (m *Bus) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)
	d.Set("geometry", string(m.AssetParams.Geometry))

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

	d.Set("nominal_voltage", []map[string]any{m.NominalVoltage.ToMap()})

	return nil
}
