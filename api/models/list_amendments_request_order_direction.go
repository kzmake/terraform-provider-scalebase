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

// ListAmendmentsRequestOrderDirection 順序
//
// - DIRECTION_ASCENDING: 昇順
//   - DIRECTION_DESCENDING: 降順
//
// swagger:model ListAmendmentsRequestOrderDirection
type ListAmendmentsRequestOrderDirection string

func NewListAmendmentsRequestOrderDirection(value ListAmendmentsRequestOrderDirection) *ListAmendmentsRequestOrderDirection {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ListAmendmentsRequestOrderDirection.
func (m ListAmendmentsRequestOrderDirection) Pointer() *ListAmendmentsRequestOrderDirection {
	return &m
}

const (

	// ListAmendmentsRequestOrderDirectionDIRECTIONASCENDING captures enum value "DIRECTION_ASCENDING"
	ListAmendmentsRequestOrderDirectionDIRECTIONASCENDING ListAmendmentsRequestOrderDirection = "DIRECTION_ASCENDING"

	// ListAmendmentsRequestOrderDirectionDIRECTIONDESCENDING captures enum value "DIRECTION_DESCENDING"
	ListAmendmentsRequestOrderDirectionDIRECTIONDESCENDING ListAmendmentsRequestOrderDirection = "DIRECTION_DESCENDING"
)

// for schema
var listAmendmentsRequestOrderDirectionEnum []interface{}

func init() {
	var res []ListAmendmentsRequestOrderDirection
	if err := json.Unmarshal([]byte(`["DIRECTION_ASCENDING","DIRECTION_DESCENDING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listAmendmentsRequestOrderDirectionEnum = append(listAmendmentsRequestOrderDirectionEnum, v)
	}
}

func (m ListAmendmentsRequestOrderDirection) validateListAmendmentsRequestOrderDirectionEnum(path, location string, value ListAmendmentsRequestOrderDirection) error {
	if err := validate.EnumCase(path, location, value, listAmendmentsRequestOrderDirectionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this list amendments request order direction
func (m ListAmendmentsRequestOrderDirection) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateListAmendmentsRequestOrderDirectionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this list amendments request order direction based on context it is used
func (m ListAmendmentsRequestOrderDirection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}