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

	componentConfigInterface := d.Get("config").([]interface{})
	componentConfigInterfaceList := make([]map[string]interface{}, len(componentConfigInterface))
	for i, componentConfigInterfaceItem := range componentConfigInterface {
		componentConfigInterfaceList[i] = componentConfigInterfaceItem.(map[string]interface{})
	}
	componentConfig := make([]client.ComponentRoutineConfigParam, len(componentConfigInterfaceList))
	for i, componentConfigItem := range componentConfigInterfaceList {
		componentConfig[i] = client.ComponentRoutineConfigParam{
			Name:        componentConfigItem["name"].(string),
			Description: componentConfigItem["description"].(string),
			Type:        componentConfigItem["type"].(string),
			Value:       json.RawMessage(componentConfigItem["value"].(string)),
			Multiple:    componentConfigItem["multiple"].(bool),
			Required:    componentConfigItem["required"].(bool),
			Sensitive:   componentConfigItem["sensitive"].(bool),
		}
	}

	componentOutputInterface := d.Get("output").([]interface{})
	componentOutputInterfaceList := make([]map[string]interface{}, len(componentOutputInterface))
	for i, componentOutputInterfaceItem := range componentOutputInterface {
		componentOutputInterfaceList[i] = componentOutputInterfaceItem.(map[string]interface{})
	}
	componentOutput := make([]client.ComponentRoutineIOParam, len(componentOutputInterfaceList))
	for i, componentOutputItem := range componentOutputInterfaceList {
		outputValue := client.ComponentRoutineDataAddress{}
		err := json.Unmarshal([]byte(componentOutputItem["value"].(string)), &outputValue)
		if err != nil {
			return err
		}
		componentOutput[i] = client.ComponentRoutineIOParam{
			Name:        componentOutputItem["name"].(string),
			Description: componentOutputItem["description"].(string),
			Type:        componentOutputItem["type"].(string),
			Multiple:    componentOutputItem["multiple"].(bool),
			Required:    componentOutputItem["required"].(bool),
			ValueType:   componentOutputItem["value_type"].(string),
			Value:       outputValue,
		}
	}

	componentInputInterface := d.Get("input").([]interface{})
	componentInputInterfaceList := make([]map[string]interface{}, len(componentInputInterface))
	for i, componentInputInterfaceItem := range componentInputInterface {
		componentInputInterfaceList[i] = componentInputInterfaceItem.(map[string]interface{})
	}
	componentInput := make([]client.ComponentRoutineIOParam, len(componentInputInterfaceList))
	for i, componentInputItem := range componentInputInterfaceList {
		inputValue := client.ComponentRoutineDataAddress{}
		err := json.Unmarshal([]byte(componentInputItem["value"].(string)), &inputValue)
		if err != nil {
			return err
		}
		componentInput[i] = client.ComponentRoutineIOParam{
			Name:        componentInputItem["name"].(string),
			Description: componentInputItem["description"].(string),
			Type:        componentInputItem["type"].(string),
			Multiple:    componentInputItem["multiple"].(bool),
			Required:    componentInputItem["required"].(bool),
			ValueType:   componentInputItem["value_type"].(string),
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

	componentConfigInterface := d.Get("config").([]interface{})
	componentConfigInterfaceList := make([]map[string]interface{}, len(componentConfigInterface))
	for i, componentConfigInterfaceItem := range componentConfigInterface {
		componentConfigInterfaceList[i] = componentConfigInterfaceItem.(map[string]interface{})
	}
	componentConfig := make([]client.ComponentRoutineConfigParam, len(componentConfigInterfaceList))
	for i, componentConfigItem := range componentConfigInterfaceList {
		componentConfig[i] = client.ComponentRoutineConfigParam{
			Name:        componentConfigItem["name"].(string),
			Description: componentConfigItem["description"].(string),
			Type:        componentConfigItem["type"].(string),
			Value:       json.RawMessage(componentConfigItem["value"].(string)),
			Multiple:    componentConfigItem["multiple"].(bool),
			Required:    componentConfigItem["required"].(bool),
			Sensitive:   componentConfigItem["sensitive"].(bool),
		}
	}

	componentOutputInterface := d.Get("output").([]interface{})
	componentOutputInterfaceList := make([]map[string]interface{}, len(componentOutputInterface))
	for i, componentOutputInterfaceItem := range componentOutputInterface {
		componentOutputInterfaceList[i] = componentOutputInterfaceItem.(map[string]interface{})
	}
	componentOutput := make([]client.ComponentRoutineIOParam, len(componentOutputInterfaceList))
	for i, componentOutputItem := range componentOutputInterfaceList {
		outputValue := client.ComponentRoutineDataAddress{}
		err := json.Unmarshal([]byte(componentOutputItem["value"].(string)), &outputValue)
		if err != nil {
			return err
		}
		componentOutput[i] = client.ComponentRoutineIOParam{
			Name:        componentOutputItem["name"].(string),
			Description: componentOutputItem["description"].(string),
			Type:        componentOutputItem["type"].(string),
			Multiple:    componentOutputItem["multiple"].(bool),
			Required:    componentOutputItem["required"].(bool),
			ValueType:   componentOutputItem["value_type"].(string),
			Value:       outputValue,
		}
	}

	componentInputInterface := d.Get("input").([]interface{})
	componentInputInterfaceList := make([]map[string]interface{}, len(componentInputInterface))
	for i, componentInputInterfaceItem := range componentInputInterface {
		componentInputInterfaceList[i] = componentInputInterfaceItem.(map[string]interface{})
	}
	componentInput := make([]client.ComponentRoutineIOParam, len(componentInputInterfaceList))
	for i, componentInputItem := range componentInputInterfaceList {
		inputValue := client.ComponentRoutineDataAddress{}
		err := json.Unmarshal([]byte(componentInputItem["value"].(string)), &inputValue)
		if err != nil {
			return err
		}
		componentInput[i] = client.ComponentRoutineIOParam{
			Name:        componentInputItem["name"].(string),
			Description: componentInputItem["description"].(string),
			Type:        componentInputItem["type"].(string),
			Multiple:    componentInputItem["multiple"].(bool),
			Required:    componentInputItem["required"].(bool),
			ValueType:   componentInputItem["value_type"].(string),
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
		} else {
			return fmt.Errorf("error finding ComponentRoutine with ID %s", itemId)
		}
	}

	configDict := make([]map[interface{}]interface{}, len(retrievedComponentRoutine.Config))
	for i, configDictItem := range retrievedComponentRoutine.Config {
		configDict[i] = map[interface{}]interface{}{
			"name":        configDictItem.Name,
			"description": configDictItem.Description,
			"multiple":    configDictItem.Multiple,
			"required":    configDictItem.Required,
			"sensitive":   configDictItem.Sensitive,
			"type":        configDictItem.Type,
			"value":       json.RawMessage(configDictItem.Value),
		}
	}
	outputDict := make([]map[interface{}]interface{}, len(retrievedComponentRoutine.Output))
	for i, outputDictItem := range retrievedComponentRoutine.Output {
		outputValue, _ := json.Marshal(outputDictItem.Value)
		// return fmt.Errorf("%s", outputValue)
		outputDict[i] = map[interface{}]interface{}{
			"name":        outputDictItem.Name,
			"description": outputDictItem.Description,
			"multiple":    outputDictItem.Multiple,
			"required":    outputDictItem.Required,
			"type":        outputDictItem.Type,
			"value_type":  outputDictItem.ValueType,
			"value":       string(outputValue),
		}
	}
	inputDict := make([]map[interface{}]interface{}, len(retrievedComponentRoutine.Input))
	for i, inputDictItem := range retrievedComponentRoutine.Input {
		inputValue, _ := json.Marshal(inputDictItem.Value)
		inputDict[i] = map[interface{}]interface{}{
			"name":        inputDictItem.Name,
			"description": inputDictItem.Description,
			"multiple":    inputDictItem.Multiple,
			"required":    inputDictItem.Required,
			"type":        inputDictItem.Type,
			"value_type":  inputDictItem.ValueType,
			"value":       string(inputValue),
		}
	}

	d.SetId(retrievedComponentRoutine.ID)
	d.Set("name", retrievedComponentRoutine.Name)
	d.Set("description", retrievedComponentRoutine.Description)
	d.Set("type", retrievedComponentRoutine.Type)
	d.Set("component_id", retrievedComponentRoutine.ComponentId)
	d.Set("config", configDict)
	d.Set("input", inputDict)
	d.Set("output", outputDict)
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
