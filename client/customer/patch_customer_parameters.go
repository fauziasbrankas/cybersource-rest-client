// Code generated by go-swagger; DO NOT EDIT.

package customer

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
)

// NewPatchCustomerParams creates a new PatchCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchCustomerParams() *PatchCustomerParams {
	return &PatchCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCustomerParamsWithTimeout creates a new PatchCustomerParams object
// with the ability to set a timeout on a request.
func NewPatchCustomerParamsWithTimeout(timeout time.Duration) *PatchCustomerParams {
	return &PatchCustomerParams{
		timeout: timeout,
	}
}

// NewPatchCustomerParamsWithContext creates a new PatchCustomerParams object
// with the ability to set a context for a request.
func NewPatchCustomerParamsWithContext(ctx context.Context) *PatchCustomerParams {
	return &PatchCustomerParams{
		Context: ctx,
	}
}

// NewPatchCustomerParamsWithHTTPClient creates a new PatchCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchCustomerParamsWithHTTPClient(client *http.Client) *PatchCustomerParams {
	return &PatchCustomerParams{
		HTTPClient: client,
	}
}

/* PatchCustomerParams contains all the parameters to send to the API endpoint
   for the patch customer operation.

   Typically these are written to a http.Request.
*/
type PatchCustomerParams struct {

	/* CustomerTokenID.

	   The TokenId of a customer.
	*/
	CustomerTokenID string

	/* IfMatch.

	   Contains an ETag value from a GET request to make the request conditional.
	*/
	IfMatch *string

	// PatchCustomerRequest.
	PatchCustomerRequest PatchCustomerBody

	/* ProfileID.

	   The id of a profile containing user specific TMS configuration.
	*/
	ProfileID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCustomerParams) WithDefaults() *PatchCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch customer params
func (o *PatchCustomerParams) WithTimeout(timeout time.Duration) *PatchCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch customer params
func (o *PatchCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch customer params
func (o *PatchCustomerParams) WithContext(ctx context.Context) *PatchCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch customer params
func (o *PatchCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch customer params
func (o *PatchCustomerParams) WithHTTPClient(client *http.Client) *PatchCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch customer params
func (o *PatchCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerTokenID adds the customerTokenID to the patch customer params
func (o *PatchCustomerParams) WithCustomerTokenID(customerTokenID string) *PatchCustomerParams {
	o.SetCustomerTokenID(customerTokenID)
	return o
}

// SetCustomerTokenID adds the customerTokenId to the patch customer params
func (o *PatchCustomerParams) SetCustomerTokenID(customerTokenID string) {
	o.CustomerTokenID = customerTokenID
}

// WithIfMatch adds the ifMatch to the patch customer params
func (o *PatchCustomerParams) WithIfMatch(ifMatch *string) *PatchCustomerParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the patch customer params
func (o *PatchCustomerParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithPatchCustomerRequest adds the patchCustomerRequest to the patch customer params
func (o *PatchCustomerParams) WithPatchCustomerRequest(patchCustomerRequest PatchCustomerBody) *PatchCustomerParams {
	o.SetPatchCustomerRequest(patchCustomerRequest)
	return o
}

// SetPatchCustomerRequest adds the patchCustomerRequest to the patch customer params
func (o *PatchCustomerParams) SetPatchCustomerRequest(patchCustomerRequest PatchCustomerBody) {
	o.PatchCustomerRequest = patchCustomerRequest
}

// WithProfileID adds the profileID to the patch customer params
func (o *PatchCustomerParams) WithProfileID(profileID *string) *PatchCustomerParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the patch customer params
func (o *PatchCustomerParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerTokenId
	if err := r.SetPathParam("customerTokenId", o.CustomerTokenID); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.PatchCustomerRequest); err != nil {
		return err
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
