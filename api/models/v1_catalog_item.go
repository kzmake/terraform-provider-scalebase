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

// V1CatalogItem カタログアイテム
//
// swagger:model v1CatalogItem
type V1CatalogItem struct {

	// カタログ課金項目
	CatalogChargeItems []*Publicv1CatalogChargeItem `json:"catalogChargeItems"`

	// カタログアイテムID
	ID string `json:"id,omitempty"`

	// カタログアイテム名
	Name string `json:"name,omitempty"`

	// カタログアイテム管理ID
	OptionalID string `json:"optionalId,omitempty"`

	// プロダクトID
	ProductID string `json:"productId,omitempty"`

	// ステータス
	Status V1CatalogItemStatus `json:"status,omitempty"`
}

// Validate validates this v1 catalog item
func (m *V1CatalogItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogChargeItems(formats); err != nil {
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

func (m *V1CatalogItem) validateCatalogChargeItems(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogChargeItems) { // not required
		return nil
	}

	for i := 0; i < len(m.CatalogChargeItems); i++ {
		if swag.IsZero(m.CatalogChargeItems[i]) { // not required
			continue
		}

		if m.CatalogChargeItems[i] != nil {
			if err := m.CatalogChargeItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalogChargeItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("catalogChargeItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CatalogItem) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 catalog item based on the context it is used
func (m *V1CatalogItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCatalogChargeItems(ctx, formats); err != nil {
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

func (m *V1CatalogItem) contextValidateCatalogChargeItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CatalogChargeItems); i++ {

		if m.CatalogChargeItems[i] != nil {

			if swag.IsZero(m.CatalogChargeItems[i]) { // not required
				return nil
			}

			if err := m.CatalogChargeItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalogChargeItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("catalogChargeItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CatalogItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

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
func (m *V1CatalogItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CatalogItem) UnmarshalBinary(b []byte) error {
	var res V1CatalogItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
