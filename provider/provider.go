package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/splightplatform/splight-terraform-provider/api/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.splight-ai.com",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"spl_function":          resourceFunction(),
			"spl_alert":             resourceAlert(),
			"spl_asset":             resourceAsset(),
			"spl_asset_attribute":   resourceAssetAttribute(),
			"spl_asset_metadata":    resourceAssetMetadata(),
			"spl_component":         resourceComponent(),
			"spl_component_routine": resourceComponentRoutine(),
			"spl_dashboard":         resourceDashboard(),
			"spl_dashboard_tab":     resourceDashboardTab(),
			"spl_dashboard_chart":   resourceDashboardChart(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("hostname").(string)
	token := d.Get("token").(string)
	return client.NewClient(hostname, token), nil

}
