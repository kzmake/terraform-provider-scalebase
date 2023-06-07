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

// V1SRNType リソース種別
//
// swagger:model v1SRNType
type V1SRNType string

func NewV1SRNType(value V1SRNType) *V1SRNType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1SRNType.
func (m V1SRNType) Pointer() *V1SRNType {
	return &m
}

const (

	// V1SRNTypeTYPECUSTOMER captures enum value "TYPE_CUSTOMER"
	V1SRNTypeTYPECUSTOMER V1SRNType = "TYPE_CUSTOMER"

	// V1SRNTypeTYPECONTRACT captures enum value "TYPE_CONTRACT"
	V1SRNTypeTYPECONTRACT V1SRNType = "TYPE_CONTRACT"
)

// for schema
var v1SRNTypeEnum []interface{}

func init() {
	var res []V1SRNType
	if err := json.Unmarshal([]byte(`["TYPE_CUSTOMER","TYPE_CONTRACT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SRNTypeEnum = append(v1SRNTypeEnum, v)
	}
}

func (m V1SRNType) validateV1SRNTypeEnum(path, location string, value V1SRNType) error {
	if err := validate.EnumCase(path, location, value, v1SRNTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 s r n type
func (m V1SRNType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SRNTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 s r n type based on context it is used
func (m V1SRNType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
