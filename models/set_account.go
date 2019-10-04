package models

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetAccount set account
// swagger:model set-account
type SetAccount struct {
	// Account ID
	AcctID string `json:"acctId,omitempty"`
}

// Validate validates this set account
func (m *SetAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAccount) UnmarshalBinary(b []byte) error {
	var res SetAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
