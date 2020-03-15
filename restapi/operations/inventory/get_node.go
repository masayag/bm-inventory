// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNodeHandlerFunc turns a function with the right signature into a get node handler
type GetNodeHandlerFunc func(GetNodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNodeHandlerFunc) Handle(params GetNodeParams) middleware.Responder {
	return fn(params)
}

// GetNodeHandler interface for that can handle valid get node params
type GetNodeHandler interface {
	Handle(GetNodeParams) middleware.Responder
}

// NewGetNode creates a new http.Handler for the get node operation
func NewGetNode(ctx *middleware.Context, handler GetNodeHandler) *GetNode {
	return &GetNode{Context: ctx, Handler: handler}
}

/*GetNode swagger:route GET /nodes/{node_id} inventory getNode

Retrieve OpenShift bare metal node information

*/
type GetNode struct {
	Context *middleware.Context
	Handler GetNodeHandler
}

func (o *GetNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
