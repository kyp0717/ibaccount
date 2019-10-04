// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Summary summary
// swagger:model summary
type Summary struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// is null
	IsNull bool `json:"isNull,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this summary
func (m *Summary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Summary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Summary) UnmarshalBinary(b []byte) error {
	var res Summary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}