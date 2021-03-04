// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Device device
//
// swagger:model Device
type Device struct {

	// i ps
	IPs []string `json:"ips"`

	// ports
	Ports []int64 `json:"ports"`

	// wireguard
	Wireguard *WireguardSpec `json:"wireguard,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWireguard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateWireguard(formats strfmt.Registry) error {
	if swag.IsZero(m.Wireguard) { // not required
		return nil
	}

	if m.Wireguard != nil {
		if err := m.Wireguard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireguard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device based on the context it is used
func (m *Device) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWireguard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) contextValidateWireguard(ctx context.Context, formats strfmt.Registry) error {

	if m.Wireguard != nil {
		if err := m.Wireguard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireguard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
