package scalebase

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CatalogItemModel struct {
	ID types.String `tfsdk:"id"`

	OptionalID types.String `tfsdk:"optional_id"`
	Name       types.String `tfsdk:"name"`
	ProductID  types.String `tfsdk:"product_id"`
	Status     types.String `tfsdk:"status"`
}
