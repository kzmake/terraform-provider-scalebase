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

// ContractItemTermUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model ContractItemTermUnit
type ContractItemTermUnit string

func NewContractItemTermUnit(value ContractItemTermUnit) *ContractItemTermUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ContractItemTermUnit.
func (m ContractItemTermUnit) Pointer() *ContractItemTermUnit {
	return &m
}

const (

	// ContractItemTermUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	ContractItemTermUnitUNITMONTHLY ContractItemTermUnit = "UNIT_MONTHLY"
)

// for schema
var contractItemTermUnitEnum []interface{}

func init() {
	var res []ContractItemTermUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contractItemTermUnitEnum = append(contractItemTermUnitEnum, v)
	}
}

func (m ContractItemTermUnit) validateContractItemTermUnitEnum(path, location string, value ContractItemTermUnit) error {
	if err := validate.EnumCase(path, location, value, contractItemTermUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this contract item term unit
func (m ContractItemTermUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContractItemTermUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this contract item term unit based on context it is used
func (m ContractItemTermUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
