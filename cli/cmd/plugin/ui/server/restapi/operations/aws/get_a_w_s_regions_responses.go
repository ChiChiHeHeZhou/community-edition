// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// GetAWSRegionsOKCode is the HTTP code returned for type GetAWSRegionsOK
const GetAWSRegionsOKCode int = 200

/*GetAWSRegionsOK Successful retrieval of AWS regions

swagger:response getAWSRegionsOK
*/
type GetAWSRegionsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetAWSRegionsOK creates GetAWSRegionsOK with default headers values
func NewGetAWSRegionsOK() *GetAWSRegionsOK {

	return &GetAWSRegionsOK{}
}

// WithPayload adds the payload to the get a w s regions o k response
func (o *GetAWSRegionsOK) WithPayload(payload []string) *GetAWSRegionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s regions o k response
func (o *GetAWSRegionsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAWSRegionsBadRequestCode is the HTTP code returned for type GetAWSRegionsBadRequest
const GetAWSRegionsBadRequestCode int = 400

/*GetAWSRegionsBadRequest Bad request

swagger:response getAWSRegionsBadRequest
*/
type GetAWSRegionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSRegionsBadRequest creates GetAWSRegionsBadRequest with default headers values
func NewGetAWSRegionsBadRequest() *GetAWSRegionsBadRequest {

	return &GetAWSRegionsBadRequest{}
}

// WithPayload adds the payload to the get a w s regions bad request response
func (o *GetAWSRegionsBadRequest) WithPayload(payload *models.Error) *GetAWSRegionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s regions bad request response
func (o *GetAWSRegionsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSRegionsUnauthorizedCode is the HTTP code returned for type GetAWSRegionsUnauthorized
const GetAWSRegionsUnauthorizedCode int = 401

/*GetAWSRegionsUnauthorized Incorrect credentials

swagger:response getAWSRegionsUnauthorized
*/
type GetAWSRegionsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSRegionsUnauthorized creates GetAWSRegionsUnauthorized with default headers values
func NewGetAWSRegionsUnauthorized() *GetAWSRegionsUnauthorized {

	return &GetAWSRegionsUnauthorized{}
}

// WithPayload adds the payload to the get a w s regions unauthorized response
func (o *GetAWSRegionsUnauthorized) WithPayload(payload *models.Error) *GetAWSRegionsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s regions unauthorized response
func (o *GetAWSRegionsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAWSRegionsInternalServerErrorCode is the HTTP code returned for type GetAWSRegionsInternalServerError
const GetAWSRegionsInternalServerErrorCode int = 500

/*GetAWSRegionsInternalServerError Internal server error

swagger:response getAWSRegionsInternalServerError
*/
type GetAWSRegionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAWSRegionsInternalServerError creates GetAWSRegionsInternalServerError with default headers values
func NewGetAWSRegionsInternalServerError() *GetAWSRegionsInternalServerError {

	return &GetAWSRegionsInternalServerError{}
}

// WithPayload adds the payload to the get a w s regions internal server error response
func (o *GetAWSRegionsInternalServerError) WithPayload(payload *models.Error) *GetAWSRegionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get a w s regions internal server error response
func (o *GetAWSRegionsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAWSRegionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
