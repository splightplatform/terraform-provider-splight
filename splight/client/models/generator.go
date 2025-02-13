package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GeneratorParams struct {
	AssetParams
	ActivePower          *AssetAttribute `json:"active_power"`
	ReactivePower        *AssetAttribute `json:"reactive_power"`
	DailyEnergy          *AssetAttribute `json:"daily_energy"`
	DailyEmissionAvoided *AssetAttribute `json:"daily_emission_avoided"`
	MonthlyEnergy        *AssetAttribute `json:"monthly_energy"`
	CO2Coefficient       AssetMetadata   `json:"CO2_coefficient"`
}

type Generator struct {
	GeneratorParams
	Id string `json:"id"`
}

func (m *Generator) GetId() string {
	return m.Id
}

func (m *Generator) GetParams() Params {
	return &m.GeneratorParams
}

func (m *Generator) ResourcePath() string {
	return "v3/engine/asset/generators/"
}

func (m *Generator) FromSchema(d *schema.ResourceData) error {
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

	m.GeneratorParams = GeneratorParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       geometry,
			CustomTimezone: timezone,
			Tags:           tags,
			Kind:           kind,
		},
	}

	CO2_coefficient, err := convertAssetMetadata(d.Get("co2_coefficient").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid CO2 coefficient metadata: %w", err)
	}
	if CO2_coefficient.Type == "" {
		CO2_coefficient.Type = "Number"
	}
	if CO2_coefficient.Name == "" {
		CO2_coefficient.Name = "co2_coefficient"
	}
	m.GeneratorParams.CO2Coefficient = *CO2_coefficient

	return nil
}

func (m *Generator) ToSchema(d *schema.ResourceData) error {
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

	d.Set("active_power", []map[string]any{m.ActivePower.ToMap()})
	d.Set("reactive_power", []map[string]any{m.ReactivePower.ToMap()})
	d.Set("daily_energy", []map[string]any{m.DailyEnergy.ToMap()})
	d.Set("daily_emission_avoided", []map[string]any{m.DailyEmissionAvoided.ToMap()})
	d.Set("monthly_energy", []map[string]any{m.MonthlyEnergy.ToMap()})
	d.Set("co2_coefficient", []map[string]any{m.CO2Coefficient.ToMap()})

	return nil
}
