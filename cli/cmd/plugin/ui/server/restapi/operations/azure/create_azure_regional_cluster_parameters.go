// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// NewCreateAzureRegionalClusterParams creates a new CreateAzureRegionalClusterParams object
// no default values defined in spec.
func NewCreateAzureRegionalClusterParams() CreateAzureRegionalClusterParams {

	return CreateAzureRegionalClusterParams{}
}

// CreateAzureRegionalClusterParams contains all the bound params for the create azure regional cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters createAzureRegionalCluster
type CreateAzureRegionalClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*parameters to create a regional cluster
	  Required: true
	  In: body
	*/
	Params *models.AzureRegionalClusterParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateAzureRegionalClusterParams() beforehand.
func (o *CreateAzureRegionalClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AzureRegionalClusterParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("params", "body"))
			} else {
				res = append(res, errors.NewParseError("params", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Params = &body
			}
		}
	} else {
		res = append(res, errors.Required("params", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
