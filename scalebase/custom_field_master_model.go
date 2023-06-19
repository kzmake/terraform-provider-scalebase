package scalebase

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomFieldMasterModel struct {
	Name      types.String `tfsdk:"name"`
	FieldType types.String `tfsdk:"field_type"`

	ID       types.String `tfsdk:"id"`
	DataType types.String `tfsdk:"data_type"`
}
