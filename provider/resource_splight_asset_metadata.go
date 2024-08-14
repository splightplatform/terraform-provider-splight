package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
	"github.com/splightplatform/terraform-provider-splight/utils"
)

func resourceAssetMetadata() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaAssetMetadata(),
		CreateContext: resourceCreateAssetMetadata,
		ReadContext:   resourceReadAssetMetadata,
		UpdateContext: resourceUpdateAssetMetadata,
		DeleteContext: resourceDeleteAssetMetadata,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toAssetMetadata(d *schema.ResourceData) *client.AssetMetadataParams {
	return &client.AssetMetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
		Unit:  utils.ValidateNullableString(d.Get("unit").(string)),
	}
}

func saveAssetMetadataToState(d *schema.ResourceData, assetMetadata *client.AssetMetadata) {
	d.SetId(assetMetadata.ID)

	d.Set("asset", assetMetadata.Asset)
	d.Set("name", assetMetadata.Name)
	d.Set("type", assetMetadata.Type)
	d.Set("value", assetMetadata.Value)
	d.Set("unit", assetMetadata.Unit)
}

func resourceCreateAssetMetadata(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toAssetMetadata(d)

	createdAssetMetadata, err := apiClient.CreateAssetMetadata(item)

	if err != nil {
		return diag.Errorf("Error creating AssetMetadata: %s", err)
	}

	saveAssetMetadataToState(d, createdAssetMetadata)

	return nil
}

func resourceUpdateAssetMetadata(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := toAssetMetadata(d)

	updatedAssetMetadata, err := apiClient.UpdateAssetMetadata(itemId, item)
	if err != nil {
		return diag.Errorf("Error updating AssetMetadata with ID '%s': %s", itemId, err)
	}

	saveAssetMetadataToState(d, updatedAssetMetadata)

	return nil
}

func resourceReadAssetMetadata(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedAssetMetadata, err := apiClient.RetrieveAssetMetadata(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("Error finding AssetMetadata with ID '%s': %s", itemId, err)
		}
	}

	saveAssetMetadataToState(d, retrievedAssetMetadata)

	return nil
}

func resourceDeleteAssetMetadata(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAssetMetadata(itemId)
	if err != nil {
		return diag.Errorf("Error deleting AssetMetadata with ID '%s': %s", itemId, err)
	}

	d.SetId("")

	return nil
}
