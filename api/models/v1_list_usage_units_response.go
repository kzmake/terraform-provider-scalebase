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

// V1ListUsageUnitsResponse 使用量単位一覧の取得レスポンス
//
// swagger:model v1ListUsageUnitsResponse
type V1ListUsageUnitsResponse struct {

	// 次の一覧取得に使用するトークン
	NextPageToken string `json:"nextPageToken,omitempty"`

	// usage units
	UsageUnits []*V1UsageUnit `json:"usageUnits"`
}

// Validate validates this v1 list usage units response
func (m *V1ListUsageUnitsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsageUnits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListUsageUnitsResponse) validateUsageUnits(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageUnits) { // not required
		return nil
	}

	for i := 0; i < len(m.UsageUnits); i++ {
		if swag.IsZero(m.UsageUnits[i]) { // not required
			continue
		}

		if m.UsageUnits[i] != nil {
			if err := m.UsageUnits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usageUnits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usageUnits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 list usage units response based on the context it is used
func (m *V1ListUsageUnitsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsageUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListUsageUnitsResponse) contextValidateUsageUnits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UsageUnits); i++ {

		if m.UsageUnits[i] != nil {
			if err := m.UsageUnits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usageUnits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usageUnits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ListUsageUnitsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListUsageUnitsResponse) UnmarshalBinary(b []byte) error {
	var res V1ListUsageUnitsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
