// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ExpectedBilling 請求予定
//
// swagger:model v1ExpectedBilling
type V1ExpectedBilling struct {

	// 請求日(RFC 3339 format)
	// Format: date-time
	BillingDate strfmt.DateTime `json:"billingDate,omitempty"`

	// 期限日(RFC 3339 format)
	// Format: date-time
	DueDate strfmt.DateTime `json:"dueDate,omitempty"`

	// 請求項目
	Items []*V1ExpectedBillingItem `json:"items"`

	// 合計(税込)。請求項目の料金の合計
	Total float64 `json:"total,omitempty"`

	// 税額
	TotalTaxAmount float64 `json:"totalTaxAmount,omitempty"`
}

// Validate validates this v1 expected billing
func (m *V1ExpectedBilling) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ExpectedBilling) validateBillingDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("billingDate", "body", "date-time", m.BillingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ExpectedBilling) validateDueDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dueDate", "body", "date-time", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ExpectedBilling) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 expected billing based on the context it is used
func (m *V1ExpectedBilling) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ExpectedBilling) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {

			if swag.IsZero(m.Items[i]) { // not required
				return nil
			}

			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ExpectedBilling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ExpectedBilling) UnmarshalBinary(b []byte) error {
	var res V1ExpectedBilling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
