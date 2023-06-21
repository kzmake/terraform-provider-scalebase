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

// V1BillingOrderField 対象
//
// - FIELD_ID: 請求ID
//   - FIELD_BILLING_DATE: 請求日
//
// swagger:model v1BillingOrderField
type V1BillingOrderField string

func NewV1BillingOrderField(value V1BillingOrderField) *V1BillingOrderField {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1BillingOrderField.
func (m V1BillingOrderField) Pointer() *V1BillingOrderField {
	return &m
}

const (

	// V1BillingOrderFieldFIELDID captures enum value "FIELD_ID"
	V1BillingOrderFieldFIELDID V1BillingOrderField = "FIELD_ID"

	// V1BillingOrderFieldFIELDBILLINGDATE captures enum value "FIELD_BILLING_DATE"
	V1BillingOrderFieldFIELDBILLINGDATE V1BillingOrderField = "FIELD_BILLING_DATE"
)

// for schema
var v1BillingOrderFieldEnum []interface{}

func init() {
	var res []V1BillingOrderField
	if err := json.Unmarshal([]byte(`["FIELD_ID","FIELD_BILLING_DATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1BillingOrderFieldEnum = append(v1BillingOrderFieldEnum, v)
	}
}

func (m V1BillingOrderField) validateV1BillingOrderFieldEnum(path, location string, value V1BillingOrderField) error {
	if err := validate.EnumCase(path, location, value, v1BillingOrderFieldEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 billing order field
func (m V1BillingOrderField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1BillingOrderFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 billing order field based on context it is used
func (m V1BillingOrderField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}