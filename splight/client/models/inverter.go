package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type InverterParams struct {
	AssetParams
	AccumulatedEnergy     *AssetAttribute `json:"accumulated_energy,omitempty"`
	ActivePower           *AssetAttribute `json:"active_power,omitempty"`
	DailyEnergy           *AssetAttribute `json:"daily_energy,omitempty"`
	RawDailyEnergy        *AssetAttribute `json:"raw_daily_energy,omitempty"`
	Temperature           *AssetAttribute `json:"temperature,omitempty"`
	Make                  *AssetMetadata  `json:"make,omitempty"`
	Model                 *AssetMetadata  `json:"model,omitempty"`
	SerialNumber          *AssetMetadata  `json:"serial_number,omitempty"`
	MaxActivePower        *AssetMetadata  `json:"max_active_power,omitempty"`
	EnergyMeasurementType *AssetMetadata  `json:"energy_measurement_type,omitempty"`
}

type Inverter struct {
	InverterParams
	Id string `json:"id"`
}

func (m *Inverter) GetId() string {
	return m.Id
}

func (m *Inverter) GetParams() Params {
	return &m.InverterParams
}

func (m *Inverter) ResourcePath() string {
	return "v2/engine/asset/inverters/"
}

func (m *Inverter) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.InverterParams = InverterParams{
		AssetParams: AssetParams{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Geometry:    json.RawMessage(d.Get("geometry").(string)),
			Tags:        tags,
			Kind:        kind,
		},
	}

	make := convertAssetMetadata(d.Get("make").(*schema.Set).List())
	if make != nil {
		make.Type = "String"
		make.Name = "Make"
	}
	m.InverterParams.Make = make

	model := convertAssetMetadata(d.Get("model").(*schema.Set).List())
	if model != nil {
		model.Type = "String"
		model.Name = "Model"
	}
	m.InverterParams.Model = model

	serial_number := convertAssetMetadata(d.Get("serial_number").(*schema.Set).List())
	if serial_number != nil {
		serial_number.Type = "Number"
		serial_number.Name = "SerialNumber"
	}
	m.InverterParams.SerialNumber = serial_number

	max_active_power := convertAssetMetadata(d.Get("max_active_power").(*schema.Set).List())
	if max_active_power != nil {
		max_active_power.Type = "Number"
		max_active_power.Name = "MaxActivePower"
	}
	m.InverterParams.MaxActivePower = max_active_power

	energy_measurement_type := convertAssetMetadata(d.Get("energy_measurement_type").(*schema.Set).List())
	if energy_measurement_type != nil {
		energy_measurement_type.Type = "String"
		energy_measurement_type.Name = "EnergyMeasurementType"
	}
	m.InverterParams.EnergyMeasurementType = energy_measurement_type

	return nil
}

func (m *Inverter) ToSchema(d *schema.ResourceData) error {
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

	d.Set("accumulated_energy", []map[string]any{m.AccumulatedEnergy.ToMap()})
	d.Set("daily_energy", []map[string]any{m.DailyEnergy.ToMap()})
	d.Set("raw_daily_energy", []map[string]any{m.RawDailyEnergy.ToMap()})
	d.Set("active_power", []map[string]any{m.ActivePower.ToMap()})
	d.Set("temperature", []map[string]any{m.Temperature.ToMap()})
	d.Set("make", []map[string]any{m.Make.ToMap()})
	d.Set("model", []map[string]any{m.Model.ToMap()})
	d.Set("serial_number", []map[string]any{m.SerialNumber.ToMap()})
	d.Set("max_active_power", []map[string]any{m.MaxActivePower.ToMap()})
	d.Set("energy_measurement_type", []map[string]any{m.EnergyMeasurementType.ToMap()})

	return nil
}
