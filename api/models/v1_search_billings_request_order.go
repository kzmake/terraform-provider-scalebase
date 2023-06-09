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

// V1SearchBillingsRequestOrder ソート条件
//
// swagger:model v1SearchBillingsRequestOrder
type V1SearchBillingsRequestOrder struct {

	// direction
	Direction V1SearchBillingsRequestOrderDirection `json:"direction,omitempty"`

	// field
	Field V1SearchBillingsRequestOrderField `json:"field,omitempty"`
}

// Validate validates this v1 search billings request order
func (m *V1SearchBillingsRequestOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchBillingsRequestOrder) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *V1SearchBillingsRequestOrder) validateField(formats strfmt.Registry) error {
	if swag.IsZero(m.Field) { // not required
		return nil
	}

	if err := m.Field.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("field")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("field")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 search billings request order based on the context it is used
func (m *V1SearchBillingsRequestOrder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateField(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchBillingsRequestOrder) contextValidateDirection(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Direction.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *V1SearchBillingsRequestOrder) contextValidateField(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Field.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("field")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("field")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchBillingsRequestOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchBillingsRequestOrder) UnmarshalBinary(b []byte) error {
	var res V1SearchBillingsRequestOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
