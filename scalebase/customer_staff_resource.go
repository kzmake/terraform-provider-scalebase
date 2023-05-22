package scalebase

import (
	"context"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
	"github.com/kzmake/terraform-provider-scalebase/api/client/customer_staff_service"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

var (
	_ resource.Resource                = &customerStaffResource{}
	_ resource.ResourceWithConfigure   = &customerStaffResource{}
	_ resource.ResourceWithImportState = &customerStaffResource{}
)

func NewCustomerStaffResource() resource.Resource {
	return &customerStaffResource{}
}

type customerStaffResource struct {
	client *apiclient.API
}

func (r *customerStaffResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*apiclient.API)
}

func (r *customerStaffResource) Metadata(_ context.Context, req resource.MetadataRequest, res *resource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_customer_staff"
}

func (r *customerStaffResource) Schema(_ context.Context, _ resource.SchemaRequest, res *resource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required: true,
			},
			"optional_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"first_name": schema.StringAttribute{
						Optional: true,
					},
					"last_name": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"email_address": schema.StringAttribute{
				Optional: true,
			},
			"phone_number": schema.StringAttribute{
				Optional: true,
			},
			"department": schema.StringAttribute{
				Optional: true,
			},
			"title": schema.StringAttribute{
				Optional: true,
			},
			"address": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zip_code": schema.StringAttribute{
						Optional: true,
					},
					"country": schema.StringAttribute{
						Optional: true,
					},
					"prefecture": schema.StringAttribute{
						Optional: true,
					},
					"city": schema.StringAttribute{
						Optional: true,
					},
					"address_lines": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (r *customerStaffResource) ImportState(ctx context.Context, req resource.ImportStateRequest, res *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("optional_id"), req, res)
}

func (r *customerStaffResource) Create(ctx context.Context, req resource.CreateRequest, res *resource.CreateResponse) {
	tflog.Debug(ctx, "Call api: CreateCustomerStaff")

	var plan CustomerStaffModel
	diags := req.Plan.Get(ctx, &plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	var name *models.V1CreateCustomerStaffRequestName
	if plan.Name != nil {
		name = &models.V1CreateCustomerStaffRequestName{
			FirstName: plan.Name.FirstName.ValueString(),
			LastName:  plan.Name.LastName.ValueString(),
		}
	}

	var address *models.V1CreateCustomerStaffRequestAddress
	if plan.Address != nil {
		addressLines := []string{}
		for _, v := range plan.Address.AddressLines.Elements() {
			addressLines = append(addressLines, v.(types.String).ValueString())
		}

		address = &models.V1CreateCustomerStaffRequestAddress{
			ZipCode:      *plan.Address.ZipCode.ValueStringPointer(),
			Country:      models.V1CreateCustomerStaffRequestAddressCountry(plan.Address.Country.ValueString()),
			Prefecture:   models.V1CreateCustomerStaffRequestAddressPrefecture(plan.Address.Prefecture.ValueString()),
			City:         plan.Address.City.ValueString(),
			AddressLines: addressLines,
		}
	}

	in := &models.V1CreateCustomerStaffRequest{
		CustomerID:   plan.CustomerID.ValueStringPointer(),
		OptionalID:   plan.OptionalID.ValueString(),
		Name:         name,
		EmailAddress: plan.EmailAddress.ValueString(),
		PhoneNumber:  plan.PhoneNumber.ValueString(),
		Department:   plan.Department.ValueString(),
		Title:        plan.Title.ValueString(),
		Address:      address,
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerStaffService.CustomerStaffServiceCreateCustomerStaff(customer_staff_service.NewCustomerStaffServiceCreateCustomerStaffParams().WithBody(in))
	if err != nil {
		errRes, ok := err.(*customer_staff_service.CustomerStaffServiceCreateCustomerStaffDefault)
		if ok {
			res.Diagnostics.AddError(
				"Error creating customerStaff",
				"Could not create customerStaff, unexpected error: "+spew.Sdump(errRes.Payload.Details),
			)
		}
		res.Diagnostics.AddError(
			"Error creating customerStaff",
			"Could not create customerStaff, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.CustomerStaff

	println(spew.Sdump(v))

	plan.OptionalID = types.StringValue(v.OptionalID)
	plan.ID = types.StringValue(v.ID)

	var pName *CustomerStaffName
	if v.Name != nil {
		pName = &CustomerStaffName{
			FirstName: types.StringValue(v.Name.FirstName),
			LastName:  types.StringValue(v.Name.LastName),
		}
	}
	plan.Name = pName
	plan.EmailAddress = types.StringValue(v.EmailAddress)
	plan.PhoneNumber = types.StringValue(v.PhoneNumber)
	plan.Department = types.StringValue(v.Department)
	plan.Title = types.StringValue(v.Title)
	var pAddress *CustomerStaffAddress
	if v.Address != nil {
		pAddressLines := []attr.Value{}
		for _, line := range v.Address.AddressLines {
			pAddressLines = append(pAddressLines, types.StringValue(line))
		}
		pAddress = &CustomerStaffAddress{
			ZipCode:      types.StringValue(v.Address.ZipCode),
			Country:      types.StringValue(string(v.Address.Country)),
			Prefecture:   types.StringValue(string(v.Address.Prefecture)),
			City:         types.StringValue(v.Address.City),
			AddressLines: types.ListValueMust(types.StringType, pAddressLines),
		}
	}
	plan.Address = pAddress

	diags = res.State.Set(ctx, plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *customerStaffResource) Read(ctx context.Context, req resource.ReadRequest, res *resource.ReadResponse) {
	tflog.Debug(ctx, "Call api: GetCustomerStaff")

	var state CustomerStaffModel
	diags := req.State.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.V1GetCustomerStaffRequest{
		CustomerID: state.CustomerID.ValueStringPointer(),
		OptionalID: state.OptionalID.ValueString(),
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerStaffService.CustomerStaffServiceGetCustomerStaff(customer_staff_service.NewCustomerStaffServiceGetCustomerStaffParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error reading scalebase_customerStaff",
			"Could not read scalebase_customerStaff OptionalID "+state.OptionalID.ValueString()+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.CustomerStaff

	state.OptionalID = types.StringValue(v.OptionalID)
	state.ID = types.StringValue(v.ID)
	var sName *CustomerStaffName
	if v.Name != nil {
		sName = &CustomerStaffName{
			FirstName: types.StringValue(v.Name.FirstName),
			LastName:  types.StringValue(v.Name.LastName),
		}
	}
	state.Name = sName
	state.EmailAddress = types.StringValue(v.EmailAddress)
	state.PhoneNumber = types.StringValue(v.PhoneNumber)
	state.Department = types.StringValue(v.Department)
	state.Title = types.StringValue(v.Title)
	var sAddress *CustomerStaffAddress
	if v.Address != nil {
		sAddressLines := []attr.Value{}
		for _, line := range v.Address.AddressLines {
			sAddressLines = append(sAddressLines, types.StringValue(line))
		}
		sAddress = &CustomerStaffAddress{
			ZipCode:      types.StringValue(v.Address.ZipCode),
			Country:      types.StringValue(string(v.Address.Country)),
			Prefecture:   types.StringValue(string(v.Address.Prefecture)),
			City:         types.StringValue(v.Address.City),
			AddressLines: types.ListValueMust(types.StringType, sAddressLines),
		}
	}
	state.Address = sAddress

	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *customerStaffResource) Update(ctx context.Context, req resource.UpdateRequest, res *resource.UpdateResponse) {
	tflog.Debug(ctx, "Call api: UpdateCustomerStaff")

	var plan CustomerStaffModel
	diags := req.Plan.Get(ctx, &plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	var name *models.Publicv1CustomerStaffName
	if plan.Name != nil {
		name = &models.Publicv1CustomerStaffName{
			FirstName: plan.Name.FirstName.ValueString(),
			LastName:  plan.Name.LastName.ValueString(),
		}
	}
	var address *models.Publicv1CustomerStaffAddress
	if plan.Address != nil {
		addressLines := []string{}
		for _, v := range plan.Address.AddressLines.Elements() {
			addressLines = append(addressLines, v.(types.String).ValueString())
		}

		address = &models.Publicv1CustomerStaffAddress{
			ZipCode:      *plan.Address.ZipCode.ValueStringPointer(),
			Country:      models.Publicv1CustomerStaffAddressCountry(plan.Address.Country.ValueString()),
			Prefecture:   models.Publicv1CustomerStaffAddressPrefecture(plan.Address.Prefecture.ValueString()),
			City:         plan.Address.City.ValueString(),
			AddressLines: addressLines,
		}
	}

	in := &models.Publicv1CustomerStaff{
		CustomerID: plan.CustomerID.ValueString(),
		ID:         plan.ID.ValueString(),
		OptionalID: plan.OptionalID.ValueString(),

		Name:         name,
		EmailAddress: plan.EmailAddress.ValueString(),
		PhoneNumber:  plan.PhoneNumber.ValueString(),
		Department:   plan.Department.ValueString(),
		Title:        plan.Title.ValueString(),
		Address:      address,
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerStaffService.CustomerStaffServiceUpdateCustomerStaff(customer_staff_service.NewCustomerStaffServiceUpdateCustomerStaffParams().WithCustomerStaff(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error updating customerStaff",
			"Could not update customerStaff, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  resonse: "+spew.Sdump(out.Payload))

	v := out.Payload.CustomerStaff

	plan.OptionalID = types.StringValue(v.OptionalID)
	plan.ID = types.StringValue(v.ID)
	var pName *CustomerStaffName
	if v.Name != nil {
		pName = &CustomerStaffName{
			FirstName: types.StringValue(v.Name.FirstName),
			LastName:  types.StringValue(v.Name.LastName),
		}
	}
	plan.Name = pName
	plan.EmailAddress = types.StringValue(v.EmailAddress)
	plan.PhoneNumber = types.StringValue(v.PhoneNumber)
	plan.Department = types.StringValue(v.Department)
	plan.Title = types.StringValue(v.Title)
	var pAddress *CustomerStaffAddress
	if v.Address != nil {
		pAddressLines := []attr.Value{}
		for _, line := range v.Address.AddressLines {
			pAddressLines = append(pAddressLines, types.StringValue(line))
		}
		pAddress = &CustomerStaffAddress{
			ZipCode:      types.StringValue(v.Address.ZipCode),
			Country:      types.StringValue(string(v.Address.Country)),
			Prefecture:   types.StringValue(string(v.Address.Prefecture)),
			City:         types.StringValue(v.Address.City),
			AddressLines: types.ListValueMust(types.StringType, pAddressLines),
		}
	}
	plan.Address = pAddress

	diags = res.State.Set(ctx, plan)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}

func (r *customerStaffResource) Delete(ctx context.Context, req resource.DeleteRequest, res *resource.DeleteResponse) {
	tflog.Debug(ctx, "Call api: DeleteCustomerStaff")

	var plan CustomerStaffModel

	// Read Terraform prior state data into the model
	res.Diagnostics.Append(req.State.Get(ctx, &plan)...)

	in := &models.V1DeleteCustomerStaffRequest{
		CustomerID: plan.CustomerID.ValueStringPointer(),
		OptionalID: plan.OptionalID.ValueString(),
	}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := r.client.CustomerStaffService.CustomerStaffServiceDeleteCustomerStaff(customer_staff_service.NewCustomerStaffServiceDeleteCustomerStaffParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Error deleting customerStaff",
			"Could not delete customerStaff, unexpected error: "+err.Error(),
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
