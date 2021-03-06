// Code generated by go-swagger; DO NOT EDIT.

package fish

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListFishesHandlerFunc turns a function with the right signature into a list fishes handler
type ListFishesHandlerFunc func(ListFishesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListFishesHandlerFunc) Handle(params ListFishesParams) middleware.Responder {
	return fn(params)
}

// ListFishesHandler interface for that can handle valid list fishes params
type ListFishesHandler interface {
	Handle(ListFishesParams) middleware.Responder
}

// NewListFishes creates a new http.Handler for the list fishes operation
func NewListFishes(ctx *middleware.Context, handler ListFishesHandler) *ListFishes {
	return &ListFishes{Context: ctx, Handler: handler}
}

/*ListFishes swagger:route GET /catalog/fishes fish listFishes

ListFishes list fishes API

*/
type ListFishes struct {
	Context *middleware.Context
	Handler ListFishesHandler
}

func (o *ListFishes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListFishesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
