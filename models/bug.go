// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Bug bug
//
// swagger:model Bug
type Bug struct {
	idField *int64

	nameField *Name

	// availability
	Availability *Availability `json:"availability,omitempty"`

	// price
	Price int64 `json:"price,omitempty"`

	// shadow
	Shadow string `json:"shadow,omitempty"`
}

// ID gets the id of this subtype
func (m *Bug) ID() *int64 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *Bug) SetID(val *int64) {
	m.idField = val
}

// Name gets the name of this subtype
func (m *Bug) Name() *Name {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *Bug) SetName(val *Name) {
	m.nameField = val
}

// ResourceType gets the resource type of this subtype
func (m *Bug) ResourceType() string {
	return "Bug"
}

// SetResourceType sets the resource type of this subtype
func (m *Bug) SetResourceType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Bug) UnmarshalJSON(raw []byte) error {
	var data struct {

		// availability
		Availability *Availability `json:"availability,omitempty"`

		// price
		Price int64 `json:"price,omitempty"`

		// shadow
		Shadow string `json:"shadow,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ID *int64 `json:"id"`

		Name *Name `json:"name"`

		ResourceType string `json:"resource_type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result Bug

	result.idField = base.ID

	result.nameField = base.Name

	if base.ResourceType != result.ResourceType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid resource_type value: %q", base.ResourceType)
	}

	result.Availability = data.Availability
	result.Price = data.Price
	result.Shadow = data.Shadow

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Bug) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// availability
		Availability *Availability `json:"availability,omitempty"`

		// price
		Price int64 `json:"price,omitempty"`

		// shadow
		Shadow string `json:"shadow,omitempty"`
	}{

		Availability: m.Availability,

		Price: m.Price,

		Shadow: m.Shadow,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ID *int64 `json:"id"`

		Name *Name `json:"name"`

		ResourceType string `json:"resource_type"`
	}{

		ID: m.ID(),

		Name: m.Name(),

		ResourceType: m.ResourceType(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this bug
func (m *Bug) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Bug) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

func (m *Bug) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	if m.Name() != nil {
		if err := m.Name().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *Bug) validateAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	if m.Availability != nil {
		if err := m.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability")
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
