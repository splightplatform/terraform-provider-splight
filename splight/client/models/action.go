package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Setpoint struct {
	ID        string          `json:"id,omitempty"`
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
	ID string `json:"id"`
}

func (m *Action) GetID() string {
	return m.ID
}

func (m *Action) GetParams() Params {
	return &m.ActionParams
}

func (m *Action) ResourcePath() string {
	return "v2/engine/asset/actions/"
}

func (m *Action) FromSchema(d *schema.ResourceData) error {
	setpoints := convertSetpoints(d.Get("setpoints").(*schema.Set).List())
	asset := d.Get("asset").(*schema.Set).List()[0].(map[string]interface{})

	m.ActionParams = ActionParams{
		Name: d.Get("name").(string),
		Asset: QueryFilter{
			Id:   asset["id"].(string),
			Name: asset["name"].(string),
		},
		Setpoints: setpoints,
	}

	return nil
}

func convertSetpoints(setpointsInterface []interface{}) []Setpoint {
	setpoints := make([]Setpoint, len(setpointsInterface))

	for i, item := range setpointsInterface {
		setpoint := item.(map[string]interface{})
		attribute := setpoint["attribute"].(*schema.Set).List()[0].(map[string]interface{})
		setpoints[i] = Setpoint{
			ID:    setpoint["id"].(string),
			Name:  "setpoint",
			Value: json.RawMessage(setpoint["value"].(string)),
			Attribute: QueryFilter{
				Id:   attribute["id"].(string),
				Name: attribute["name"].(string),
			},
		}

	}

	return setpoints
}

func (m *Action) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)

	d.Set("name", m.Name)

	// Remember this is a Set in the schema
	d.Set("asset", []map[string]string{
		{
			"id":   m.Asset.Id,
			"name": m.Asset.Name,
		},
	})

	setpointsInterface := make([]map[string]interface{}, len(m.Setpoints))
	for i, setpoint := range m.Setpoints {
		setpointsInterface[i] = map[string]interface{}{
			"id":    setpoint.ID,
			"name":  setpoint.Name,
			"value": setpoint.Value,
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
