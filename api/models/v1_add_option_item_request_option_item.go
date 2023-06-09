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

// V1AddOptionItemRequestOptionItem アイテムカスタマイズ情報
//
// swagger:model v1AddOptionItemRequestOptionItem
type V1AddOptionItemRequestOptionItem struct {

	// カタログアイテムID
	// Required: true
	CatalogItemID *string `json:"catalogItemId"`

	// 課金項目
	ChargeItems []*V1AddOptionItemRequestChargeItem `json:"chargeItems"`

	// 基本契約期間(省略時はカタログの契約期間から適用)
	DefaultContractTerm *V1AddOptionItemRequestTerm `json:"defaultContractTerm,omitempty"`

	// 契約終了日(省略時はカタログの契約期間から算出(RFC 3339 format))
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// アイテム名。指定しなければカタログアイテムに従います。
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 add option item request option item
func (m *V1AddOptionItemRequestOptionItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultContractTerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AddOptionItemRequestOptionItem) validateCatalogItemID(formats strfmt.Registry) error {

	if err := validate.Required("catalogItemId", "body", m.CatalogItemID); err != nil {
		return err
	}

	return nil
}

func (m *V1AddOptionItemRequestOptionItem) validateChargeItems(formats strfmt.Registry) error {
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

func (m *V1AddOptionItemRequestOptionItem) validateDefaultContractTerm(formats strfmt.Registry) error {
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

func (m *V1AddOptionItemRequestOptionItem) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 add option item request option item based on the context it is used
func (m *V1AddOptionItemRequestOptionItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChargeItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultContractTerm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AddOptionItemRequestOptionItem) contextValidateChargeItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChargeItems); i++ {

		if m.ChargeItems[i] != nil {

			if swag.IsZero(m.ChargeItems[i]) { // not required
				return nil
			}

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

func (m *V1AddOptionItemRequestOptionItem) contextValidateDefaultContractTerm(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultContractTerm != nil {

		if swag.IsZero(m.DefaultContractTerm) { // not required
			return nil
		}

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

// MarshalBinary interface implementation
func (m *V1AddOptionItemRequestOptionItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AddOptionItemRequestOptionItem) UnmarshalBinary(b []byte) error {
	var res V1AddOptionItemRequestOptionItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
