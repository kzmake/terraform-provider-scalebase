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

// ChangeMainItemRequestItemChargeItemTimingUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model ChangeMainItemRequestItemChargeItemTimingUnit
type ChangeMainItemRequestItemChargeItemTimingUnit string

func NewChangeMainItemRequestItemChargeItemTimingUnit(value ChangeMainItemRequestItemChargeItemTimingUnit) *ChangeMainItemRequestItemChargeItemTimingUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ChangeMainItemRequestItemChargeItemTimingUnit.
func (m ChangeMainItemRequestItemChargeItemTimingUnit) Pointer() *ChangeMainItemRequestItemChargeItemTimingUnit {
	return &m
}

const (

	// ChangeMainItemRequestItemChargeItemTimingUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	ChangeMainItemRequestItemChargeItemTimingUnitUNITMONTHLY ChangeMainItemRequestItemChargeItemTimingUnit = "UNIT_MONTHLY"
)

// for schema
var changeMainItemRequestItemChargeItemTimingUnitEnum []interface{}

func init() {
	var res []ChangeMainItemRequestItemChargeItemTimingUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeMainItemRequestItemChargeItemTimingUnitEnum = append(changeMainItemRequestItemChargeItemTimingUnitEnum, v)
	}
}

func (m ChangeMainItemRequestItemChargeItemTimingUnit) validateChangeMainItemRequestItemChargeItemTimingUnitEnum(path, location string, value ChangeMainItemRequestItemChargeItemTimingUnit) error {
	if err := validate.EnumCase(path, location, value, changeMainItemRequestItemChargeItemTimingUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this change main item request item charge item timing unit
func (m ChangeMainItemRequestItemChargeItemTimingUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateChangeMainItemRequestItemChargeItemTimingUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this change main item request item charge item timing unit based on context it is used
func (m ChangeMainItemRequestItemChargeItemTimingUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
