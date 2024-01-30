package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
	"github.com/splightplatform/splight-terraform-provider/verify"
)

func resourceMetadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"asset": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
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
			"unit": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
		Create: resourceCreateMetadata,
		Read:   resourceReadMetadata,
		Update: resourceUpdateMetadata,
		Delete: resourceDeleteMetadata,
		Exists: resourceExistsMetadata,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.MetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
		Unit:  verify.ValidateNullableString(d.Get("unit").(string)),
	}
	createdMetadata, err := apiClient.CreateMetadata(&item)
	if err != nil {
		return fmt.Errorf("Error creating asset %s", err)
	}
	d.SetId(createdMetadata.ID)
	return nil
}

func resourceUpdateMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	itemId := d.Id()
	item := client.MetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
		Unit:  verify.ValidateNullableString(d.Get("unit").(string)),
	}
	updatedMetadata, err := apiClient.UpdateMetadata(itemId, &item)

	if err != nil {
		return err
	}
	d.Set("asset", updatedMetadata.Asset)
	d.Set("name", updatedMetadata.Name)
	d.Set("type", updatedMetadata.Type)
	d.Set("value", updatedMetadata.Value)
	d.Set("unit", updatedMetadata.Unit)
	return nil
}

func resourceReadMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedMetadata, err := apiClient.RetrieveMetadata(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Metadata with ID %s", itemId)
		}
	}

	d.SetId(retrievedMetadata.ID)
	d.Set("asset", retrievedMetadata.Asset)
	d.Set("name", retrievedMetadata.Name)
	d.Set("type", retrievedMetadata.Type)
	d.Set("value", retrievedMetadata.Value)
	d.Set("unit", retrievedMetadata.Unit)
	return nil
}

func resourceDeleteMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteMetadata(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsMetadata(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveMetadata(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
