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

// Steps steps
//
// swagger:model steps
type Steps struct {

	// instructions
	Instructions []*Step `json:"instructions"`

	// next instruction seconds
	NextInstructionSeconds *int64 `json:"next_instruction_seconds,omitempty"`
}

// Validate validates this steps
func (m *Steps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstructions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Steps) validateInstructions(formats strfmt.Registry) error {

	if swag.IsZero(m.Instructions) { // not required
		return nil
	}

	for i := 0; i < len(m.Instructions); i++ {
		if swag.IsZero(m.Instructions[i]) { // not required
			continue
		}

		if m.Instructions[i] != nil {
			if err := m.Instructions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instructions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Steps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Steps) UnmarshalBinary(b []byte) error {
	var res Steps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
