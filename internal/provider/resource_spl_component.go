package provider

import (
	"context"
	"errors"
	"fmt"
	"math/big"

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
var _ resource.Resource = &AssetResource{}
var _ resource.ResourceWithImportState = &AssetResource{}

func NewComponentResource() resource.Resource {
	return &ComponentResource{}
}

type ComponentResource struct {
	client *client.Client
}

func (r *ComponentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset"
}

func (r *ComponentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description: "the name of the component to be created",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "optinal description to add details of the resource",
			},
			"version": schema.StringAttribute{
				Required:    true,
				Description: "[NAME-VERSION] the version of the hub component",
			},
			"input": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
					},
					"description": schema.StringAttribute{
						Required: true,
					},
					"multiple": schema.BoolAttribute{
						Required: true,
					},
					"required": schema.BoolAttribute{
						Required: true,
					},
					"sensitive": schema.BoolAttribute{
						Required: true,
					},
					"type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf("Number", "Boolean", "String"),
						},
					},
					"value": schema.DynamicAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

type ComponentResourceInputParam struct {
	Name        types.String  `tfsdk:"name"`
	Description types.String  `tfsdk:"description"`
	Multiple    types.Bool    `tfsdk:"multiple"`
	Required    types.Bool    `tfsdk:"required"`
	Sensitive   types.Bool    `tfsdk:"sensitive"`
	Type        types.String  `tfsdk:"type"`
	Value       types.Dynamic `tfsdk:"value"`
}

type ComponentResourceParams struct {
	Id          types.String                  `tfsdk:"id"`
	Name        types.String                  `tfsdk:"name"`
	Description types.String                  `tfsdk:"description"`
	Version     types.String                  `tfsdk:"version"`
	Input       []ComponentResourceInputParam `tfsdk:"input"`
}

func (data *ComponentResourceParams) ToComponentParams() (*client.ComponentParams, error) {
	var inputParams []client.ComponentInputParam = make([]client.ComponentInputParam, len(data.Input))

	for _, resourceInputParam := range data.Input {

		var value interface{}

		switch tfValue := resourceInputParam.Value.UnderlyingValue().(type) {
		case types.Bool:
			value = tfValue.ValueBool()
		case types.Number:
			value = tfValue.ValueBigFloat()
		case types.String:
			value = tfValue.ValueString()
		default:
			return nil, errors.New("Metadata 'value' must be one of types [bool, str, float, int]")
		}

		inputParams = append(inputParams, client.ComponentInputParam{
			Name:        resourceInputParam.Name.ValueString(),
			Description: resourceInputParam.Description.ValueString(),
			Multiple:    resourceInputParam.Multiple.ValueBool(),
			Required:    resourceInputParam.Required.ValueBool(),
			Sensitive:   resourceInputParam.Sensitive.ValueBool(),
			Type:        resourceInputParam.Type.ValueString(),
			Value:       value,
		})

	}

	item := client.ComponentParams{
		Name:        data.Name.ValueString(),
		Description: data.Description.ValueString(),
		Version:     data.Description.ValueString(),
		Input:       inputParams,
	}

	return &item, nil
}

func (data *ComponentResourceParams) FromComponent(ctx context.Context, component *client.Component) error {

	var componentResourceInputParams []ComponentResourceInputParam = make([]ComponentResourceInputParam, len(component.Input))

	for _, componentInput := range component.Input {

		// Check the type of the decoded value
		var value types.Dynamic
		switch v := componentInput.Value.(type) {
		case string:
			value = types.DynamicValue(types.StringValue(v))
		case float64:
			value = types.DynamicValue(types.NumberValue(big.NewFloat(v)))
		case bool:
			value = types.DynamicValue(types.BoolValue(v))
		case []string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, v)
			value = types.DynamicValue(listValue)
		case []bool:
			listValue, _ := types.ListValueFrom(ctx, types.BoolType, v)
			value = types.DynamicValue(listValue)
		case []float64:
			listValue, _ := types.ListValueFrom(ctx, types.Float64Type, v)
			value = types.DynamicValue(listValue)
		default:
			return fmt.Errorf("unsupported value type: %T", v)
		}

		componentResourceInputParams = append(componentResourceInputParams, ComponentResourceInputParam{
			Name:        types.StringValue(componentInput.Name),
			Description: types.StringValue(componentInput.Description),
			Multiple:    types.BoolValue(componentInput.Multiple),
			Required:    types.BoolValue(componentInput.Required),
			Sensitive:   types.BoolValue(componentInput.Sensitive),
			Type:        types.StringValue(componentInput.Type),
			Value:       value,
		})

	}

	data.Name = types.StringValue(component.Name)
	data.Description = types.StringValue(component.Description)
	data.Version = types.StringValue(component.Version)
	data.Input = componentResourceInputParams

	return nil
}

func (r *ComponentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ComponentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ComponentResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	item, err := data.ToComponentParams()
	if err != nil {
		resp.Diagnostics.AddError("Argument error", fmt.Sprintf("Error while serializing to client: %s", err))
		return
	}

	createdComponent, err := r.client.CreateComponent(item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Component, got error: %s", err))
		return
	}

	data.Id = types.StringValue(createdComponent.Id)

	tflog.Trace(ctx, "created an Component")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComponentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ComponentResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()

	retrievedComponent, err := r.client.RetrieveComponent(id)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to retrieve Component, got error: %s", err))
		return
	}

	err = data.FromComponent(ctx, retrievedComponent)
	if err != nil {
		resp.Diagnostics.AddError("Argument error", fmt.Sprintf("Error while deserializing from client: %s", err))
		return
	}

	tflog.Trace(ctx, "retrieved an Component")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComponentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ComponentResourceParams

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()
	item, err := data.ToComponentParams()
	if err != nil {
		resp.Diagnostics.AddError("Argument error", fmt.Sprintf("Error while serializing to client: %s", err))
		return
	}
	_, err = r.client.UpdateComponent(id, item)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Component, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "updated an Component")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComponentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ComponentResourceParams

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.Id.ValueString()

	err := r.client.DeleteComponent(id)

	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("Unable to delete Component with id '%s': %s", id, err))
		return
	}
}

func (r *ComponentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
