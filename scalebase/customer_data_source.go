package scalebase

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
	"github.com/kzmake/terraform-provider-scalebase/api/client/customer_service"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

var (
	_ datasource.DataSource              = &customerDataSource{}
	_ datasource.DataSourceWithConfigure = &customerDataSource{}
)

func NewCustomerDataSource() datasource.DataSource {
	return &customerDataSource{}
}

type customerDataSource struct {
	client *apiclient.API
}

type customerModel struct {
	ID         types.String `tfsdk:"id"`
	OptionalID types.String `tfsdk:"optional_id"`
	Name       types.String `tfsdk:"name"`
}

func (d *customerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*apiclient.API)
}

func (d *customerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_customer"
}

func (d *customerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"optional_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *customerDataSource) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {
	tflog.Debug(ctx, "Call api: GetCustomer")

	target := new(customerModel)
	req.Config.Get(ctx, target)
	in := &models.V1GetCustomerRequest{OptionalID: target.OptionalID.ValueString()}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := d.client.CustomerService.CustomerServiceGetCustomer(customer_service.NewCustomerServiceGetCustomerParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Unable to read scalebase customer",
			err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  response: "+spew.Sdump(out.Payload))

	v := out.Payload.Customer
	state := customerModel{
		ID:         types.StringValue(v.ID),
		OptionalID: types.StringValue(v.OptionalID),
		Name:       types.StringValue(v.Name),
	}

	diags := res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}
