package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceComponent() *schema.Resource {
	return &schema.Resource{
		Schema: schemaComponent(),
		Create: resourceCreateComponent,
		Read:   resourceReadComponent,
		Update: resourceUpdateComponent,
		Delete: resourceDeleteComponent,
		Exists: resourceExistsComponent,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ToComponent(d *schema.ResourceData) *client.ComponentParams {
	componentInputInterface := d.Get("input").(*schema.Set).List()
	componentInputInterfaceList := make([]map[string]interface{}, len(componentInputInterface))
	for i, componentInputInterfaceItem := range componentInputInterface {
		componentInputInterfaceList[i] = componentInputInterfaceItem.(map[string]interface{})
	}
	componentInput := make([]client.ComponentInputParam, len(componentInputInterfaceList))
	for i, componentInputItem := range componentInputInterfaceList {
		componentInput[i] = client.ComponentInputParam{
			Name:        componentInputItem["name"].(string),
			Description: componentInputItem["description"].(string),
			Multiple:    componentInputItem["multiple"].(bool),
			Required:    componentInputItem["required"].(bool),
			Sensitive:   componentInputItem["sensitive"].(bool),
			Type:        componentInputItem["type"].(string),
		}
		if componentInputItem["value"] != nil && componentInputItem["value"] != "" {
			value := json.RawMessage(componentInputItem["value"].(string))
			componentInput[i].Value = &value
		}
	}

	return &client.ComponentParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       componentInput,
	}
}

func resourceCreateComponent(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	item := ToComponent(d)

	createdComponent, err := apiClient.CreateComponent(item)
	if err != nil {
		return err
	}

	d.SetId(createdComponent.ID)
	d.Set("name", createdComponent.Name)
	d.Set("description", createdComponent.Description)
	d.Set("version", createdComponent.Version)

	// We need to initialize the memory for nested elements
	// Needed because d.Set() can not handle properly json.RawMessage
	if _, ok := d.GetOk("input"); !ok {
		d.Set("input", []interface{}{})
	}

	inputInterface := make([]map[string]interface{}, len(createdComponent.Input))
	for i, inputItem := range createdComponent.Input {
		inputInterface[i] = map[string]interface{}{
			"name":        inputItem.Name,
			"description": inputItem.Description,
			"multiple":    inputItem.Multiple,
			"required":    inputItem.Required,
			"sensitive":   inputItem.Sensitive,
			"type":        inputItem.Type,
			"value":       inputItem.Value,
		}
	}
	d.Set("input", inputInterface)

	return nil
}

func resourceUpdateComponent(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := ToComponent(d)

	updatedComponent, err := apiClient.UpdateComponent(itemId, item)
	if err != nil {
		return err
	}

	d.Set("name", updatedComponent.Name)
	d.Set("description", updatedComponent.Description)
	d.Set("version", updatedComponent.Version)
	d.Set("input", updatedComponent.Input)
	return nil
}

func resourceReadComponent(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedComponent, err := apiClient.RetrieveComponent(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Component with ID %s", itemId)
		}
	}

	d.SetId(retrievedComponent.ID)
	d.Set("name", retrievedComponent.Name)
	d.Set("description", retrievedComponent.Description)
	d.Set("version", retrievedComponent.Version)
	d.Set("input", retrievedComponent.Input)
	return nil
}

func resourceDeleteComponent(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteComponent(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsComponent(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveComponent(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
