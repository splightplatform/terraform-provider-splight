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
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
		ResourcesMap:         buildResourceMap(),
		DataSourcesMap:       buildDataSourceMap(),
		ConfigureContextFunc: providerConfigure,
	}
}

func buildResourceMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"splight_asset":                       resourceForType[*models.Asset](schemas.SchemaAsset),
		"splight_asset_relation":              resourceForType[*models.AssetRelation](schemas.SchemaAssetRelation),
		"splight_asset_attribute":             resourceForType[*models.AssetAttribute](schemas.SchemaAssetAttribute),
		"splight_asset_metadata":              resourceForType[*models.AssetMetadata](schemas.SchemaAssetMetadata),
		"splight_grid":                        resourceForType[*models.Grid](schemas.SchemaGrid),
		"splight_bus":                         resourceForType[*models.Bus](schemas.SchemaBus),
		"splight_line":                        resourceForType[*models.Line](schemas.SchemaLine),
		"splight_slack_line":                  resourceForType[*models.SlackLine](schemas.SchemaSlackLine),
		"splight_segment":                     resourceForType[*models.Segment](schemas.SchemaSegment),
		"splight_generator":                   resourceForType[*models.Generator](schemas.SchemaGenerator),
		"splight_slack_generator":             resourceForType[*models.SlackGenerator](schemas.SchemaSlackGenerator),
		"splight_inverter":                    resourceForType[*models.Inverter](schemas.SchemaInverter),
		"splight_tag":                         resourceForType[*models.Tag](schemas.SchemaTag),
		"splight_alert":                       resourceForType[*models.Alert](schemas.SchemaAlert),
		"splight_function":                    resourceForType[*models.Function](schemas.SchemaFunction),
		"splight_action":                      resourceForType[*models.Action](schemas.SchemaAction),
		"splight_command":                     resourceForType[*models.Command](schemas.SchemaCommand),
		"splight_algorithm":                   resourceForType[*models.Algorithm](schemas.SchemaAlgorithm),
		"splight_server":                      resourceForType[*models.Server](schemas.SchemaServer),
		"splight_connector":                   resourceForType[*models.Connector](schemas.SchemaConnector),
		"splight_component":                   resourceForType[*models.Component](schemas.SchemaComponent),
		"splight_component_routine":           resourceForType[*models.ComponentRoutine](schemas.SchemaComponentRoutine),
		"splight_dashboard":                   resourceForType[*models.Dashboard](schemas.SchemaDashboard),
		"splight_dashboard_tab":               resourceForType[*models.DashboardTab](schemas.SchemaDashboardTab),
		"splight_dashboard_table_chart":       resourceForType[*models.DashboardTableChart](schemas.SchemaDashboardTableChart),
		"splight_dashboard_timeseries_chart":  resourceForType[*models.DashboardTimeseriesChart](schemas.SchemaDashboardTimeseriesChart),
		"splight_dashboard_histogram_chart":   resourceForType[*models.DashboardHistogramChart](schemas.SchemaDashboardHistogramChart),
		"splight_dashboard_image_chart":       resourceForType[*models.DashboardImageChart](schemas.SchemaDashboardImageChart),
		"splight_dashboard_bargauge_chart":    resourceForType[*models.DashboardBarGaugeChart](schemas.SchemaDashboardBarGaugeChart),
		"splight_dashboard_alertevents_chart": resourceForType[*models.DashboardAlertEventsChart](schemas.SchemaDashboardAlertEventsChart),
		"splight_dashboard_commandlist_chart": resourceForType[*models.DashboardCommandListChart](schemas.SchemaDashboardCommandListChart),
		"splight_dashboard_text_chart":        resourceForType[*models.DashboardTextChart](schemas.SchemaDashboardTextChart),
		"splight_dashboard_bar_chart":         resourceForType[*models.DashboardBarChart](schemas.SchemaDashboardBarChart),
		"splight_dashboard_stat_chart":        resourceForType[*models.DashboardStatChart](schemas.SchemaDashboardStatChart),
		"splight_dashboard_gauge_chart":       resourceForType[*models.DashboardGaugeChart](schemas.SchemaDashboardGaugeChart),
		"splight_dashboard_alertlist_chart":   resourceForType[*models.DashboardAlertListChart](schemas.SchemaDashboardAlertListChart),
		"splight_dashboard_assetlist_chart":   resourceForType[*models.DashboardAssetListChart](schemas.SchemaDashboardAssetListChart),
		"splight_dashboard_actionlist_chart":  resourceForType[*models.DashboardActionListChart](schemas.SchemaDashboardActionListChart),
		"splight_external_grid":               resourceForType[*models.ExternalGrid](schemas.SchemaExternalGrid),
		"splight_file":                        resourceForType[*models.File](schemas.SchemaFile),
		"splight_file_folder":                 resourceForType[*models.FileFolder](schemas.SchemaFileFolder),
		"splight_secret":                      resourceForType[*models.Secret](schemas.SchemaSecret),
		"splight_node":                        resourceForType[*models.Node](schemas.SchemaNode, ResourceMethods{methods: NoUpdate}),
		"splight_transformer":                 resourceForType[*models.Transformer](schemas.SchemaTransformer),
	}
}

func buildDataSourceMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"splight_asset_kinds": dataSourceForType[*models.AssetKinds](schemas.SchemaAssetKinds),
		"splight_tags":        dataSourceForType[*models.Tags](schemas.SchemaTags),
		"splight_grids":       dataSourceForType[*models.Grid](schemas.SchemaTags),
		"splight_buses":       dataSourceForType[*models.Bus](schemas.SchemaTags),
		"splight_lines":       dataSourceForType[*models.Line](schemas.SchemaTags),
		"splight_generators":  dataSourceForType[*models.Generator](schemas.SchemaTags),
	}
}
