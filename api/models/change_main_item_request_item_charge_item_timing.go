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

// ChangeMainItemRequestItemChargeItemTiming タイミング
//
// swagger:model ChangeMainItemRequestItemChargeItemTiming
type ChangeMainItemRequestItemChargeItemTiming struct {

	// UNIT_MONTHLY指定時のオプション
	MonthlyOption ChangeMainItemRequestItemChargeItemTimingMonthlyOption `json:"monthlyOption,omitempty"`

	// 単位
	Unit ChangeMainItemRequestItemChargeItemTimingUnit `json:"unit,omitempty"`

	// 数(負数の場合は前を表す。例: { "value": -1, "unit": "UNIT_MONTHLY" } => 先月)
	// Required: true
	Value *int32 `json:"value"`
}

// Validate validates this change main item request item charge item timing
func (m *ChangeMainItemRequestItemChargeItemTiming) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonthlyOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeMainItemRequestItemChargeItemTiming) validateMonthlyOption(formats strfmt.Registry) error {
	if swag.IsZero(m.MonthlyOption) { // not required
		return nil
	}

	if err := m.MonthlyOption.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("monthlyOption")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("monthlyOption")
		}
		return err
	}

	return nil
}

func (m *ChangeMainItemRequestItemChargeItemTiming) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if err := m.Unit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("unit")
		}
		return err
	}

	return nil
}

func (m *ChangeMainItemRequestItemChargeItemTiming) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this change main item request item charge item timing based on the context it is used
func (m *ChangeMainItemRequestItemChargeItemTiming) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonthlyOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeMainItemRequestItemChargeItemTiming) contextValidateMonthlyOption(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MonthlyOption.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("monthlyOption")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("monthlyOption")
		}
		return err
	}

	return nil
}

func (m *ChangeMainItemRequestItemChargeItemTiming) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Unit.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("unit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeMainItemRequestItemChargeItemTiming) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeMainItemRequestItemChargeItemTiming) UnmarshalBinary(b []byte) error {
	var res ChangeMainItemRequestItemChargeItemTiming
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}