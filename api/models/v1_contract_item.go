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
	"github.com/go-openapi/validate"
)

// V1ContractItem 契約アイテム
//
// swagger:model v1ContractItem
type V1ContractItem struct {

	// 契約課金項目
	ChargeItems []*V1ContractChargeItem `json:"chargeItems"`

	// 基本契約期間
	DefaultContractTerm *ContractItemTerm `json:"defaultContractTerm,omitempty"`

	// 契約終了日(RFC 3339 format)
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// 契約アイテムID
	ID string `json:"id,omitempty"`

	// アイテム名
	Name string `json:"name,omitempty"`

	// オプション契約アイテム（自身がオプションアイテムの場合は空）
	OptionItems []*V1ContractItem `json:"optionItems"`

	// 契約開始日(RFC 3339 format)
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// 契約アイテムステータス
	Status ContractItemStatus `json:"status,omitempty"`
}

// Validate validates this v1 contract item
func (m *V1ContractItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargeItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultContractTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContractItem) validateChargeItems(formats strfmt.Registry) error {
	if swag.IsZero(m.ChargeItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ChargeItems); i++ {
		if swag.IsZero(m.ChargeItems[i]) { // not required
			continue
		}

		if m.ChargeItems[i] != nil {
			if err := m.ChargeItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chargeItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chargeItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ContractItem) validateDefaultContractTerm(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultContractTerm) { // not required
		return nil
	}

	if m.DefaultContractTerm != nil {
		if err := m.DefaultContractTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultContractTerm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultContractTerm")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContractItem) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ContractItem) validateOptionItems(formats strfmt.Registry) error {
	if swag.IsZero(m.OptionItems) { // not required
		return nil
	}

	for i := 0; i < len(m.OptionItems); i++ {
		if swag.IsZero(m.OptionItems[i]) { // not required
			continue
		}

		if m.OptionItems[i] != nil {
			if err := m.OptionItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("optionItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("optionItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ContractItem) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ContractItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 contract item based on the context it is used
func (m *V1ContractItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChargeItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultContractTerm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptionItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContractItem) contextValidateChargeItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChargeItems); i++ {

		if m.ChargeItems[i] != nil {
			if err := m.ChargeItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chargeItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chargeItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ContractItem) contextValidateDefaultContractTerm(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultContractTerm != nil {
		if err := m.DefaultContractTerm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultContractTerm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultContractTerm")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContractItem) contextValidateOptionItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OptionItems); i++ {

		if m.OptionItems[i] != nil {
			if err := m.OptionItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("optionItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("optionItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ContractItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ContractItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContractItem) UnmarshalBinary(b []byte) error {
	var res V1ContractItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}