package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LineParams struct {
	AssetParams
	ActivePower                  *AssetAttribute `json:"active_power"`
	ActivePowerEnd               *AssetAttribute `json:"active_power_end"`
	Ampacity                     *AssetAttribute `json:"ampacity"`
	Current                      *AssetAttribute `json:"current"`
	CurrentR                     *AssetAttribute `json:"current_r"`
	CurrentS                     *AssetAttribute `json:"current_s"`
	CurrentT                     *AssetAttribute `json:"current_t"`
	Energy                       *AssetAttribute `json:"energy"`
	MaxTemperature               *AssetAttribute `json:"max_temperature"`
	ReactivePower                *AssetAttribute `json:"reactive_power"`
	VoltageRS                    *AssetAttribute `json:"voltage_rs"`
	VoltageST                    *AssetAttribute `json:"voltage_st"`
	VoltageTR                    *AssetAttribute `json:"voltage_tr"`
	Contingency                  *AssetAttribute `json:"contingency"`
	SwitchStatusStart            *AssetAttribute `json:"switch_status_start"`
	SwitchStatusEnd              *AssetAttribute `json:"switch_status_end"`
	Absorptivity                 AssetMetadata   `json:"absorptivity"`
	Atmosphere                   AssetMetadata   `json:"atmosphere"`
	Capacitance                  AssetMetadata   `json:"capacitance"`
	Conductance                  AssetMetadata   `json:"conductance"`
	Diameter                     AssetMetadata   `json:"diameter"`
	Emissivity                   AssetMetadata   `json:"emissivity"`
	Length                       AssetMetadata   `json:"length"`
	MaximumAllowedCurrent        AssetMetadata   `json:"maximum_allowed_current"`
	MaximumAllowedPower          AssetMetadata   `json:"maximum_allowed_power"`
	MaximumAllowedTemperature    AssetMetadata   `json:"maximum_allowed_temperature"`
	MaximumAllowedTemperatureLTE AssetMetadata   `json:"maximum_allowed_temperature_lte"`
	MaximumAllowedTemperatureSTE AssetMetadata   `json:"maximum_allowed_temperature_ste"`
	NumberOfConductors           AssetMetadata   `json:"number_of_conductors"`
	Reactance                    AssetMetadata   `json:"reactance"`
	ReferenceResistance          AssetMetadata   `json:"reference_resistance"`
	Resistance                   AssetMetadata   `json:"resistance"`
	SafetyMarginForPower         AssetMetadata   `json:"safety_margin_for_power"`
	Susceptance                  AssetMetadata   `json:"susceptance"`
	TemperatureCoeffResistance   AssetMetadata   `json:"temperature_coeff_resistance"`
	SpecificHeat                 AssetMetadata   `json:"specific_heat"`
	ConductorMass                AssetMetadata   `json:"conductor_mass"`
	ThermalElongationCoef        AssetMetadata   `json:"thermal_elongation_coef"`
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
	return "v3/engine/asset/lines/"
}

func (m *Line) FromSchema(d *schema.ResourceData) error {
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

	m.LineParams = LineParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       geometry,
			CustomTimezone: timezone,
			Tags:           tags,
			Kind:           kind,
		},
	}

	diameter, err := convertAssetMetadata(d.Get("diameter").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid diameter metadata: %w", err)
	}
	if diameter.Type == "" {
		diameter.Type = "Number"
	}
	if diameter.Name == "" {
		diameter.Name = "diameter"
	}
	m.LineParams.Diameter = *diameter

	absorptivity, err := convertAssetMetadata(d.Get("absorptivity").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid absorptivity metadata: %w", err)
	}
	if absorptivity.Type == "" {
		absorptivity.Type = "Number"
	}
	if absorptivity.Name == "" {
		absorptivity.Name = "absorptivity"
	}
	m.LineParams.Absorptivity = *absorptivity

	atmosphere, err := convertAssetMetadata(d.Get("atmosphere").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid atmosphere metadata: %w", err)
	}
	if atmosphere.Type == "" {
		atmosphere.Type = "Number"
	}
	if atmosphere.Name == "" {
		atmosphere.Name = "atmosphere"
	}
	m.LineParams.Atmosphere = *atmosphere

	capacitance, err := convertAssetMetadata(d.Get("capacitance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid capacitance metadata: %w", err)
	}
	if capacitance.Type == "" {
		capacitance.Type = "Number"
	}
	if capacitance.Name == "" {
		capacitance.Name = "capacitance"
	}
	m.LineParams.Capacitance = *capacitance

	conductance, err := convertAssetMetadata(d.Get("conductance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid conductance metadata: %w", err)
	}
	if conductance.Type == "" {
		conductance.Type = "Number"
	}
	if conductance.Name == "" {
		conductance.Name = "conductance"
	}
	m.LineParams.Conductance = *conductance

	emissivity, err := convertAssetMetadata(d.Get("emissivity").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid emissivity metadata: %w", err)
	}
	if emissivity.Type == "" {
		emissivity.Type = "Number"
	}
	if emissivity.Name == "" {
		emissivity.Name = "emissivity"
	}
	m.LineParams.Emissivity = *emissivity

	length, err := convertAssetMetadata(d.Get("length").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid length metadata: %w", err)
	}
	if length.Type == "" {
		length.Type = "Number"
	}
	if length.Name == "" {
		length.Name = "length"
	}
	m.LineParams.Length = *length

	maximumAllowedCurrent, err := convertAssetMetadata(d.Get("maximum_allowed_current").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximum_allowed_current metadata: %w", err)
	}
	if maximumAllowedCurrent.Type == "" {
		maximumAllowedCurrent.Type = "Number"
	}
	if maximumAllowedCurrent.Name == "" {
		maximumAllowedCurrent.Name = "maximum_allowed_current"
	}
	m.LineParams.MaximumAllowedCurrent = *maximumAllowedCurrent

	maximumAllowedPower, err := convertAssetMetadata(d.Get("maximum_allowed_power").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximum_allowed_power metadata: %w", err)
	}
	if maximumAllowedPower.Type == "" {
		maximumAllowedPower.Type = "Number"
	}
	if maximumAllowedPower.Name == "" {
		maximumAllowedPower.Name = "maximum_allowed_power"
	}
	m.LineParams.MaximumAllowedPower = *maximumAllowedPower

	maximumAllowedTemperature, err := convertAssetMetadata(d.Get("maximum_allowed_temperature").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximum_allowed_temperature metadata: %w", err)
	}
	if maximumAllowedTemperature.Type == "" {
		maximumAllowedTemperature.Type = "Number"
	}
	if maximumAllowedTemperature.Name == "" {
		maximumAllowedTemperature.Name = "maximum_allowed_temperature"
	}
	m.LineParams.MaximumAllowedTemperature = *maximumAllowedTemperature

	maximumAllowedTemperatureLTE, err := convertAssetMetadata(d.Get("maximum_allowed_temperature_lte").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximum_allowed_temperature_lte metadata: %w", err)
	}
	if maximumAllowedTemperatureLTE.Type == "" {
		maximumAllowedTemperatureLTE.Type = "Number"
	}
	if maximumAllowedTemperatureLTE.Name == "" {
		maximumAllowedTemperatureLTE.Name = "maximum_allowed_temperature_lte"
	}
	m.LineParams.MaximumAllowedTemperatureLTE = *maximumAllowedTemperatureLTE

	maximumAllowedTemperatureSTE, err := convertAssetMetadata(d.Get("maximum_allowed_temperature_ste").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid maximum_allowed_temperature_ste metadata: %w", err)
	}
	if maximumAllowedTemperatureSTE.Type == "" {
		maximumAllowedTemperatureSTE.Type = "Number"
	}
	if maximumAllowedTemperatureSTE.Name == "" {
		maximumAllowedTemperatureSTE.Name = "maximum_allowed_temperature_ste"
	}
	m.LineParams.MaximumAllowedTemperatureSTE = *maximumAllowedTemperatureSTE

	numberOfConductors, err := convertAssetMetadata(d.Get("number_of_conductors").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid number_of_conductors metadata: %w", err)
	}
	if numberOfConductors.Type == "" {
		numberOfConductors.Type = "Number"
	}
	if numberOfConductors.Name == "" {
		numberOfConductors.Name = "number_of_conductors"
	}
	m.LineParams.NumberOfConductors = *numberOfConductors

	reactance, err := convertAssetMetadata(d.Get("reactance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid reactance metadata: %w", err)
	}
	if reactance.Type == "" {
		reactance.Type = "Number"
	}
	if reactance.Name == "" {
		reactance.Name = "reactance"
	}
	m.LineParams.Reactance = *reactance

	referenceResistance, err := convertAssetMetadata(d.Get("reference_resistance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid reference_resistance metadata: %w", err)
	}
	if referenceResistance.Type == "" {
		referenceResistance.Type = "Number"
	}
	if referenceResistance.Name == "" {
		referenceResistance.Name = "reference_resistance"
	}
	m.LineParams.ReferenceResistance = *referenceResistance

	resistance, err := convertAssetMetadata(d.Get("resistance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid resistance metadata: %w", err)
	}
	if resistance.Type == "" {
		resistance.Type = "Number"
	}
	if resistance.Name == "" {
		resistance.Name = "resistance"
	}
	m.LineParams.Resistance = *resistance

	safetyMarginForPower, err := convertAssetMetadata(d.Get("safety_margin_for_power").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid safety_margin_for_power metadata: %w", err)
	}
	if safetyMarginForPower.Type == "" {
		safetyMarginForPower.Type = "Number"
	}
	if safetyMarginForPower.Name == "" {
		safetyMarginForPower.Name = "safety_margin_for_power"
	}
	m.LineParams.SafetyMarginForPower = *safetyMarginForPower

	susceptance, err := convertAssetMetadata(d.Get("susceptance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid susceptance metadata: %w", err)
	}
	if susceptance.Type == "" {
		susceptance.Type = "Number"
	}
	if susceptance.Name == "" {
		susceptance.Name = "susceptance"
	}
	m.LineParams.Susceptance = *susceptance

	temperatureCoeffResistance, err := convertAssetMetadata(d.Get("temperature_coeff_resistance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid temperature_coeff_resistance metadata: %w", err)
	}
	if temperatureCoeffResistance.Type == "" {
		temperatureCoeffResistance.Type = "Number"
	}
	if temperatureCoeffResistance.Name == "" {
		temperatureCoeffResistance.Name = "temperature_coeff_resistance"
	}
	m.LineParams.TemperatureCoeffResistance = *temperatureCoeffResistance

	specificHeat, err := convertAssetMetadata(d.Get("specific_heat").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid specific_heat metadata: %w", err)
	}
	if specificHeat.Type == "" {
		specificHeat.Type = "Number"
	}
	if specificHeat.Name == "" {
		specificHeat.Name = "specific_heat"
	}
	m.LineParams.SpecificHeat = *specificHeat

	conductorMass, err := convertAssetMetadata(d.Get("conductor_mass").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid conductor_mass metadata: %w", err)
	}
	if conductorMass.Type == "" {
		conductorMass.Type = "Number"
	}
	if conductorMass.Name == "" {
		conductorMass.Name = "conductor_mass"
	}
	m.LineParams.ConductorMass = *conductorMass

	thermalElongationCoef, err := convertAssetMetadata(d.Get("thermal_elongation_coef").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid thermal_elongation_coef metadata: %w", err)
	}
	if thermalElongationCoef.Type == "" {
		thermalElongationCoef.Type = "Number"
	}
	if thermalElongationCoef.Name == "" {
		thermalElongationCoef.Name = "thermal_elongation_coef"
	}
	m.LineParams.ThermalElongationCoef = *thermalElongationCoef

	return nil
}

func (m *Line) ToSchema(d *schema.ResourceData) error {
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
	d.Set("contingency", []map[string]any{m.Contingency.ToMap()})
	d.Set("switch_status_start", []map[string]any{m.SwitchStatusStart.ToMap()})
	d.Set("switch_status_end", []map[string]any{m.SwitchStatusEnd.ToMap()})
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
	d.Set("specific_heat", []map[string]any{m.SpecificHeat.ToMap()})
	d.Set("conductor_mass", []map[string]any{m.ConductorMass.ToMap()})
	d.Set("thermal_elongation_coef", []map[string]any{m.ThermalElongationCoef.ToMap()})

	return nil
}
