package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
	"github.com/splightplatform/splight-terraform-provider/verify"
)

func resourceAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"asset": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"unit": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
		Create: resourceCreateAttribute,
		Read:   resourceReadAttribute,
		Update: resourceUpdateAttribute,
		Delete: resourceDeleteAttribute,
		Exists: resourceExistsAttribute,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.AttributeParams{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Asset: d.Get("asset").(string),
		Unit:  verify.ValidateNullableString(d.Get("unit").(string)),
	}
	createdAttribute, err := apiClient.CreateAttribute(&item)
	if err != nil {
		return fmt.Errorf("Error creating asset %s", err)
	}
	d.SetId(createdAttribute.ID)
	return nil
}

func resourceUpdateAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	itemId := d.Id()
	item := client.AttributeParams{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Asset: d.Get("asset").(string),
		Unit:  verify.ValidateNullableString(d.Get("unit").(string)),
	}
	updatedAttribute, err := apiClient.UpdateAttribute(itemId, &item)

	if err != nil {
		return err
	}
	d.Set("name", updatedAttribute.Name)
	d.Set("type", updatedAttribute.Type)
	d.Set("asset", updatedAttribute.Asset)
	d.Set("unit", updatedAttribute.Unit)
	return nil
}

func resourceReadAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAttribute, err := apiClient.RetrieveAttribute(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Attribute with ID %s", itemId)
		}
	}

	d.SetId(retrievedAttribute.ID)
	d.Set("name", retrievedAttribute.Name)
	d.Set("type", retrievedAttribute.Type)
	d.Set("unit", retrievedAttribute.Unit)
	d.Set("asset", retrievedAttribute.Asset)
	return nil
}

func resourceDeleteAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAttribute(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsAttribute(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveAttribute(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
