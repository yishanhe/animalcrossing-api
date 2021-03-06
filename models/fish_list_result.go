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

// FishListResult fish list result
//
// swagger:model FishListResult
type FishListResult struct {
	ListResult

	// results
	Results []*Fish `json:"results"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FishListResult) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ListResult
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ListResult = aO0

	// AO1
	var dataAO1 struct {
		Results []*Fish `json:"results"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Results = dataAO1.Results

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FishListResult) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ListResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Results []*Fish `json:"results"`
	}

	dataAO1.Results = m.Results

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this fish list result
func (m *FishListResult) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ListResult
	if err := m.ListResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FishListResult) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FishListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FishListResult) UnmarshalBinary(b []byte) error {
	var res FishListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
