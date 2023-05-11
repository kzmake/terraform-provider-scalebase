package scalebase

import (
	"context"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
	"github.com/kzmake/terraform-provider-scalebase/api/client/customer_service"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

var (
	_ resource.Resource                = &customerResource{}
	_ resource.ResourceWithConfigure   = &customerResource{}
	_ resource.ResourceWithImportState = &customerResource{}
)

func NewCustomerResource() resource.Resource {
	return &customerResource{}
}

type customerResource struct {
	client *apiclient.API
}

func (r *customerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*apiclient.API)
}

func (r *customerResource) Metadata(_ context.Context, req resource.MetadataRequest, res *resource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_customer"
}

func (r *customerResource) Schema(_ context.Context, _ resource.SchemaRequest, res *resource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"optional_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *customerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, res *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("optional_id"), req, res)
}

func (r *customerResource) Create(ctx context.Context, req resource.CreateRequest, res *resource.CreateResponse) {
	tflog.Debug(ctx, "Call api: CreateCustomer")

	var plan CustomerModel
	diags := req.Plan.Get(ctx, &plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.V1CreateCustomerRequest{
		OptionalID: plan.OptionalID.ValueString(),
		Name:       plan.Name.ValueStringPointer(),
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerService.CustomerServiceCreateCustomer(customer_service.NewCustomerServiceCreateCustomerParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error creating customer",
			"Could not create customer, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.Customer

	plan.OptionalID = types.StringValue(v.OptionalID)
	plan.ID = types.StringValue(v.ID)
	plan.Name = types.StringValue(v.Name)

	diags = res.State.Set(ctx, plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *customerResource) Read(ctx context.Context, req resource.ReadRequest, res *resource.ReadResponse) {
	tflog.Debug(ctx, "Call api: GetCustomer")

	var state CustomerModel
	diags := req.State.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.V1GetCustomerRequest{OptionalID: state.OptionalID.ValueString()}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerService.CustomerServiceGetCustomer(customer_service.NewCustomerServiceGetCustomerParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error reading scalebase_customer",
			"Could not read scalebase_customer OptionalID "+state.OptionalID.ValueString()+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.Customer

	state.OptionalID = types.StringValue(v.OptionalID)
	state.ID = types.StringValue(v.ID)
	state.Name = types.StringValue(v.Name)

	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *customerResource) Update(ctx context.Context, req resource.UpdateRequest, res *resource.UpdateResponse) {
	tflog.Debug(ctx, "Call api: UpdateCustomer")

	var plan CustomerModel
	diags := req.Plan.Get(ctx, &plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.Publicv1Customer{
		OptionalID: plan.OptionalID.ValueString(),

		Name: plan.Name.ValueString(),
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerService.CustomerServiceUpdateCustomer(customer_service.NewCustomerServiceUpdateCustomerParams().WithCustomer(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error updating customer",
			"Could not update customer, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.Customer

	plan.OptionalID = types.StringValue(v.OptionalID)
	plan.ID = types.StringValue(v.ID)
	plan.Name = types.StringValue(v.Name)

	diags = res.State.Set(ctx, plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *customerResource) Delete(ctx context.Context, req resource.DeleteRequest, res *resource.DeleteResponse) {
	tflog.Debug(ctx, "Call api: DeleteCustomer")

	var plan CustomerModel

	// Read Terraform prior state data into the model
	res.Diagnostics.Append(req.State.Get(ctx, &plan)...)

	in := &models.V1DeleteCustomerRequest{
		OptionalID: plan.OptionalID.ValueString(),
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerService.CustomerServiceDeleteCustomer(customer_service.NewCustomerServiceDeleteCustomerParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error deleting customer",
			"Could not delete customer, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	// Return error if the HTTP status code is not 200 OK or 404 Not Found
	if !out.IsCode(http.StatusNotFound) && !out.IsSuccess() {
		res.Diagnostics.AddError(
			"Unable to delete resource",
			"An unexpected error occurred while attempting to delete the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+string(rune(out.Code())),
		)

		return
	}
}
