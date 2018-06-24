// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: telemetry.proto

/*
Package telemetry is a generated protocol buffer package.

It is generated from these files:
	telemetry.proto

It has these top-level messages:
	Empty
	AllVehiclesResponse
*/
package telemetry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Telemetry service

type TelemetryService interface {
	AllVehicles(ctx context.Context, in *Empty, opts ...client.CallOption) (*AllVehiclesResponse, error)
}

type telemetryService struct {
	c    client.Client
	name string
}

func NewTelemetryService(name string, c client.Client) TelemetryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "telemetry"
	}
	return &telemetryService{
		c:    c,
		name: name,
	}
}

func (c *telemetryService) AllVehicles(ctx context.Context, in *Empty, opts ...client.CallOption) (*AllVehiclesResponse, error) {
	req := c.c.NewRequest(c.name, "Telemetry.AllVehicles", in)
	out := new(AllVehiclesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Telemetry service

type TelemetryHandler interface {
	AllVehicles(context.Context, *Empty, *AllVehiclesResponse) error
}

func RegisterTelemetryHandler(s server.Server, hdlr TelemetryHandler, opts ...server.HandlerOption) {
	type telemetry interface {
		AllVehicles(ctx context.Context, in *Empty, out *AllVehiclesResponse) error
	}
	type Telemetry struct {
		telemetry
	}
	h := &telemetryHandler{hdlr}
	s.Handle(s.NewHandler(&Telemetry{h}, opts...))
}

type telemetryHandler struct {
	TelemetryHandler
}

func (h *telemetryHandler) AllVehicles(ctx context.Context, in *Empty, out *AllVehiclesResponse) error {
	return h.TelemetryHandler.AllVehicles(ctx, in, out)
}
