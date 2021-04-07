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

// DevicePeer device peer
//
// swagger:model DevicePeer
type DevicePeer struct {

	// wireguard
	Wireguard *WireGuardPeer `json:"wireguard,omitempty"`
}

// Validate validates this device peer
func (m *DevicePeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWireguard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicePeer) validateWireguard(formats strfmt.Registry) error {
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

// ContextValidate validate this device peer based on the context it is used
func (m *DevicePeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWireguard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicePeer) contextValidateWireguard(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DevicePeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevicePeer) UnmarshalBinary(b []byte) error {
	var res DevicePeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
