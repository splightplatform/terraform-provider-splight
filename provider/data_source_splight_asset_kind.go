package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

func dataSourceAssetKind() *schema.Resource {
	return &schema.Resource{
		Description: "Data source to fetch all asset kinds defined in the platform",
		ReadContext: dataSourceKindRead,
		Schema: map[string]*schema.Schema{
			"kinds": {
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

func dataSourceKindRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apiClient := meta.(*client.Client)

	assetKinds, err := apiClient.ListAssetKinds()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("kinds", assetKinds); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("kinds")

	return nil
}
