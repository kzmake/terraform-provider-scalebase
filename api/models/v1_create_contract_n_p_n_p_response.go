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

// V1CreateContractNPNPResponse 契約作成のレスポンス
//
// swagger:model v1CreateContractNPNPResponse
type V1CreateContractNPNPResponse struct {

	// 契約
	Contract *Publicv1Contract `json:"contract,omitempty"`
}

// Validate validates this v1 create contract n p n p response
func (m *V1CreateContractNPNPResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateContractNPNPResponse) validateContract(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 create contract n p n p response based on the context it is used
func (m *V1CreateContractNPNPResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContract(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateContractNPNPResponse) contextValidateContract(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1CreateContractNPNPResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateContractNPNPResponse) UnmarshalBinary(b []byte) error {
	var res V1CreateContractNPNPResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
