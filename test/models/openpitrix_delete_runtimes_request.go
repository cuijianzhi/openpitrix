// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteRuntimesRequest openpitrix delete runtimes request
// swagger:model openpitrixDeleteRuntimesRequest
type OpenpitrixDeleteRuntimesRequest struct {

	// required, ids of runtime to delete
	RuntimeID []string `json:"runtime_id"`
}

// Validate validates this openpitrix delete runtimes request
func (m *OpenpitrixDeleteRuntimesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuntimeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteRuntimesRequest) validateRuntimeID(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteRuntimesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteRuntimesRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteRuntimesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
