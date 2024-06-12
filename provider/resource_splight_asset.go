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

	// Initialize the kind variable
	var kind map[string]interface{}
	kinds := d.Get("kind").(*schema.Set).List()
	if len(kinds) > 0 {
		kind = kinds[0].(map[string]interface{})
	}

	// Prepare the item with a check on kind
	item := client.AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Geometry:      json.RawMessage(d.Get("geometry").(string)),
		RelatedAssets: assetRelatedAssets,
	}

	if kind != nil {
		item.Kind = &client.AssetKind{
			ID:   kind["id"].(string),
			Name: kind["name"].(string),
		}
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
	if kind != nil {
		d.Set("kind", []map[string]interface{}{
			{
				"id":   createdAsset.Kind.ID,
				"name": createdAsset.Kind.Name,
			},
		})
	}

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

	// Initialize the kind variable
	var kind map[string]interface{}
	kinds := d.Get("kind").(*schema.Set).List()
	if len(kinds) > 0 {
		kind = kinds[0].(map[string]interface{})
	}

	// Prepare the item
	item := client.AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Geometry:      json.RawMessage(d.Get("geometry").(string)),
		RelatedAssets: assetRelatedAssets,
	}

	// Only set Kind if kind is not nil
	if kind != nil {
		item.Kind = &client.AssetKind{
			ID:   kind["id"].(string),
			Name: kind["name"].(string),
		}
	}

	updatedAsset, err := apiClient.UpdateAsset(itemId, &item)
	if err != nil {
		return err
	}

	d.Set("name", updatedAsset.Name)
	d.Set("description", updatedAsset.Description)
	d.Set("related_assets", updatedAsset.RelatedAssets)
	d.Set("geometry", string(updatedAsset.Geometry))

	// Set kind as a set if it's not nil
	if kind != nil {
		d.Set("kind", []map[string]interface{}{
			{
				"id":   updatedAsset.Kind.ID,
				"name": updatedAsset.Kind.Name,
			},
		})
	} else {
		// Clear the kind field if it was not provided
		d.Set("kind", nil)
	}

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
		return nil
	}

	d.SetId(retrievedAsset.ID)
	d.Set("name", retrievedAsset.Name)
	d.Set("description", retrievedAsset.Description)
	d.Set("related_assets", retrievedAsset.RelatedAssets)
	d.Set("geometry", string(retrievedAsset.Geometry))

	// Check if kind is present in the retrieved asset
	if retrievedAsset.Kind != nil {
		d.Set("kind", []map[string]interface{}{
			{
				"id":   retrievedAsset.Kind.ID,
				"name": retrievedAsset.Kind.Name,
			},
		})
	} else {
		// Clear the kind field if it was not provided
		d.Set("kind", nil)
	}

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
