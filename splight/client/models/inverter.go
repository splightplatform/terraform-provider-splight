package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type InverterParams struct {
	AssetParams
	AccumulatedEnergy     *AssetAttribute `json:"accumulated_energy"`
	ActivePower           *AssetAttribute `json:"active_power"`
	DailyEnergy           *AssetAttribute `json:"daily_energy"`
	RawDailyEnergy        *AssetAttribute `json:"raw_daily_energy"`
	Temperature           *AssetAttribute `json:"temperature"`
	SwitchStatus          *AssetAttribute `json:"switch_status"`
	Make                  AssetMetadata   `json:"make"`
	Model                 AssetMetadata   `json:"model"`
	SerialNumber          AssetMetadata   `json:"serial_number"`
	MaxActivePower        AssetMetadata   `json:"max_active_power"`
	EnergyMeasurementType AssetMetadata   `json:"energy_measurement_type"`
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
	return "v3/engine/asset/inverters/"
}

func (m *Inverter) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()
	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Validate geometry JSON
	geometryStr := d.Get("geometry").(string)
	if err := validateJSONString(geometryStr); err != nil {
		return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
	}

	geometry := json.RawMessage(geometryStr)
	m.InverterParams = InverterParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       &geometry,
			CustomTimezone: d.Get("timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

	make, err := convertAssetMetadata(d.Get("make").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid make metadata: %w", err)
	}
	if make.Type == "" {
		make.Type = "String"
	}
	if make.Name == "" {
		make.Name = "Make"
	}
	m.InverterParams.Make = *make

	model, err := convertAssetMetadata(d.Get("model").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid model metadata: %w", err)
	}
	if model.Type == "" {
		model.Type = "String"
	}
	if model.Name == "" {
		model.Name = "Model"
	}
	m.InverterParams.Model = *model

	serialNumber, err := convertAssetMetadata(d.Get("serial_number").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid serial number metadata: %w", err)
	}
	if serialNumber.Type == "" {
		serialNumber.Type = "Number"
	}
	if serialNumber.Name == "" {
		serialNumber.Name = "SerialNumber"
	}
	m.InverterParams.SerialNumber = *serialNumber

	maxActivePower, err := convertAssetMetadata(d.Get("max_active_power").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid max active power metadata: %w", err)
	}
	if maxActivePower.Type == "" {
		maxActivePower.Type = "Number"
	}
	if maxActivePower.Name == "" {
		maxActivePower.Name = "MaxActivePower"
	}
	m.InverterParams.MaxActivePower = *maxActivePower

	energyMeasurementType, err := convertAssetMetadata(d.Get("energy_measurement_type").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid energy measurement type metadata: %w", err)
	}
	if energyMeasurementType.Type == "" {
		energyMeasurementType.Type = "String"
	}
	if energyMeasurementType.Name == "" {
		energyMeasurementType.Name = "EnergyMeasurementType"
	}
	m.InverterParams.EnergyMeasurementType = *energyMeasurementType

	return nil
}

func (m *Inverter) ToSchema(d *schema.ResourceData) error {
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

	d.Set("accumulated_energy", []map[string]any{m.AccumulatedEnergy.ToMap()})
	d.Set("daily_energy", []map[string]any{m.DailyEnergy.ToMap()})
	d.Set("raw_daily_energy", []map[string]any{m.RawDailyEnergy.ToMap()})
	d.Set("active_power", []map[string]any{m.ActivePower.ToMap()})
	d.Set("temperature", []map[string]any{m.Temperature.ToMap()})
	d.Set("switch_status", []map[string]any{m.SwitchStatus.ToMap()})
	d.Set("make", []map[string]any{m.Make.ToMap()})
	d.Set("model", []map[string]any{m.Model.ToMap()})
	d.Set("serial_number", []map[string]any{m.SerialNumber.ToMap()})
	d.Set("max_active_power", []map[string]any{m.MaxActivePower.ToMap()})
	d.Set("energy_measurement_type", []map[string]any{m.EnergyMeasurementType.ToMap()})

	return nil
}
