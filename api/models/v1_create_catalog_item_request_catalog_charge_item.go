// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateCatalogItemRequestCatalogChargeItem v1 create catalog item request catalog charge item
//
// swagger:model v1CreateCatalogItemRequestCatalogChargeItem
type V1CreateCatalogItemRequestCatalogChargeItem struct {

	// 一括請求オプション（定期課金のみ指定できるオプションです）
	BillAllAtOnce bool `json:"billAllAtOnce,omitempty"`

	// 請求周期
	BillingCycle *V1CreateCatalogItemRequestTerm `json:"billingCycle,omitempty"`

	// 支払期限タイミング
	BillingDueTiming *CreateCatalogItemRequestCatalogChargeItemTiming `json:"billingDueTiming,omitempty"`

	// 請求タイミング(ex: 翌々月 初日)
	BillingTiming *CreateCatalogItemRequestCatalogChargeItemTiming `json:"billingTiming,omitempty"`

	// 終了月日割持設定
	EndDailyRate *CreateCatalogItemRequestCatalogChargeItemDailyRate `json:"endDailyRate,omitempty"`

	// 課金項目名
	Name string `json:"name,omitempty"`

	// 料金モデル
	PricingModel *CatalogChargeItemPricingModel `json:"pricingModel,omitempty"`

	// 端数処理の方法
	Rounding CatalogChargeItemRounding `json:"rounding,omitempty"`

	// 端数処理の精度
	Scale int32 `json:"scale,omitempty"`

	// 開始月日割持設定
	StartDailyRate *CreateCatalogItemRequestCatalogChargeItemDailyRate `json:"startDailyRate,omitempty"`

	// 課金項目タイプ
	Type CreateCatalogItemRequestCatalogChargeItemType `json:"type,omitempty"`
}

// Validate validates this v1 create catalog item request catalog charge item
func (m *V1CreateCatalogItemRequestCatalogChargeItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingDueTiming(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingTiming(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDailyRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRounding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDailyRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateBillingCycle(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingCycle) { // not required
		return nil
	}

	if m.BillingCycle != nil {
		if err := m.BillingCycle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingCycle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingCycle")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateBillingDueTiming(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingDueTiming) { // not required
		return nil
	}

	if m.BillingDueTiming != nil {
		if err := m.BillingDueTiming.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingDueTiming")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingDueTiming")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateBillingTiming(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingTiming) { // not required
		return nil
	}

	if m.BillingTiming != nil {
		if err := m.BillingTiming.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingTiming")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingTiming")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateEndDailyRate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDailyRate) { // not required
		return nil
	}

	if m.EndDailyRate != nil {
		if err := m.EndDailyRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endDailyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endDailyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validatePricingModel(formats strfmt.Registry) error {
	if swag.IsZero(m.PricingModel) { // not required
		return nil
	}

	if m.PricingModel != nil {
		if err := m.PricingModel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricingModel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricingModel")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateRounding(formats strfmt.Registry) error {
	if swag.IsZero(m.Rounding) { // not required
		return nil
	}

	if err := m.Rounding.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rounding")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("rounding")
		}
		return err
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateStartDailyRate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDailyRate) { // not required
		return nil
	}

	if m.StartDailyRate != nil {
		if err := m.StartDailyRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startDailyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startDailyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 create catalog item request catalog charge item based on the context it is used
func (m *V1CreateCatalogItemRequestCatalogChargeItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingCycle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBillingDueTiming(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBillingTiming(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndDailyRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePricingModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRounding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartDailyRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateBillingCycle(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingCycle != nil {

		if swag.IsZero(m.BillingCycle) { // not required
			return nil
		}

		if err := m.BillingCycle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingCycle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingCycle")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateBillingDueTiming(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingDueTiming != nil {

		if swag.IsZero(m.BillingDueTiming) { // not required
			return nil
		}

		if err := m.BillingDueTiming.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingDueTiming")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingDueTiming")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateBillingTiming(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingTiming != nil {

		if swag.IsZero(m.BillingTiming) { // not required
			return nil
		}

		if err := m.BillingTiming.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingTiming")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingTiming")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateEndDailyRate(ctx context.Context, formats strfmt.Registry) error {

	if m.EndDailyRate != nil {

		if swag.IsZero(m.EndDailyRate) { // not required
			return nil
		}

		if err := m.EndDailyRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endDailyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endDailyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidatePricingModel(ctx context.Context, formats strfmt.Registry) error {

	if m.PricingModel != nil {

		if swag.IsZero(m.PricingModel) { // not required
			return nil
		}

		if err := m.PricingModel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricingModel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricingModel")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateRounding(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Rounding) { // not required
		return nil
	}

	if err := m.Rounding.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rounding")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("rounding")
		}
		return err
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateStartDailyRate(ctx context.Context, formats strfmt.Registry) error {

	if m.StartDailyRate != nil {

		if swag.IsZero(m.StartDailyRate) { // not required
			return nil
		}

		if err := m.StartDailyRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startDailyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startDailyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1CreateCatalogItemRequestCatalogChargeItem) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateCatalogItemRequestCatalogChargeItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateCatalogItemRequestCatalogChargeItem) UnmarshalBinary(b []byte) error {
	var res V1CreateCatalogItemRequestCatalogChargeItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
