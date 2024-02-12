package provider

import (
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

func resourceCreateComponent(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	componentInputInterface := d.Get("input").([]interface{})
	componentInputInterfaceList := make([]map[string]interface{}, len(componentInputInterface))
	for i, componentInputInterfaceItem := range componentInputInterface {
		componentInputInterfaceList[i] = componentInputInterfaceItem.(map[string]interface{})
	}
	componentInput := make([]client.ComponentInputParam, len(componentInputInterfaceList))
	for i, componentInputItem := range componentInputInterfaceList {
		componentInput[i] = client.ComponentInputParam{
			Name:        componentInputItem["name"].(string),
			Description: componentInputItem["description"].(string),
			Type:        componentInputItem["type"].(string),
			Value:       componentInputItem["value"].(string),
			Multiple:    componentInputItem["multiple"].(bool),
			Required:    componentInputItem["required"].(bool),
			Sensitive:   componentInputItem["sensitive"].(bool),
		}
	}

	item := client.ComponentParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       componentInput,
	}

	createdComponent, err := apiClient.CreateComponent(&item)
	if err != nil {
		return err
	}

	d.SetId(createdComponent.ID)
	d.Set("name", createdComponent.Name)
	d.Set("description", createdComponent.Description)
	d.Set("version", createdComponent.Version)
	d.Set("input", createdComponent.Input)
	return nil
}

func resourceUpdateComponent(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	componentInputInterface := d.Get("input").([]interface{})
	componentInputInterfaceList := make([]map[string]interface{}, len(componentInputInterface))
	for i, componentInputInterfaceItem := range componentInputInterface {
		componentInputInterfaceList[i] = componentInputInterfaceItem.(map[string]interface{})
	}
	componentInput := make([]client.ComponentInputParam, len(componentInputInterfaceList))
	for i, componentInputItem := range componentInputInterfaceList {
		componentInput[i] = client.ComponentInputParam{
			Name:        componentInputItem["name"].(string),
			Description: componentInputItem["description"].(string),
			Type:        componentInputItem["type"].(string),
			Value:       componentInputItem["value"].(string),
			Multiple:    componentInputItem["multiple"].(bool),
			Required:    componentInputItem["required"].(bool),
			Sensitive:   componentInputItem["sensitive"].(bool),
		}
	}

	item := client.ComponentParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Version:     d.Get("version").(string),
		Input:       componentInput,
	}

	updatedComponent, err := apiClient.UpdateComponent(itemId, &item)
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
