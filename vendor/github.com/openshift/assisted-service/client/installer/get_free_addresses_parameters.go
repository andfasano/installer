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
	"github.com/go-openapi/swag"
)

// NewGetFreeAddressesParams creates a new GetFreeAddressesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFreeAddressesParams() *GetFreeAddressesParams {
	return &GetFreeAddressesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFreeAddressesParamsWithTimeout creates a new GetFreeAddressesParams object
// with the ability to set a timeout on a request.
func NewGetFreeAddressesParamsWithTimeout(timeout time.Duration) *GetFreeAddressesParams {
	return &GetFreeAddressesParams{
		timeout: timeout,
	}
}

// NewGetFreeAddressesParamsWithContext creates a new GetFreeAddressesParams object
// with the ability to set a context for a request.
func NewGetFreeAddressesParamsWithContext(ctx context.Context) *GetFreeAddressesParams {
	return &GetFreeAddressesParams{
		Context: ctx,
	}
}

// NewGetFreeAddressesParamsWithHTTPClient creates a new GetFreeAddressesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFreeAddressesParamsWithHTTPClient(client *http.Client) *GetFreeAddressesParams {
	return &GetFreeAddressesParams{
		HTTPClient: client,
	}
}

/* GetFreeAddressesParams contains all the parameters to send to the API endpoint
   for the get free addresses operation.

   Typically these are written to a http.Request.
*/
type GetFreeAddressesParams struct {

	/* ClusterID.

	   The cluster to return free addresses for.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* Limit.

	   The maximum number of free addresses to return.

	   Default: 8000
	*/
	Limit *int64

	/* Network.

	   The cluster network to return free addresses for.
	*/
	Network string

	/* Prefix.

	   A prefix for the free addresses to return.
	*/
	Prefix *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get free addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFreeAddressesParams) WithDefaults() *GetFreeAddressesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get free addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFreeAddressesParams) SetDefaults() {
	var (
		limitDefault = int64(8000)
	)

	val := GetFreeAddressesParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get free addresses params
func (o *GetFreeAddressesParams) WithTimeout(timeout time.Duration) *GetFreeAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get free addresses params
func (o *GetFreeAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get free addresses params
func (o *GetFreeAddressesParams) WithContext(ctx context.Context) *GetFreeAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get free addresses params
func (o *GetFreeAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get free addresses params
func (o *GetFreeAddressesParams) WithHTTPClient(client *http.Client) *GetFreeAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get free addresses params
func (o *GetFreeAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get free addresses params
func (o *GetFreeAddressesParams) WithClusterID(clusterID strfmt.UUID) *GetFreeAddressesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get free addresses params
func (o *GetFreeAddressesParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithLimit adds the limit to the get free addresses params
func (o *GetFreeAddressesParams) WithLimit(limit *int64) *GetFreeAddressesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get free addresses params
func (o *GetFreeAddressesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNetwork adds the network to the get free addresses params
func (o *GetFreeAddressesParams) WithNetwork(network string) *GetFreeAddressesParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the get free addresses params
func (o *GetFreeAddressesParams) SetNetwork(network string) {
	o.Network = network
}

// WithPrefix adds the prefix to the get free addresses params
func (o *GetFreeAddressesParams) WithPrefix(prefix *string) *GetFreeAddressesParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the get free addresses params
func (o *GetFreeAddressesParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetFreeAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// query param network
	qrNetwork := o.Network
	qNetwork := qrNetwork
	if qNetwork != "" {

		if err := r.SetQueryParam("network", qNetwork); err != nil {
			return err
		}
	}

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string

		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {

			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}