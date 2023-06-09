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

// V1CatalogItemStatus v1 catalog item status
//
// swagger:model v1CatalogItemStatus
type V1CatalogItemStatus string

func NewV1CatalogItemStatus(value V1CatalogItemStatus) *V1CatalogItemStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1CatalogItemStatus.
func (m V1CatalogItemStatus) Pointer() *V1CatalogItemStatus {
	return &m
}

const (

	// V1CatalogItemStatusSTATUSONSALE captures enum value "STATUS_ON_SALE"
	V1CatalogItemStatusSTATUSONSALE V1CatalogItemStatus = "STATUS_ON_SALE"

	// V1CatalogItemStatusSTATUSDISCONTINUED captures enum value "STATUS_DISCONTINUED"
	V1CatalogItemStatusSTATUSDISCONTINUED V1CatalogItemStatus = "STATUS_DISCONTINUED"
)

// for schema
var v1CatalogItemStatusEnum []interface{}

func init() {
	var res []V1CatalogItemStatus
	if err := json.Unmarshal([]byte(`["STATUS_ON_SALE","STATUS_DISCONTINUED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CatalogItemStatusEnum = append(v1CatalogItemStatusEnum, v)
	}
}

func (m V1CatalogItemStatus) validateV1CatalogItemStatusEnum(path, location string, value V1CatalogItemStatus) error {
	if err := validate.EnumCase(path, location, value, v1CatalogItemStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 catalog item status
func (m V1CatalogItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1CatalogItemStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 catalog item status based on context it is used
func (m V1CatalogItemStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
