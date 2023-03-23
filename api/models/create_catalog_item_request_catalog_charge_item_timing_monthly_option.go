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

// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption UNIT_MONTHLY指定時のオプション(1~31の数字 or Enumの文字列)
//
// swagger:model CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption
type CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption string

func NewCreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption(value CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption) *CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption.
func (m CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption) Pointer() *CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption {
	return &m
}

const (

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY1 captures enum value "MONTHLY_OPTION_DAY_1"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY1 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_1"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY2 captures enum value "MONTHLY_OPTION_DAY_2"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY2 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_2"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY3 captures enum value "MONTHLY_OPTION_DAY_3"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY3 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_3"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY4 captures enum value "MONTHLY_OPTION_DAY_4"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY4 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_4"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY5 captures enum value "MONTHLY_OPTION_DAY_5"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY5 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_5"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY6 captures enum value "MONTHLY_OPTION_DAY_6"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY6 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_6"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY7 captures enum value "MONTHLY_OPTION_DAY_7"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY7 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_7"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY8 captures enum value "MONTHLY_OPTION_DAY_8"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY8 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_8"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY9 captures enum value "MONTHLY_OPTION_DAY_9"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY9 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_9"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY10 captures enum value "MONTHLY_OPTION_DAY_10"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY10 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_10"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY11 captures enum value "MONTHLY_OPTION_DAY_11"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY11 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_11"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY12 captures enum value "MONTHLY_OPTION_DAY_12"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY12 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_12"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY13 captures enum value "MONTHLY_OPTION_DAY_13"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY13 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_13"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY14 captures enum value "MONTHLY_OPTION_DAY_14"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY14 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_14"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY15 captures enum value "MONTHLY_OPTION_DAY_15"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY15 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_15"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY16 captures enum value "MONTHLY_OPTION_DAY_16"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY16 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_16"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY17 captures enum value "MONTHLY_OPTION_DAY_17"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY17 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_17"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY18 captures enum value "MONTHLY_OPTION_DAY_18"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY18 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_18"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY19 captures enum value "MONTHLY_OPTION_DAY_19"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY19 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_19"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY20 captures enum value "MONTHLY_OPTION_DAY_20"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY20 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_20"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY21 captures enum value "MONTHLY_OPTION_DAY_21"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY21 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_21"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY22 captures enum value "MONTHLY_OPTION_DAY_22"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY22 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_22"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY23 captures enum value "MONTHLY_OPTION_DAY_23"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY23 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_23"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY24 captures enum value "MONTHLY_OPTION_DAY_24"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY24 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_24"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY25 captures enum value "MONTHLY_OPTION_DAY_25"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY25 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_25"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY26 captures enum value "MONTHLY_OPTION_DAY_26"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY26 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_26"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY27 captures enum value "MONTHLY_OPTION_DAY_27"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY27 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_27"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY28 captures enum value "MONTHLY_OPTION_DAY_28"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY28 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_28"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY29 captures enum value "MONTHLY_OPTION_DAY_29"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY29 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_29"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY30 captures enum value "MONTHLY_OPTION_DAY_30"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY30 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_30"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY31 captures enum value "MONTHLY_OPTION_DAY_31"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONDAY31 CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_DAY_31"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONFIRSTDAYOFMONTH captures enum value "MONTHLY_OPTION_FIRST_DAY_OF_MONTH"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONFIRSTDAYOFMONTH CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_FIRST_DAY_OF_MONTH"

	// CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONLASTDAYOFMONTH captures enum value "MONTHLY_OPTION_LAST_DAY_OF_MONTH"
	CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionMONTHLYOPTIONLASTDAYOFMONTH CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption = "MONTHLY_OPTION_LAST_DAY_OF_MONTH"
)

// for schema
var createCatalogItemRequestCatalogChargeItemTimingMonthlyOptionEnum []interface{}

func init() {
	var res []CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption
	if err := json.Unmarshal([]byte(`["MONTHLY_OPTION_DAY_1","MONTHLY_OPTION_DAY_2","MONTHLY_OPTION_DAY_3","MONTHLY_OPTION_DAY_4","MONTHLY_OPTION_DAY_5","MONTHLY_OPTION_DAY_6","MONTHLY_OPTION_DAY_7","MONTHLY_OPTION_DAY_8","MONTHLY_OPTION_DAY_9","MONTHLY_OPTION_DAY_10","MONTHLY_OPTION_DAY_11","MONTHLY_OPTION_DAY_12","MONTHLY_OPTION_DAY_13","MONTHLY_OPTION_DAY_14","MONTHLY_OPTION_DAY_15","MONTHLY_OPTION_DAY_16","MONTHLY_OPTION_DAY_17","MONTHLY_OPTION_DAY_18","MONTHLY_OPTION_DAY_19","MONTHLY_OPTION_DAY_20","MONTHLY_OPTION_DAY_21","MONTHLY_OPTION_DAY_22","MONTHLY_OPTION_DAY_23","MONTHLY_OPTION_DAY_24","MONTHLY_OPTION_DAY_25","MONTHLY_OPTION_DAY_26","MONTHLY_OPTION_DAY_27","MONTHLY_OPTION_DAY_28","MONTHLY_OPTION_DAY_29","MONTHLY_OPTION_DAY_30","MONTHLY_OPTION_DAY_31","MONTHLY_OPTION_FIRST_DAY_OF_MONTH","MONTHLY_OPTION_LAST_DAY_OF_MONTH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCatalogItemRequestCatalogChargeItemTimingMonthlyOptionEnum = append(createCatalogItemRequestCatalogChargeItemTimingMonthlyOptionEnum, v)
	}
}

func (m CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption) validateCreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionEnum(path, location string, value CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption) error {
	if err := validate.EnumCase(path, location, value, createCatalogItemRequestCatalogChargeItemTimingMonthlyOptionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this create catalog item request catalog charge item timing monthly option
func (m CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCreateCatalogItemRequestCatalogChargeItemTimingMonthlyOptionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this create catalog item request catalog charge item timing monthly option based on context it is used
func (m CreateCatalogItemRequestCatalogChargeItemTimingMonthlyOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
