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

// V1GetCatalogItemResponse カタログアイテム取得のレスポンス
//
// swagger:model v1GetCatalogItemResponse
type V1GetCatalogItemResponse struct {

	// catalog item
	CatalogItem *V1CatalogItem `json:"catalogItem,omitempty"`
}

// Validate validates this v1 get catalog item response
func (m *V1GetCatalogItemResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetCatalogItemResponse) validateCatalogItem(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogItem) { // not required
		return nil
	}

	if m.CatalogItem != nil {
		if err := m.CatalogItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("catalogItem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 get catalog item response based on the context it is used
func (m *V1GetCatalogItemResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCatalogItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetCatalogItemResponse) contextValidateCatalogItem(ctx context.Context, formats strfmt.Registry) error {

	if m.CatalogItem != nil {
		if err := m.CatalogItem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("catalogItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("catalogItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetCatalogItemResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetCatalogItemResponse) UnmarshalBinary(b []byte) error {
	var res V1GetCatalogItemResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}