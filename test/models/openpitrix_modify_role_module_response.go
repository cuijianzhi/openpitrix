// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRoleModuleResponse openpitrix modify role module response
// swagger:model openpitrixModifyRoleModuleResponse
type OpenpitrixModifyRoleModuleResponse struct {

	// role id used to modify role module
	RoleID string `json:"role_id,omitempty"`
}

// Validate validates this openpitrix modify role module response
func (m *OpenpitrixModifyRoleModuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRoleModuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRoleModuleResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRoleModuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
