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
)

// V1ListAmendmentsByContractResponse 契約に紐づく改定の一覧取得レスポンス
//
// swagger:model v1ListAmendmentsByContractResponse
type V1ListAmendmentsByContractResponse struct {

	// 改定
	Amendments []*V1Amendment `json:"amendments"`

	// 次の一覧取得に使用するトークン
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Validate validates this v1 list amendments by contract response
func (m *V1ListAmendmentsByContractResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmendments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListAmendmentsByContractResponse) validateAmendments(formats strfmt.Registry) error {
	if swag.IsZero(m.Amendments) { // not required
		return nil
	}

	for i := 0; i < len(m.Amendments); i++ {
		if swag.IsZero(m.Amendments[i]) { // not required
			continue
		}

		if m.Amendments[i] != nil {
			if err := m.Amendments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amendments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amendments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 list amendments by contract response based on the context it is used
func (m *V1ListAmendmentsByContractResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmendments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ListAmendmentsByContractResponse) contextValidateAmendments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Amendments); i++ {

		if m.Amendments[i] != nil {

			if swag.IsZero(m.Amendments[i]) { // not required
				return nil
			}

			if err := m.Amendments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amendments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amendments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ListAmendmentsByContractResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListAmendmentsByContractResponse) UnmarshalBinary(b []byte) error {
	var res V1ListAmendmentsByContractResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
