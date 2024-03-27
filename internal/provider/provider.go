package provider

import (
	"context"
	"terraform-provider-hashicups/api/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure SplightProvider satisfies various provider interfaces.
var _ provider.Provider = &SplightProvider{}

/*
SplightProvider defines the provider implementation.
Version is set to:
- the provider version on release
- "dev" when the provider is built and ran locally
- "test" when running acceptance testing.
*/
type SplightProvider struct {
	version string
}

func (p *SplightProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// TODO: fill
		NewExampleResource,
	}
}

func (p *SplightProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// TODO: fill
		NewExampleDataSource,
	}
}

// This is the prefix for each resource and data source
func (p *SplightProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "spl"
	resp.Version = p.version
}

// This is the provider configuration
func (p *SplightProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Splight API host",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "Splight <ACCESS_ID> <SECRET_KEY>",
				Optional:            true,
			},
		},
	}
}

type ProviderConfig struct {
	Hostname types.String `tfsdk:"hostname"`
	Token    types.String `tfsdk:"token"`
}

func (p *SplightProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ProviderConfig

	var hostname string
	var token string

	// Load provider configuration from the terraform file
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Hostname.IsNull() {
		var err error

		hostname, err = LoadSplightHostname()
		if err != nil {
			tflog.Error(ctx, "Unable to configure splight hostname")
			return
		}

	} else {
		hostname = data.Hostname.String()
	}

	if data.Token.IsNull() {
		var err error

		token, err = LoadSplightToken()
		if err != nil {
			tflog.Error(ctx, "Unable to configure splight token")
			return
		}

	} else {
		token = data.Token.String()
	}

	client := client.NewClient(hostname, token)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SplightProvider{
			version: version,
		}
	}
}
