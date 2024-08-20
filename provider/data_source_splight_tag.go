package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func dataSourceTag() *schema.Resource {
	return &schema.Resource{
		Description: "Data source to fetch all tags defined in the organization account",
		ReadContext: dataSourceTagRead,
		Schema: map[string]*schema.Schema{
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the resource",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of the resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*client.Client)
	tags, err := apiClient.ListTags()

	if err != nil {
		return diag.FromErr(err)
	}

	var tagsMap []map[string]string

	for _, tag := range tags {
		tagMap := map[string]string{
			"id":   tag.ID,
			"name": tag.Name,
		}
		tagsMap = append(tagsMap, tagMap)
	}

	if err := d.Set("tags", tagsMap); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("tags")

	return nil
}
