// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Device device
//
// swagger:model Device
type Device struct {

	// i ps
	IPs []IP `json:"ips"`

	// dns
	DNS []DNS `json:"dns"`

	// external key
	ExternalKey string `json:"external_key,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// last active at
	// Format: date-time
	LastActiveAt strfmt.DateTime `json:"last_active_at,omitempty"`

	// locations
	Locations []*Location `json:"locations"`

	// name
	Name string `json:"name,omitempty"`

	// ports
	Ports []Port `json:"ports"`

	// stats
	Stats *DeviceStats `json:"stats,omitempty"`

	// wireguard
	Wireguard *WireGuard `json:"wireguard,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActiveAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWireguard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateIPs(formats strfmt.Registry) error {
	if swag.IsZero(m.IPs) { // not required
		return nil
	}

	for i := 0; i < len(m.IPs); i++ {

		if err := m.IPs[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ips" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Device) validateDNS(formats strfmt.Registry) error {
	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	for i := 0; i < len(m.DNS); i++ {

		if err := m.DNS[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Device) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateLastActiveAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActiveAt) { // not required
		return nil
	}

	if err := validate.FormatOf("last_active_at", "body", "date-time", m.LastActiveAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	for i := 0; i < len(m.Locations); i++ {
		if swag.IsZero(m.Locations[i]) { // not required
			continue
		}

		if m.Locations[i] != nil {
			if err := m.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {

		if err := m.Ports[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Device) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
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

	if err := m.contextValidateIPs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWireguard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) contextValidateIPs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPs); i++ {

		if err := m.IPs[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ips" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Device) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNS); i++ {

		if err := m.DNS[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Device) contextValidateLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Locations); i++ {

		if m.Locations[i] != nil {
			if err := m.Locations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ports" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Device) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {
		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
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
