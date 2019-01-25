// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "golang-foo": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/golang-starters/golang-rest-http/design
// --out=$(GOPATH)/src/github.com/golang-starters/golang-rest-http
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// ShowStatusContext provides the status show action context.
type ShowStatusContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowStatusContext parses the incoming request URL and body, performs validations and creates the
// context used by the status controller show action.
func NewShowStatusContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowStatusContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowStatusContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowStatusContext) OK(r *Status) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.status+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ServiceUnavailable sends a HTTP response with status code 503.
func (ctx *ShowStatusContext) ServiceUnavailable(r *Status) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.status+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 503, r)
}
