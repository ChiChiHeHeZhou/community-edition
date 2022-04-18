// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ApplyTKGConfigForAzureHandlerFunc turns a function with the right signature into a apply t k g config for azure handler
type ApplyTKGConfigForAzureHandlerFunc func(ApplyTKGConfigForAzureParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ApplyTKGConfigForAzureHandlerFunc) Handle(params ApplyTKGConfigForAzureParams) middleware.Responder {
	return fn(params)
}

// ApplyTKGConfigForAzureHandler interface for that can handle valid apply t k g config for azure params
type ApplyTKGConfigForAzureHandler interface {
	Handle(ApplyTKGConfigForAzureParams) middleware.Responder
}

// NewApplyTKGConfigForAzure creates a new http.Handler for the apply t k g config for azure operation
func NewApplyTKGConfigForAzure(ctx *middleware.Context, handler ApplyTKGConfigForAzureHandler) *ApplyTKGConfigForAzure {
	return &ApplyTKGConfigForAzure{Context: ctx, Handler: handler}
}

/*ApplyTKGConfigForAzure swagger:route POST /api/provider/azure/tkgconfig azure applyTKGConfigForAzure

Apply the changes to TKG configuration file for Azure"

*/
type ApplyTKGConfigForAzure struct {
	Context *middleware.Context
	Handler ApplyTKGConfigForAzureHandler
}

func (o *ApplyTKGConfigForAzure) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewApplyTKGConfigForAzureParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
