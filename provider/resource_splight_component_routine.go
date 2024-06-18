package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceComponentRoutine() *schema.Resource {
	return &schema.Resource{
		Schema: schemaComponentRoutine(),
		Create: resourceCreateComponentRoutine,
		Read:   resourceReadComponentRoutine,
		Update: resourceUpdateComponentRoutine,
		Delete: resourceDeleteComponentRoutine,
		Exists: resourceExistsComponentRoutine,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func ToComponentRoutine(d *schema.ResourceData) *client.ComponentRoutineParams {
	// Handling config
	componentConfigInterface := d.Get("config").(*schema.Set).List()
	componentConfig := make([]client.ComponentRoutineConfigParam, len(componentConfigInterface))
	for i, item := range componentConfigInterface {
		data := item.(map[string]interface{})
		componentConfig[i] = client.ComponentRoutineConfigParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			Sensitive:   data["sensitive"].(bool),
			Type:        data["type"].(string),
		}
		if data["value"] != nil && data["value"] != "" {
			value := json.RawMessage(data["value"].(string))
			componentConfig[i].Value = &value
		}
	}

	// Handling output
	componentOutputInterface := d.Get("output").(*schema.Set).List()
	componentOutput := make([]client.ComponentRoutineIOParam, len(componentOutputInterface))
	for i, item := range componentOutputInterface {
		data := item.(map[string]interface{})
		componentOutput[i] = client.ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
		}
		if data["value"] != nil {
			value := data["value"].(*schema.Set).List()
			for _, value := range value {
				valueData := value.(map[string]interface{})
				componentOutput[i].Value = append(componentOutput[i].Value, client.ComponentRoutineDataAddress{
					Asset:     valueData["asset"].(string),
					Attribute: valueData["attribute"].(string),
				})
			}
		}
	}

	// Handling input
	componentInputInterface := d.Get("input").(*schema.Set).List()
	componentInput := make([]client.ComponentRoutineIOParam, len(componentInputInterface))
	for i, item := range componentInputInterface {
		data := item.(map[string]interface{})
		componentInput[i] = client.ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
		}
		if data["value"] != nil {
			value := data["value"].(*schema.Set).List()
			for _, value := range value {
				valueData := value.(map[string]interface{})
				componentOutput[i].Value = append(componentOutput[i].Value, client.ComponentRoutineDataAddress{
					Asset:     valueData["asset"].(string),
					Attribute: valueData["attribute"].(string),
				})
			}
		}
	}

	item := client.ComponentRoutineParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		ComponentId: d.Get("component_id").(string),
		Config:      componentConfig,
		Input:       componentInput,
		Output:      componentOutput,
	}

	return &item
}

func resourceCreateComponentRoutine(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	item := ToComponentRoutine(d)

	createdComponentRoutine, err := apiClient.CreateComponentRoutine(item)
	if err != nil {
		return err
	}

	d.SetId(createdComponentRoutine.ID)
	d.Set("name", createdComponentRoutine.Name)
	d.Set("description", createdComponentRoutine.Description)
	d.Set("type", createdComponentRoutine.Type)
	d.Set("component_id", createdComponentRoutine.ComponentId)

	// We need to initialize the memory for nested elements
	// Needed because d.Set() can not handle properly json.RawMessage
	if _, ok := d.GetOk("config"); !ok {
		d.Set("config", []interface{}{})
	}

	routineConfigInterface := make([]map[string]interface{}, len(createdComponentRoutine.Config))
	for i, configItem := range createdComponentRoutine.Config {
		routineConfigInterface[i] = map[string]interface{}{
			"name":        configItem.Name,
			"description": configItem.Description,
			"multiple":    configItem.Multiple,
			"required":    configItem.Required,
			"sensitive":   configItem.Sensitive,
			"type":        configItem.Type,
			"value":       configItem.Value,
		}
	}
	d.Set("config", routineConfigInterface)

	d.Set("input", createdComponentRoutine.Input)
	d.Set("output", createdComponentRoutine.Output)
	return nil
}

func resourceUpdateComponentRoutine(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := ToComponentRoutine(d)

	updatedComponentRoutine, err := apiClient.UpdateComponentRoutine(itemId, item)
	if err != nil {
		return err
	}

	d.Set("name", updatedComponentRoutine.Name)
	d.Set("description", updatedComponentRoutine.Description)
	d.Set("type", updatedComponentRoutine.Type)
	d.Set("component_id", updatedComponentRoutine.ComponentId)
	d.Set("config", updatedComponentRoutine.Config)
	d.Set("input", updatedComponentRoutine.Input)
	d.Set("output", updatedComponentRoutine.Output)
	return nil
}

func resourceReadComponentRoutine(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedComponentRoutine, err := apiClient.RetrieveComponentRoutine(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding ComponentRoutine with ID %s: %s", itemId, err)
	}

	// Converting Config
	configList := make([]interface{}, len(retrievedComponentRoutine.Config))
	for i, configItem := range retrievedComponentRoutine.Config {
		configList[i] = map[string]interface{}{
			"name":        configItem.Name,
			"description": configItem.Description,
			"multiple":    configItem.Multiple,
			"required":    configItem.Required,
			"sensitive":   configItem.Sensitive,
			"type":        configItem.Type,
		}
		if configItem.Value != nil {
			configList[i].(map[string]interface{})["value"] = string(*configItem.Value)
		}
	}

	// Converting Output
	outputList := make([]interface{}, len(retrievedComponentRoutine.Output))
	for i, outputItem := range retrievedComponentRoutine.Output {
		outputList[i] = map[string]interface{}{
			"name":        outputItem.Name,
			"description": outputItem.Description,
			"multiple":    outputItem.Multiple,
			"required":    outputItem.Required,
			"type":        outputItem.Type,
			"value_type":  outputItem.ValueType,
		}
		if outputItem.Value != nil {
			outputList[i].(map[string]interface{})["value"] = []interface{}{}
			for _, value := range outputItem.Value {
				outputList[i].(map[string]interface{})["value"] = append(
					outputList[i].(map[string]interface{})["value"].([]interface{}), map[string]interface{}{
						"asset":     value.Asset,
						"attribute": value.Attribute,
					})
			}
		}
	}

	// Converting Input
	inputList := make([]interface{}, len(retrievedComponentRoutine.Input))
	for i, inputItem := range retrievedComponentRoutine.Input {
		inputList[i] = map[string]interface{}{
			"name":        inputItem.Name,
			"description": inputItem.Description,
			"multiple":    inputItem.Multiple,
			"required":    inputItem.Required,
			"type":        inputItem.Type,
			"value_type":  inputItem.ValueType,
		}
		if inputItem.Value != nil {
			inputList[i].(map[string]interface{})["value"] = []interface{}{}
			for _, value := range inputItem.Value {
				inputList[i].(map[string]interface{})["value"] = append(
					inputList[i].(map[string]interface{})["value"].([]interface{}), map[string]interface{}{
						"asset":     value.Asset,
						"attribute": value.Attribute,
					})
			}
		}
	}

	d.SetId(retrievedComponentRoutine.ID)
	d.Set("name", retrievedComponentRoutine.Name)
	d.Set("description", retrievedComponentRoutine.Description)
	d.Set("type", retrievedComponentRoutine.Type)
	d.Set("component_id", retrievedComponentRoutine.ComponentId)
	d.Set("config", configList)
	d.Set("input", inputList)
	d.Set("output", outputList)

	return nil
}

func resourceDeleteComponentRoutine(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteComponentRoutine(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsComponentRoutine(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveComponentRoutine(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
