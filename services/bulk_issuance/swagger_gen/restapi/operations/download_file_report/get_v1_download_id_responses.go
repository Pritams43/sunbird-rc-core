// Code generated by go-swagger; DO NOT EDIT.

package download_file_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bulk_issuance/swagger_gen/models"
)

// GetV1DownloadIDOKCode is the HTTP code returned for type GetV1DownloadIDOK
const GetV1DownloadIDOKCode int = 200

/*GetV1DownloadIDOK OK

swagger:response getV1DownloadIdOK
*/
type GetV1DownloadIDOK struct {
	/*

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload models.FileDownload `json:"body,omitempty"`
}

// NewGetV1DownloadIDOK creates GetV1DownloadIDOK with default headers values
func NewGetV1DownloadIDOK() *GetV1DownloadIDOK {

	return &GetV1DownloadIDOK{}
}

// WithContentDisposition adds the contentDisposition to the get v1 download Id o k response
func (o *GetV1DownloadIDOK) WithContentDisposition(contentDisposition string) *GetV1DownloadIDOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get v1 download Id o k response
func (o *GetV1DownloadIDOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the get v1 download Id o k response
func (o *GetV1DownloadIDOK) WithPayload(payload models.FileDownload) *GetV1DownloadIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 download Id o k response
func (o *GetV1DownloadIDOK) SetPayload(payload models.FileDownload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1DownloadIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1DownloadIDForbiddenCode is the HTTP code returned for type GetV1DownloadIDForbidden
const GetV1DownloadIDForbiddenCode int = 403

/*GetV1DownloadIDForbidden Forbidden

swagger:response getV1DownloadIdForbidden
*/
type GetV1DownloadIDForbidden struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetV1DownloadIDForbidden creates GetV1DownloadIDForbidden with default headers values
func NewGetV1DownloadIDForbidden() *GetV1DownloadIDForbidden {

	return &GetV1DownloadIDForbidden{}
}

// WithPayload adds the payload to the get v1 download Id forbidden response
func (o *GetV1DownloadIDForbidden) WithPayload(payload string) *GetV1DownloadIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 download Id forbidden response
func (o *GetV1DownloadIDForbidden) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1DownloadIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1DownloadIDNotFoundCode is the HTTP code returned for type GetV1DownloadIDNotFound
const GetV1DownloadIDNotFoundCode int = 404

/*GetV1DownloadIDNotFound Not found

swagger:response getV1DownloadIdNotFound
*/
type GetV1DownloadIDNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetV1DownloadIDNotFound creates GetV1DownloadIDNotFound with default headers values
func NewGetV1DownloadIDNotFound() *GetV1DownloadIDNotFound {

	return &GetV1DownloadIDNotFound{}
}

// WithPayload adds the payload to the get v1 download Id not found response
func (o *GetV1DownloadIDNotFound) WithPayload(payload string) *GetV1DownloadIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 download Id not found response
func (o *GetV1DownloadIDNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1DownloadIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
