package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		Schema:        schemaTag(),
		CreateContext: resourceCreateTag,
		ReadContext:   resourceReadTag,
		UpdateContext: resourceUpdateTag,
		DeleteContext: resourceDeleteTag,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func toTag(d *schema.ResourceData) *client.TagParams {
	return &client.TagParams{
		Name: d.Get("name").(string),
	}
}

func saveTagToState(d *schema.ResourceData, tag *client.Tag) {
	d.SetId(tag.ID)

	d.Set("name", tag.Name)
}

func resourceCreateTag(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := toTag(d)

	createdTag, err := apiClient.CreateTag(item)

	if err != nil {
		return diag.Errorf("error creating Tag: %s", err.Error())
	}

	saveTagToState(d, createdTag)

	return nil
}

func resourceUpdateTag(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	item := toTag(d)

	updatedTag, err := apiClient.UpdateTag(itemId, item)

	if err != nil {
		return diag.Errorf("error updating Tag with ID '%s': %s", itemId, err.Error())
	}

	saveTagToState(d, updatedTag)

	return nil
}

func resourceReadTag(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	retrievedTag, err := apiClient.RetrieveTag(itemId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		} else {
			return diag.Errorf("error reading Tag with ID '%s': %s", itemId, err.Error())
		}
	}

	saveTagToState(d, retrievedTag)

	return nil
}

func resourceDeleteTag(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteTag(itemId)

	if err != nil {
		return diag.Errorf("error deleting Tag with ID '%s': %s", itemId, err.Error())
	}

	d.SetId("")

	return nil
}
