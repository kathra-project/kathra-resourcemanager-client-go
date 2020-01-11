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

// CatalogEntry catalog entry
// swagger:model CatalogEntry
type CatalogEntry struct {
	Resource

	// Description
	Description string `json:"description,omitempty"`

	// package template
	// Enum: [REST_SERVICE]
	PackageTemplate interface{} `json:"packageTemplate,omitempty"`

	// packages
	Packages []*CatalogEntryPackage `json:"packages"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CatalogEntry) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Resource = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		PackageTemplate interface{} `json:"packageTemplate,omitempty"`

		Packages []*CatalogEntryPackage `json:"packages"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.PackageTemplate = dataAO1.PackageTemplate

	m.Packages = dataAO1.Packages

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CatalogEntry) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		PackageTemplate interface{} `json:"packageTemplate,omitempty"`

		Packages []*CatalogEntryPackage `json:"packages"`
	}

	dataAO1.Description = m.Description

	dataAO1.PackageTemplate = m.PackageTemplate

	dataAO1.Packages = m.Packages

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this catalog entry
func (m *CatalogEntry) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Resource
	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogEntry) validatePackages(formats strfmt.Registry) error {

	if swag.IsZero(m.Packages) { // not required
		return nil
	}

	for i := 0; i < len(m.Packages); i++ {
		if swag.IsZero(m.Packages[i]) { // not required
			continue
		}

		if m.Packages[i] != nil {
			if err := m.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogEntry) UnmarshalBinary(b []byte) error {
	var res CatalogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}