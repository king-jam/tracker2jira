// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchConfigHandlerFunc turns a function with the right signature into a patch config handler
type PatchConfigHandlerFunc func(PatchConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchConfigHandlerFunc) Handle(params PatchConfigParams) middleware.Responder {
	return fn(params)
}

// PatchConfigHandler interface for that can handle valid patch config params
type PatchConfigHandler interface {
	Handle(PatchConfigParams) middleware.Responder
}

// NewPatchConfig creates a new http.Handler for the patch config operation
func NewPatchConfig(ctx *middleware.Context, handler PatchConfigHandler) *PatchConfig {
	return &PatchConfig{Context: ctx, Handler: handler}
}

/*PatchConfig swagger:route PATCH /config general patchConfig

Update the current running configuration

*/
type PatchConfig struct {
	Context *middleware.Context
	Handler PatchConfigHandler
}

func (o *PatchConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
