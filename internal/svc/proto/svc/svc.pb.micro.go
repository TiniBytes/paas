// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/svc/svc.proto

package svc

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Svc service

func NewSvcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Svc service

type SvcService interface {
	AddSvc(ctx context.Context, in *SvcInfo, opts ...client.CallOption) (*Response, error)
	DeleteSvc(ctx context.Context, in *SvcID, opts ...client.CallOption) (*Response, error)
	UpdateSvc(ctx context.Context, in *SvcInfo, opts ...client.CallOption) (*Response, error)
	FindSvcByID(ctx context.Context, in *SvcID, opts ...client.CallOption) (*SvcInfo, error)
	FindAllSvc(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllSvc, error)
}

type svcService struct {
	c    client.Client
	name string
}

func NewSvcService(name string, c client.Client) SvcService {
	return &svcService{
		c:    c,
		name: name,
	}
}

func (c *svcService) AddSvc(ctx context.Context, in *SvcInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Svc.AddSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) DeleteSvc(ctx context.Context, in *SvcID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Svc.DeleteSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) UpdateSvc(ctx context.Context, in *SvcInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Svc.UpdateSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) FindSvcByID(ctx context.Context, in *SvcID, opts ...client.CallOption) (*SvcInfo, error) {
	req := c.c.NewRequest(c.name, "Svc.FindSvcByID", in)
	out := new(SvcInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) FindAllSvc(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllSvc, error) {
	req := c.c.NewRequest(c.name, "Svc.FindAllSvc", in)
	out := new(AllSvc)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Svc service

type SvcHandler interface {
	AddSvc(context.Context, *SvcInfo, *Response) error
	DeleteSvc(context.Context, *SvcID, *Response) error
	UpdateSvc(context.Context, *SvcInfo, *Response) error
	FindSvcByID(context.Context, *SvcID, *SvcInfo) error
	FindAllSvc(context.Context, *FindAll, *AllSvc) error
}

func RegisterSvcHandler(s server.Server, hdlr SvcHandler, opts ...server.HandlerOption) error {
	type svc interface {
		AddSvc(ctx context.Context, in *SvcInfo, out *Response) error
		DeleteSvc(ctx context.Context, in *SvcID, out *Response) error
		UpdateSvc(ctx context.Context, in *SvcInfo, out *Response) error
		FindSvcByID(ctx context.Context, in *SvcID, out *SvcInfo) error
		FindAllSvc(ctx context.Context, in *FindAll, out *AllSvc) error
	}
	type Svc struct {
		svc
	}
	h := &svcHandler{hdlr}
	return s.Handle(s.NewHandler(&Svc{h}, opts...))
}

type svcHandler struct {
	SvcHandler
}

func (h *svcHandler) AddSvc(ctx context.Context, in *SvcInfo, out *Response) error {
	return h.SvcHandler.AddSvc(ctx, in, out)
}

func (h *svcHandler) DeleteSvc(ctx context.Context, in *SvcID, out *Response) error {
	return h.SvcHandler.DeleteSvc(ctx, in, out)
}

func (h *svcHandler) UpdateSvc(ctx context.Context, in *SvcInfo, out *Response) error {
	return h.SvcHandler.UpdateSvc(ctx, in, out)
}

func (h *svcHandler) FindSvcByID(ctx context.Context, in *SvcID, out *SvcInfo) error {
	return h.SvcHandler.FindSvcByID(ctx, in, out)
}

func (h *svcHandler) FindAllSvc(ctx context.Context, in *FindAll, out *AllSvc) error {
	return h.SvcHandler.FindAllSvc(ctx, in, out)
}