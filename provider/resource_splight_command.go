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
	actions := convertQueryFilters(d.Get("actions").(*schema.Set).List())
	assetInterface := d.Get("asset").(*schema.Set).List()

	var asset *client.QueryFilter
	if len(assetInterface) > 0 {
		assetInterface := assetInterface[0].(map[string]interface{})
		asset = &client.QueryFilter{
			Id:   assetInterface["id"].(string),
			Name: assetInterface["name"].(string),
		}
	}

	return &client.CommandParams{
		Name:    d.Get("name").(string),
		Asset:   asset,
		Actions: actions,
	}
}

func convertQueryFilters(queryFiltersInterface []interface{}) []client.QueryFilter {
	queryFilters := make([]client.QueryFilter, len(queryFiltersInterface))

	for i, item := range queryFiltersInterface {
		queryFilter := item.(map[string]interface{})
		queryFilters[i] = client.QueryFilter{
			Id:   queryFilter["id"].(string),
			Name: queryFilter["id"].(string),
		}

	}

	return queryFilters

}

func saveCommandToState(d *schema.ResourceData, command *client.Command) {
	d.SetId(command.ID)

	d.Set("name", command.Name)

	// Remember this is a Set in the schema
	if command.Asset != nil {
		d.Set("asset", []map[string]string{
			{
				"id":   command.Asset.Id,
				"name": command.Asset.Name,
			},
		})
	}

	actionsInterface := make([]map[string]interface{}, len(command.Actions))
	for i, action := range command.Actions {
		actionsInterface[i] = map[string]interface{}{
			"id":   action.Id,
			"name": action.Name,
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
