package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/provider/schemas"
	"github.com/splightplatform/terraform-provider-splight/splight/client"
	"github.com/splightplatform/terraform-provider-splight/splight/client/models"
	"github.com/splightplatform/terraform-provider-splight/splight/settings"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string = "dev"

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	var diags diag.Diagnostics

	hostname := d.Get("hostname")
	token := d.Get("token")

	// Prepare overrides for the Splight configuration
	options := &settings.SplightConfigOverrides{
		HostnameOverride: hostname.(string),
		TokenOverride:    token.(string),
	}

	// Load configuration with possible overrides
	_, err := settings.LoadSplightConfig(options)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	userAgentOptions := client.UserAgent{
		ProductName:    "terraform-provider-splight",
		ProductVersion: Version,
	}

	client, err := client.NewClient(ctx, userAgentOptions)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"splight_asset":           resourceForType[*models.Asset](schemas.SchemaAsset),
			"splight_asset_attribute": resourceForType[*models.AssetAttribute](schemas.SchemaAssetAttribute),
			"splight_asset_metadata":  resourceForType[*models.AssetMetadata](schemas.SchemaAssetMetadata),
			"splight_tag":             resourceForType[*models.Tag](schemas.SchemaTag),
			"splight_alert":           resourceForType[*models.Alert](schemas.SchemaAlert),
			"splight_function":        resourceForType[*models.Function](schemas.SchemaFunction),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"splight_asset_kinds": dataSourceForType[*models.AssetKinds](schemas.SchemaAssetKinds),
			"splight_tags":        dataSourceForType[*models.Tags](schemas.SchemaTags),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
