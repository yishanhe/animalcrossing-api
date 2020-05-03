// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Availability availability
//
// swagger:model Availability
type Availability struct {

	// availability by hour range
	Hours [][]string `json:"hours"`

	// location
	Location string `json:"location,omitempty"`

	// months
	Months *Months `json:"months,omitempty"`

	// rarity
	Rarity Rarity `json:"rarity,omitempty"`

	// weather
	Weather []Weather `json:"weather"`
}

// Validate validates this availability
func (m *Availability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHours(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRarity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeather(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Availability) validateHours(formats strfmt.Registry) error {

	if swag.IsZero(m.Hours) { // not required
		return nil
	}

	for i := 0; i < len(m.Hours); i++ {

	}

	return nil
}

func (m *Availability) validateMonths(formats strfmt.Registry) error {

	if swag.IsZero(m.Months) { // not required
		return nil
	}

	if m.Months != nil {
		if err := m.Months.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("months")
			}
			return err
		}
	}

	return nil
}

func (m *Availability) validateRarity(formats strfmt.Registry) error {

	if swag.IsZero(m.Rarity) { // not required
		return nil
	}

	if err := m.Rarity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rarity")
		}
		return err
	}

	return nil
}

func (m *Availability) validateWeather(formats strfmt.Registry) error {

	if swag.IsZero(m.Weather) { // not required
		return nil
	}

	for i := 0; i < len(m.Weather); i++ {

		if err := m.Weather[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weather" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Availability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Availability) UnmarshalBinary(b []byte) error {
	var res Availability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}