package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TransformerParams struct {
	AssetParams
	ActivePowerHV         *AssetAttribute `json:"active_power_hv"`
	ActivePowerLV         *AssetAttribute `json:"active_power_lv"`
	ReactivePowerHV       *AssetAttribute `json:"reactive_power_hv"`
	ReactivePowerLV       *AssetAttribute `json:"reactive_power_lv"`
	ActivePowerLoss       *AssetAttribute `json:"active_power_loss"`
	ReactivePowerLoss     *AssetAttribute `json:"reactive_power_loss"`
	CurrentHV             *AssetAttribute `json:"current_hv"`
	CurrentLV             *AssetAttribute `json:"current_lv"`
	VoltageHV             *AssetAttribute `json:"voltage_hv"`
	VoltageLV             *AssetAttribute `json:"voltage_lv"`
	Contingency           *AssetAttribute `json:"contingency"`
	SwitchStatusHV        *AssetAttribute `json:"switch_status_hv"`
	SwitchStatusLV        *AssetAttribute `json:"switch_status_lv"`
	TapPos                AssetMetadata   `json:"tap_pos"`
	XnOhm                 AssetMetadata   `json:"xn_ohm"`
	StandardType          AssetMetadata   `json:"standard_type"`
	Capacitance           AssetMetadata   `json:"capacitance"`
	Conductance           AssetMetadata   `json:"conductance"`
	MaximumAllowedCurrent AssetMetadata   `json:"maximum_allowed_current"`
	MaximumAllowedPower   AssetMetadata   `json:"maximum_allowed_power"`
	Reactance             AssetMetadata   `json:"reactance"`
	Resistance            AssetMetadata   `json:"resistance"`
	SafetyMarginForPower  AssetMetadata   `json:"safety_margin_for_power"`
}

type Transformer struct {
	TransformerParams
	Id string `json:"id"`
}

func (m *Transformer) GetId() string {
	return m.Id
}

func (m *Transformer) GetParams() Params {
	return &m.TransformerParams
}

func (m *Transformer) ResourcePath() string {
	return "v3/engine/asset/transformers/"
}

func (m *Transformer) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Get values of custom_timezone and geometry
	custom_timezone := d.Get("custom_timezone").(string)
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

	m.TransformerParams = TransformerParams{
		AssetParams: AssetParams{
			Name:              d.Get("name").(string),
			Description:       d.Get("description").(string),
			Geometry:          geometry,
			CustomTimezone:    custom_timezone,
			UseCustomTimezone: custom_timezone != "",
			Tags:              tags,
			Kind:              kind,
		},
	}

	tapPos, err := convertAssetMetadata(d.Get("tap_pos").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid tapPos metadata: %w", err)
	}
	if tapPos.Type == "" {
		tapPos.Type = "String"
	}
	if tapPos.Name == "" {
		tapPos.Name = "tap_pos"
	}
	m.TransformerParams.TapPos = *tapPos

	xnOhm, err := convertAssetMetadata(d.Get("xn_ohm").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid xnOhm metadata: %w", err)
	}
	if xnOhm.Type == "" {
		xnOhm.Type = "String"
	}
	if xnOhm.Name == "" {
		xnOhm.Name = "xn_ohm"
	}
	m.TransformerParams.XnOhm = *xnOhm

	standardType, err := convertAssetMetadata(d.Get("standard_type").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid standardType metadata: %w", err)
	}
	if standardType.Type == "" {
		standardType.Type = "String"
	}
	if standardType.Name == "" {
		standardType.Name = "standard_type"
	}
	m.TransformerParams.StandardType = *standardType

	capacitance, err := convertAssetMetadata(d.Get("capacitance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid capacitance metadata: %w", err)
	}
	if capacitance.Type == "" {
		capacitance.Type = "String"
	}
	if capacitance.Name == "" {
		capacitance.Name = "capacitance"
	}
	m.TransformerParams.Capacitance = *capacitance

	conductance, err := convertAssetMetadata(d.Get("conductance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid conductance metadata: %w", err)
	}
	if conductance.Type == "" {
		conductance.Type = "String"
	}
	if conductance.Name == "" {
		conductance.Name = "conductance"
	}
	m.TransformerParams.Conductance = *conductance

	maximumAllowedCurrent, err := convertAssetMetadata(d.Get("maximum_allowed_current").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximumAllowedCurrent metadata: %w", err)
	}
	if maximumAllowedCurrent.Type == "" {
		maximumAllowedCurrent.Type = "String"
	}
	if maximumAllowedCurrent.Name == "" {
		maximumAllowedCurrent.Name = "maximum_allowed_current"
	}
	m.TransformerParams.MaximumAllowedCurrent = *maximumAllowedCurrent

	maximumAllowedPower, err := convertAssetMetadata(d.Get("maximum_allowed_power").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximumAllowedPower metadata: %w", err)
	}
	if maximumAllowedPower.Type == "" {
		maximumAllowedPower.Type = "String"
	}
	if maximumAllowedPower.Name == "" {
		maximumAllowedPower.Name = "maximum_allowed_power"
	}
	m.TransformerParams.MaximumAllowedPower = *maximumAllowedPower

	reactance, err := convertAssetMetadata(d.Get("reactance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid reactance metadata: %w", err)
	}
	if reactance.Type == "" {
		reactance.Type = "String"
	}
	if reactance.Name == "" {
		reactance.Name = "reactance"
	}
	m.TransformerParams.Reactance = *reactance

	resistance, err := convertAssetMetadata(d.Get("resistance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid resistance metadata: %w", err)
	}
	if resistance.Type == "" {
		resistance.Type = "String"
	}
	if resistance.Name == "" {
		resistance.Name = "resistance"
	}
	m.TransformerParams.Resistance = *resistance

	safetyMarginForPower, err := convertAssetMetadata(d.Get("safety_margin_for_power").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid safetyMarginForPower metadata: %w", err)
	}
	if safetyMarginForPower.Type == "" {
		safetyMarginForPower.Type = "String"
	}
	if safetyMarginForPower.Name == "" {
		safetyMarginForPower.Name = "safety_margin_for_power"
	}
	m.TransformerParams.SafetyMarginForPower = *safetyMarginForPower

	return nil
}

func (m *Transformer) ToSchema(d *schema.ResourceData) error {
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

	d.Set("timezone", m.Timezone)
	d.Set("custom_timezone", m.CustomTimezone)

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

	d.Set("active_power_hv", []map[string]any{
		m.ActivePowerHV.ToMap(),
	})

	d.Set("active_power_lv", []map[string]any{
		m.ActivePowerLV.ToMap(),
	})

	d.Set("reactive_power_hv", []map[string]any{
		m.ReactivePowerHV.ToMap(),
	})

	d.Set("reactive_power_lv", []map[string]any{
		m.ReactivePowerLV.ToMap(),
	})

	d.Set("active_power_loss", []map[string]any{
		m.ActivePowerLoss.ToMap(),
	})

	d.Set("reactive_power_loss", []map[string]any{
		m.ReactivePowerLoss.ToMap(),
	})

	d.Set("current_hv", []map[string]any{
		m.CurrentHV.ToMap(),
	})

	d.Set("current_lv", []map[string]any{
		m.CurrentLV.ToMap(),
	})

	d.Set("voltage_hv", []map[string]any{
		m.VoltageHV.ToMap(),
	})

	d.Set("voltage_lv", []map[string]any{
		m.VoltageLV.ToMap(),
	})

	d.Set("contingency", []map[string]any{
		m.Contingency.ToMap(),
	})

	d.Set("switch_status_hv", []map[string]any{
		m.SwitchStatusHV.ToMap(),
	})

	d.Set("switch_status_lv", []map[string]any{
		m.SwitchStatusLV.ToMap(),
	})

	d.Set("xn_ohm", []map[string]any{
		m.XnOhm.ToMap(),
	})

	d.Set("standard_type", []map[string]any{
		m.StandardType.ToMap(),
	})

	d.Set("capacitance", []map[string]any{
		m.Capacitance.ToMap(),
	})

	d.Set("conductance", []map[string]any{
		m.Conductance.ToMap(),
	})

	d.Set("maximum_allowed_current", []map[string]any{
		m.MaximumAllowedCurrent.ToMap(),
	})

	d.Set("maximum_allowed_power", []map[string]any{
		m.MaximumAllowedPower.ToMap(),
	})

	d.Set("reactance", []map[string]any{
		m.Reactance.ToMap(),
	})

	d.Set("resistance", []map[string]any{
		m.Resistance.ToMap(),
	})

	d.Set("safety_margin_for_power", []map[string]any{
		m.SafetyMarginForPower.ToMap(),
	})

	return nil
}
