package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Setpoint struct {
	Id        string          `json:"id,omitempty"`
	Name      string          `json:"name"`
	Value     json.RawMessage `json:"value"`
	Attribute QueryFilter     `json:"attribute"`
}

type ActionParams struct {
	Asset     QueryFilter `json:"asset"`
	Name      string      `json:"name"`
	Setpoints []Setpoint  `json:"setpoints"`
}

type Action struct {
	ActionParams
	Id string `json:"id"`
}

func (m *Action) GetId() string {
	return m.Id
}

func (m *Action) GetParams() Params {
	return &m.ActionParams
}

func (m *Action) ResourcePath() string {
	return "v3/engine/asset/actions/"
}

func (m *Action) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	setpoints, err := convertSetpoints(d.Get("setpoints").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("error converting action: %w", err)
	}

	asset := convertSingleQueryFilter(d.Get("asset").(*schema.Set).List())

	m.ActionParams = ActionParams{
		Name:      d.Get("name").(string),
		Asset:     *asset,
		Setpoints: setpoints,
	}

	return nil
}

func convertSetpoints(setpointsInterface []interface{}) ([]Setpoint, error) {
	setpoints := make([]Setpoint, len(setpointsInterface))

	for i, item := range setpointsInterface {

		setpoint := item.(map[string]interface{})

		// Validate value JSON
		valueStr := setpoint["value"].(string)
		if err := validateJSONString(valueStr); err != nil {
			return nil, fmt.Errorf("setpoint value must be JSON encoded")
		}

		attribute := convertSingleQueryFilter(setpoint["attribute"].(*schema.Set).List())
		setpoints[i] = Setpoint{
			Id:        setpoint["id"].(string),
			Name:      "setpoint",
			Value:     json.RawMessage(valueStr),
			Attribute: *attribute,
		}

	}

	return setpoints, nil
}

func (m *Action) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)

	// Remember this is a Set in the schema.
	// This is always set.
	d.Set("asset", []map[string]string{
		{
			"id":   m.Asset.Id,
			"name": m.Asset.Name,
		},
	})

	setpointsInterface := make([]map[string]interface{}, len(m.Setpoints))
	for i, setpoint := range m.Setpoints {
		setpointsInterface[i] = map[string]interface{}{
			"id":    setpoint.Id,
			"name":  setpoint.Name,
			"value": string(setpoint.Value),
			"attribute": []map[string]string{
				{
					"id":   setpoint.Attribute.Id,
					"name": setpoint.Attribute.Name,
				},
			},
		}
	}
	d.Set("setpoints", setpointsInterface)

	return nil
}
