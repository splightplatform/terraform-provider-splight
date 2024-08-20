package provider

import (
	"context"
	"encoding/json"
	"runtime"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaAsset(),
		CreateContext: resourceCreateAsset,
		ReadContext:   resourceReadAsset,
		UpdateContext: resourceUpdateAsset,
		DeleteContext: resourceDeleteAsset,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toAsset(d *schema.ResourceData) *client.AssetParams {
	relatedAssetsSet := d.Get("related_assets").(*schema.Set).List()
	relatedAssets := make([]client.RelatedAsset, len(relatedAssetsSet))
	for i, relatedAsset := range relatedAssetsSet {
		relatedAssets[i] = client.RelatedAsset{
			Id: relatedAsset.(string),
		}
	}

	var kind *client.AssetKind
	kinds := d.Get("kind").(*schema.Set).List()
	if len(kinds) > 0 {
		kindMap := kinds[0].(map[string]interface{})
		kind = &client.AssetKind{
			ID:   kindMap["id"].(string),
			Name: kindMap["name"].(string),
		}
	}

	runtime.Breakpoint()
	tagsInterface := d.Get("tags").(*schema.Set).List()
	tags := make([]client.Tag, len(tagsInterface))
	for i, item := range tagsInterface {
		tagItem := item.(map[string]interface{})
		tags[i] = client.Tag{
			ID: tagItem["id"].(string),
			TagParams: client.TagParams{
				Name: tagItem["name"].(string),
			},
		}
	}

	return &client.AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Geometry:      json.RawMessage(d.Get("geometry").(string)),
		RelatedAssets: relatedAssets,
		Tags:          tags,
		Kind:          kind,
	}
}

func saveAssetToState(d *schema.ResourceData, asset *client.Asset) {
	d.SetId(asset.ID)
	d.Set("name", asset.Name)
	d.Set("description", asset.Description)
	d.Set("related_assets", asset.RelatedAssets)
	d.Set("geometry", string(asset.Geometry))
	d.Set("tags", asset.Tags)

	if asset.Kind != nil {
		d.Set("kind", []map[string]interface{}{
			{
				"id":   asset.Kind.ID,
				"name": asset.Kind.Name,
			},
		})
	} else {
		d.Set("kind", nil)
	}
}

func resourceCreateAsset(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toAsset(d)

	createdAsset, err := apiClient.CreateAsset(item)
	if err != nil {
		return diag.Errorf("error creating Asset: %s", err.Error())
	}

	saveAssetToState(d, createdAsset)

	return nil
}

func resourceUpdateAsset(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item := toAsset(d)

	updatedAsset, err := apiClient.UpdateAsset(itemId, item)
	if err != nil {
		return diag.Errorf("error updating Asset with ID '%s': %s", itemId, err.Error())
	}

	saveAssetToState(d, updatedAsset)

	return nil
}

func resourceReadAsset(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	retrievedAsset, err := apiClient.RetrieveAsset(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading Asset with ID '%s': %s", itemId, err.Error())
		}
	}

	saveAssetToState(d, retrievedAsset)

	return nil
}

func resourceDeleteAsset(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteAsset(itemId)
	if err != nil {
		return diag.Errorf("error deleting Asset with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
