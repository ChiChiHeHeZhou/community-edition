// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateUnmanagedClusterParams create unmanaged cluster params
// swagger:model CreateUnmanagedClusterParams
type CreateUnmanagedClusterParams struct {

	// cni
	Cni string `json:"cni,omitempty"`

	// controlplanecount
	Controlplanecount int64 `json:"controlplanecount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// podcidr
	Podcidr string `json:"podcidr,omitempty"`

	// portmappings
	Portmappings []string `json:"portmappings"`

	// provider
	Provider string `json:"provider,omitempty"`

	// servicecidr
	Servicecidr string `json:"servicecidr,omitempty"`

	// workernodecount
	Workernodecount int64 `json:"workernodecount,omitempty"`
}

// Validate validates this create unmanaged cluster params
func (m *CreateUnmanagedClusterParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUnmanagedClusterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUnmanagedClusterParams) UnmarshalBinary(b []byte) error {
	var res CreateUnmanagedClusterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
