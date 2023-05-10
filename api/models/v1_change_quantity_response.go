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

// V1ChangeQuantityResponse 契約数変更レスポンス
//
// swagger:model v1ChangeQuantityResponse
type V1ChangeQuantityResponse struct {

	// 対象の契約
	Contract *V1ChangeQuantityResponseContract `json:"contract,omitempty"`
}

// Validate validates this v1 change quantity response
func (m *V1ChangeQuantityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ChangeQuantityResponse) validateContract(formats strfmt.Registry) error {
	if swag.IsZero(m.Contract) { // not required
		return nil
	}

	if m.Contract != nil {
		if err := m.Contract.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contract")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contract")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 change quantity response based on the context it is used
func (m *V1ChangeQuantityResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContract(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ChangeQuantityResponse) contextValidateContract(ctx context.Context, formats strfmt.Registry) error {

	if m.Contract != nil {
		if err := m.Contract.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contract")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contract")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ChangeQuantityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ChangeQuantityResponse) UnmarshalBinary(b []byte) error {
	var res V1ChangeQuantityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
