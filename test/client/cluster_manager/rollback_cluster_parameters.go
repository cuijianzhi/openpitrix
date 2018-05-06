// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewRollbackClusterParams creates a new RollbackClusterParams object
// with the default values initialized.
func NewRollbackClusterParams() *RollbackClusterParams {
	var ()
	return &RollbackClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackClusterParamsWithTimeout creates a new RollbackClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRollbackClusterParamsWithTimeout(timeout time.Duration) *RollbackClusterParams {
	var ()
	return &RollbackClusterParams{

		timeout: timeout,
	}
}

// NewRollbackClusterParamsWithContext creates a new RollbackClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewRollbackClusterParamsWithContext(ctx context.Context) *RollbackClusterParams {
	var ()
	return &RollbackClusterParams{

		Context: ctx,
	}
}

// NewRollbackClusterParamsWithHTTPClient creates a new RollbackClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRollbackClusterParamsWithHTTPClient(client *http.Client) *RollbackClusterParams {
	var ()
	return &RollbackClusterParams{
		HTTPClient: client,
	}
}

/*RollbackClusterParams contains all the parameters to send to the API endpoint
for the rollback cluster operation typically these are written to a http.Request
*/
type RollbackClusterParams struct {

	/*Body*/
	Body *models.OpenpitrixRollbackClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rollback cluster params
func (o *RollbackClusterParams) WithTimeout(timeout time.Duration) *RollbackClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback cluster params
func (o *RollbackClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback cluster params
func (o *RollbackClusterParams) WithContext(ctx context.Context) *RollbackClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback cluster params
func (o *RollbackClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback cluster params
func (o *RollbackClusterParams) WithHTTPClient(client *http.Client) *RollbackClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback cluster params
func (o *RollbackClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rollback cluster params
func (o *RollbackClusterParams) WithBody(body *models.OpenpitrixRollbackClusterRequest) *RollbackClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rollback cluster params
func (o *RollbackClusterParams) SetBody(body *models.OpenpitrixRollbackClusterRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}