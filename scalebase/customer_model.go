package scalebase

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomerModel struct {
	OptionalID types.String `tfsdk:"optional_id"`

	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
