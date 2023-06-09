// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListBillingsResponse 請求の一覧取得のレスポンス
//
// swagger:model v1ListBillingsResponse
type V1ListBillingsResponse struct {

	// 請求
	Billings []*V1Billing `json:"billings"`

	// 次の一覧取得に使用するトークン
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Validate validates this v1 list billings response
func (m *V1ListBillingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListBillingsResponse) validateBillings(formats strfmt.Registry) error {
	if swag.IsZero(m.Billings) { // not required
		return nil
	}

	for i := 0; i < len(m.Billings); i++ {
		if swag.IsZero(m.Billings[i]) { // not required
			continue
		}

		if m.Billings[i] != nil {
			if err := m.Billings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 list billings response based on the context it is used
func (m *V1ListBillingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListBillingsResponse) contextValidateBillings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Billings); i++ {

		if m.Billings[i] != nil {

			if swag.IsZero(m.Billings[i]) { // not required
				return nil
			}

			if err := m.Billings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("billings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ListBillingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListBillingsResponse) UnmarshalBinary(b []byte) error {
	var res V1ListBillingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
