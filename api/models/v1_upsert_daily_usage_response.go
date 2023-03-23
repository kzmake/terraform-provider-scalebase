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

// V1UpsertDailyUsageResponse 使用量を作成or更新のレスポンス
//
// swagger:model v1UpsertDailyUsageResponse
type V1UpsertDailyUsageResponse struct {

	// usage
	Usage *V1DailyUsage `json:"usage,omitempty"`
}

// Validate validates this v1 upsert daily usage response
func (m *V1UpsertDailyUsageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpsertDailyUsageResponse) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 upsert daily usage response based on the context it is used
func (m *V1UpsertDailyUsageResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpsertDailyUsageResponse) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {
		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UpsertDailyUsageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpsertDailyUsageResponse) UnmarshalBinary(b []byte) error {
	var res V1UpsertDailyUsageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
