package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LineParams struct {
	AssetParams
	ActivePower                  AssetAttribute `json:"active_power"`
	ActivePowerEnd               AssetAttribute `json:"active_power_end"`
	Ampacity                     AssetAttribute `json:"ampacity"`
	Current                      AssetAttribute `json:"current"`
	CurrentR                     AssetAttribute `json:"current_r"`
	CurrentS                     AssetAttribute `json:"current_s"`
	CurrentT                     AssetAttribute `json:"current_t"`
	Energy                       AssetAttribute `json:"energy"`
	MaxTemperature               AssetAttribute `json:"max_temperature"`
	ReactivePower                AssetAttribute `json:"reactive_power"`
	VoltageRS                    AssetAttribute `json:"voltage_rs"`
	VoltageST                    AssetAttribute `json:"voltage_st"`
	VoltageTR                    AssetAttribute `json:"voltage_tr"`
	Absorptivity                 AssetMetadata  `json:"absorptivity"`
	Atmosphere                   AssetMetadata  `json:"atmosphere"`
	Capacitance                  AssetMetadata  `json:"capacitance"`
	Conductance                  AssetMetadata  `json:"conductance"`
	Diameter                     AssetMetadata  `json:"diameter"`
	Emissivity                   AssetMetadata  `json:"emissivity"`
	Length                       AssetMetadata  `json:"length"`
	MaximumAllowedCurrent        AssetMetadata  `json:"maximum_allowed_current"`
	MaximumAllowedPower          AssetMetadata  `json:"maximum_allowed_power"`
	MaximumAllowedTemperature    AssetMetadata  `json:"maximum_allowed_temperature"`
	MaximumAllowedTemperatureLTE AssetMetadata  `json:"maximum_allowed_temperature_lte"`
	MaximumAllowedTemperatureSTE AssetMetadata  `json:"maximum_allowed_temperature_ste"`
	NumberOfConductors           AssetMetadata  `json:"number_of_conductors"`
	Reactance                    AssetMetadata  `json:"reactance"`
	ReferenceResistance          AssetMetadata  `json:"reference_resistance"`
	Resistance                   AssetMetadata  `json:"resistance"`
	SafetyMarginForPower         AssetMetadata  `json:"safety_margin_for_power"`
	Susceptance                  AssetMetadata  `json:"susceptance"`
	TemperatureCoeffResistance   AssetMetadata  `json:"temperature_coeff_resistance"`
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

	// Validate geometry JSON
	geometryStr := d.Get("geometry").(string)
	if err := validateJSONString(geometryStr); err != nil {
		return fmt.Errorf("geometry field contains %w", err)
	}

	m.LineParams = LineParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       json.RawMessage(geometryStr),
			CustomTimezone: d.Get("timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

	// Asset Attributes with consistent nil handling
	attributes := map[string]*AssetAttribute{
		"active_power":     convertAssetAttribute(d.Get("active_power").(*schema.Set).List()),
		"active_power_end": convertAssetAttribute(d.Get("active_power_end").(*schema.Set).List()),
		"ampacity":         convertAssetAttribute(d.Get("ampacity").(*schema.Set).List()),
		"current":          convertAssetAttribute(d.Get("current").(*schema.Set).List()),
		"current_r":        convertAssetAttribute(d.Get("current_r").(*schema.Set).List()),
		"current_s":        convertAssetAttribute(d.Get("current_s").(*schema.Set).List()),
		"current_t":        convertAssetAttribute(d.Get("current_t").(*schema.Set).List()),
		"energy":           convertAssetAttribute(d.Get("energy").(*schema.Set).List()),
		"max_temperature":  convertAssetAttribute(d.Get("max_temperature").(*schema.Set).List()),
		"reactive_power":   convertAssetAttribute(d.Get("reactive_power").(*schema.Set).List()),
		"voltage_rs":       convertAssetAttribute(d.Get("voltage_rs").(*schema.Set).List()),
		"voltage_st":       convertAssetAttribute(d.Get("voltage_st").(*schema.Set).List()),
		"voltage_tr":       convertAssetAttribute(d.Get("voltage_tr").(*schema.Set).List()),
	}

	for name, attr := range attributes {
		if attr == nil {
			attr = &AssetAttribute{
				AssetAttributeParams: AssetAttributeParams{
					Type: "Number",
					Name: name,
				},
			}
		}
		switch name {
		case "active_power":
			m.LineParams.ActivePower = *attr
		case "active_power_end":
			m.LineParams.ActivePowerEnd = *attr
		case "ampacity":
			m.LineParams.Ampacity = *attr
		case "current":
			m.LineParams.Current = *attr
		case "current_r":
			m.LineParams.CurrentR = *attr
		case "current_s":
			m.LineParams.CurrentS = *attr
		case "current_t":
			m.LineParams.CurrentT = *attr
		case "energy":
			m.LineParams.Energy = *attr
		case "max_temperature":
			m.LineParams.MaxTemperature = *attr
		case "reactive_power":
			m.LineParams.ReactivePower = *attr
		case "voltage_rs":
			m.LineParams.VoltageRS = *attr
		case "voltage_st":
			m.LineParams.VoltageST = *attr
		case "voltage_tr":
			m.LineParams.VoltageTR = *attr
		}
	}

	// Metadata fields with error handling and type/name defaults
	metadataFields := []struct {
		name        string
		defaultType string
		defaultName string
		destination *AssetMetadata
	}{
		{"diameter", "Number", "diameter", &m.LineParams.Diameter},
		{"absorptivity", "Number", "absorptivity", &m.LineParams.Absorptivity},
		{"atmosphere", "Number", "atmosphere", &m.LineParams.Atmosphere},
		{"capacitance", "Number", "capacitance", &m.LineParams.Capacitance},
		{"conductance", "Number", "conductance", &m.LineParams.Conductance},
		{"emissivity", "Number", "emissivity", &m.LineParams.Emissivity},
		{"length", "Number", "length", &m.LineParams.Length},
		{"maximum_allowed_current", "Number", "maximum_allowed_current", &m.LineParams.MaximumAllowedCurrent},
		{"maximum_allowed_power", "Number", "maximum_allowed_power", &m.LineParams.MaximumAllowedPower},
		{"maximum_allowed_temperature", "Number", "maximum_allowed_temperature", &m.LineParams.MaximumAllowedTemperature},
		{"maximum_allowed_temperature_lte", "Number", "maximum_allowed_temperature_lte", &m.LineParams.MaximumAllowedTemperatureLTE},
		{"maximum_allowed_temperature_ste", "Number", "maximum_allowed_temperature_ste", &m.LineParams.MaximumAllowedTemperatureSTE},
		{"number_of_conductors", "Number", "number_of_conductors", &m.LineParams.NumberOfConductors},
		{"reactance", "Number", "reactance", &m.LineParams.Reactance},
		{"reference_resistance", "Number", "reference_resistance", &m.LineParams.ReferenceResistance},
		{"resistance", "Number", "resistance", &m.LineParams.Resistance},
		{"safety_margin_for_power", "Number", "safety_margin_for_power", &m.LineParams.SafetyMarginForPower},
		{"susceptance", "Number", "susceptance", &m.LineParams.Susceptance},
		{"temperature_coeff_resistance", "Number", "temperature_coeff_resistance", &m.LineParams.TemperatureCoeffResistance},
	}

	for _, field := range metadataFields {
		metadata, err := convertAssetMetadata(d.Get(field.name).(*schema.Set).List())
		if err != nil {
			return fmt.Errorf("invalid %s metadata: %w", field.name, err)
		}
		if metadata.Type == "" {
			metadata.Type = field.defaultType
		}
		if metadata.Name == "" {
			metadata.Name = field.defaultName
		}
		*field.destination = *metadata
	}

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
