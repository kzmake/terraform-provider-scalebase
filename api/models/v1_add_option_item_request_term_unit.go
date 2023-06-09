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

// V1AddOptionItemRequestTermUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model v1AddOptionItemRequestTermUnit
type V1AddOptionItemRequestTermUnit string

func NewV1AddOptionItemRequestTermUnit(value V1AddOptionItemRequestTermUnit) *V1AddOptionItemRequestTermUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1AddOptionItemRequestTermUnit.
func (m V1AddOptionItemRequestTermUnit) Pointer() *V1AddOptionItemRequestTermUnit {
	return &m
}

const (

	// V1AddOptionItemRequestTermUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	V1AddOptionItemRequestTermUnitUNITMONTHLY V1AddOptionItemRequestTermUnit = "UNIT_MONTHLY"
)

// for schema
var v1AddOptionItemRequestTermUnitEnum []interface{}

func init() {
	var res []V1AddOptionItemRequestTermUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AddOptionItemRequestTermUnitEnum = append(v1AddOptionItemRequestTermUnitEnum, v)
	}
}

func (m V1AddOptionItemRequestTermUnit) validateV1AddOptionItemRequestTermUnitEnum(path, location string, value V1AddOptionItemRequestTermUnit) error {
	if err := validate.EnumCase(path, location, value, v1AddOptionItemRequestTermUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 add option item request term unit
func (m V1AddOptionItemRequestTermUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1AddOptionItemRequestTermUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 add option item request term unit based on context it is used
func (m V1AddOptionItemRequestTermUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
