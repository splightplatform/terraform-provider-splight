package models

import (
	"encoding/json"

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

	m.LineParams = LineParams{
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
	m.LineParams.ActivePower = *activePower

	activePowerEnd := convertAssetAttribute(d.Get("active_power_end").(*schema.Set).List())
	if activePowerEnd == nil {
		activePowerEnd = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "active_power_end",
			},
		}
	}
	m.LineParams.ActivePowerEnd = *activePowerEnd

	ampacity := convertAssetAttribute(d.Get("ampacity").(*schema.Set).List())
	if ampacity == nil {
		ampacity = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "ampacity",
			},
		}
	}
	m.LineParams.Ampacity = *ampacity

	current := convertAssetAttribute(d.Get("current").(*schema.Set).List())
	if current == nil {
		current = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "current",
			},
		}
	}
	m.LineParams.Current = *current

	currentR := convertAssetAttribute(d.Get("current_r").(*schema.Set).List())
	if currentR == nil {
		currentR = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "current_r",
			},
		}
	}
	m.LineParams.CurrentR = *currentR

	currentS := convertAssetAttribute(d.Get("current_s").(*schema.Set).List())
	if currentS == nil {
		currentS = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "current_s",
			},
		}
	}
	m.LineParams.CurrentS = *currentS

	currentT := convertAssetAttribute(d.Get("current_t").(*schema.Set).List())
	if currentT == nil {
		currentT = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "current_t",
			},
		}
	}
	m.LineParams.CurrentT = *currentT

	energy := convertAssetAttribute(d.Get("energy").(*schema.Set).List())
	if energy == nil {
		energy = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "energy",
			},
		}
	}
	m.LineParams.Energy = *energy

	maxTemperature := convertAssetAttribute(d.Get("max_temperature").(*schema.Set).List())
	if maxTemperature == nil {
		maxTemperature = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "max_temperature",
			},
		}
	}
	m.LineParams.MaxTemperature = *maxTemperature

	reactivePower := convertAssetAttribute(d.Get("reactive_power").(*schema.Set).List())
	if reactivePower == nil {
		reactivePower = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "reactive_power",
			},
		}
	}
	m.LineParams.ReactivePower = *reactivePower

	voltageRs := convertAssetAttribute(d.Get("voltage_rs").(*schema.Set).List())
	if voltageRs == nil {
		voltageRs = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "voltage_rs",
			},
		}
	}
	m.LineParams.VoltageRS = *voltageRs

	voltageSt := convertAssetAttribute(d.Get("voltage_st").(*schema.Set).List())
	if voltageSt == nil {
		voltageSt = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "voltage_st",
			},
		}
	}
	m.LineParams.VoltageST = *voltageSt

	voltageTr := convertAssetAttribute(d.Get("voltage_tr").(*schema.Set).List())
	if voltageTr == nil {
		voltageTr = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "voltage_tr",
			},
		}
	}
	m.LineParams.VoltageTR = *voltageTr

	diameter := convertAssetMetadata(d.Get("diameter").(*schema.Set).List())
	if diameter.Type == "" {
		diameter.Type = "Number"
	}
	if diameter.Name == "" {
		diameter.Name = "diameter"
	}
	m.LineParams.Diameter = *diameter

	absorptivity := convertAssetMetadata(d.Get("absorptivity").(*schema.Set).List())
	if absorptivity.Type == "" {
		absorptivity.Type = "Number"
	}
	if absorptivity.Name == "" {
		absorptivity.Name = "absorptivity"
	}
	m.LineParams.Absorptivity = *absorptivity

	atmosphere := convertAssetMetadata(d.Get("atmosphere").(*schema.Set).List())
	if atmosphere.Type == "" {
		atmosphere.Type = "Number"
	}
	if atmosphere.Name == "" {
		atmosphere.Name = "atmosphere"
	}
	m.LineParams.Atmosphere = *atmosphere

	capacitance := convertAssetMetadata(d.Get("capacitance").(*schema.Set).List())
	if capacitance.Type == "" {
		capacitance.Type = "Number"
	}
	if capacitance.Name == "" {
		capacitance.Name = "capacitance"
	}
	m.LineParams.Capacitance = *capacitance

	conductance := convertAssetMetadata(d.Get("conductance").(*schema.Set).List())
	if conductance.Type == "" {
		conductance.Type = "Number"
	}
	if conductance.Name == "" {
		conductance.Name = "conductance"
	}
	m.LineParams.Conductance = *conductance

	emissivity := convertAssetMetadata(d.Get("emissivity").(*schema.Set).List())
	if emissivity.Type == "" {
		emissivity.Type = "Number"
	}
	if emissivity.Name == "" {
		emissivity.Name = "emissivity"
	}
	m.LineParams.Emissivity = *emissivity

	length := convertAssetMetadata(d.Get("length").(*schema.Set).List())
	if length.Type == "" {
		length.Type = "Number"
	}
	if length.Name == "" {
		length.Name = "length"
	}
	m.LineParams.Length = *length

	maximumAllowedCurrent := convertAssetMetadata(d.Get("maximum_allowed_current").(*schema.Set).List())
	if maximumAllowedCurrent.Type == "" {
		maximumAllowedCurrent.Type = "Number"
	}
	if maximumAllowedCurrent.Name == "" {
		maximumAllowedCurrent.Name = "maximum_allowed_current"
	}
	m.LineParams.MaximumAllowedCurrent = *maximumAllowedCurrent

	maximumAllowedPower := convertAssetMetadata(d.Get("maximum_allowed_power").(*schema.Set).List())
	if maximumAllowedPower.Type == "" {
		maximumAllowedPower.Type = "Number"
	}
	if maximumAllowedPower.Name == "" {
		maximumAllowedPower.Name = "maximum_allowed_power"
	}
	m.LineParams.MaximumAllowedPower = *maximumAllowedPower

	maximumAllowedTemperature := convertAssetMetadata(d.Get("maximum_allowed_temperature").(*schema.Set).List())
	if maximumAllowedTemperature.Type == "" {
		maximumAllowedTemperature.Type = "Number"
	}
	if maximumAllowedTemperature.Name == "" {
		maximumAllowedTemperature.Name = "maximum_allowed_temperature"
	}
	m.LineParams.MaximumAllowedTemperature = *maximumAllowedTemperature

	maximumAllowedTemperatureLTE := convertAssetMetadata(d.Get("maximum_allowed_temperature_lte").(*schema.Set).List())
	if maximumAllowedTemperatureLTE.Type == "" {
		maximumAllowedTemperatureLTE.Type = "Number"
	}
	if maximumAllowedTemperatureLTE.Name == "" {
		maximumAllowedTemperatureLTE.Name = "maximum_allowed_temperature_lte"
	}
	m.LineParams.MaximumAllowedTemperatureLTE = *maximumAllowedTemperatureLTE

	maximumAllowedTemperatureSTE := convertAssetMetadata(d.Get("maximum_allowed_temperature_ste").(*schema.Set).List())
	if maximumAllowedTemperatureSTE.Type == "" {
		maximumAllowedTemperatureSTE.Type = "Number"
	}
	if maximumAllowedTemperatureSTE.Name == "" {
		maximumAllowedTemperatureSTE.Name = "maximum_allowed_temperature_ste"
	}
	m.LineParams.MaximumAllowedTemperatureSTE = *maximumAllowedTemperatureSTE

	numberOfConductors := convertAssetMetadata(d.Get("number_of_conductors").(*schema.Set).List())
	if numberOfConductors.Type == "" {
		numberOfConductors.Type = "Number"
	}
	if numberOfConductors.Name == "" {
		numberOfConductors.Name = "number_of_conductors"
	}
	m.LineParams.NumberOfConductors = *numberOfConductors

	reactance := convertAssetMetadata(d.Get("reactance").(*schema.Set).List())
	if reactance.Type == "" {
		reactance.Type = "Number"
	}
	if reactance.Name == "" {
		reactance.Name = "reactance"
	}
	m.LineParams.Reactance = *reactance

	referenceResistance := convertAssetMetadata(d.Get("reference_resistance").(*schema.Set).List())
	if referenceResistance.Type == "" {
		referenceResistance.Type = "Number"
	}
	if referenceResistance.Name == "" {
		referenceResistance.Name = "reference_resistance"
	}
	m.LineParams.ReferenceResistance = *referenceResistance

	resistance := convertAssetMetadata(d.Get("resistance").(*schema.Set).List())
	if resistance.Type == "" {
		resistance.Type = "Number"
	}
	if resistance.Name == "" {
		resistance.Name = "resistance"
	}
	m.LineParams.Resistance = *resistance

	safetyMarginForPower := convertAssetMetadata(d.Get("safety_margin_for_power").(*schema.Set).List())
	if safetyMarginForPower.Type == "" {
		safetyMarginForPower.Type = "Number"
	}
	if safetyMarginForPower.Name == "" {
		safetyMarginForPower.Name = "safety_margin_for_power"
	}
	m.LineParams.SafetyMarginForPower = *safetyMarginForPower

	susceptance := convertAssetMetadata(d.Get("susceptance").(*schema.Set).List())
	if susceptance.Type == "" {
		susceptance.Type = "Number"
	}
	if susceptance.Name == "" {
		susceptance.Name = "susceptance"
	}
	m.LineParams.Susceptance = *susceptance

	temperatureCoeffResistance := convertAssetMetadata(d.Get("temperature_coeff_resistance").(*schema.Set).List())
	if temperatureCoeffResistance.Type == "" {
		temperatureCoeffResistance.Type = "Number"
	}
	if temperatureCoeffResistance.Name == "" {
		temperatureCoeffResistance.Name = "temperature_coeff_resistance"
	}
	m.LineParams.TemperatureCoeffResistance = *temperatureCoeffResistance

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
