// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeAppClustersResponse openpitrix describe app clusters response
// swagger:model openpitrixDescribeAppClustersResponse
type OpenpitrixDescribeAppClustersResponse struct {

	// cluster set
	ClusterSet OpenpitrixDescribeAppClustersResponseClusterSet `json:"cluster_set"`

	// total count of cluster of app
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe app clusters response
func (m *OpenpitrixDescribeAppClustersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeAppClustersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeAppClustersResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeAppClustersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
