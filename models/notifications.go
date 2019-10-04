// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Notifications notifications
// swagger:model notifications
type Notifications []*NotificationsItems0

// Validate validates this notifications
func (m Notifications) Validate(formats strfmt.Registry) error {
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

// NotificationsItems0 notification
// swagger:model NotificationsItems0
type NotificationsItems0 struct {

	// notification date
	D string `json:"D,omitempty"`

	// FYI code, we can use it to find whether the disclaimer is accepted or not in settings
	FC string `json:"FC,omitempty"`

	// unique way to reference this notification
	ID string `json:"ID,omitempty"`

	// content of notification
	MD string `json:"MD,omitempty"`

	// title of notification
	MS string `json:"MS,omitempty"`

	// 0-unread, 1-read
	R string `json:"R,omitempty"`
}

// Validate validates this notifications items0
func (m *NotificationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationsItems0) UnmarshalBinary(b []byte) error {
	var res NotificationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}