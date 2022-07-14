// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewUpdateDiscoveryIgnitionParams creates a new UpdateDiscoveryIgnitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDiscoveryIgnitionParams() *UpdateDiscoveryIgnitionParams {
	return &UpdateDiscoveryIgnitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDiscoveryIgnitionParamsWithTimeout creates a new UpdateDiscoveryIgnitionParams object
// with the ability to set a timeout on a request.
func NewUpdateDiscoveryIgnitionParamsWithTimeout(timeout time.Duration) *UpdateDiscoveryIgnitionParams {
	return &UpdateDiscoveryIgnitionParams{
		timeout: timeout,
	}
}

// NewUpdateDiscoveryIgnitionParamsWithContext creates a new UpdateDiscoveryIgnitionParams object
// with the ability to set a context for a request.
func NewUpdateDiscoveryIgnitionParamsWithContext(ctx context.Context) *UpdateDiscoveryIgnitionParams {
	return &UpdateDiscoveryIgnitionParams{
		Context: ctx,
	}
}

// NewUpdateDiscoveryIgnitionParamsWithHTTPClient creates a new UpdateDiscoveryIgnitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDiscoveryIgnitionParamsWithHTTPClient(client *http.Client) *UpdateDiscoveryIgnitionParams {
	return &UpdateDiscoveryIgnitionParams{
		HTTPClient: client,
	}
}

/* UpdateDiscoveryIgnitionParams contains all the parameters to send to the API endpoint
   for the update discovery ignition operation.

   Typically these are written to a http.Request.
*/
type UpdateDiscoveryIgnitionParams struct {

	/* ClusterID.

	   The cluster for which the discovery ignition config should be updated.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* DiscoveryIgnitionParams.

	   Overrides for the discovery ignition config.
	*/
	DiscoveryIgnitionParams *models.DiscoveryIgnitionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update discovery ignition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDiscoveryIgnitionParams) WithDefaults() *UpdateDiscoveryIgnitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update discovery ignition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDiscoveryIgnitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) WithTimeout(timeout time.Duration) *UpdateDiscoveryIgnitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) WithContext(ctx context.Context) *UpdateDiscoveryIgnitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) WithHTTPClient(client *http.Client) *UpdateDiscoveryIgnitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) WithClusterID(clusterID strfmt.UUID) *UpdateDiscoveryIgnitionParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithDiscoveryIgnitionParams adds the discoveryIgnitionParams to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) WithDiscoveryIgnitionParams(discoveryIgnitionParams *models.DiscoveryIgnitionParams) *UpdateDiscoveryIgnitionParams {
	o.SetDiscoveryIgnitionParams(discoveryIgnitionParams)
	return o
}

// SetDiscoveryIgnitionParams adds the discoveryIgnitionParams to the update discovery ignition params
func (o *UpdateDiscoveryIgnitionParams) SetDiscoveryIgnitionParams(discoveryIgnitionParams *models.DiscoveryIgnitionParams) {
	o.DiscoveryIgnitionParams = discoveryIgnitionParams
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDiscoveryIgnitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}
	if o.DiscoveryIgnitionParams != nil {
		if err := r.SetBodyParam(o.DiscoveryIgnitionParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}