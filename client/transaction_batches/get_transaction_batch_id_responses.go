// Code generated by go-swagger; DO NOT EDIT.

package transaction_batches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// GetTransactionBatchIDReader is a Reader for the GetTransactionBatchID structure.
type GetTransactionBatchIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionBatchIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionBatchIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransactionBatchIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTransactionBatchIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTransactionBatchIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTransactionBatchIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetTransactionBatchIDBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionBatchIDOK creates a GetTransactionBatchIDOK with default headers values
func NewGetTransactionBatchIDOK() *GetTransactionBatchIDOK {
	return &GetTransactionBatchIDOK{}
}

/*GetTransactionBatchIDOK handles this case with default header values.

OK
*/
type GetTransactionBatchIDOK struct {
	Payload *GetTransactionBatchIDOKBody
}

func (o *GetTransactionBatchIDOK) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batches/{id}][%d] getTransactionBatchIdOK  %+v", 200, o.Payload)
}

func (o *GetTransactionBatchIDOK) GetPayload() *GetTransactionBatchIDOKBody {
	return o.Payload
}

func (o *GetTransactionBatchIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchIDBadRequest creates a GetTransactionBatchIDBadRequest with default headers values
func NewGetTransactionBatchIDBadRequest() *GetTransactionBatchIDBadRequest {
	return &GetTransactionBatchIDBadRequest{}
}

/*GetTransactionBatchIDBadRequest handles this case with default header values.

Bad Request
*/
type GetTransactionBatchIDBadRequest struct {
	Payload *GetTransactionBatchIDBadRequestBody
}

func (o *GetTransactionBatchIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batches/{id}][%d] getTransactionBatchIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionBatchIDBadRequest) GetPayload() *GetTransactionBatchIDBadRequestBody {
	return o.Payload
}

func (o *GetTransactionBatchIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchIDUnauthorized creates a GetTransactionBatchIDUnauthorized with default headers values
func NewGetTransactionBatchIDUnauthorized() *GetTransactionBatchIDUnauthorized {
	return &GetTransactionBatchIDUnauthorized{}
}

/*GetTransactionBatchIDUnauthorized handles this case with default header values.

Not Authorized
*/
type GetTransactionBatchIDUnauthorized struct {
	Payload *GetTransactionBatchIDUnauthorizedBody
}

func (o *GetTransactionBatchIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batches/{id}][%d] getTransactionBatchIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransactionBatchIDUnauthorized) GetPayload() *GetTransactionBatchIDUnauthorizedBody {
	return o.Payload
}

func (o *GetTransactionBatchIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchIDUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchIDForbidden creates a GetTransactionBatchIDForbidden with default headers values
func NewGetTransactionBatchIDForbidden() *GetTransactionBatchIDForbidden {
	return &GetTransactionBatchIDForbidden{}
}

/*GetTransactionBatchIDForbidden handles this case with default header values.

No Authenticated
*/
type GetTransactionBatchIDForbidden struct {
	Payload *GetTransactionBatchIDForbiddenBody
}

func (o *GetTransactionBatchIDForbidden) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batches/{id}][%d] getTransactionBatchIdForbidden  %+v", 403, o.Payload)
}

func (o *GetTransactionBatchIDForbidden) GetPayload() *GetTransactionBatchIDForbiddenBody {
	return o.Payload
}

func (o *GetTransactionBatchIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchIDNotFound creates a GetTransactionBatchIDNotFound with default headers values
func NewGetTransactionBatchIDNotFound() *GetTransactionBatchIDNotFound {
	return &GetTransactionBatchIDNotFound{}
}

/*GetTransactionBatchIDNotFound handles this case with default header values.

No Reports Found
*/
type GetTransactionBatchIDNotFound struct {
	Payload *GetTransactionBatchIDNotFoundBody
}

func (o *GetTransactionBatchIDNotFound) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batches/{id}][%d] getTransactionBatchIdNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionBatchIDNotFound) GetPayload() *GetTransactionBatchIDNotFoundBody {
	return o.Payload
}

func (o *GetTransactionBatchIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchIDBadGateway creates a GetTransactionBatchIDBadGateway with default headers values
func NewGetTransactionBatchIDBadGateway() *GetTransactionBatchIDBadGateway {
	return &GetTransactionBatchIDBadGateway{}
}

/*GetTransactionBatchIDBadGateway handles this case with default header values.

Bad Gateway
*/
type GetTransactionBatchIDBadGateway struct {
	Payload *GetTransactionBatchIDBadGatewayBody
}

func (o *GetTransactionBatchIDBadGateway) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batches/{id}][%d] getTransactionBatchIdBadGateway  %+v", 502, o.Payload)
}

func (o *GetTransactionBatchIDBadGateway) GetPayload() *GetTransactionBatchIDBadGatewayBody {
	return o.Payload
}

func (o *GetTransactionBatchIDBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchIDBadGatewayBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTransactionBatchIDBadGatewayBody ptsV1TransactionBatchesGet502Response
swagger:model GetTransactionBatchIDBadGatewayBody
*/
type GetTransactionBatchIDBadGatewayBody struct {

	// error information
	ErrorInformation *GetTransactionBatchIDBadGatewayBodyErrorInformation `json:"errorInformation,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this get transaction batch ID bad gateway body
func (o *GetTransactionBatchIDBadGatewayBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDBadGatewayBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchIdBadGateway" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDBadGatewayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDBadGatewayBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDBadGatewayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDBadGatewayBodyErrorInformation get transaction batch ID bad gateway body error information
swagger:model GetTransactionBatchIDBadGatewayBodyErrorInformation
*/
type GetTransactionBatchIDBadGatewayBodyErrorInformation struct {

	// The detailed message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of status
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch ID bad gateway body error information
func (o *GetTransactionBatchIDBadGatewayBodyErrorInformation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDBadGatewayBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDBadGatewayBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDBadGatewayBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDBadRequestBody ptsV1TransactionBatchesGet400Response
swagger:model GetTransactionBatchIDBadRequestBody
*/
type GetTransactionBatchIDBadRequestBody struct {

	// error information
	ErrorInformation *GetTransactionBatchIDBadRequestBodyErrorInformation `json:"errorInformation,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this get transaction batch ID bad request body
func (o *GetTransactionBatchIDBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDBadRequestBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchIdBadRequest" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDBadRequestBodyErrorInformation get transaction batch ID bad request body error information
swagger:model GetTransactionBatchIDBadRequestBodyErrorInformation
*/
type GetTransactionBatchIDBadRequestBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch ID bad request body error information
func (o *GetTransactionBatchIDBadRequestBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDBadRequestBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTransactionBatchIdBadRequest" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDBadRequestBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDBadRequestBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDBadRequestBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0 get transaction batch ID bad request body error information details items0
swagger:model GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch ID bad request body error information details items0
func (o *GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDBadRequestBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDForbiddenBody ptsV1TransactionBatchesGet403Response
swagger:model GetTransactionBatchIDForbiddenBody
*/
type GetTransactionBatchIDForbiddenBody struct {

	// error information
	ErrorInformation *GetTransactionBatchIDForbiddenBodyErrorInformation `json:"errorInformation,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this get transaction batch ID forbidden body
func (o *GetTransactionBatchIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDForbiddenBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchIdForbidden" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDForbiddenBodyErrorInformation get transaction batch ID forbidden body error information
swagger:model GetTransactionBatchIDForbiddenBodyErrorInformation
*/
type GetTransactionBatchIDForbiddenBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch ID forbidden body error information
func (o *GetTransactionBatchIDForbiddenBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDForbiddenBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTransactionBatchIdForbidden" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDForbiddenBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDForbiddenBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDForbiddenBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0 get transaction batch ID forbidden body error information details items0
swagger:model GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch ID forbidden body error information details items0
func (o *GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDForbiddenBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDNotFoundBody ptsV1TransactionBatchesGet404Response
swagger:model GetTransactionBatchIDNotFoundBody
*/
type GetTransactionBatchIDNotFoundBody struct {

	// error information
	ErrorInformation *GetTransactionBatchIDNotFoundBodyErrorInformation `json:"errorInformation,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this get transaction batch ID not found body
func (o *GetTransactionBatchIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDNotFoundBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchIdNotFound" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDNotFoundBodyErrorInformation get transaction batch ID not found body error information
swagger:model GetTransactionBatchIDNotFoundBodyErrorInformation
*/
type GetTransactionBatchIDNotFoundBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch ID not found body error information
func (o *GetTransactionBatchIDNotFoundBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDNotFoundBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTransactionBatchIdNotFound" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDNotFoundBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDNotFoundBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDNotFoundBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0 get transaction batch ID not found body error information details items0
swagger:model GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch ID not found body error information details items0
func (o *GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDNotFoundBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDOKBody ptsV1TransactionBatchesIdGet200Response
swagger:model GetTransactionBatchIDOKBody
*/
type GetTransactionBatchIDOKBody struct {

	// links
	Links *GetTransactionBatchIDOKBodyLinks `json:"_links,omitempty"`

	// Number of transactions accepted.
	AcceptedTransactionCount int64 `json:"acceptedTransactionCount,omitempty"`

	// The date when the batch template processing completed.
	CompletionDate string `json:"completionDate,omitempty"`

	// Unique identifier assigned to the batch file.
	// Max Length: 8
	// Min Length: 1
	// Pattern: ^[a-zA-Z0-9_+-]*$
	ID string `json:"id,omitempty"`

	// Number of transactions rejected.
	RejectedTransactionCount string `json:"rejectedTransactionCount,omitempty"`

	// The status of you batch template processing.
	Status string `json:"status,omitempty"`

	// Number of transactions in the transaction.
	TransactionCount int64 `json:"transactionCount,omitempty"`

	// Date when the batch template was update.
	UploadDate string `json:"uploadDate,omitempty"`
}

// Validate validates this get transaction batch ID o k body
func (o *GetTransactionBatchIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDOKBody) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchIdOK" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *GetTransactionBatchIDOKBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MinLength("getTransactionBatchIdOK"+"."+"id", "body", string(o.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("getTransactionBatchIdOK"+"."+"id", "body", string(o.ID), 8); err != nil {
		return err
	}

	if err := validate.Pattern("getTransactionBatchIdOK"+"."+"id", "body", string(o.ID), `^[a-zA-Z0-9_+-]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDOKBodyLinks get transaction batch ID o k body links
swagger:model GetTransactionBatchIDOKBodyLinks
*/
type GetTransactionBatchIDOKBodyLinks struct {

	// transactions
	Transactions []*GetTransactionBatchIDOKBodyLinksTransactionsItems0 `json:"transactions"`
}

// Validate validates this get transaction batch ID o k body links
func (o *GetTransactionBatchIDOKBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDOKBodyLinks) validateTransactions(formats strfmt.Registry) error {

	if swag.IsZero(o.Transactions) { // not required
		return nil
	}

	for i := 0; i < len(o.Transactions); i++ {
		if swag.IsZero(o.Transactions[i]) { // not required
			continue
		}

		if o.Transactions[i] != nil {
			if err := o.Transactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTransactionBatchIdOK" + "." + "_links" + "." + "transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDOKBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDOKBodyLinks) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDOKBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDOKBodyLinksTransactionsItems0 get transaction batch ID o k body links transactions items0
swagger:model GetTransactionBatchIDOKBodyLinksTransactionsItems0
*/
type GetTransactionBatchIDOKBodyLinksTransactionsItems0 struct {

	// Self link for this request
	Href string `json:"href,omitempty"`

	// method
	Method string `json:"method,omitempty"`
}

// Validate validates this get transaction batch ID o k body links transactions items0
func (o *GetTransactionBatchIDOKBodyLinksTransactionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDOKBodyLinksTransactionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDOKBodyLinksTransactionsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDOKBodyLinksTransactionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDUnauthorizedBody ptsV1TransactionBatchesGet401Response
swagger:model GetTransactionBatchIDUnauthorizedBody
*/
type GetTransactionBatchIDUnauthorizedBody struct {

	// error information
	ErrorInformation *GetTransactionBatchIDUnauthorizedBodyErrorInformation `json:"errorInformation,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this get transaction batch ID unauthorized body
func (o *GetTransactionBatchIDUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDUnauthorizedBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchIdUnauthorized" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDUnauthorizedBodyErrorInformation get transaction batch ID unauthorized body error information
swagger:model GetTransactionBatchIDUnauthorizedBodyErrorInformation
*/
type GetTransactionBatchIDUnauthorizedBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch ID unauthorized body error information
func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getTransactionBatchIdUnauthorized" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDUnauthorizedBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0 get transaction batch ID unauthorized body error information details items0
swagger:model GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch ID unauthorized body error information details items0
func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchIDUnauthorizedBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}