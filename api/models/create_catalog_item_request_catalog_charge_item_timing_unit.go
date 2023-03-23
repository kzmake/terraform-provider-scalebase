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

// CreateCatalogItemRequestCatalogChargeItemTimingUnit 期間単位
//
// - UNIT_MONTHLY: 月
//
// swagger:model CreateCatalogItemRequestCatalogChargeItemTimingUnit
type CreateCatalogItemRequestCatalogChargeItemTimingUnit string

func NewCreateCatalogItemRequestCatalogChargeItemTimingUnit(value CreateCatalogItemRequestCatalogChargeItemTimingUnit) *CreateCatalogItemRequestCatalogChargeItemTimingUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CreateCatalogItemRequestCatalogChargeItemTimingUnit.
func (m CreateCatalogItemRequestCatalogChargeItemTimingUnit) Pointer() *CreateCatalogItemRequestCatalogChargeItemTimingUnit {
	return &m
}

const (

	// CreateCatalogItemRequestCatalogChargeItemTimingUnitUNITMONTHLY captures enum value "UNIT_MONTHLY"
	CreateCatalogItemRequestCatalogChargeItemTimingUnitUNITMONTHLY CreateCatalogItemRequestCatalogChargeItemTimingUnit = "UNIT_MONTHLY"
)

// for schema
var createCatalogItemRequestCatalogChargeItemTimingUnitEnum []interface{}

func init() {
	var res []CreateCatalogItemRequestCatalogChargeItemTimingUnit
	if err := json.Unmarshal([]byte(`["UNIT_MONTHLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCatalogItemRequestCatalogChargeItemTimingUnitEnum = append(createCatalogItemRequestCatalogChargeItemTimingUnitEnum, v)
	}
}

func (m CreateCatalogItemRequestCatalogChargeItemTimingUnit) validateCreateCatalogItemRequestCatalogChargeItemTimingUnitEnum(path, location string, value CreateCatalogItemRequestCatalogChargeItemTimingUnit) error {
	if err := validate.EnumCase(path, location, value, createCatalogItemRequestCatalogChargeItemTimingUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this create catalog item request catalog charge item timing unit
func (m CreateCatalogItemRequestCatalogChargeItemTimingUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCreateCatalogItemRequestCatalogChargeItemTimingUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this create catalog item request catalog charge item timing unit based on context it is used
func (m CreateCatalogItemRequestCatalogChargeItemTimingUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
