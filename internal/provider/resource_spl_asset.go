package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/splightplatform/terraform-provider-splight/api/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AssetResource{}
var _ resource.ResourceWithImportState = &AssetResource{}

func NewAssetResource() resource.Resource {
	return &AssetResource{}
}

type AssetResource struct {
	client *client.Client
}

func (r *AssetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset"
}

func (r *AssetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Asset resource",
		Attributes: map[string]schema.Attribute{
			// Read only
			"id": schema.StringAttribute{
				MarkdownDescription: "id of the resource",
				Required:            false,
				Optional:            false,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "name of the resource",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "description of the resource",
				Optional:            true,
			},
			"geometry": schema.StringAttribute{
				CustomType:          jsontypes.NormalizedType{},
				Optional:            true,
				MarkdownDescription: "geojson compliant geometry collection JSON (see: https://datatracker.ietf.org/doc/html/rfc7946#section-3.1.8)",
				Validators: []validator.String{
					geoJSONGeometryCollectionValidator{},
				},
			},
			"related_assets": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "related assets ids",
			},
		},
	}
}

type AssetResourceParams struct {
	Id            types.String         `tfsdk:"id"`
	Name          types.String         `tfsdk:"name"`
	Description   types.String         `tfsdk:"description"`
	RelatedAssets types.Set            `tfsdk:"related_assets"`
	Geometry      jsontypes.Normalized `tfsdk:"geometry"`
}

func (data *AssetResourceParams) ToAssetParams(ctx context.Context) *client.AssetParams {
	// Convert related assets input set to the API format
	// from: {id1, id2, id3}
	// to:
	// [
	//  {
	//    id: <id>
	//   },
	//   ...
	// ]
	var relatedAssetsSet []types.String
	data.RelatedAssets.ElementsAs(ctx, &relatedAssetsSet, false)
	assetRelatedAssets := make([]client.RelatedAsset, len(relatedAssetsSet))
	for i, relatedAsset := range relatedAssetsSet {
		assetRelatedAssets[i] = client.RelatedAsset{
			Id: relatedAsset.ValueString(),
		}
	}

	item := client.AssetParams{
		Name:          data.Name.ValueString(),
		Description:   data.Description.ValueString(),
		Geometry:      json.RawMessage(data.Geometry.ValueString()),
		RelatedAssets: assetRelatedAssets,
	}

	return &item
}

func (data *AssetResourceParams) FromAsset(ctx context.Context, params *client.Asset) diag.Diagnostics {
	var diags diag.Diagnostics

	data.Id = types.StringValue(params.Id)
	data.Name = types.StringValue(params.Name)
	data.Description = types.StringValue(params.Description)
	relatedAssets := make([]string, len(params.RelatedAssets))
	for _, relatedAsset := range params.RelatedAssets {
		relatedAssets = append(relatedAssets, relatedAsset.Id)

	}
	data.RelatedAssets, diags = types.SetValueFrom(ctx, types.StringType, relatedAssets)
	data.Geometry = jsontypes.NewNormalizedValue(string(params.Geometry))

	return diags
}

func (r *AssetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	// Get client from provider
	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Client error",
			fmt.Sprintf("Unable to retrieve client for Splight API: %s", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AssetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AssetResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	item := data.ToAssetParams(ctx)
	createdAsset, err := r.client.CreateAsset(item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Asset, got error: %s", err))
		return
	}

	data.Id = types.StringValue(createdAsset.Id)

	tflog.Trace(ctx, "created an Asset")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AssetResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()

	retrievedAsset, err := r.client.RetrieveAsset(id)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to retrieve Asset, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(data.FromAsset(ctx, retrievedAsset)...)

	tflog.Trace(ctx, "retrieved an Asset")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AssetResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	item := data.ToAssetParams(ctx)
	_, err := r.client.UpdateAsset(id, item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Asset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated an Asset")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AssetResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	err := r.client.DeleteAsset(id)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("Unable to delete Asset with id '%s': %s", id, err))
		return
	}
}

func (r *AssetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
