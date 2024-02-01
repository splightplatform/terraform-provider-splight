package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
	"github.com/splightplatform/splight-terraform-provider/verify"
)

func resourceComponentRoutine() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"component_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"multiple": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"sensitive": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"output": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"multiple": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"value": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: verify.JSONStringEqualSupressFunc,
						},
					},
				},
			},
		},
		Create: resourceCreateComponentRoutine,
		Read:   resourceReadComponentRoutine,
		Update: resourceUpdateComponentRoutine,
		Delete: resourceDeleteComponentRoutine,
		Exists: resourceExistsComponentRoutine,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			Value:       componentConfigItem["value"].(string),
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
			ValueType:   componentOutputItem["value_type"].(string),
			Multiple:    componentOutputItem["multiple"].(bool),
			Required:    componentOutputItem["required"].(bool),
			Value:       outputValue,
		}
	}

	item := client.ComponentRoutineParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		ComponentId: d.Get("component_id").(string),
		Config:      componentConfig,
		Input:       make([]client.ComponentRoutineIOParam, 0), //TODO
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
	d.Set("config", createdComponentRoutine.Config)
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
			Value:       componentConfigItem["value"].(string),
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
			ValueType:   componentOutputItem["value_type"].(string),
			Multiple:    componentOutputItem["multiple"].(bool),
			Required:    componentOutputItem["required"].(bool),
			Value:       outputValue,
		}
	}

	item := client.ComponentRoutineParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		ComponentId: d.Get("component_id").(string),
		Config:      componentConfig,
		Input:       make([]client.ComponentRoutineIOParam, 0), //TODO
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
	configDict := make([]map[interface{}]interface{}, len(retrievedComponentRoutine.Config))
	for i, configDictItem := range retrievedComponentRoutine.Config {
		configDict[i] = map[interface{}]interface{}{
			"name":        configDictItem.Name,
			"description": configDictItem.Description,
			"multiple":    configDictItem.Multiple,
			"required":    configDictItem.Required,
			"sensitive":   configDictItem.Sensitive,
			"type":        configDictItem.Type,
			"value":       configDictItem.Value,
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
