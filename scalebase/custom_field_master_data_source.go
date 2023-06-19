package scalebase

import (
	"context"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
	"github.com/kzmake/terraform-provider-scalebase/api/client/custom_field_service"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

var (
	_ datasource.DataSource              = &customFieldMasterDataSource{}
	_ datasource.DataSourceWithConfigure = &customFieldMasterDataSource{}
)

func NewCustomFieldMasterDataSource() datasource.DataSource {
	return &customFieldMasterDataSource{}
}

type customFieldMasterDataSource struct {
	client *apiclient.API
}

func (d *customFieldMasterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*apiclient.API)
}

func (d *customFieldMasterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_custom_field_master"
}

func (d *customFieldMasterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"field_type": schema.StringAttribute{
				Required: true,
			},

			"id": schema.StringAttribute{
				Computed: true,
			},
			"data_type": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *customFieldMasterDataSource) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {
	tflog.Debug(ctx, "Call api: GetCustomFieldMaster")

	var config CustomFieldMasterModel
	diags := req.Config.Get(ctx, &config)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.V1ListCustomFieldMastersRequest{PageSize: func(a int32) *int32 { return &a }(100)}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := d.client.CustomFieldService.CustomFieldServiceListCustomFieldMasters(custom_field_service.NewCustomFieldServiceListCustomFieldMastersParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Unable to read scalebase custom field master",
			err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  response: "+spew.Sdump(out.Payload))

	var v *models.V1CustomFieldMaster
	for _, x := range out.Payload.CustomFieldMasters {
		if x.Name == config.Name.ValueString() && strings.ToLower(strings.TrimPrefix(string(x.FieldType), "FIELD_TYPE_")) == config.FieldType.ValueString() {
			v = &models.V1CustomFieldMaster{
				Name:      x.Name,
				FieldType: x.FieldType,
				DataType:  x.DataType,
				ID:        x.ID,
			}
		}
	}
	state := CustomFieldMasterModel{
		Name:      types.StringValue(v.Name),
		FieldType: types.StringValue(strings.ToLower(strings.TrimPrefix(string(v.FieldType), "FIELD_TYPE_"))),

		ID:       types.StringValue(v.ID),
		DataType: types.StringValue(string(v.DataType)),
	}

	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}
