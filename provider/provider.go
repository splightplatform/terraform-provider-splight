package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/settings"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string = "dev"

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
			"splight_asset": resourceAsset(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	var diags diag.Diagnostics

	hostname := d.Get("hostname")
	token := d.Get("token")

	// Validate that either both or neither are set
	if (hostname != "" && token == "") || (hostname == "" && token != "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid Configuration",
			Detail:   "Both 'hostname' and 'token' must be set or neither.",
		})
		return nil, diags
	}

	// Prepare overrides for the Splight configuration
	options := &settings.SplightConfigOverrides{
		HostnameOverride: hostname.(string),
		TokenOverride:    token.(string),
	}

	// Load configuration with optional overrides
	_, err := settings.LoadSplightConfig(options)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// Save info in the context
	ctx = context.WithValue(ctx, "ProductName", "terraform-provider-splight")
	ctx = context.WithValue(ctx, "Version", Version)

	return ctx, diags
}
