// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateGroupRequest openpitrix create group request
// swagger:model openpitrixCreateGroupRequest
type OpenpitrixCreateGroupRequest struct {

	// group description
	Description string `json:"description,omitempty"`

	// required, group name
	Name string `json:"name,omitempty"`

	// required, parent group id
	ParentGroupID string `json:"parent_group_id,omitempty"`
}

// Validate validates this openpitrix create group request
func (m *OpenpitrixCreateGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateGroupRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
