package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
	"github.com/splightplatform/terraform-provider-splight/utils"
)

func resourceAssetAttribute() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaAssetAttribute(),
		CreateContext: resourceCreateAssetAttribute,
		ReadContext:   resourceReadAssetAttribute,
		UpdateContext: resourceUpdateAssetAttribute,
		DeleteContext: resourceDeleteAssetAttribute,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toAssetAttribute(d *schema.ResourceData) *client.AssetAttributeParams {
	return &client.AssetAttributeParams{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Asset: d.Get("asset").(string),
		Unit:  utils.ValidateNullableString(d.Get("unit").(string)),
	}
}

func saveAssetAttributeToState(d *schema.ResourceData, assetAttribute *client.AssetAttribute) {
	d.SetId(assetAttribute.ID)

	d.Set("name", assetAttribute.Name)
	d.Set("type", assetAttribute.Type)
	d.Set("asset", assetAttribute.Asset)
	d.Set("unit", assetAttribute.Unit)
}

func resourceCreateAssetAttribute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toAssetAttribute(d)

	createdAssetAttribute, err := apiClient.CreateAssetAttribute(item)

	if err != nil {
		return diag.Errorf("error creating AssetAttribute: %s", err.Error())
	}

	saveAssetAttributeToState(d, createdAssetAttribute)

	return nil
}

func resourceUpdateAssetAttribute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := toAssetAttribute(d)

	updatedAssetAttribute, err := apiClient.UpdateAssetAttribute(itemId, item)

	if err != nil {
		return diag.Errorf("error updating AssetAttribute with ID '%s': %s", itemId, err.Error())
	}

	saveAssetAttributeToState(d, updatedAssetAttribute)

	return nil
}

func resourceReadAssetAttribute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedAssetAttribute, err := apiClient.RetrieveAssetAttribute(itemId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading AssetAttribute with ID '%s': %s", itemId, err.Error())
		}
	}

	saveAssetAttributeToState(d, retrievedAssetAttribute)

	return nil
}

func resourceDeleteAssetAttribute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAssetAttribute(itemId)

	if err != nil {
		return diag.Errorf("error deleting AssetAttribute with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
