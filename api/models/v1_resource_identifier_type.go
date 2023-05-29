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

// V1ResourceIdentifierType リソース種別
//
// swagger:model v1ResourceIdentifierType
type V1ResourceIdentifierType string

func NewV1ResourceIdentifierType(value V1ResourceIdentifierType) *V1ResourceIdentifierType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1ResourceIdentifierType.
func (m V1ResourceIdentifierType) Pointer() *V1ResourceIdentifierType {
	return &m
}

const (

	// V1ResourceIdentifierTypeTYPECUSTOMER captures enum value "TYPE_CUSTOMER"
	V1ResourceIdentifierTypeTYPECUSTOMER V1ResourceIdentifierType = "TYPE_CUSTOMER"

	// V1ResourceIdentifierTypeTYPECONTRACT captures enum value "TYPE_CONTRACT"
	V1ResourceIdentifierTypeTYPECONTRACT V1ResourceIdentifierType = "TYPE_CONTRACT"
)

// for schema
var v1ResourceIdentifierTypeEnum []interface{}

func init() {
	var res []V1ResourceIdentifierType
	if err := json.Unmarshal([]byte(`["TYPE_CUSTOMER","TYPE_CONTRACT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ResourceIdentifierTypeEnum = append(v1ResourceIdentifierTypeEnum, v)
	}
}

func (m V1ResourceIdentifierType) validateV1ResourceIdentifierTypeEnum(path, location string, value V1ResourceIdentifierType) error {
	if err := validate.EnumCase(path, location, value, v1ResourceIdentifierTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 resource identifier type
func (m V1ResourceIdentifierType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ResourceIdentifierTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 resource identifier type based on context it is used
func (m V1ResourceIdentifierType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
