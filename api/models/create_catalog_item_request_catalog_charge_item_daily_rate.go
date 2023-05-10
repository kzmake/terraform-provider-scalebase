// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateCatalogItemRequestCatalogChargeItemDailyRate create catalog item request catalog charge item daily rate
//
// swagger:model CreateCatalogItemRequestCatalogChargeItemDailyRate
type CreateCatalogItemRequestCatalogChargeItemDailyRate struct {

	// FIXED_RATE指定時の値
	FixedRateValue float64 `json:"fixedRateValue,omitempty"`

	// type
	Type CreateCatalogItemRequestCatalogChargeItemDailyRateType `json:"type,omitempty"`
}

// Validate validates this create catalog item request catalog charge item daily rate
func (m *CreateCatalogItemRequestCatalogChargeItemDailyRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCatalogItemRequestCatalogChargeItemDailyRate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this create catalog item request catalog charge item daily rate based on the context it is used
func (m *CreateCatalogItemRequestCatalogChargeItemDailyRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCatalogItemRequestCatalogChargeItemDailyRate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCatalogItemRequestCatalogChargeItemDailyRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCatalogItemRequestCatalogChargeItemDailyRate) UnmarshalBinary(b []byte) error {
	var res CreateCatalogItemRequestCatalogChargeItemDailyRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
