// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteNamespacedNotebookHandlerFunc turns a function with the right signature into a delete namespaced notebook handler
type DeleteNamespacedNotebookHandlerFunc func(DeleteNamespacedNotebookParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNamespacedNotebookHandlerFunc) Handle(params DeleteNamespacedNotebookParams) middleware.Responder {
	return fn(params)
}

// DeleteNamespacedNotebookHandler interface for that can handle valid delete namespaced notebook params
type DeleteNamespacedNotebookHandler interface {
	Handle(DeleteNamespacedNotebookParams) middleware.Responder
}

// NewDeleteNamespacedNotebook creates a new http.Handler for the delete namespaced notebook operation
func NewDeleteNamespacedNotebook(ctx *middleware.Context, handler DeleteNamespacedNotebookHandler) *DeleteNamespacedNotebook {
	return &DeleteNamespacedNotebook{Context: ctx, Handler: handler}
}

/*DeleteNamespacedNotebook swagger:route DELETE /aide/v1/namespaces/{namespace}/notebooks/{notebook} deleteNamespacedNotebook

Delete a Notebook in the given Namespace

Delete Notebook.

*/
type DeleteNamespacedNotebook struct {
	Context *middleware.Context
	Handler DeleteNamespacedNotebookHandler
}

func (o *DeleteNamespacedNotebook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNamespacedNotebookParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}