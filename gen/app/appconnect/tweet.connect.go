// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: app/tweet.proto

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
	// TweetServiceName is the fully-qualified name of the TweetService service.
	TweetServiceName = "proto.TweetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TweetServiceCreateTweetProcedure is the fully-qualified name of the TweetService's CreateTweet
	// RPC.
	TweetServiceCreateTweetProcedure = "/proto.TweetService/CreateTweet"
	// TweetServiceListTweetsProcedure is the fully-qualified name of the TweetService's ListTweets RPC.
	TweetServiceListTweetsProcedure = "/proto.TweetService/ListTweets"
	// TweetServiceDeleteTweetProcedure is the fully-qualified name of the TweetService's DeleteTweet
	// RPC.
	TweetServiceDeleteTweetProcedure = "/proto.TweetService/DeleteTweet"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tweetServiceServiceDescriptor           = app.File_app_tweet_proto.Services().ByName("TweetService")
	tweetServiceCreateTweetMethodDescriptor = tweetServiceServiceDescriptor.Methods().ByName("CreateTweet")
	tweetServiceListTweetsMethodDescriptor  = tweetServiceServiceDescriptor.Methods().ByName("ListTweets")
	tweetServiceDeleteTweetMethodDescriptor = tweetServiceServiceDescriptor.Methods().ByName("DeleteTweet")
)

// TweetServiceClient is a client for the proto.TweetService service.
type TweetServiceClient interface {
	// Add a new tweet
	CreateTweet(context.Context, *connect.Request[app.CreateTweetRequest]) (*connect.Response[app.Tweet], error)
	// listing tweets of a particular author
	ListTweets(context.Context, *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error)
	// delete a tweet
	DeleteTweet(context.Context, *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewTweetServiceClient constructs a client for the proto.TweetService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTweetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TweetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tweetServiceClient{
		createTweet: connect.NewClient[app.CreateTweetRequest, app.Tweet](
			httpClient,
			baseURL+TweetServiceCreateTweetProcedure,
			connect.WithSchema(tweetServiceCreateTweetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listTweets: connect.NewClient[app.ListTweetsRequest, app.ListTweetsResponse](
			httpClient,
			baseURL+TweetServiceListTweetsProcedure,
			connect.WithSchema(tweetServiceListTweetsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTweet: connect.NewClient[app.DeleteTweetRequest, emptypb.Empty](
			httpClient,
			baseURL+TweetServiceDeleteTweetProcedure,
			connect.WithSchema(tweetServiceDeleteTweetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tweetServiceClient implements TweetServiceClient.
type tweetServiceClient struct {
	createTweet *connect.Client[app.CreateTweetRequest, app.Tweet]
	listTweets  *connect.Client[app.ListTweetsRequest, app.ListTweetsResponse]
	deleteTweet *connect.Client[app.DeleteTweetRequest, emptypb.Empty]
}

// CreateTweet calls proto.TweetService.CreateTweet.
func (c *tweetServiceClient) CreateTweet(ctx context.Context, req *connect.Request[app.CreateTweetRequest]) (*connect.Response[app.Tweet], error) {
	return c.createTweet.CallUnary(ctx, req)
}

// ListTweets calls proto.TweetService.ListTweets.
func (c *tweetServiceClient) ListTweets(ctx context.Context, req *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error) {
	return c.listTweets.CallUnary(ctx, req)
}

// DeleteTweet calls proto.TweetService.DeleteTweet.
func (c *tweetServiceClient) DeleteTweet(ctx context.Context, req *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteTweet.CallUnary(ctx, req)
}

// TweetServiceHandler is an implementation of the proto.TweetService service.
type TweetServiceHandler interface {
	// Add a new tweet
	CreateTweet(context.Context, *connect.Request[app.CreateTweetRequest]) (*connect.Response[app.Tweet], error)
	// listing tweets of a particular author
	ListTweets(context.Context, *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error)
	// delete a tweet
	DeleteTweet(context.Context, *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewTweetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTweetServiceHandler(svc TweetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tweetServiceCreateTweetHandler := connect.NewUnaryHandler(
		TweetServiceCreateTweetProcedure,
		svc.CreateTweet,
		connect.WithSchema(tweetServiceCreateTweetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tweetServiceListTweetsHandler := connect.NewUnaryHandler(
		TweetServiceListTweetsProcedure,
		svc.ListTweets,
		connect.WithSchema(tweetServiceListTweetsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tweetServiceDeleteTweetHandler := connect.NewUnaryHandler(
		TweetServiceDeleteTweetProcedure,
		svc.DeleteTweet,
		connect.WithSchema(tweetServiceDeleteTweetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.TweetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TweetServiceCreateTweetProcedure:
			tweetServiceCreateTweetHandler.ServeHTTP(w, r)
		case TweetServiceListTweetsProcedure:
			tweetServiceListTweetsHandler.ServeHTTP(w, r)
		case TweetServiceDeleteTweetProcedure:
			tweetServiceDeleteTweetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTweetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTweetServiceHandler struct{}

func (UnimplementedTweetServiceHandler) CreateTweet(context.Context, *connect.Request[app.CreateTweetRequest]) (*connect.Response[app.Tweet], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.TweetService.CreateTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) ListTweets(context.Context, *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.TweetService.ListTweets is not implemented"))
}

func (UnimplementedTweetServiceHandler) DeleteTweet(context.Context, *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.TweetService.DeleteTweet is not implemented"))
}
