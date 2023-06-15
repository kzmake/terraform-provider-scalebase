// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1AddOptionItemRequestTerm 基本契約期間
//
// swagger:model v1AddOptionItemRequestTerm
type V1AddOptionItemRequestTerm struct {

	// 単位
	Unit V1AddOptionItemRequestTermUnit `json:"unit,omitempty"`

	// 期間数
	// Required: true
	Value *int32 `json:"value"`
}

// Validate validates this v1 add option item request term
func (m *V1AddOptionItemRequestTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AddOptionItemRequestTerm) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if err := m.Unit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("unit")
		}
		return err
	}

	return nil
}

func (m *V1AddOptionItemRequestTerm) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 add option item request term based on the context it is used
func (m *V1AddOptionItemRequestTerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AddOptionItemRequestTerm) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if err := m.Unit.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("unit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AddOptionItemRequestTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AddOptionItemRequestTerm) UnmarshalBinary(b []byte) error {
	var res V1AddOptionItemRequestTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
