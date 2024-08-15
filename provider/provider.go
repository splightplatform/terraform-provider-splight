package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/terraform-provider-splight/api/client"
	"github.com/splightplatform/terraform-provider-splight/utils"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string = "dev"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: utils.HostnameDefaultFunc(),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: utils.TokenDefaultFunc(),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"splight_alert":                       resourceAlert(),
			"splight_asset":                       resourceAsset(),
			"splight_asset_attribute":             resourceAssetAttribute(),
			"splight_asset_metadata":              resourceAssetMetadata(),
			"splight_component":                   resourceComponent(),
			"splight_component_routine":           resourceComponentRoutine(),
			"splight_dashboard":                   resourceDashboard(),
			"splight_dashboard_tab":               resourceDashboardTab(),
			"splight_dashboard_table_chart":       resourceDashboardTableChart(),
			"splight_dashboard_timeseries_chart":  resourceDashboardTimeseriesChart(),
			"splight_dashboard_histogram_chart":   resourceDashboardHistogramChart(),
			"splight_dashboard_image_chart":       resourceDashboardImageChart(),
			"splight_dashboard_bargauge_chart":    resourceDashboardBarGaugeChart(),
			"splight_dashboard_alertevents_chart": resourceDashboardAlertEventsChart(),
			"splight_dashboard_commandlist_chart": resourceDashboardCommandListChart(),
			"splight_dashboard_text_chart":        resourceDashboardTextChart(),
			"splight_dashboard_bar_chart":         resourceDashboardBarChart(),
			"splight_dashboard_stat_chart":        resourceDashboardStatChart(),
			"splight_dashboard_gauge_chart":       resourceDashboardGaugeChart(),
			"splight_dashboard_alertlist_chart":   resourceDashboardAlertListChart(),
			"splight_dashboard_assetlist_chart":   resourceDashboardAssetListChart(),
			"splight_dashboard_actionlist_chart":  resourceDashboardActionListChart(),
			"splight_file":                        resourceFile(),
			"splight_file_folder":                 resourceFileFolder(),
			"splight_function":                    resourceFunction(),
			"splight_secret":                      resourceSecret(),
			"splight_node":                        resourceNode(),
			"splight_tag":                         resourceTag(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"splight_asset_kinds": dataSourceAssetKind(),
			"splight_tags":        dataSourceTag(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	var diags diag.Diagnostics

	hostname := d.Get("hostname").(string)
	token := d.Get("token").(string)

	userAgentOptions := client.UserAgent{
		ProductName:    "terraform-provider-splight",
		ProductVersion: Version,
	}

	client, err := client.NewClient(hostname, token, ctx, userAgentOptions)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to create client",
			Detail:   fmt.Sprintf("Error creating client: %v", err),
		})
		return nil, diags
	}

	// Return the client and no diagnostics if successful
	return client, diags
}
