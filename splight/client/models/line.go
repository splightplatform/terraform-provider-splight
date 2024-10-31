package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LineParams struct {
	AssetParams
	ActivePower                  *AssetAttribute `json:"active_power,omitempty"`
	ActivePowerEnd               *AssetAttribute `json:"active_power_end,omitempty"`
	Ampacity                     *AssetAttribute `json:"ampacity,omitempty"`
	Current                      *AssetAttribute `json:"current,omitempty"`
	CurrentR                     *AssetAttribute `json:"current_r,omitempty"`
	CurrentS                     *AssetAttribute `json:"current_s,omitempty"`
	CurrentT                     *AssetAttribute `json:"current_t,omitempty"`
	Energy                       *AssetAttribute `json:"energy,omitempty"`
	MaxTemperature               *AssetAttribute `json:"max_temperature,omitempty"`
	ReactivePower                *AssetAttribute `json:"reactive_power,omitempty"`
	VoltageRS                    *AssetAttribute `json:"voltage_rs,omitempty"`
	VoltageST                    *AssetAttribute `json:"voltage_st,omitempty"`
	VoltageTR                    *AssetAttribute `json:"voltage_tr,omitempty"`
	Absorptivity                 *AssetMetadata  `json:"absorptivity,omitempty"`
	Atmosphere                   *AssetMetadata  `json:"atmosphere,omitempty"`
	Capacitance                  *AssetMetadata  `json:"capacitance,omitempty"`
	Conductance                  *AssetMetadata  `json:"conductance,omitempty"`
	Diameter                     *AssetMetadata  `json:"diameter,omitempty"`
	Emissivity                   *AssetMetadata  `json:"emissivity,omitempty"`
	Length                       *AssetMetadata  `json:"length,omitempty"`
	MaximumAllowedCurrent        *AssetMetadata  `json:"maximum_allowed_current,omitempty"`
	MaximumAllowedPower          *AssetMetadata  `json:"maximum_allowed_power,omitempty"`
	MaximumAllowedTemperature    *AssetMetadata  `json:"maximum_allowed_temperature,omitempty"`
	MaximumAllowedTemperatureLTE *AssetMetadata  `json:"maximum_allowed_temperature_lte,omitempty"`
	MaximumAllowedTemperatureSTE *AssetMetadata  `json:"maximum_allowed_temperature_ste,omitempty"`
	NumberOfConductors           *AssetMetadata  `json:"number_of_conductors,omitempty"`
	Reactance                    *AssetMetadata  `json:"reactance,omitempty"`
	ReferenceResistance          *AssetMetadata  `json:"reference_resistance,omitempty"`
	Resistance                   *AssetMetadata  `json:"resistance,omitempty"`
	SafetyMarginForPower         *AssetMetadata  `json:"safety_margin_for_power,omitempty"`
	Susceptance                  *AssetMetadata  `json:"susceptance,omitempty"`
	TemperatureCoeffResistance   *AssetMetadata  `json:"temperature_coeff_resistance,omitempty"`
}

type Line struct {
	LineParams
	Id string `json:"id"`
}

func (m *Line) GetId() string {
	return m.Id
}

func (m *Line) GetParams() Params {
	return &m.LineParams
}

func (m *Line) ResourcePath() string {
	return "v2/engine/asset/lines/"
}

func (m *Line) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.LineParams = LineParams{
		AssetParams: AssetParams{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Geometry:    json.RawMessage(d.Get("geometry").(string)),
			Tags:        tags,
			Kind:        kind,
		},
	}

	diameter := convertAssetMetadata(d.Get("diameter").(*schema.Set).List())
	if diameter != nil {
		diameter.Type = "Number"
		diameter.Name = "diameter"
	}
	m.LineParams.Diameter = diameter

	absorptivity := convertAssetMetadata(d.Get("absorptivity").(*schema.Set).List())
	if absorptivity != nil {
		absorptivity.Type = "Number"
		absorptivity.Name = "absorptivity"
	}
	m.LineParams.Absorptivity = absorptivity

	atmosphere := convertAssetMetadata(d.Get("atmosphere").(*schema.Set).List())
	if atmosphere != nil {
		atmosphere.Type = "Number"
		atmosphere.Name = "atmosphere"
	}
	m.LineParams.Atmosphere = atmosphere

	capacitance := convertAssetMetadata(d.Get("capacitance").(*schema.Set).List())
	if capacitance != nil {
		capacitance.Type = "Number"
		capacitance.Name = "capacitance"
	}
	m.LineParams.Capacitance = capacitance

	conductance := convertAssetMetadata(d.Get("conductance").(*schema.Set).List())
	if conductance != nil {
		conductance.Type = "Number"
		conductance.Name = "conductance"
	}
	m.LineParams.Conductance = conductance

	emissivity := convertAssetMetadata(d.Get("emissivity").(*schema.Set).List())
	if emissivity != nil {
		emissivity.Type = "Number"
		emissivity.Name = "emissivity"
	}
	m.LineParams.Emissivity = emissivity

	length := convertAssetMetadata(d.Get("length").(*schema.Set).List())
	if length != nil {
		length.Type = "Number"
		length.Name = "length"
	}
	m.LineParams.Length = length

	maximumAllowedCurrent := convertAssetMetadata(d.Get("maximum_allowed_current").(*schema.Set).List())
	if maximumAllowedCurrent != nil {
		maximumAllowedCurrent.Type = "Number"
		maximumAllowedCurrent.Name = "maximum_allowed_current"
	}
	m.LineParams.MaximumAllowedCurrent = maximumAllowedCurrent

	maximumAllowedPower := convertAssetMetadata(d.Get("maximum_allowed_power").(*schema.Set).List())
	if maximumAllowedPower != nil {
		maximumAllowedPower.Type = "Number"
		maximumAllowedPower.Name = "maximum_allowed_power"
	}
	m.LineParams.MaximumAllowedPower = maximumAllowedPower

	maximumAllowedTemperature := convertAssetMetadata(d.Get("maximum_allowed_temperature").(*schema.Set).List())
	if maximumAllowedTemperature != nil {
		maximumAllowedTemperature.Type = "Number"
		maximumAllowedTemperature.Name = "maximum_allowed_temperature"
	}
	m.LineParams.MaximumAllowedTemperature = maximumAllowedTemperature

	maximumAllowedTemperatureLTE := convertAssetMetadata(d.Get("maximum_allowed_temperature_lte").(*schema.Set).List())
	if maximumAllowedTemperatureLTE != nil {
		maximumAllowedTemperatureLTE.Type = "Number"
		maximumAllowedTemperatureLTE.Name = "maximum_allowed_temperature_lte"
	}
	m.LineParams.MaximumAllowedTemperatureLTE = maximumAllowedTemperatureLTE

	maximumAllowedTemperatureSTE := convertAssetMetadata(d.Get("maximum_allowed_temperature_ste").(*schema.Set).List())
	if maximumAllowedTemperatureSTE != nil {
		maximumAllowedTemperatureSTE.Type = "Number"
		maximumAllowedTemperatureSTE.Name = "maximum_allowed_temperature_ste"
	}
	m.LineParams.MaximumAllowedTemperatureSTE = maximumAllowedTemperatureSTE

	numberOfConductors := convertAssetMetadata(d.Get("number_of_conductors").(*schema.Set).List())
	if numberOfConductors != nil {
		numberOfConductors.Type = "Number"
		numberOfConductors.Name = "number_of_conductors"
	}
	m.LineParams.NumberOfConductors = numberOfConductors

	reactance := convertAssetMetadata(d.Get("reactance").(*schema.Set).List())
	if reactance != nil {
		reactance.Type = "Number"
		reactance.Name = "reactance"
	}
	m.LineParams.Reactance = reactance

	referenceResistance := convertAssetMetadata(d.Get("reference_resistance").(*schema.Set).List())
	if referenceResistance != nil {
		referenceResistance.Type = "Number"
		referenceResistance.Name = "reference_resistance"
	}
	m.LineParams.ReferenceResistance = referenceResistance

	resistance := convertAssetMetadata(d.Get("resistance").(*schema.Set).List())
	if resistance != nil {
		resistance.Type = "Number"
		resistance.Name = "resistance"
	}
	m.LineParams.Resistance = resistance

	safetyMarginForPower := convertAssetMetadata(d.Get("safety_margin_for_power").(*schema.Set).List())
	if safetyMarginForPower != nil {
		safetyMarginForPower.Type = "Number"
		safetyMarginForPower.Name = "safety_margin_for_power"
	}
	m.LineParams.SafetyMarginForPower = safetyMarginForPower

	susceptance := convertAssetMetadata(d.Get("susceptance").(*schema.Set).List())
	if susceptance != nil {
		susceptance.Type = "Number"
		susceptance.Name = "susceptance"
	}
	m.LineParams.Susceptance = susceptance

	temperatureCoeffResistance := convertAssetMetadata(d.Get("temperature_coeff_resistance").(*schema.Set).List())
	if temperatureCoeffResistance != nil {
		temperatureCoeffResistance.Type = "Number"
		temperatureCoeffResistance.Name = "temperature_coeff_resistance"
	}
	m.LineParams.TemperatureCoeffResistance = temperatureCoeffResistance

	return nil
}

func (m *Line) ToSchema(d *schema.ResourceData) error {
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
	d.Set("active_power_end", []map[string]any{m.ActivePowerEnd.ToMap()})
	d.Set("ampacity", []map[string]any{m.Ampacity.ToMap()})
	d.Set("current", []map[string]any{m.Current.ToMap()})
	d.Set("current_r", []map[string]any{m.CurrentR.ToMap()})
	d.Set("current_s", []map[string]any{m.CurrentS.ToMap()})
	d.Set("current_t", []map[string]any{m.CurrentT.ToMap()})
	d.Set("energy", []map[string]any{m.Energy.ToMap()})
	d.Set("max_temperature", []map[string]any{m.MaxTemperature.ToMap()})
	d.Set("reactive_power", []map[string]any{m.ReactivePower.ToMap()})
	d.Set("voltage_rs", []map[string]any{m.VoltageRS.ToMap()})
	d.Set("voltage_st", []map[string]any{m.VoltageST.ToMap()})
	d.Set("voltage_tr", []map[string]any{m.VoltageTR.ToMap()})
	d.Set("diameter", []map[string]any{m.Diameter.ToMap()})
	d.Set("absorptivity", []map[string]any{m.Absorptivity.ToMap()})
	d.Set("atmosphere", []map[string]any{m.Atmosphere.ToMap()})
	d.Set("capacitance", []map[string]any{m.Capacitance.ToMap()})
	d.Set("conductance", []map[string]any{m.Conductance.ToMap()})
	d.Set("emissivity", []map[string]any{m.Emissivity.ToMap()})
	d.Set("length", []map[string]any{m.Length.ToMap()})
	d.Set("maximum_allowed_current", []map[string]any{m.MaximumAllowedCurrent.ToMap()})
	d.Set("maximum_allowed_power", []map[string]any{m.MaximumAllowedPower.ToMap()})
	d.Set("maximum_allowed_temperature", []map[string]any{m.MaximumAllowedTemperature.ToMap()})
	d.Set("maximum_allowed_temperature_lte", []map[string]any{m.MaximumAllowedTemperatureLTE.ToMap()})
	d.Set("maximum_allowed_temperature_ste", []map[string]any{m.MaximumAllowedTemperatureSTE.ToMap()})
	d.Set("number_of_conductors", []map[string]any{m.NumberOfConductors.ToMap()})
	d.Set("reactance", []map[string]any{m.Reactance.ToMap()})
	d.Set("reference_resistance", []map[string]any{m.ReferenceResistance.ToMap()})
	d.Set("resistance", []map[string]any{m.Resistance.ToMap()})
	d.Set("safety_margin_for_power", []map[string]any{m.SafetyMarginForPower.ToMap()})
	d.Set("susceptance", []map[string]any{m.Susceptance.ToMap()})
	d.Set("temperature_coeff_resistance", []map[string]any{m.TemperatureCoeffResistance.ToMap()})

	return nil
}
