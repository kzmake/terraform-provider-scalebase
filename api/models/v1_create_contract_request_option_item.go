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

// V1CreateContractRequestOptionItem オプションとなるアイテム
//
// swagger:model v1CreateContractRequestOptionItem
type V1CreateContractRequestOptionItem struct {

	// カタログアイテムID
	// Required: true
	CatalogItemID *string `json:"catalogItemId"`

	// カタログアイテム管理ID(カタログアイテムIDの代わりに指定可)
	CatalogItemOptionalID string `json:"catalogItemOptionalId,omitempty"`

	// 課金項目パラメータ
	ChargeItems []*V1CreateContractRequestChargeItem `json:"chargeItems"`

	// 基本契約期間(省略時はカタログの契約期間から適用)
	DefaultContractTerm *V1CreateContractRequestTerm `json:"defaultContractTerm,omitempty"`

	// 契約終了日(省略時はカタログの契約期間から算出(RFC 3339 format))
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// アイテム名。指定しなければカタログアイテムに従います。
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 create contract request option item
func (m *V1CreateContractRequestOptionItem) Validate(formats strfmt.Registry) error {
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

func (m *V1CreateContractRequestOptionItem) validateCatalogItemID(formats strfmt.Registry) error {

	if err := validate.Required("catalogItemId", "body", m.CatalogItemID); err != nil {
		return err
	}

	return nil
}

func (m *V1CreateContractRequestOptionItem) validateChargeItems(formats strfmt.Registry) error {
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

func (m *V1CreateContractRequestOptionItem) validateDefaultContractTerm(formats strfmt.Registry) error {
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

func (m *V1CreateContractRequestOptionItem) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 create contract request option item based on the context it is used
func (m *V1CreateContractRequestOptionItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *V1CreateContractRequestOptionItem) contextValidateChargeItems(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1CreateContractRequestOptionItem) contextValidateDefaultContractTerm(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1CreateContractRequestOptionItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateContractRequestOptionItem) UnmarshalBinary(b []byte) error {
	var res V1CreateContractRequestOptionItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
