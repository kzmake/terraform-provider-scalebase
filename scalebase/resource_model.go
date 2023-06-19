package scalebase

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceModel struct {
	SRN SRN `tfsdk:"srn"`

	CustomFields []*CustomField `tfsdk:"custom_fields"`
}

type SRN struct {
	ID   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type CustomField struct {
	ID types.String `tfsdk:"id"`

	String types.String `tfsdk:"string"`
	Select types.String `tfsdk:"select"`
	Date   types.String `tfsdk:"date"`
}

func (r *SRN) String() string { return fmt.Sprintf("%s:%s", r.Type, r.ID) }
