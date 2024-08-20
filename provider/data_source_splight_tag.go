package provider

import (
	"context"
	"runtime"

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
							Required:    true,
							Description: "ID of the resource",
						},
						"name": {
							Type:        schema.TypeString,
							Description: "name of the resource",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*client.Client)
	runtime.Breakpoint()
	tags, err := apiClient.ListTags()

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("tags", tags); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("tags")

	return nil
}
