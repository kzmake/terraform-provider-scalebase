// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CreateCatalogItemRequestCatalogChargeItemType create catalog item request catalog charge item type
//
// swagger:model CreateCatalogItemRequestCatalogChargeItemType
type CreateCatalogItemRequestCatalogChargeItemType string

func NewCreateCatalogItemRequestCatalogChargeItemType(value CreateCatalogItemRequestCatalogChargeItemType) *CreateCatalogItemRequestCatalogChargeItemType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CreateCatalogItemRequestCatalogChargeItemType.
func (m CreateCatalogItemRequestCatalogChargeItemType) Pointer() *CreateCatalogItemRequestCatalogChargeItemType {
	return &m
}

const (

	// CreateCatalogItemRequestCatalogChargeItemTypeTYPERECURRING captures enum value "TYPE_RECURRING"
	CreateCatalogItemRequestCatalogChargeItemTypeTYPERECURRING CreateCatalogItemRequestCatalogChargeItemType = "TYPE_RECURRING"

	// CreateCatalogItemRequestCatalogChargeItemTypeTYPEONETIME captures enum value "TYPE_ONE_TIME"
	CreateCatalogItemRequestCatalogChargeItemTypeTYPEONETIME CreateCatalogItemRequestCatalogChargeItemType = "TYPE_ONE_TIME"
)

// for schema
var createCatalogItemRequestCatalogChargeItemTypeEnum []interface{}

func init() {
	var res []CreateCatalogItemRequestCatalogChargeItemType
	if err := json.Unmarshal([]byte(`["TYPE_RECURRING","TYPE_ONE_TIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCatalogItemRequestCatalogChargeItemTypeEnum = append(createCatalogItemRequestCatalogChargeItemTypeEnum, v)
	}
}

func (m CreateCatalogItemRequestCatalogChargeItemType) validateCreateCatalogItemRequestCatalogChargeItemTypeEnum(path, location string, value CreateCatalogItemRequestCatalogChargeItemType) error {
	if err := validate.EnumCase(path, location, value, createCatalogItemRequestCatalogChargeItemTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this create catalog item request catalog charge item type
func (m CreateCatalogItemRequestCatalogChargeItemType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCreateCatalogItemRequestCatalogChargeItemTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this create catalog item request catalog charge item type based on context it is used
func (m CreateCatalogItemRequestCatalogChargeItemType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
