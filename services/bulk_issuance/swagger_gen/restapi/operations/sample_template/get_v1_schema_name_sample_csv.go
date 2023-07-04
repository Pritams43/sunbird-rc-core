// Code generated by go-swagger; DO NOT EDIT.

package sample_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"bulk_issuance/swagger_gen/models"
)

// GetV1SchemaNameSampleCsvHandlerFunc turns a function with the right signature into a get v1 schema name sample csv handler
type GetV1SchemaNameSampleCsvHandlerFunc func(GetV1SchemaNameSampleCsvParams, *models.JWTClaimBody) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1SchemaNameSampleCsvHandlerFunc) Handle(params GetV1SchemaNameSampleCsvParams, principal *models.JWTClaimBody) middleware.Responder {
	return fn(params, principal)
}

// GetV1SchemaNameSampleCsvHandler interface for that can handle valid get v1 schema name sample csv params
type GetV1SchemaNameSampleCsvHandler interface {
	Handle(GetV1SchemaNameSampleCsvParams, *models.JWTClaimBody) middleware.Responder
}

// NewGetV1SchemaNameSampleCsv creates a new http.Handler for the get v1 schema name sample csv operation
func NewGetV1SchemaNameSampleCsv(ctx *middleware.Context, handler GetV1SchemaNameSampleCsvHandler) *GetV1SchemaNameSampleCsv {
	return &GetV1SchemaNameSampleCsv{Context: ctx, Handler: handler}
}

/*
	GetV1SchemaNameSampleCsv swagger:route GET /v1/{schemaName}/sample-csv sampleTemplate getV1SchemaNameSampleCsv

get sample template
*/
type GetV1SchemaNameSampleCsv struct {
	Context *middleware.Context
	Handler GetV1SchemaNameSampleCsvHandler
}

func (o *GetV1SchemaNameSampleCsv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetV1SchemaNameSampleCsvParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.JWTClaimBody
	if uprinc != nil {
		principal = uprinc.(*models.JWTClaimBody) // this is really a models.JWTClaimBody, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
