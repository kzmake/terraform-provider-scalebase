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

// V1ListCustomersRequest 顧客の一覧取得のリクエスト
//
// swagger:model v1ListCustomersRequest
type V1ListCustomersRequest struct {

	// 一覧取得する最大数
	// Required: true
	PageSize *int32 `json:"pageSize"`

	// 一覧取得に使用するトークン
	PageToken string `json:"pageToken,omitempty"`
}

// Validate validates this v1 list customers request
func (m *V1ListCustomersRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListCustomersRequest) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 list customers request based on context it is used
func (m *V1ListCustomersRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListCustomersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListCustomersRequest) UnmarshalBinary(b []byte) error {
	var res V1ListCustomersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
