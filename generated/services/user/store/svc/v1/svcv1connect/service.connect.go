// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/user/store/svc/v1/service.proto

package svcv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserStoreSvcName is the fully-qualified name of the UserStoreSvc service.
	UserStoreSvcName = "services.user.store.svc.v1.UserStoreSvc"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserStoreSvcCreateUserProcedure is the fully-qualified name of the UserStoreSvc's CreateUser RPC.
	UserStoreSvcCreateUserProcedure = "/services.user.store.svc.v1.UserStoreSvc/CreateUser"
	// UserStoreSvcGetUserByIDProcedure is the fully-qualified name of the UserStoreSvc's GetUserByID
	// RPC.
	UserStoreSvcGetUserByIDProcedure = "/services.user.store.svc.v1.UserStoreSvc/GetUserByID"
	// UserStoreSvcGetUserByEmailProcedure is the fully-qualified name of the UserStoreSvc's
	// GetUserByEmail RPC.
	UserStoreSvcGetUserByEmailProcedure = "/services.user.store.svc.v1.UserStoreSvc/GetUserByEmail"
	// UserStoreSvcGetUserByUsernameProcedure is the fully-qualified name of the UserStoreSvc's
	// GetUserByUsername RPC.
	UserStoreSvcGetUserByUsernameProcedure = "/services.user.store.svc.v1.UserStoreSvc/GetUserByUsername"
)

// UserStoreSvcClient is a client for the services.user.store.svc.v1.UserStoreSvc service.
type UserStoreSvcClient interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	GetUserByID(context.Context, *connect_go.Request[v1.GetUserByIDRequest]) (*connect_go.Response[v1.GetUserByIDResponse], error)
	GetUserByEmail(context.Context, *connect_go.Request[v1.GetUserByEmailRequest]) (*connect_go.Response[v1.GetUserByEmailResponse], error)
	GetUserByUsername(context.Context, *connect_go.Request[v1.GetUserByUsernameRequest]) (*connect_go.Response[v1.GetUserByUsernameResponse], error)
}

// NewUserStoreSvcClient constructs a client for the services.user.store.svc.v1.UserStoreSvc
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserStoreSvcClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserStoreSvcClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userStoreSvcClient{
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+UserStoreSvcCreateUserProcedure,
			opts...,
		),
		getUserByID: connect_go.NewClient[v1.GetUserByIDRequest, v1.GetUserByIDResponse](
			httpClient,
			baseURL+UserStoreSvcGetUserByIDProcedure,
			opts...,
		),
		getUserByEmail: connect_go.NewClient[v1.GetUserByEmailRequest, v1.GetUserByEmailResponse](
			httpClient,
			baseURL+UserStoreSvcGetUserByEmailProcedure,
			opts...,
		),
		getUserByUsername: connect_go.NewClient[v1.GetUserByUsernameRequest, v1.GetUserByUsernameResponse](
			httpClient,
			baseURL+UserStoreSvcGetUserByUsernameProcedure,
			opts...,
		),
	}
}

// userStoreSvcClient implements UserStoreSvcClient.
type userStoreSvcClient struct {
	createUser        *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	getUserByID       *connect_go.Client[v1.GetUserByIDRequest, v1.GetUserByIDResponse]
	getUserByEmail    *connect_go.Client[v1.GetUserByEmailRequest, v1.GetUserByEmailResponse]
	getUserByUsername *connect_go.Client[v1.GetUserByUsernameRequest, v1.GetUserByUsernameResponse]
}

// CreateUser calls services.user.store.svc.v1.UserStoreSvc.CreateUser.
func (c *userStoreSvcClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// GetUserByID calls services.user.store.svc.v1.UserStoreSvc.GetUserByID.
func (c *userStoreSvcClient) GetUserByID(ctx context.Context, req *connect_go.Request[v1.GetUserByIDRequest]) (*connect_go.Response[v1.GetUserByIDResponse], error) {
	return c.getUserByID.CallUnary(ctx, req)
}

// GetUserByEmail calls services.user.store.svc.v1.UserStoreSvc.GetUserByEmail.
func (c *userStoreSvcClient) GetUserByEmail(ctx context.Context, req *connect_go.Request[v1.GetUserByEmailRequest]) (*connect_go.Response[v1.GetUserByEmailResponse], error) {
	return c.getUserByEmail.CallUnary(ctx, req)
}

// GetUserByUsername calls services.user.store.svc.v1.UserStoreSvc.GetUserByUsername.
func (c *userStoreSvcClient) GetUserByUsername(ctx context.Context, req *connect_go.Request[v1.GetUserByUsernameRequest]) (*connect_go.Response[v1.GetUserByUsernameResponse], error) {
	return c.getUserByUsername.CallUnary(ctx, req)
}

// UserStoreSvcHandler is an implementation of the services.user.store.svc.v1.UserStoreSvc service.
type UserStoreSvcHandler interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	GetUserByID(context.Context, *connect_go.Request[v1.GetUserByIDRequest]) (*connect_go.Response[v1.GetUserByIDResponse], error)
	GetUserByEmail(context.Context, *connect_go.Request[v1.GetUserByEmailRequest]) (*connect_go.Response[v1.GetUserByEmailResponse], error)
	GetUserByUsername(context.Context, *connect_go.Request[v1.GetUserByUsernameRequest]) (*connect_go.Response[v1.GetUserByUsernameResponse], error)
}

// NewUserStoreSvcHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserStoreSvcHandler(svc UserStoreSvcHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userStoreSvcCreateUserHandler := connect_go.NewUnaryHandler(
		UserStoreSvcCreateUserProcedure,
		svc.CreateUser,
		opts...,
	)
	userStoreSvcGetUserByIDHandler := connect_go.NewUnaryHandler(
		UserStoreSvcGetUserByIDProcedure,
		svc.GetUserByID,
		opts...,
	)
	userStoreSvcGetUserByEmailHandler := connect_go.NewUnaryHandler(
		UserStoreSvcGetUserByEmailProcedure,
		svc.GetUserByEmail,
		opts...,
	)
	userStoreSvcGetUserByUsernameHandler := connect_go.NewUnaryHandler(
		UserStoreSvcGetUserByUsernameProcedure,
		svc.GetUserByUsername,
		opts...,
	)
	return "/services.user.store.svc.v1.UserStoreSvc/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserStoreSvcCreateUserProcedure:
			userStoreSvcCreateUserHandler.ServeHTTP(w, r)
		case UserStoreSvcGetUserByIDProcedure:
			userStoreSvcGetUserByIDHandler.ServeHTTP(w, r)
		case UserStoreSvcGetUserByEmailProcedure:
			userStoreSvcGetUserByEmailHandler.ServeHTTP(w, r)
		case UserStoreSvcGetUserByUsernameProcedure:
			userStoreSvcGetUserByUsernameHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserStoreSvcHandler returns CodeUnimplemented from all methods.
type UnimplementedUserStoreSvcHandler struct{}

func (UnimplementedUserStoreSvcHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.user.store.svc.v1.UserStoreSvc.CreateUser is not implemented"))
}

func (UnimplementedUserStoreSvcHandler) GetUserByID(context.Context, *connect_go.Request[v1.GetUserByIDRequest]) (*connect_go.Response[v1.GetUserByIDResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.user.store.svc.v1.UserStoreSvc.GetUserByID is not implemented"))
}

func (UnimplementedUserStoreSvcHandler) GetUserByEmail(context.Context, *connect_go.Request[v1.GetUserByEmailRequest]) (*connect_go.Response[v1.GetUserByEmailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.user.store.svc.v1.UserStoreSvc.GetUserByEmail is not implemented"))
}

func (UnimplementedUserStoreSvcHandler) GetUserByUsername(context.Context, *connect_go.Request[v1.GetUserByUsernameRequest]) (*connect_go.Response[v1.GetUserByUsernameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.user.store.svc.v1.UserStoreSvc.GetUserByUsername is not implemented"))
}
