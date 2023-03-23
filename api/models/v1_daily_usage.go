// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DailyUsage 使用量(日毎)
//
// swagger:model v1DailyUsage
type V1DailyUsage struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this v1 daily usage
func (m *V1DailyUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 daily usage based on context it is used
func (m *V1DailyUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DailyUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DailyUsage) UnmarshalBinary(b []byte) error {
	var res V1DailyUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}