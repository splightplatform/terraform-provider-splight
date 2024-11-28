package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BusParams struct {
	AssetParams
	ActivePower   AssetAttribute `json:"active_power"`
	ReactivePower AssetAttribute `json:"reactive_power"`
	NominalVoltageKV AssetMetadata `json:"nominal_voltage_kv"`
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

	// Validate geometry JSON
	geometryStr := d.Get("geometry").(string)
	if err := validateJSONString(geometryStr); err != nil {
		return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
	}

	m.BusParams = BusParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       json.RawMessage(geometryStr),
			CustomTimezone: d.Get("timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

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
	d.Set("geometry", string(m.AssetParams.Geometry))
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
