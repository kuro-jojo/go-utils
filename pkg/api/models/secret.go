// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Secret secret
//
// swagger:model Secret
type Secret struct {

	// data
	// Required: true
	Data map[string]string `json:"data"`

	// The name of the secret
	// Required: true
	Name *string `json:"name"`

	// The scope of the secret
	// Required: true
	Scope *string `json:"scope"`
}

// Validate validates this secret
func (m *Secret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secret) validateData(formats strfmt.Registry) error {

	return nil
}

func (m *Secret) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Secret) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secret) UnmarshalBinary(b []byte) error {
	var res Secret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}