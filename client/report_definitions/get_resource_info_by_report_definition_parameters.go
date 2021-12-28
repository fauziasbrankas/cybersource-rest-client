// Code generated by go-swagger; DO NOT EDIT.

package report_definitions

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

// NewGetResourceInfoByReportDefinitionParams creates a new GetResourceInfoByReportDefinitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceInfoByReportDefinitionParams() *GetResourceInfoByReportDefinitionParams {
	return &GetResourceInfoByReportDefinitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceInfoByReportDefinitionParamsWithTimeout creates a new GetResourceInfoByReportDefinitionParams object
// with the ability to set a timeout on a request.
func NewGetResourceInfoByReportDefinitionParamsWithTimeout(timeout time.Duration) *GetResourceInfoByReportDefinitionParams {
	return &GetResourceInfoByReportDefinitionParams{
		timeout: timeout,
	}
}

// NewGetResourceInfoByReportDefinitionParamsWithContext creates a new GetResourceInfoByReportDefinitionParams object
// with the ability to set a context for a request.
func NewGetResourceInfoByReportDefinitionParamsWithContext(ctx context.Context) *GetResourceInfoByReportDefinitionParams {
	return &GetResourceInfoByReportDefinitionParams{
		Context: ctx,
	}
}

// NewGetResourceInfoByReportDefinitionParamsWithHTTPClient creates a new GetResourceInfoByReportDefinitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceInfoByReportDefinitionParamsWithHTTPClient(client *http.Client) *GetResourceInfoByReportDefinitionParams {
	return &GetResourceInfoByReportDefinitionParams{
		HTTPClient: client,
	}
}

/* GetResourceInfoByReportDefinitionParams contains all the parameters to send to the API endpoint
   for the get resource info by report definition operation.

   Typically these are written to a http.Request.
*/
type GetResourceInfoByReportDefinitionParams struct {

	/* OrganizationID.

	   Valid Organization Id
	*/
	OrganizationID *string

	/* ReportDefinitionName.

	   Name of the Report definition to retrieve
	*/
	ReportDefinitionName string

	/* ReportMimeType.

	     The format for which the report definition is required. By default the value will be CSV.
	Valid Values:
	- application/xml
	- text/csv

	*/
	ReportMimeType *string

	/* SubscriptionType.

	     The subscription type for which report definition is required. By default the type will be CUSTOM.
	Valid Values:
	- CLASSIC
	- CUSTOM
	- STANDARD

	*/
	SubscriptionType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource info by report definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceInfoByReportDefinitionParams) WithDefaults() *GetResourceInfoByReportDefinitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource info by report definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceInfoByReportDefinitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithTimeout(timeout time.Duration) *GetResourceInfoByReportDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithContext(ctx context.Context) *GetResourceInfoByReportDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithHTTPClient(client *http.Client) *GetResourceInfoByReportDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithOrganizationID(organizationID *string) *GetResourceInfoByReportDefinitionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithReportDefinitionName adds the reportDefinitionName to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithReportDefinitionName(reportDefinitionName string) *GetResourceInfoByReportDefinitionParams {
	o.SetReportDefinitionName(reportDefinitionName)
	return o
}

// SetReportDefinitionName adds the reportDefinitionName to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetReportDefinitionName(reportDefinitionName string) {
	o.ReportDefinitionName = reportDefinitionName
}

// WithReportMimeType adds the reportMimeType to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithReportMimeType(reportMimeType *string) *GetResourceInfoByReportDefinitionParams {
	o.SetReportMimeType(reportMimeType)
	return o
}

// SetReportMimeType adds the reportMimeType to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetReportMimeType(reportMimeType *string) {
	o.ReportMimeType = reportMimeType
}

// WithSubscriptionType adds the subscriptionType to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) WithSubscriptionType(subscriptionType *string) *GetResourceInfoByReportDefinitionParams {
	o.SetSubscriptionType(subscriptionType)
	return o
}

// SetSubscriptionType adds the subscriptionType to the get resource info by report definition params
func (o *GetResourceInfoByReportDefinitionParams) SetSubscriptionType(subscriptionType *string) {
	o.SubscriptionType = subscriptionType
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceInfoByReportDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param reportDefinitionName
	if err := r.SetPathParam("reportDefinitionName", o.ReportDefinitionName); err != nil {
		return err
	}

	if o.ReportMimeType != nil {

		// query param reportMimeType
		var qrReportMimeType string

		if o.ReportMimeType != nil {
			qrReportMimeType = *o.ReportMimeType
		}
		qReportMimeType := qrReportMimeType
		if qReportMimeType != "" {

			if err := r.SetQueryParam("reportMimeType", qReportMimeType); err != nil {
				return err
			}
		}
	}

	if o.SubscriptionType != nil {

		// query param subscriptionType
		var qrSubscriptionType string

		if o.SubscriptionType != nil {
			qrSubscriptionType = *o.SubscriptionType
		}
		qSubscriptionType := qrSubscriptionType
		if qSubscriptionType != "" {

			if err := r.SetQueryParam("subscriptionType", qSubscriptionType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
