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

// V1CreateContractRequestChargeItemTimingUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model v1CreateContractRequestChargeItemTimingUnit
type V1CreateContractRequestChargeItemTimingUnit string

func NewV1CreateContractRequestChargeItemTimingUnit(value V1CreateContractRequestChargeItemTimingUnit) *V1CreateContractRequestChargeItemTimingUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1CreateContractRequestChargeItemTimingUnit.
func (m V1CreateContractRequestChargeItemTimingUnit) Pointer() *V1CreateContractRequestChargeItemTimingUnit {
	return &m
}

const (

	// V1CreateContractRequestChargeItemTimingUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	V1CreateContractRequestChargeItemTimingUnitUNITMONTHLY V1CreateContractRequestChargeItemTimingUnit = "UNIT_MONTHLY"
)

// for schema
var v1CreateContractRequestChargeItemTimingUnitEnum []interface{}

func init() {
	var res []V1CreateContractRequestChargeItemTimingUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CreateContractRequestChargeItemTimingUnitEnum = append(v1CreateContractRequestChargeItemTimingUnitEnum, v)
	}
}

func (m V1CreateContractRequestChargeItemTimingUnit) validateV1CreateContractRequestChargeItemTimingUnitEnum(path, location string, value V1CreateContractRequestChargeItemTimingUnit) error {
	if err := validate.EnumCase(path, location, value, v1CreateContractRequestChargeItemTimingUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 create contract request charge item timing unit
func (m V1CreateContractRequestChargeItemTimingUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1CreateContractRequestChargeItemTimingUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 create contract request charge item timing unit based on context it is used
func (m V1CreateContractRequestChargeItemTimingUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}