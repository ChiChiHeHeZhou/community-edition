// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVPCsHandlerFunc turns a function with the right signature into a get v p cs handler
type GetVPCsHandlerFunc func(GetVPCsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVPCsHandlerFunc) Handle(params GetVPCsParams) middleware.Responder {
	return fn(params)
}

// GetVPCsHandler interface for that can handle valid get v p cs params
type GetVPCsHandler interface {
	Handle(GetVPCsParams) middleware.Responder
}

// NewGetVPCs creates a new http.Handler for the get v p cs operation
func NewGetVPCs(ctx *middleware.Context, handler GetVPCsHandler) *GetVPCs {
	return &GetVPCs{Context: ctx, Handler: handler}
}

/*GetVPCs swagger:route GET /api/provider/aws/vpc aws getVPCs

Retrieve AWS VPCs

*/
type GetVPCs struct {
	Context *middleware.Context
	Handler GetVPCsHandler
}

func (o *GetVPCs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVPCsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
