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

// V1ContractItemTermUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model v1ContractItemTermUnit
type V1ContractItemTermUnit string

func NewV1ContractItemTermUnit(value V1ContractItemTermUnit) *V1ContractItemTermUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ContractItemTermUnit.
func (m V1ContractItemTermUnit) Pointer() *V1ContractItemTermUnit {
	return &m
}

const (

	// V1ContractItemTermUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	V1ContractItemTermUnitUNITMONTHLY V1ContractItemTermUnit = "UNIT_MONTHLY"
)

// for schema
var v1ContractItemTermUnitEnum []interface{}

func init() {
	var res []V1ContractItemTermUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ContractItemTermUnitEnum = append(v1ContractItemTermUnitEnum, v)
	}
}

func (m V1ContractItemTermUnit) validateV1ContractItemTermUnitEnum(path, location string, value V1ContractItemTermUnit) error {
	if err := validate.EnumCase(path, location, value, v1ContractItemTermUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 contract item term unit
func (m V1ContractItemTermUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ContractItemTermUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 contract item term unit based on context it is used
func (m V1ContractItemTermUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}