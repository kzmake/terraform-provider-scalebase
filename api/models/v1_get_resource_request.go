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

// V1GetResourceRequest 対象リソースのカスタムフィールド取得のリクエスト
//
// swagger:model v1GetResourceRequest
type V1GetResourceRequest struct {

	// リソースの識別子
	Srn *V1SRN `json:"srn,omitempty"`
}

// Validate validates this v1 get resource request
func (m *V1GetResourceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetResourceRequest) validateSrn(formats strfmt.Registry) error {
	if swag.IsZero(m.Srn) { // not required
		return nil
	}

	if m.Srn != nil {
		if err := m.Srn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("srn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("srn")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 get resource request based on the context it is used
func (m *V1GetResourceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSrn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetResourceRequest) contextValidateSrn(ctx context.Context, formats strfmt.Registry) error {

	if m.Srn != nil {

		if swag.IsZero(m.Srn) { // not required
			return nil
		}

		if err := m.Srn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("srn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("srn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetResourceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetResourceRequest) UnmarshalBinary(b []byte) error {
	var res V1GetResourceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
