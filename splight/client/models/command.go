package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type CommandParams struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Actions     []Action `json:"actions"`
}

type Command struct {
	CommandParams
	Id string `json:"id"`
}

func (m *Command) GetId() string {
	return m.Id
}

func (m *Command) GetParams() Params {
	return &m.CommandParams
}

func (m *Command) ResourcePath() string {
	return "v2/engine/command/commands/"
}

func (m *Command) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	actions := convertActions(d.Get("actions").(*schema.Set).List())

	m.CommandParams = CommandParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Actions:     actions,
	}

	return nil
}

func convertActions(actionInterface []interface{}) []Action {
	actions := make([]Action, len(actionInterface))

	for i, item := range actionInterface {
		action := item.(map[string]interface{})
		asset := convertSingleQueryFilter(action["asset"].(*schema.Set).List())
		actions[i] = Action{
			Id: action["id"].(string),
			ActionParams: ActionParams{
				Name:  "setpoint",
				Asset: *asset,
			},
		}
	}

	return actions
}

func (m *Command) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)

	actionsInterface := make([]map[string]interface{}, len(m.Actions))
	for i, action := range m.Actions {

		// Remember this is a Set in the schema
		asset := []map[string]string{
			{
				"id":   action.Asset.Id,
				"name": action.Asset.Name,
			},
		}

		actionsInterface[i] = map[string]interface{}{
			"id":    action.Id,
			"name":  action.Name,
			"asset": asset,
		}
	}
	d.Set("setpoints", actionsInterface)

	return nil
}
