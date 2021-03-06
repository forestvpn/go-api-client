// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// DeviceBindings device bindings
//
// swagger:model DeviceBindings
type DeviceBindings []string

// Validate validates this device bindings
func (m DeviceBindings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device bindings based on context it is used
func (m DeviceBindings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
