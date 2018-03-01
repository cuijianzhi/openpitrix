// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRuntimeEnvCredential openpitrix runtime env credential
// swagger:model openpitrixRuntimeEnvCredential
type OpenpitrixRuntimeEnvCredential struct {

	// content
	Content map[string]string `json:"content,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// runtime env credential id
	RuntimeEnvCredentialID string `json:"runtime_env_credential_id,omitempty"`

	// runtime env ids
	RuntimeEnvIds []string `json:"runtime_env_ids"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`
}

// Validate validates this openpitrix runtime env credential
func (m *OpenpitrixRuntimeEnvCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuntimeEnvIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixRuntimeEnvCredential) validateRuntimeEnvIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeEnvIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRuntimeEnvCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRuntimeEnvCredential) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRuntimeEnvCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
