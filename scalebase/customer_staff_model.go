package scalebase

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomerStaffModel struct {
	CustomerID types.String `tfsdk:"customer_id"`
	OptionalID types.String `tfsdk:"optional_id"`

	ID           types.String          `tfsdk:"id"`
	Name         *CustomerStaffName    `tfsdk:"name"`
	EmailAddress types.String          `tfsdk:"email_address"`
	PhoneNumber  types.String          `tfsdk:"phone_number"`
	Department   types.String          `tfsdk:"department"`
	Title        types.String          `tfsdk:"title"`
	Address      *CustomerStaffAddress `tfsdk:"address"`
}

type CustomerStaffName struct {
	FirstName types.String `tfsdk:"first_name"`
	LastName  types.String `tfsdk:"last_name"`
}

type CustomerStaffAddress struct {
	ZipCode      types.String `tfsdk:"zip_code"`
	Country      types.String `tfsdk:"country"`
	Prefecture   types.String `tfsdk:"prefecture"`
	City         types.String `tfsdk:"city"`
	AddressLines types.List   `tfsdk:"address_lines"`
}
