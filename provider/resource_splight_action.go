package provider

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceAction() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaAction(),
		CreateContext: resourceCreateAction,
		ReadContext:   resourceReadAction,
		UpdateContext: resourceUpdateAction,
		DeleteContext: resourceDeleteAction,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toAction(d *schema.ResourceData) *client.ActionParams {
	setpoints := convertSetpoints(d.Get("setpoints").(*schema.Set).List())
	asset := d.Get("asset").(*schema.Set).List()[0].(map[string]interface{})

	return &client.ActionParams{
		Name: d.Get("name").(string),
		Asset: client.QueryFilter{
			Id:   asset["id"].(string),
			Name: asset["name"].(string),
		},
		Setpoints: setpoints,
	}
}

func convertSetpoints(setpointsInterface []interface{}) []client.Setpoint {
	setpoints := make([]client.Setpoint, len(setpointsInterface))

	for i, item := range setpointsInterface {
		setpoint := item.(map[string]interface{})
		attribute := setpoint["attribute"].(*schema.Set).List()[0].(map[string]interface{})
		setpoints[i] = client.Setpoint{
			ID:    setpoint["id"].(string),
			Name:  "setpoint",
			Value: json.RawMessage(setpoint["value"].(string)),
			Attribute: client.QueryFilter{
				Id:   attribute["id"].(string),
				Name: attribute["name"].(string),
			},
		}

	}

	return setpoints
}

func saveActionToState(d *schema.ResourceData, action *client.Action) {
	d.SetId(action.ID)

	d.Set("name", action.Name)

	// Remember this is a Set in the schema
	d.Set("asset", []map[string]string{
		{
			"id":   action.Asset.Id,
			"name": action.Asset.Name,
		},
	})

	setpointsInterface := make([]map[string]interface{}, len(action.Setpoints))
	for i, setpoint := range action.Setpoints {
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

}

func resourceCreateAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toAction(d)

	createdAction, err := apiClient.CreateAction(item)
	if err != nil {
		return diag.Errorf("error creating Action: %s", err.Error())
	}

	saveActionToState(d, createdAction)

	return nil
}

func resourceUpdateAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := toAction(d)

	updatedAction, err := apiClient.UpdateAction(itemId, item)
	if err != nil {
		return diag.Errorf("error updating Action with ID '%s': %s", itemId, err.Error())
	}

	saveActionToState(d, updatedAction)

	return nil
}

func resourceReadAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedAction, err := apiClient.RetrieveAction(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading Action with ID '%s': %s", itemId, err.Error())
		}
	}

	saveActionToState(d, retrievedAction)

	return nil
}

func resourceDeleteAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAction(itemId)
	if err != nil {
		return diag.Errorf("error deleting Action with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
