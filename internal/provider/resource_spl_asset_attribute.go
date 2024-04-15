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
var _ resource.Resource = &AssetAttributeResource{}
var _ resource.ResourceWithImportState = &AssetAttributeResource{}

func NewAssetAttributeResource() resource.Resource {
	return &AssetAttributeResource{}
}

type AssetAttributeResource struct {
	client *client.Client
}

func (r *AssetAttributeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset_attribute"
}

func (r *AssetAttributeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description: "[String|Boolean|Number] type of the data to be ingested in this attribute",
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
		},
	}
}

type AssetAttributeResourceParams struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Type  types.String `tfsdk:"type"`
	Unit  types.String `tfsdk:"unit"`
	Asset types.String `tfsdk:"asset"`
}

func (data *AssetAttributeResourceParams) ToAssetAttributeParams() *client.AssetAttributeParams {
	item := client.AssetAttributeParams{
		Name:  data.Name.ValueString(),
		Type:  data.Type.ValueString(),
		Unit:  data.Unit.ValueString(),
		Asset: data.Asset.ValueString(),
	}

	return &item
}

func (data *AssetAttributeResourceParams) FromAssetAttribute(assetAttribute client.AssetAttribute) {
	data.Id = types.StringValue(assetAttribute.Id)
	data.Name = types.StringValue(assetAttribute.Name)
	data.Type = types.StringValue(assetAttribute.Type)
	data.Unit = types.StringValue(assetAttribute.Unit)
	data.Asset = types.StringValue(assetAttribute.Asset)
}

func (r *AssetAttributeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AssetAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AssetAttributeResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	item := data.ToAssetAttributeParams()
	createdAttribute, err := r.client.CreateAssetAttribute(item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Asset, got error: %s", err))
		return
	}

	data.Id = types.StringValue(createdAttribute.Id)

	tflog.Trace(ctx, "created an AssetAttribute")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AssetAttributeResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	retrievedAssetAttribute, err := r.client.RetrieveAssetAttribute(id)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to retrieve AssetAttribute, got error: %s", err))
		return
	}

	data.FromAssetAttribute(*retrievedAssetAttribute)

	tflog.Trace(ctx, "retrieved an AssetAttribute")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AssetAttributeResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	item := data.ToAssetAttributeParams()
	_, err := r.client.UpdateAssetAttribute(id, item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AssetAttribute, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated an AssetAttribute")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AssetAttributeResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	err := r.client.DeleteAssetAttribute(id)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("Unable to delete AssetAttribute with id '%s': %s", id, err))
		return
	}
}

func (r *AssetAttributeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
