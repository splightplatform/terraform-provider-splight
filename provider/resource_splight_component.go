package provider

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceComponent() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaComponent(),
		CreateContext: resourceCreateComponent,
		ReadContext:   resourceReadComponent,
		UpdateContext: resourceUpdateComponent,
		DeleteContext: resourceDeleteComponent,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toComponent(d *schema.ResourceData) *client.ComponentParams {
	// Convert component inputs
	componentInput := convertComponentInputs(d.Get("input").(*schema.Set).List())

	// Create the ComponentParams object
	item := client.ComponentParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       componentInput,
	}

	return &item
}

func convertComponentInputs(inputInterface []interface{}) []client.ComponentInputParam {
	componentInputs := make([]client.ComponentInputParam, len(inputInterface))
	for i, item := range inputInterface {
		inputItem := item.(map[string]interface{})
		componentInputs[i] = client.ComponentInputParam{
			Name:        inputItem["name"].(string),
			Description: inputItem["description"].(string),
			Multiple:    inputItem["multiple"].(bool),
			Required:    inputItem["required"].(bool),
			Sensitive:   inputItem["sensitive"].(bool),
			Type:        inputItem["type"].(string),
		}
		if inputItem["value"] != nil && inputItem["value"] != "" {
			value := json.RawMessage(inputItem["value"].(string))
			componentInputs[i].Value = &value
		}
	}
	return componentInputs
}

func saveComponentToState(d *schema.ResourceData, component *client.Component) {
	d.SetId(component.ID)

	d.Set("name", component.Name)
	d.Set("description", component.Description)
	d.Set("version", component.Version)
	d.Set("input", component.Input)
}

func resourceCreateComponent(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toComponent(d)

	createdComponent, err := apiClient.CreateComponent(item)

	if err != nil {
		return diag.Errorf("error creating Component: %s", err.Error())
	}

	saveComponentToState(d, createdComponent)

	return nil
}

func resourceUpdateComponent(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := toComponent(d)

	updatedComponent, err := apiClient.UpdateComponent(itemId, item)

	if err != nil {
		return diag.Errorf("error updating Component with ID '%s': %s", itemId, err.Error())
	}

	saveComponentToState(d, updatedComponent)

	return nil
}

func resourceReadComponent(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedComponent, err := apiClient.RetrieveComponent(itemId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading Component with ID '%s': %s", itemId, err.Error())
		}
	}

	saveComponentToState(d, retrievedComponent)

	return nil
}

func resourceDeleteComponent(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteComponent(itemId)

	if err != nil {
		return diag.Errorf("error deleting Component with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
