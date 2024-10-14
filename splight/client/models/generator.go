package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GeneratorParams struct {
	AssetParams
	ActivePower          AssetAttribute `json:"active_power"`
	ReactivePower        AssetAttribute `json:"reactive_power"`
	DailyEnergy          AssetAttribute `json:"daily_energy"`
	DailyEmissionAvoided AssetAttribute `json:"daily_emission_avoided"`
	CO2Coefficient       AssetMetadata  `json:"CO2_coefficient"`
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
	return "v2/engine/asset/generators/"
}

func (m *Generator) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.GeneratorParams = GeneratorParams{
		AssetParams: AssetParams{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Geometry:    json.RawMessage(d.Get("geometry").(string)),
			Tags:        tags,
			Kind:        kind,
		},
	}

	// TODO: remove ALL of these sets when API fixes its contract
	activePower := convertAssetAttribute(d.Get("active_power").(*schema.Set).List())
	if activePower == nil {
		activePower = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "active_power",
			},
		}
	}
	m.GeneratorParams.ActivePower = *activePower

	reactivePower := convertAssetAttribute(d.Get("reactive_power").(*schema.Set).List())
	if reactivePower == nil {
		reactivePower = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "reactive_power",
			},
		}
	}
	m.GeneratorParams.ReactivePower = *reactivePower

	dailyEnergy := convertAssetAttribute(d.Get("daily_energy").(*schema.Set).List())
	if dailyEnergy == nil {
		dailyEnergy = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "daily_energy",
			},
		}
	}
	m.GeneratorParams.DailyEnergy = *dailyEnergy

	dailyEmissionAvoided := convertAssetAttribute(d.Get("daily_emission_avoided").(*schema.Set).List())
	if dailyEmissionAvoided == nil {
		dailyEmissionAvoided = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "daily_emission_avoided",
			},
		}
	}
	m.GeneratorParams.DailyEmissionAvoided = *dailyEmissionAvoided

	CO2_coefficient := convertAssetMetadata(d.Get("co2_coefficient").(*schema.Set).List())
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

	d.Set("active_power", []map[string]any{m.ActivePower.ToMap()})
	d.Set("reactive_power", []map[string]any{m.ReactivePower.ToMap()})
	d.Set("daily_energy", []map[string]any{m.DailyEnergy.ToMap()})
	d.Set("daily_emission_avoided", []map[string]any{m.DailyEmissionAvoided.ToMap()})
	d.Set("co2_coefficient", []map[string]any{m.CO2Coefficient.ToMap()})

	return nil
}
