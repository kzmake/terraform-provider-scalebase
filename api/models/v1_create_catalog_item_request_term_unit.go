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

// V1CreateCatalogItemRequestTermUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model v1CreateCatalogItemRequestTermUnit
type V1CreateCatalogItemRequestTermUnit string

func NewV1CreateCatalogItemRequestTermUnit(value V1CreateCatalogItemRequestTermUnit) *V1CreateCatalogItemRequestTermUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1CreateCatalogItemRequestTermUnit.
func (m V1CreateCatalogItemRequestTermUnit) Pointer() *V1CreateCatalogItemRequestTermUnit {
	return &m
}

const (

	// V1CreateCatalogItemRequestTermUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	V1CreateCatalogItemRequestTermUnitUNITMONTHLY V1CreateCatalogItemRequestTermUnit = "UNIT_MONTHLY"
)

// for schema
var v1CreateCatalogItemRequestTermUnitEnum []interface{}

func init() {
	var res []V1CreateCatalogItemRequestTermUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CreateCatalogItemRequestTermUnitEnum = append(v1CreateCatalogItemRequestTermUnitEnum, v)
	}
}

func (m V1CreateCatalogItemRequestTermUnit) validateV1CreateCatalogItemRequestTermUnitEnum(path, location string, value V1CreateCatalogItemRequestTermUnit) error {
	if err := validate.EnumCase(path, location, value, v1CreateCatalogItemRequestTermUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 create catalog item request term unit
func (m V1CreateCatalogItemRequestTermUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1CreateCatalogItemRequestTermUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 create catalog item request term unit based on context it is used
func (m V1CreateCatalogItemRequestTermUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
