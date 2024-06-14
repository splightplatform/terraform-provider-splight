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

func resourceCreateComponentRoutine(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

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
			Value:       json.RawMessage(data["value"].(string)),
		}
	}

	// Handling output
	componentOutputInterface := d.Get("output").(*schema.Set).List()
	componentOutput := make([]client.ComponentRoutineIOParam, len(componentOutputInterface))
	for i, item := range componentOutputInterface {
		data := item.(map[string]interface{})
		var outputValue *client.ComponentRoutineDataAddress
		if data["value"] != nil {
			valueData := data["value"].(*schema.Set).List()[0].(map[string]interface{})
			outputValue = &client.ComponentRoutineDataAddress{
				Asset:     valueData["asset"].(string),
				Attribute: valueData["attribute"].(string),
			}
		}
		componentOutput[i] = client.ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
			Value:       outputValue,
		}
	}

	// Handling input
	componentInputInterface := d.Get("input").(*schema.Set).List()
	componentInput := make([]client.ComponentRoutineIOParam, len(componentInputInterface))
	for i, item := range componentInputInterface {
		data := item.(map[string]interface{})
		var inputValue *client.ComponentRoutineDataAddress
		listData := data["value"].(*schema.Set).List()
		// fmt.Println("List Data", listData)
		if len(listData) > 0 {
			valueData := listData[0].(map[string]interface{})
			inputValue = &client.ComponentRoutineDataAddress{
				Asset:     valueData["asset"].(string),
				Attribute: valueData["attribute"].(string),
			}
		}
		componentInput[i] = client.ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
			Value:       inputValue,
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

	createdComponentRoutine, err := apiClient.CreateComponentRoutine(&item)
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
			Value:       json.RawMessage(data["value"].(string)),
		}
	}

	// Handling output
	componentOutputInterface := d.Get("output").(*schema.Set).List()
	componentOutput := make([]client.ComponentRoutineIOParam, len(componentOutputInterface))
	for i, item := range componentOutputInterface {
		data := item.(map[string]interface{})
		var outputValue *client.ComponentRoutineDataAddress
		listData := data["value"].(*schema.Set).List()
		if len(listData) > 0 {
			valueData := listData[0].(map[string]interface{})
			outputValue = &client.ComponentRoutineDataAddress{
				Asset:     valueData["asset"].(string),
				Attribute: valueData["attribute"].(string),
			}
		}
		componentOutput[i] = client.ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
			Value:       outputValue,
		}
	}

	// Handling input
	componentInputInterface := d.Get("input").(*schema.Set).List()
	componentInput := make([]client.ComponentRoutineIOParam, len(componentInputInterface))
	for i, item := range componentInputInterface {
		data := item.(map[string]interface{})
		var inputValue *client.ComponentRoutineDataAddress
		listData := data["value"].(*schema.Set).List()
		if len(listData) > 0 {
			valueData := listData[0].(map[string]interface{})
			inputValue = &client.ComponentRoutineDataAddress{
				Asset:     valueData["asset"].(string),
				Attribute: valueData["attribute"].(string),
			}
		}
		componentInput[i] = client.ComponentRoutineIOParam{
			Name:        data["name"].(string),
			Description: data["description"].(string),
			Type:        data["type"].(string),
			Multiple:    data["multiple"].(bool),
			Required:    data["required"].(bool),
			ValueType:   data["value_type"].(string),
			Value:       inputValue,
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

	updatedComponentRoutine, err := apiClient.UpdateComponentRoutine(itemId, &item)
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
			"value":       string(configItem.Value),
		}
	}

	// Converting Output
	outputList := make([]interface{}, len(retrievedComponentRoutine.Output))
	for i, outputItem := range retrievedComponentRoutine.Output {
		var value *[]interface{}
		if outputItem.Value != nil {
			value = &[]interface{}{
				map[string]interface{}{
					"asset":     outputItem.Value.Asset,
					"attribute": outputItem.Value.Attribute,
				},
			}
		}
		outputList[i] = map[string]interface{}{
			"name":        outputItem.Name,
			"description": outputItem.Description,
			"multiple":    outputItem.Multiple,
			"required":    outputItem.Required,
			"type":        outputItem.Type,
			"value_type":  outputItem.ValueType,
			"value":       value,
			// "value": []interface{}{
			// 	map[string]interface{}{
			// 		"asset":     outputItem.Value.Asset,
			// 		"attribute": outputItem.Value.Attribute,
			// 	},
			// },
		}
	}

	// Converting Input
	inputList := make([]interface{}, len(retrievedComponentRoutine.Input))
	for i, inputItem := range retrievedComponentRoutine.Input {
		var value *[]interface{}
		if inputItem.Value != nil {
			value = &[]interface{}{
				map[string]interface{}{
					"asset":     inputItem.Value.Asset,
					"attribute": inputItem.Value.Attribute,
				},
			}
		}
		inputList[i] = map[string]interface{}{
			"name":        inputItem.Name,
			"description": inputItem.Description,
			"multiple":    inputItem.Multiple,
			"required":    inputItem.Required,
			"type":        inputItem.Type,
			"value_type":  inputItem.ValueType,
			"value":       value,
			// "value": []interface{}{
			// 	map[string]interface{}{
			// 		"asset":     inputItem.Value.Asset,
			// 		"attribute": inputItem.Value.Attribute,
			// 	},
			// },
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
