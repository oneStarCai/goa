// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen goa.design/goa/examples/basic/design -o
// $(GOPATH)/src/goa.design/goa/examples/basic

package calcsvc

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "calc" service client.
type Client struct {
	AddEndpoint    goa.Endpoint
	ConcatEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(add, concat goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:    add,
		ConcatEndpoint: concat,
	}
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Concat calls the "concat" endpoint of the "calc" service.
func (c *Client) Concat(ctx context.Context, p *ConcatPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.ConcatEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
