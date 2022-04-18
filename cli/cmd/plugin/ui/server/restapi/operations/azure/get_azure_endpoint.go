// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAzureEndpointHandlerFunc turns a function with the right signature into a get azure endpoint handler
type GetAzureEndpointHandlerFunc func(GetAzureEndpointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAzureEndpointHandlerFunc) Handle(params GetAzureEndpointParams) middleware.Responder {
	return fn(params)
}

// GetAzureEndpointHandler interface for that can handle valid get azure endpoint params
type GetAzureEndpointHandler interface {
	Handle(GetAzureEndpointParams) middleware.Responder
}

// NewGetAzureEndpoint creates a new http.Handler for the get azure endpoint operation
func NewGetAzureEndpoint(ctx *middleware.Context, handler GetAzureEndpointHandler) *GetAzureEndpoint {
	return &GetAzureEndpoint{Context: ctx, Handler: handler}
}

/*GetAzureEndpoint swagger:route GET /api/provider/azure azure getAzureEndpoint

Retrieve azure account params from environment variables

*/
type GetAzureEndpoint struct {
	Context *middleware.Context
	Handler GetAzureEndpointHandler
}

func (o *GetAzureEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAzureEndpointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
