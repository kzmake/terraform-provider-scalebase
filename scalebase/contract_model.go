package scalebase

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ContractModel struct {
	OptionalID types.String `tfsdk:"optional_id"`

	ID types.String `tfsdk:"id"`
}
