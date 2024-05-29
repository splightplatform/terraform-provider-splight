package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Schema: schemaAsset(),
		Create: resourceCreateAsset,
		Read:   resourceReadAsset,
		Update: resourceUpdateAsset,
		Delete: resourceDeleteAsset,
		Exists: resourceExistsAsset,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCreateAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	assetRelatedAssetsSet := d.Get("related_assets").(*schema.Set).List()
	assetRelatedAssets := make([]client.RelatedAsset, len(assetRelatedAssetsSet))
	for i, relatedAsset := range assetRelatedAssetsSet {
		assetRelatedAssets[i] = client.RelatedAsset{
			Id: relatedAsset.(string),
		}
	}

	// Consider only the first element of the set
	kind := d.Get("kind").(*schema.Set).List()[0].(map[string]interface{})

	item := client.AssetParams{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Geometry:    json.RawMessage(d.Get("geometry").(string)),
		Kind: client.AssetKind{
			ID:   kind["id"].(string),
			Name: kind["name"].(string),
		},
		RelatedAssets: assetRelatedAssets,
	}

	createdAsset, err := apiClient.CreateAsset(&item)
	if err != nil {
		return err
	}

	d.SetId(createdAsset.ID)
	d.Set("name", createdAsset.Name)
	d.Set("description", createdAsset.Description)
	d.Set("related_assets", createdAsset.RelatedAssets)
	d.Set("geometry", string(createdAsset.Geometry))

	// Set kind as a set
	d.Set("kind", []map[string]interface{}{
		{
			"id":   createdAsset.Kind.ID,
			"name": createdAsset.Kind.Name,
		},
	})

	return nil
}

func resourceUpdateAsset(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	assetRelatedAssetsSet := d.Get("related_assets").(*schema.Set).List()
	assetRelatedAssets := make([]client.RelatedAsset, len(assetRelatedAssetsSet))
	for i, relatedAsset := range assetRelatedAssetsSet {
		assetRelatedAssets[i] = client.RelatedAsset{
			Id: relatedAsset.(string),
		}
	}

	// Consider only the first element of the set
	kind := d.Get("kind").(*schema.Set).List()[0].(map[string]interface{})

	item := client.AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Geometry:      json.RawMessage(d.Get("geometry").(string)),
		RelatedAssets: assetRelatedAssets,
		Kind: client.AssetKind{
			ID:   kind["id"].(string),
			Name: kind["name"].(string),
		},
	}

	updatedAsset, err := apiClient.UpdateAsset(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("name", updatedAsset.Name)
	d.Set("description", updatedAsset.Description)
	d.Set("related_assets", updatedAsset.RelatedAssets)
	d.Set("geometry", string(updatedAsset.Geometry))

	// Set kind as a set
	d.Set("kind", []map[string]interface{}{
		{
			"id":   updatedAsset.Kind.ID,
			"name": updatedAsset.Kind.Name,
		},
	})

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

	d.SetId(retrievedAsset.ID)
	d.Set("name", retrievedAsset.Name)
	d.Set("description", retrievedAsset.Description)
	d.Set("related_assets", retrievedAsset.RelatedAssets)
	d.Set("geometry", string(retrievedAsset.Geometry))

	// Set kind as a set
	d.Set("kind", []map[string]interface{}{
		{
			"id":   retrievedAsset.Kind.ID,
			"name": retrievedAsset.Kind.Name,
		},
	})

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
