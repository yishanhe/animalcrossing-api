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

// Bug bug
//
// swagger:model Bug
type Bug struct {

	// catches to unlock
	CatchesToUnlock int64 `json:"catches_to_unlock,omitempty" bson:"catches_to_unlock"`

	// colors
	Colors Colors `json:"colors,omitempty" bson:"colors"`

	// entry id
	EntryID string `json:"entry_id,omitempty" bson:"entry_id"`

	// availability by hour range
	Hours string `json:"hours,omitempty" bson:"hours"`

	// id
	ID int64 `json:"id,omitempty" bson:"id"`

	// images
	Images []*Image `json:"images" bson:"images"`

	// location
	Location string `json:"location,omitempty" bson:"location"`

	// months
	Months *Months `json:"months,omitempty" bson:"months"`

	// name
	Name *Name `json:"name,omitempty" bson:"name"`

	// sell price
	SellPrice int64 `json:"sell_price,omitempty" bson:"sell_price"`

	// weather
	Weather string `json:"weather,omitempty" bson:"weather"`
}

// Validate validates this bug
func (m *Bug) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Bug) validateColors(formats strfmt.Registry) error {

	if swag.IsZero(m.Colors) { // not required
		return nil
	}

	if err := m.Colors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("colors")
		}
		return err
	}

	return nil
}

func (m *Bug) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Bug) validateMonths(formats strfmt.Registry) error {

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

func (m *Bug) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Bug) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bug) UnmarshalBinary(b []byte) error {
	var res Bug
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
