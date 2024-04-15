package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
var _ resource.Resource = &AssetMetadataResource{}
var _ resource.ResourceWithImportState = &AssetMetadataResource{}

func NewAssetMetadataResource() resource.Resource {
	return &AssetMetadataResource{}
}

type AssetMetadataResource struct {
	client *client.Client
}

func (r *AssetMetadataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset_metadata"
}

func (r *AssetMetadataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required:    true,
				Description: "name of the resource",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "[String|Boolean|Number] type of the data to be ingested in this metadata",
				Validators: []validator.String{
					stringvalidator.OneOf("Boolean", "Number", "String"),
				},
			},
			"unit": schema.StringAttribute{
				Required:    false,
				Optional:    true,
				Description: "optional reference to the unit of the measure",
			},
			"asset": schema.StringAttribute{
				Required:    true,
				Description: "reference to the asset to be linked to",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "metadata value",
			},
		},
	}
}

type AssetMetadataResourceParams struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Type  types.String `tfsdk:"type"`
	Unit  types.String `tfsdk:"unit"`
	Asset types.String `tfsdk:"asset"`
	Value types.String `tfsdk:"value"`
}

func (data *AssetMetadataResourceParams) ToAssetMetadataParams() client.AssetMetadataParams {
	item := client.AssetMetadataParams{
		Name:  data.Name.ValueString(),
		Type:  data.Type.ValueString(),
		Unit:  data.Unit.ValueString(),
		Asset: data.Asset.ValueString(),
	}

	return item
}

func (r *AssetMetadataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	// Get client from provider
	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		// FIX: change error
		resp.Diagnostics.AddError(
			"Client error",
			fmt.Sprintf("Unable to retrieve client for Splight API: %s", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AssetMetadataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AssetMetadataResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	item := data.ToAssetMetadataParams()
	createdMetadata, err := r.client.CreateAssetMetadata(&item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Asset, got error: %s", err))
		return
	}

	data.Id = types.StringValue(createdMetadata.Id)
	data.Name = types.StringValue(createdMetadata.Name)
	data.Type = types.StringValue(createdMetadata.Type)
	data.Unit = types.StringValue(createdMetadata.Unit)
	data.Asset = types.StringValue(createdMetadata.Asset)

	tflog.Trace(ctx, "created an AssetMetadata")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetMetadataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AssetMetadataResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()

	retrievedAssetMetadata, err := r.client.RetrieveAssetMetadata(id)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to retrieve AssetMetadata, got error: %s", err))
		return
	}

	data.Id = types.StringValue(retrievedAssetMetadata.Id)
	data.Name = types.StringValue(retrievedAssetMetadata.Name)
	data.Type = types.StringValue(retrievedAssetMetadata.Type)
	data.Unit = types.StringValue(retrievedAssetMetadata.Unit)
	data.Asset = types.StringValue(retrievedAssetMetadata.Asset)

	tflog.Trace(ctx, "retrieved an AssetMetadata")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetMetadataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AssetMetadataResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	item := data.ToAssetMetadataParams()
	updatedMetadata, err := r.client.UpdateAssetMetadata(id, &item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AssetMetadata, got error: %s", err))
		return
	}

	data.Id = types.StringValue(updatedMetadata.Id)
	data.Name = types.StringValue(updatedMetadata.Name)
	data.Type = types.StringValue(updatedMetadata.Type)
	data.Unit = types.StringValue(updatedMetadata.Unit)
	data.Asset = types.StringValue(updatedMetadata.Asset)

	tflog.Trace(ctx, "updated an AssetMetadata")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetMetadataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AssetMetadataResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()

	err := r.client.DeleteAssetMetadata(id)

	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("Unable to delete AssetMetadata with id '%s': %s", id, err))
		return
	}
}

func (r *AssetMetadataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
