package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceCommand() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaCommand(),
		CreateContext: resourceCreateCommand,
		ReadContext:   resourceReadCommand,
		UpdateContext: resourceUpdateCommand,
		DeleteContext: resourceDeleteCommand,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toCommand(d *schema.ResourceData) *client.CommandParams {
	actions := convertActions(d.Get("actions").(*schema.Set).List())

	return &client.CommandParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Actions:     actions,
	}
}

func convertActions(actionInterface []interface{}) []client.Action {
	actions := make([]client.Action, len(actionInterface))

	for i, item := range actionInterface {
		action := item.(map[string]interface{})
		asset := action["asset"].(*schema.Set).List()[0].(map[string]interface{})
		actions[i] = client.Action{
			ID: action["id"].(string),
			ActionParams: client.ActionParams{
				Name: "setpoint",
				Asset: client.QueryFilter{
					Id:   asset["id"].(string),
					Name: asset["name"].(string),
				},
			},
		}
	}

	return actions
}

func saveCommandToState(d *schema.ResourceData, command *client.Command) {
	d.SetId(command.ID)

	d.Set("name", command.Name)

	actionsInterface := make([]map[string]interface{}, len(command.Actions))
	for i, action := range command.Actions {

		// Remember this is a Set in the schema
		asset := []map[string]string{
			{
				"id":   action.Asset.Id,
				"name": action.Asset.Name,
			},
		}

		actionsInterface[i] = map[string]interface{}{
			"id":    action.ID,
			"name":  action.Name,
			"asset": asset,
		}
	}
	d.Set("setpoints", actionsInterface)
}

func resourceCreateCommand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toCommand(d)

	createdCommand, err := apiClient.CreateCommand(item)
	if err != nil {
		return diag.Errorf("error creating Command: %s", err.Error())
	}

	saveCommandToState(d, createdCommand)

	return nil
}

func resourceUpdateCommand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := toCommand(d)

	updatedCommand, err := apiClient.UpdateCommand(itemId, item)
	if err != nil {
		return diag.Errorf("error updating Command with ID '%s': %s", itemId, err.Error())
	}

	saveCommandToState(d, updatedCommand)

	return nil
}

func resourceReadCommand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedCommand, err := apiClient.RetrieveCommand(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading Command with ID '%s': %s", itemId, err.Error())
		}
	}

	saveCommandToState(d, retrievedCommand)

	return nil
}

func resourceDeleteCommand(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteCommand(itemId)
	if err != nil {
		return diag.Errorf("error deleting Command with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
