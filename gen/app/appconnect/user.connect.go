// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: app/user.proto

package appconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	app "github.com/Sachin24704/go-backend/gen/app"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "proto.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceCreateUserProcedure is the fully-qualified name of the UserService's CreateUser RPC.
	UserServiceCreateUserProcedure = "/proto.UserService/CreateUser"
	// UserServiceDeletUserProcedure is the fully-qualified name of the UserService's DeletUser RPC.
	UserServiceDeletUserProcedure = "/proto.UserService/DeletUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor          = app.File_app_user_proto.Services().ByName("UserService")
	userServiceCreateUserMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("CreateUser")
	userServiceDeletUserMethodDescriptor  = userServiceServiceDescriptor.Methods().ByName("DeletUser")
)

// UserServiceClient is a client for the proto.UserService service.
type UserServiceClient interface {
	// to add a new user
	CreateUser(context.Context, *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error)
	// to delete an existing user
	DeletUser(context.Context, *connect.Request[app.DeletUserRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewUserServiceClient constructs a client for the proto.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		createUser: connect.NewClient[app.CreateUserRequest, app.User](
			httpClient,
			baseURL+UserServiceCreateUserProcedure,
			connect.WithSchema(userServiceCreateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deletUser: connect.NewClient[app.DeletUserRequest, emptypb.Empty](
			httpClient,
			baseURL+UserServiceDeletUserProcedure,
			connect.WithSchema(userServiceDeletUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	createUser *connect.Client[app.CreateUserRequest, app.User]
	deletUser  *connect.Client[app.DeletUserRequest, emptypb.Empty]
}

// CreateUser calls proto.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error) {
	return c.createUser.CallUnary(ctx, req)
}

// DeletUser calls proto.UserService.DeletUser.
func (c *userServiceClient) DeletUser(ctx context.Context, req *connect.Request[app.DeletUserRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deletUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the proto.UserService service.
type UserServiceHandler interface {
	// to add a new user
	CreateUser(context.Context, *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error)
	// to delete an existing user
	DeletUser(context.Context, *connect.Request[app.DeletUserRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceCreateUserHandler := connect.NewUnaryHandler(
		UserServiceCreateUserProcedure,
		svc.CreateUser,
		connect.WithSchema(userServiceCreateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeletUserHandler := connect.NewUnaryHandler(
		UserServiceDeletUserProcedure,
		svc.DeletUser,
		connect.WithSchema(userServiceDeletUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceCreateUserProcedure:
			userServiceCreateUserHandler.ServeHTTP(w, r)
		case UserServiceDeletUserProcedure:
			userServiceDeletUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeletUser(context.Context, *connect.Request[app.DeletUserRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.DeletUser is not implemented"))
}
