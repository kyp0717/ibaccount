// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CalendarRequest calendar request
// swagger:model calendar_request
type CalendarRequest struct {

	// date
	Date *CalendarRequestDate `json:"date,omitempty"`

	// filters
	Filters *CalendarRequestFilters `json:"filters,omitempty"`
}

// Validate validates this calendar request
func (m *CalendarRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalendarRequest) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if m.Date != nil {
		if err := m.Date.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("date")
			}
			return err
		}
	}

	return nil
}

func (m *CalendarRequest) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {
		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CalendarRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalendarRequest) UnmarshalBinary(b []byte) error {
	var res CalendarRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CalendarRequestDate calendar request date
// swagger:model CalendarRequestDate
type CalendarRequestDate struct {

	// end date of a period. for example 20180808-0400
	End string `json:"end,omitempty"`

	// start date of a period. for example 20180808-0400
	Start string `json:"start,omitempty"`
}

// Validate validates this calendar request date
func (m *CalendarRequestDate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CalendarRequestDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalendarRequestDate) UnmarshalBinary(b []byte) error {
	var res CalendarRequestDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CalendarRequestFilters calendar request filters
// swagger:model CalendarRequestFilters
type CalendarRequestFilters struct {

	// value can be 'true' or 'false'.
	DivExDates string `json:"DivExDates,omitempty"`

	// value can be 'true' or 'false'.
	CorporateEarnings string `json:"corporate_earnings,omitempty"`

	// value can be 'true' or 'false'.
	CorporateEvents string `json:"corporate_events,omitempty"`

	// default is 'All'.
	Country string `json:"country,omitempty"`

	// value can be 'true' or 'false'.
	EconomicEvents string `json:"economic_events,omitempty"`

	// value can be 'true' or 'false'.
	Ipo string `json:"ipo,omitempty"`

	// default is '250'.
	Limit string `json:"limit,omitempty"`

	// default is '50'.
	LimitRegion string `json:"limit_region,omitempty"`

	// value can be 'true' or 'false'.
	OptionShowMonthly string `json:"option_show_monthly,omitempty"`

	// value can be 'true' or 'false'.
	OptionShowWeekly string `json:"option_show_weekly,omitempty"`

	// value can be 'true' or 'false'.
	RecentlyHeld string `json:"recently_held,omitempty"`

	// value can be 'true' or 'false'.
	Splits string `json:"splits,omitempty"`
}

// Validate validates this calendar request filters
func (m *CalendarRequestFilters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CalendarRequestFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalendarRequestFilters) UnmarshalBinary(b []byte) error {
	var res CalendarRequestFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
