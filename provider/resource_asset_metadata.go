package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
	"github.com/splightplatform/splight-terraform-provider/utils"
)

func resourceAssetMetadata() *schema.Resource {
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
		Create: resourceCreateAssetMetadata,
		Read:   resourceReadAssetMetadata,
		Update: resourceUpdateAssetMetadata,
		Delete: resourceDeleteAssetMetadata,
		Exists: resourceExistsAssetMetadata,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateAssetMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.AssetMetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
		Unit:  utils.ValidateNullableString(d.Get("unit").(string)),
	}
	createdAssetMetadata, err := apiClient.CreateAssetMetadata(&item)
	if err != nil {
		return fmt.Errorf("Error creating asset %s", err)
	}
	d.SetId(createdAssetMetadata.ID)
	return nil
}

func resourceUpdateAssetMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	itemId := d.Id()
	item := client.AssetMetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
		Unit:  utils.ValidateNullableString(d.Get("unit").(string)),
	}
	updatedAssetMetadata, err := apiClient.UpdateAssetMetadata(itemId, &item)

	if err != nil {
		return err
	}
	d.Set("asset", updatedAssetMetadata.Asset)
	d.Set("name", updatedAssetMetadata.Name)
	d.Set("type", updatedAssetMetadata.Type)
	d.Set("value", updatedAssetMetadata.Value)
	d.Set("unit", updatedAssetMetadata.Unit)
	return nil
}

func resourceReadAssetMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAssetMetadata, err := apiClient.RetrieveAssetMetadata(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding AssetMetadata with ID %s", itemId)
		}
	}

	d.SetId(retrievedAssetMetadata.ID)
	d.Set("asset", retrievedAssetMetadata.Asset)
	d.Set("name", retrievedAssetMetadata.Name)
	d.Set("type", retrievedAssetMetadata.Type)
	d.Set("value", retrievedAssetMetadata.Value)
	d.Set("unit", retrievedAssetMetadata.Unit)
	return nil
}

func resourceDeleteAssetMetadata(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAssetMetadata(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsAssetMetadata(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveAssetMetadata(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
