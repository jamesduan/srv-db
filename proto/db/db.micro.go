// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: db.proto

/*
Package go_micro_srv_db is a generated protocol buffer package.

It is generated from these files:
	db.proto

It has these top-level messages:
	Message
	DeleteUserRequest
	DeleteUserResponse
	GetUserRequest
	GetUserResponse
	UpdateUserRequest
	CreateUserRequest
	CreateUserResponse
	UpdateUserResponse
	Response
	StreamingRequest
	StreamingResponse
	Ping
	Pong
*/
package go_micro_srv_db

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

// Client API for DB service

type DBClient interface {
	GetUserById(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*DeleteUserResponse, error)
}

type dBClient struct {
	c           client.Client
	serviceName string
}

func NewDBClient(serviceName string, c client.Client) DBClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.db"
	}
	return &dBClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *dBClient) GetUserById(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.GetUserById", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.UpdateUser", in)
	out := new(UpdateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.CreateUser", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*DeleteUserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.DeleteUser", in)
	out := new(DeleteUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DB service

type DBHandler interface {
	GetUserById(context.Context, *GetUserRequest, *GetUserResponse) error
	UpdateUser(context.Context, *UpdateUserRequest, *UpdateUserResponse) error
	CreateUser(context.Context, *CreateUserRequest, *CreateUserResponse) error
	DeleteUser(context.Context, *DeleteUserRequest, *DeleteUserResponse) error
}

func RegisterDBHandler(s server.Server, hdlr DBHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DB{hdlr}, opts...))
}

type DB struct {
	DBHandler
}

func (h *DB) GetUserById(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.DBHandler.GetUserById(ctx, in, out)
}

func (h *DB) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error {
	return h.DBHandler.UpdateUser(ctx, in, out)
}

func (h *DB) CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error {
	return h.DBHandler.CreateUser(ctx, in, out)
}

func (h *DB) DeleteUser(ctx context.Context, in *DeleteUserRequest, out *DeleteUserResponse) error {
	return h.DBHandler.DeleteUser(ctx, in, out)
}
