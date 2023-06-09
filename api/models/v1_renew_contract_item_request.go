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

// V1RenewContractItemRequest 契約アイテム更新リクエスト
//
// swagger:model v1RenewContractItemRequest
type V1RenewContractItemRequest struct {

	// 契約ID
	// Required: true
	ContractID *string `json:"contractId"`

	// 契約アイテムID
	// Required: true
	ContractItemID *string `json:"contractItemId"`
}

// Validate validates this v1 renew contract item request
func (m *V1RenewContractItemRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractItemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RenewContractItemRequest) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contractId", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *V1RenewContractItemRequest) validateContractItemID(formats strfmt.Registry) error {

	if err := validate.Required("contractItemId", "body", m.ContractItemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 renew contract item request based on context it is used
func (m *V1RenewContractItemRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RenewContractItemRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RenewContractItemRequest) UnmarshalBinary(b []byte) error {
	var res V1RenewContractItemRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
