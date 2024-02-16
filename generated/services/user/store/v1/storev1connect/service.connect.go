// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/user/store/v1/service.proto

package storev1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/Golerplate/contracts/generated/services/user/store/v1"
	connect_go "github.com/bufbuild/connect-go"
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
	// UserStoreServiceName is the fully-qualified name of the UserStoreService service.
	UserStoreServiceName = "services.user.store.v1.UserStoreService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserStoreServiceCreateUserProcedure is the fully-qualified name of the UserStoreService's
	// CreateUser RPC.
	UserStoreServiceCreateUserProcedure = "/services.user.store.v1.UserStoreService/CreateUser"
)

// UserStoreServiceClient is a client for the services.user.store.v1.UserStoreService service.
type UserStoreServiceClient interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
}

// NewUserStoreServiceClient constructs a client for the services.user.store.v1.UserStoreService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserStoreServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserStoreServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userStoreServiceClient{
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+UserStoreServiceCreateUserProcedure,
			opts...,
		),
	}
}

// userStoreServiceClient implements UserStoreServiceClient.
type userStoreServiceClient struct {
	createUser *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
}

// CreateUser calls services.user.store.v1.UserStoreService.CreateUser.
func (c *userStoreServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// UserStoreServiceHandler is an implementation of the services.user.store.v1.UserStoreService
// service.
type UserStoreServiceHandler interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
}

// NewUserStoreServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserStoreServiceHandler(svc UserStoreServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userStoreServiceCreateUserHandler := connect_go.NewUnaryHandler(
		UserStoreServiceCreateUserProcedure,
		svc.CreateUser,
		opts...,
	)
	return "/services.user.store.v1.UserStoreService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserStoreServiceCreateUserProcedure:
			userStoreServiceCreateUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserStoreServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserStoreServiceHandler struct{}

func (UnimplementedUserStoreServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.user.store.v1.UserStoreService.CreateUser is not implemented"))
}