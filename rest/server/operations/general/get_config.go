// Code generated by go-swagger; DO NOT EDIT.

package general

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetConfigHandlerFunc turns a function with the right signature into a get config handler
type GetConfigHandlerFunc func(GetConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigHandlerFunc) Handle(params GetConfigParams) middleware.Responder {
	return fn(params)
}

// GetConfigHandler interface for that can handle valid get config params
type GetConfigHandler interface {
	Handle(GetConfigParams) middleware.Responder
}

// NewGetConfig creates a new http.Handler for the get config operation
func NewGetConfig(ctx *middleware.Context, handler GetConfigHandler) *GetConfig {
	return &GetConfig{Context: ctx, Handler: handler}
}

/*GetConfig swagger:route GET /config general getConfig

Get the current running configuration

*/
type GetConfig struct {
	Context *middleware.Context
	Handler GetConfigHandler
}

func (o *GetConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
