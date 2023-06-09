// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListBillingsRequest 請求の一覧取得のリクエスト
//
// swagger:model v1ListBillingsRequest
type V1ListBillingsRequest struct {

	// 一覧取得する最大数
	PageSize int32 `json:"pageSize,omitempty"`

	// 一覧取得に使用するトークン
	PageToken string `json:"pageToken,omitempty"`
}

// Validate validates this v1 list billings request
func (m *V1ListBillingsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 list billings request based on context it is used
func (m *V1ListBillingsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListBillingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListBillingsRequest) UnmarshalBinary(b []byte) error {
	var res V1ListBillingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
