package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
	"github.com/splightplatform/splight-terraform-provider/verify"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the resource",
			},
			"geometry": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Geometry of the resource",
				DiffSuppressFunc: verify.JSONStringEqualSupressFunc,
			},
		},
		Create: resourceCreateAsset,
		Read:   resourceReadAsset,
		Update: resourceUpdateAsset,
		Delete: resourceDeleteAsset,
		Exists: resourceExistsAsset,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	assetGeometry := client.AssetGeometry{}
	err := json.Unmarshal([]byte(d.Get("geometry").(string)), &assetGeometry)
	if err != nil {
		return err
	}
	item := client.AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		AssetGeometry: assetGeometry,
	}
	createdAsset, err := apiClient.CreateAsset(&item)
	if err != nil {
		return err
	}

	geometry, err := json.Marshal(createdAsset.AssetGeometry)
	if err != nil {
		return err
	}

	d.SetId(createdAsset.ID)
	d.Set("name", createdAsset.Name)
	d.Set("description", createdAsset.Description)
	d.Set("geometry", string(geometry))
	return nil
}

func resourceUpdateAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	assetGeometry := client.AssetGeometry{}
	err := json.Unmarshal([]byte(d.Get("geometry").(string)), &assetGeometry)
	if err != nil {
		return err
	}
	item := client.AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		AssetGeometry: assetGeometry,
	}

	updatedAsset, err := apiClient.UpdateAsset(itemId, &item)
	if err != nil {
		return err
	}

	geometry, err := json.Marshal(updatedAsset.AssetGeometry)
	if err != nil {
		return err
	}

	d.Set("name", updatedAsset.Name)
	d.Set("description", updatedAsset.Description)
	d.Set("geometry", string(geometry))
	return nil
}

func resourceReadAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAsset, err := apiClient.RetrieveAsset(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Asset with ID %s", itemId)
		}
	}

	geometry, err := json.Marshal(retrievedAsset.AssetGeometry)
	if err != nil {
		return err
	}

	d.SetId(retrievedAsset.ID)
	d.Set("name", retrievedAsset.Name)
	d.Set("description", retrievedAsset.Description)
	d.Set("geometry", string(geometry))
	return nil
}

func resourceDeleteAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAsset(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsAsset(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.RetrieveAsset(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
