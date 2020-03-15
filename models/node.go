// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Node node
//
// swagger:model node
type Node struct {
	Base

	NodeCreateParams

	// cluster id
	// Format: uuid
	ClusterID strfmt.UUID `json:"cluster_id,omitempty"`

	// status
	// Required: true
	// Enum: [tbd]
	Status *string `json:"status"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Node) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Base
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Base = aO0

	// AO1
	var aO1 NodeCreateParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NodeCreateParams = aO1

	// AO2
	var dataAO2 struct {
		ClusterID strfmt.UUID `json:"cluster_id,omitempty"`

		Status *string `json:"status"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.ClusterID = dataAO2.ClusterID

	m.Status = dataAO2.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Node) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.Base)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.NodeCreateParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		ClusterID strfmt.UUID `json:"cluster_id,omitempty"`

		Status *string `json:"status"`
	}

	dataAO2.ClusterID = m.ClusterID

	dataAO2.Status = m.Status

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Base
	if err := m.Base.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with NodeCreateParams
	if err := m.NodeCreateParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateClusterID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

var nodeTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tbd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeStatusPropEnum = append(nodeTypeStatusPropEnum, v)
	}
}

// property enum
func (m *Node) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nodeTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
