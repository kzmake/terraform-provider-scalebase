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

// V1DeleteCustomerStaffRequest 顧客担当削除のリクエスト
//
// swagger:model v1DeleteCustomerStaffRequest
type V1DeleteCustomerStaffRequest struct {

	// 顧客ID
	// Required: true
	CustomerID *string `json:"customerId"`

	// 顧客管理ID(顧客IDの代わりに指定可)
	CustomerOptionalID string `json:"customerOptionalId,omitempty"`

	// 顧客担当者ID
	// Required: true
	ID *string `json:"id"`

	// 顧客担当者管理ID(顧客担当者IDの代わりに指定可)
	OptionalID string `json:"optionalId,omitempty"`
}

// Validate validates this v1 delete customer staff request
func (m *V1DeleteCustomerStaffRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DeleteCustomerStaffRequest) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customerId", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *V1DeleteCustomerStaffRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 delete customer staff request based on context it is used
func (m *V1DeleteCustomerStaffRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DeleteCustomerStaffRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeleteCustomerStaffRequest) UnmarshalBinary(b []byte) error {
	var res V1DeleteCustomerStaffRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
