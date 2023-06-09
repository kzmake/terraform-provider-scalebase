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

// V1ChangeQuantityRequest 契約数変更リクエスト
//
// swagger:model v1ChangeQuantityRequest
type V1ChangeQuantityRequest struct {

	// 契約数変更時の日割設定
	ChangeQuantityDailyRate *V1ChangeQuantityRequestDailyRate `json:"changeQuantityDailyRate,omitempty"`

	// 契約課金項目ID
	ContractChargeItemID string `json:"contractChargeItemId,omitempty"`

	// 契約課金項目管理ID(契約課金項目IDの代わりに指定可)
	ContractChargeItemOptionalID string `json:"contractChargeItemOptionalId,omitempty"`

	// 契約ID
	ContractID string `json:"contractId,omitempty"`

	// 契約管理ID(契約IDの代わりに指定可)
	ContractOptionalID string `json:"contractOptionalId,omitempty"`

	// effective start date
	// Format: date-time
	EffectiveStartDate strfmt.DateTime `json:"effectiveStartDate,omitempty"`

	// 改定メモ
	Memo string `json:"memo,omitempty"`

	// 変更後の契約数
	Quantity int32 `json:"quantity,omitempty"`
}

// Validate validates this v1 change quantity request
func (m *V1ChangeQuantityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeQuantityDailyRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ChangeQuantityRequest) validateChangeQuantityDailyRate(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangeQuantityDailyRate) { // not required
		return nil
	}

	if m.ChangeQuantityDailyRate != nil {
		if err := m.ChangeQuantityDailyRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQuantityDailyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeQuantityDailyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1ChangeQuantityRequest) validateEffectiveStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveStartDate", "body", "date-time", m.EffectiveStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 change quantity request based on the context it is used
func (m *V1ChangeQuantityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChangeQuantityDailyRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ChangeQuantityRequest) contextValidateChangeQuantityDailyRate(ctx context.Context, formats strfmt.Registry) error {

	if m.ChangeQuantityDailyRate != nil {

		if swag.IsZero(m.ChangeQuantityDailyRate) { // not required
			return nil
		}

		if err := m.ChangeQuantityDailyRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQuantityDailyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changeQuantityDailyRate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ChangeQuantityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ChangeQuantityRequest) UnmarshalBinary(b []byte) error {
	var res V1ChangeQuantityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
