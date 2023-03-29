package scalebase

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	apiclient "github.com/kzmake/terraform-provider-scalebase/api/client"
	"github.com/kzmake/terraform-provider-scalebase/api/client/catalog_item_service"
	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

var (
	_ datasource.DataSource              = &catalogItemDataSource{}
	_ datasource.DataSourceWithConfigure = &catalogItemDataSource{}
)

func NewCatalogItemDataSource() datasource.DataSource {
	return &catalogItemDataSource{}
}

type catalogItemDataSource struct {
	client *apiclient.API
}

func (d *catalogItemDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*apiclient.API)
}

func (d *catalogItemDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_catalog_item"
}

func (d *catalogItemDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"optional_id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"product_id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *catalogItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {
	tflog.Debug(ctx, "Call api: GetCatalogItem")

	var config CatalogItemModel
	diags := req.Config.Get(ctx, &config)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	in := &models.V1GetCatalogItemRequest{ID: config.ID.ValueStringPointer()}

	tflog.Debug(ctx, "  request: "+spew.Sdump(in))

	out, err := d.client.CatalogItemService.CatalogItemServiceGetCatalogItem(catalog_item_service.NewCatalogItemServiceGetCatalogItemParams().WithBody(in))
	if err != nil {
		res.Diagnostics.AddError(
			"Unable to read scalebase_catalog_item",
			err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "  response: "+spew.Sdump(out.Payload))

	v := out.Payload.CatalogItem
	state := CatalogItemModel{
		ID:         types.StringValue(v.ID),
		OptionalID: types.StringValue(v.OptionalID),
		Name:       types.StringValue(v.Name),
		ProductID:  types.StringValue(v.ProductID),
		Status:     types.StringValue(string(v.Status)),
	}

	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}
