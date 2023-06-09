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

// CreateCatalogItemRequestCatalogChargeItemDailyRateType create catalog item request catalog charge item daily rate type
//
// swagger:model CreateCatalogItemRequestCatalogChargeItemDailyRateType
type CreateCatalogItemRequestCatalogChargeItemDailyRateType string

func NewCreateCatalogItemRequestCatalogChargeItemDailyRateType(value CreateCatalogItemRequestCatalogChargeItemDailyRateType) *CreateCatalogItemRequestCatalogChargeItemDailyRateType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CreateCatalogItemRequestCatalogChargeItemDailyRateType.
func (m CreateCatalogItemRequestCatalogChargeItemDailyRateType) Pointer() *CreateCatalogItemRequestCatalogChargeItemDailyRateType {
	return &m
}

const (

	// CreateCatalogItemRequestCatalogChargeItemDailyRateTypeTYPEFULLPRICE captures enum value "TYPE_FULL_PRICE"
	CreateCatalogItemRequestCatalogChargeItemDailyRateTypeTYPEFULLPRICE CreateCatalogItemRequestCatalogChargeItemDailyRateType = "TYPE_FULL_PRICE"

	// CreateCatalogItemRequestCatalogChargeItemDailyRateTypeTYPEDAILYRATE captures enum value "TYPE_DAILY_RATE"
	CreateCatalogItemRequestCatalogChargeItemDailyRateTypeTYPEDAILYRATE CreateCatalogItemRequestCatalogChargeItemDailyRateType = "TYPE_DAILY_RATE"

	// CreateCatalogItemRequestCatalogChargeItemDailyRateTypeTYPEFIXEDRATE captures enum value "TYPE_FIXED_RATE"
	CreateCatalogItemRequestCatalogChargeItemDailyRateTypeTYPEFIXEDRATE CreateCatalogItemRequestCatalogChargeItemDailyRateType = "TYPE_FIXED_RATE"
)

// for schema
var createCatalogItemRequestCatalogChargeItemDailyRateTypeEnum []interface{}

func init() {
	var res []CreateCatalogItemRequestCatalogChargeItemDailyRateType
	if err := json.Unmarshal([]byte(`["TYPE_FULL_PRICE","TYPE_DAILY_RATE","TYPE_FIXED_RATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCatalogItemRequestCatalogChargeItemDailyRateTypeEnum = append(createCatalogItemRequestCatalogChargeItemDailyRateTypeEnum, v)
	}
}

func (m CreateCatalogItemRequestCatalogChargeItemDailyRateType) validateCreateCatalogItemRequestCatalogChargeItemDailyRateTypeEnum(path, location string, value CreateCatalogItemRequestCatalogChargeItemDailyRateType) error {
	if err := validate.EnumCase(path, location, value, createCatalogItemRequestCatalogChargeItemDailyRateTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this create catalog item request catalog charge item daily rate type
func (m CreateCatalogItemRequestCatalogChargeItemDailyRateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCreateCatalogItemRequestCatalogChargeItemDailyRateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this create catalog item request catalog charge item daily rate type based on context it is used
func (m CreateCatalogItemRequestCatalogChargeItemDailyRateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
