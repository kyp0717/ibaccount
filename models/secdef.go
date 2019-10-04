// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"strconv"
)

// Secdef secdef
// swagger:model secdef
type Secdef []*SecdefItems0

// Validate validates this secdef
func (m Secdef) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// SecdefItems0 security definition information
// swagger:model SecdefItems0
type SecdefItems0 struct {

	// asset class
	AssetClass string `json:"assetClass,omitempty"`

	// conid
	Conid int64 `json:"conid,omitempty"`

	// expiry
	Expiry string `json:"expiry,omitempty"`

	// full name
	FullName string `json:"fullName,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// last trading day
	LastTradingDay string `json:"lastTradingDay,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// page size
	PageSize int64 `json:"pageSize,omitempty"`

	// put or call
	PutOrCall string `json:"putOrCall,omitempty"`

	// sector
	Sector string `json:"sector,omitempty"`

	// sector group
	SectorGroup string `json:"sectorGroup,omitempty"`

	// strike
	Strike float64 `json:"strike,omitempty"`

	// ticker
	Ticker string `json:"ticker,omitempty"`

	// und conid
	UndConid int64 `json:"undConid,omitempty"`
}

// Validate validates this secdef items0
func (m *SecdefItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecdefItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecdefItems0) UnmarshalBinary(b []byte) error {
	var res SecdefItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}