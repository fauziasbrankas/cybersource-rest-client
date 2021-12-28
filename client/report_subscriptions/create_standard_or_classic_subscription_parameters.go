// Code generated by go-swagger; DO NOT EDIT.

package report_subscriptions

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

// NewCreateStandardOrClassicSubscriptionParams creates a new CreateStandardOrClassicSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStandardOrClassicSubscriptionParams() *CreateStandardOrClassicSubscriptionParams {
	return &CreateStandardOrClassicSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStandardOrClassicSubscriptionParamsWithTimeout creates a new CreateStandardOrClassicSubscriptionParams object
// with the ability to set a timeout on a request.
func NewCreateStandardOrClassicSubscriptionParamsWithTimeout(timeout time.Duration) *CreateStandardOrClassicSubscriptionParams {
	return &CreateStandardOrClassicSubscriptionParams{
		timeout: timeout,
	}
}

// NewCreateStandardOrClassicSubscriptionParamsWithContext creates a new CreateStandardOrClassicSubscriptionParams object
// with the ability to set a context for a request.
func NewCreateStandardOrClassicSubscriptionParamsWithContext(ctx context.Context) *CreateStandardOrClassicSubscriptionParams {
	return &CreateStandardOrClassicSubscriptionParams{
		Context: ctx,
	}
}

// NewCreateStandardOrClassicSubscriptionParamsWithHTTPClient creates a new CreateStandardOrClassicSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStandardOrClassicSubscriptionParamsWithHTTPClient(client *http.Client) *CreateStandardOrClassicSubscriptionParams {
	return &CreateStandardOrClassicSubscriptionParams{
		HTTPClient: client,
	}
}

/* CreateStandardOrClassicSubscriptionParams contains all the parameters to send to the API endpoint
   for the create standard or classic subscription operation.

   Typically these are written to a http.Request.
*/
type CreateStandardOrClassicSubscriptionParams struct {

	/* OrganizationID.

	   Valid Organization Id
	*/
	OrganizationID *string

	/* PredefinedSubscriptionRequestBean.

	   Report subscription request payload
	*/
	PredefinedSubscriptionRequestBean CreateStandardOrClassicSubscriptionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create standard or classic subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStandardOrClassicSubscriptionParams) WithDefaults() *CreateStandardOrClassicSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create standard or classic subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStandardOrClassicSubscriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) WithTimeout(timeout time.Duration) *CreateStandardOrClassicSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) WithContext(ctx context.Context) *CreateStandardOrClassicSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) WithHTTPClient(client *http.Client) *CreateStandardOrClassicSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) WithOrganizationID(organizationID *string) *CreateStandardOrClassicSubscriptionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithPredefinedSubscriptionRequestBean adds the predefinedSubscriptionRequestBean to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) WithPredefinedSubscriptionRequestBean(predefinedSubscriptionRequestBean CreateStandardOrClassicSubscriptionBody) *CreateStandardOrClassicSubscriptionParams {
	o.SetPredefinedSubscriptionRequestBean(predefinedSubscriptionRequestBean)
	return o
}

// SetPredefinedSubscriptionRequestBean adds the predefinedSubscriptionRequestBean to the create standard or classic subscription params
func (o *CreateStandardOrClassicSubscriptionParams) SetPredefinedSubscriptionRequestBean(predefinedSubscriptionRequestBean CreateStandardOrClassicSubscriptionBody) {
	o.PredefinedSubscriptionRequestBean = predefinedSubscriptionRequestBean
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStandardOrClassicSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID string

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.PredefinedSubscriptionRequestBean); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}