package provider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
	"github.com/splightplatform/splight-terraform-provider/utils"
)

func resourceAssetAttribute() *schema.Resource {
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
		Create: resourceCreateAssetAttribute,
		Read:   resourceReadAssetAttribute,
		Update: resourceUpdateAssetAttribute,
		Delete: resourceDeleteAssetAttribute,
		Exists: resourceExistsAssetAttribute,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateAssetAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	item := client.AssetAttributeParams{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Asset: d.Get("asset").(string),
		Unit:  utils.ValidateNullableString(d.Get("unit").(string)),
	}
	createdAssetAttribute, err := apiClient.CreateAssetAttribute(&item)
	if err != nil {
		return fmt.Errorf("Error creating asset %s", err)
	}
	d.SetId(createdAssetAttribute.ID)
	return nil
}

func resourceUpdateAssetAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	itemId := d.Id()
	item := client.AssetAttributeParams{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Asset: d.Get("asset").(string),
		Unit:  utils.ValidateNullableString(d.Get("unit").(string)),
	}
	updatedAssetAttribute, err := apiClient.UpdateAssetAttribute(itemId, &item)

	if err != nil {
		return err
	}
	d.Set("name", updatedAssetAttribute.Name)
	d.Set("type", updatedAssetAttribute.Type)
	d.Set("asset", updatedAssetAttribute.Asset)
	d.Set("unit", updatedAssetAttribute.Unit)
	return nil
}

func resourceReadAssetAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAssetAttribute, err := apiClient.RetrieveAssetAttribute(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding AssetAttribute with ID %s", itemId)
		}
	}

	d.SetId(retrievedAssetAttribute.ID)
	d.Set("name", retrievedAssetAttribute.Name)
	d.Set("type", retrievedAssetAttribute.Type)
	d.Set("unit", retrievedAssetAttribute.Unit)
	d.Set("asset", retrievedAssetAttribute.Asset)
	return nil
}

func resourceDeleteAssetAttribute(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAssetAttribute(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsAssetAttribute(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveAssetAttribute(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
