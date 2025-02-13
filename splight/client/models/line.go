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

	// Validate timezone or geometry (or both equal)
	if err := validateTimezoneOrGeometry(timezone, geometryStr); err != nil {
		return err
	}

	// Validate geometry JSON if it's set
	if geometryStr != "" {
		if err := validateJSONString(geometryStr); err != nil {
			return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
		}
	}

	m.LineParams = LineParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       json.RawMessage(geometryStr),
			CustomTimezone: timezone,
			Tags:           tags,
			Kind:           kind,
		},
	}

	// Handle metadata conversion for LineParams
	// Here, you can follow the same pattern for other metadata fields as shown in your original code

	return nil
}

func (m *Line) ToSchema(d *schema.ResourceData) error {
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
