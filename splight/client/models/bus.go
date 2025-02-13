package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BusParams struct {
	AssetParams
	ActivePower      *AssetAttribute `json:"active_power"`
	ReactivePower    *AssetAttribute `json:"reactive_power"`
	NominalVoltageKV AssetMetadata   `json:"nominal_voltage_kv"`
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
	return "v3/engine/asset/buses/"
}

func (m *Bus) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Get values of timezone and geometry
	timezone := d.Get("timezone").(string)
	geometryStr := d.Get("geometry").(string)

	// Validate geometry JSON if it's set
	if geometryStr != "" {
		if err := validateJSONString(geometryStr); err != nil {
			return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
		}
	}

	// Check if geometryStr is empty and handle accordingly
	var geometry *json.RawMessage
	if geometryStr != "" {
		// Convert string to json.RawMessage
		raw := json.RawMessage(geometryStr)
		geometry = &raw
	}

	m.BusParams = BusParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       geometry,
			CustomTimezone: timezone,
			Tags:           tags,
			Kind:           kind,
		},
	}

	nominalVoltageKV, err := convertAssetMetadata(d.Get("nominal_voltage_kv").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid nominal voltage metadata: %w", err)
	}
	if nominalVoltageKV.Type == "" {
		nominalVoltageKV.Type = "Number"
	}
	if nominalVoltageKV.Name == "" {
		nominalVoltageKV.Name = "nominal_voltage_kv"
	}
	m.BusParams.NominalVoltageKV = *nominalVoltageKV

	return nil
}

func (m *Bus) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)

	var geometryStr string
	if m.Geometry != nil {
		geometryStr = string(*m.Geometry)
	} else {
		geometryStr = ""
	}
	d.Set("geometry", geometryStr)

	d.Set("timezone", m.AssetParams.CustomTimezone)

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

	d.Set("nominal_voltage_kv", []map[string]any{m.NominalVoltageKV.ToMap()})

	return nil
}
