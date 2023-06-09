// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GetCustomerRequest 顧客取得のリクエスト
//
// swagger:model v1GetCustomerRequest
type V1GetCustomerRequest struct {

	// 顧客ID
	ID string `json:"id,omitempty"`

	// 顧客管理ID(顧客IDの代わりに指定可)
	OptionalID string `json:"optionalId,omitempty"`
}

// Validate validates this v1 get customer request
func (m *V1GetCustomerRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 get customer request based on context it is used
func (m *V1GetCustomerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GetCustomerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetCustomerRequest) UnmarshalBinary(b []byte) error {
	var res V1GetCustomerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
