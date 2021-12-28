// Code generated by go-swagger; DO NOT EDIT.

package customer_shipping_address

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

// NewGetCustomerShippingAddressesListParams creates a new GetCustomerShippingAddressesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomerShippingAddressesListParams() *GetCustomerShippingAddressesListParams {
	return &GetCustomerShippingAddressesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerShippingAddressesListParamsWithTimeout creates a new GetCustomerShippingAddressesListParams object
// with the ability to set a timeout on a request.
func NewGetCustomerShippingAddressesListParamsWithTimeout(timeout time.Duration) *GetCustomerShippingAddressesListParams {
	return &GetCustomerShippingAddressesListParams{
		timeout: timeout,
	}
}

// NewGetCustomerShippingAddressesListParamsWithContext creates a new GetCustomerShippingAddressesListParams object
// with the ability to set a context for a request.
func NewGetCustomerShippingAddressesListParamsWithContext(ctx context.Context) *GetCustomerShippingAddressesListParams {
	return &GetCustomerShippingAddressesListParams{
		Context: ctx,
	}
}

// NewGetCustomerShippingAddressesListParamsWithHTTPClient creates a new GetCustomerShippingAddressesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomerShippingAddressesListParamsWithHTTPClient(client *http.Client) *GetCustomerShippingAddressesListParams {
	return &GetCustomerShippingAddressesListParams{
		HTTPClient: client,
	}
}

/* GetCustomerShippingAddressesListParams contains all the parameters to send to the API endpoint
   for the get customer shipping addresses list operation.

   Typically these are written to a http.Request.
*/
type GetCustomerShippingAddressesListParams struct {

	/* CustomerTokenID.

	   The TokenId of a customer.
	*/
	CustomerTokenID string

	/* Limit.

	   The maximum number that can be returned in the array starting from the offset record in zero-based dataset. Default is 20, maximum is 100.

	   Format: int64
	   Default: 20
	*/
	Limit *int64

	/* Offset.

	   Starting record in zero-based dataset that should be returned as the first object in the array. Default is 0.

	   Format: int64
	*/
	Offset *int64

	/* ProfileID.

	   The id of a profile containing user specific TMS configuration.
	*/
	ProfileID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get customer shipping addresses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomerShippingAddressesListParams) WithDefaults() *GetCustomerShippingAddressesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get customer shipping addresses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomerShippingAddressesListParams) SetDefaults() {
	var (
		limitDefault = int64(20)

		offsetDefault = int64(0)
	)

	val := GetCustomerShippingAddressesListParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithTimeout(timeout time.Duration) *GetCustomerShippingAddressesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithContext(ctx context.Context) *GetCustomerShippingAddressesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithHTTPClient(client *http.Client) *GetCustomerShippingAddressesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerTokenID adds the customerTokenID to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithCustomerTokenID(customerTokenID string) *GetCustomerShippingAddressesListParams {
	o.SetCustomerTokenID(customerTokenID)
	return o
}

// SetCustomerTokenID adds the customerTokenId to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetCustomerTokenID(customerTokenID string) {
	o.CustomerTokenID = customerTokenID
}

// WithLimit adds the limit to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithLimit(limit *int64) *GetCustomerShippingAddressesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithOffset(offset *int64) *GetCustomerShippingAddressesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithProfileID adds the profileID to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) WithProfileID(profileID *string) *GetCustomerShippingAddressesListParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get customer shipping addresses list params
func (o *GetCustomerShippingAddressesListParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerShippingAddressesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerTokenId
	if err := r.SetPathParam("customerTokenId", o.CustomerTokenID); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.ProfileID != nil {

		// header param profile-id
		if err := r.SetHeaderParam("profile-id", *o.ProfileID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}