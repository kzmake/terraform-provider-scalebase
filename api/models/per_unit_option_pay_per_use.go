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

// PerUnitOptionPayPerUse per unit option pay per use
//
// swagger:model PerUnitOptionPayPerUse
type PerUnitOptionPayPerUse string

func NewPerUnitOptionPayPerUse(value PerUnitOptionPayPerUse) *PerUnitOptionPayPerUse {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PerUnitOptionPayPerUse.
func (m PerUnitOptionPayPerUse) Pointer() *PerUnitOptionPayPerUse {
	return &m
}

const (

	// PerUnitOptionPayPerUsePAYPERUSEQUANTITY captures enum value "PAY_PER_USE_QUANTITY"
	PerUnitOptionPayPerUsePAYPERUSEQUANTITY PerUnitOptionPayPerUse = "PAY_PER_USE_QUANTITY"
)

// for schema
var perUnitOptionPayPerUseEnum []interface{}

func init() {
	var res []PerUnitOptionPayPerUse
	if err := json.Unmarshal([]byte(`["PAY_PER_USE_QUANTITY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		perUnitOptionPayPerUseEnum = append(perUnitOptionPayPerUseEnum, v)
	}
}

func (m PerUnitOptionPayPerUse) validatePerUnitOptionPayPerUseEnum(path, location string, value PerUnitOptionPayPerUse) error {
	if err := validate.EnumCase(path, location, value, perUnitOptionPayPerUseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this per unit option pay per use
func (m PerUnitOptionPayPerUse) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePerUnitOptionPayPerUseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this per unit option pay per use based on context it is used
func (m PerUnitOptionPayPerUse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
