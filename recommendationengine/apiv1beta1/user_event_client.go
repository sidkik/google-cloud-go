// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recommendationengine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	recommendationenginepb "cloud.google.com/go/recommendationengine/apiv1beta1/recommendationenginepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newUserEventClientHook clientHook

// UserEventCallOptions contains the retry settings for each method of UserEventClient.
type UserEventCallOptions struct {
	WriteUserEvent   []gax.CallOption
	CollectUserEvent []gax.CallOption
	ListUserEvents   []gax.CallOption
	PurgeUserEvents  []gax.CallOption
	ImportUserEvents []gax.CallOption
}

func defaultUserEventGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recommendationengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("recommendationengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("recommendationengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://recommendationengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultUserEventCallOptions() *UserEventCallOptions {
	return &UserEventCallOptions{
		WriteUserEvent: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CollectUserEvent: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PurgeUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ImportUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultUserEventRESTCallOptions() *UserEventCallOptions {
	return &UserEventCallOptions{
		WriteUserEvent: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		CollectUserEvent: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		ListUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		PurgeUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
		ImportUserEvents: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusGatewayTimeout)
			}),
		},
	}
}

// internalUserEventClient is an interface that defines the methods available from Recommendations AI.
type internalUserEventClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	WriteUserEvent(context.Context, *recommendationenginepb.WriteUserEventRequest, ...gax.CallOption) (*recommendationenginepb.UserEvent, error)
	CollectUserEvent(context.Context, *recommendationenginepb.CollectUserEventRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	ListUserEvents(context.Context, *recommendationenginepb.ListUserEventsRequest, ...gax.CallOption) *UserEventIterator
	PurgeUserEvents(context.Context, *recommendationenginepb.PurgeUserEventsRequest, ...gax.CallOption) (*PurgeUserEventsOperation, error)
	PurgeUserEventsOperation(name string) *PurgeUserEventsOperation
	ImportUserEvents(context.Context, *recommendationenginepb.ImportUserEventsRequest, ...gax.CallOption) (*ImportUserEventsOperation, error)
	ImportUserEventsOperation(name string) *ImportUserEventsOperation
}

// UserEventClient is a client for interacting with Recommendations AI.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for ingesting end user actions on the customer website.
type UserEventClient struct {
	// The internal transport-dependent client.
	internalClient internalUserEventClient

	// The call options for this service.
	CallOptions *UserEventCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *UserEventClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *UserEventClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *UserEventClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// WriteUserEvent writes a single user event.
func (c *UserEventClient) WriteUserEvent(ctx context.Context, req *recommendationenginepb.WriteUserEventRequest, opts ...gax.CallOption) (*recommendationenginepb.UserEvent, error) {
	return c.internalClient.WriteUserEvent(ctx, req, opts...)
}

// CollectUserEvent writes a single user event from the browser. This uses a GET request to
// due to browser restriction of POST-ing to a 3rd party domain.
//
// This method is used only by the Recommendations AI JavaScript pixel.
// Users should not call this method directly.
func (c *UserEventClient) CollectUserEvent(ctx context.Context, req *recommendationenginepb.CollectUserEventRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.CollectUserEvent(ctx, req, opts...)
}

// ListUserEvents gets a list of user events within a time range, with potential filtering.
func (c *UserEventClient) ListUserEvents(ctx context.Context, req *recommendationenginepb.ListUserEventsRequest, opts ...gax.CallOption) *UserEventIterator {
	return c.internalClient.ListUserEvents(ctx, req, opts...)
}

// PurgeUserEvents deletes permanently all user events specified by the filter provided.
// Depending on the number of events specified by the filter, this operation
// could take hours or days to complete. To test a filter, use the list
// command first.
func (c *UserEventClient) PurgeUserEvents(ctx context.Context, req *recommendationenginepb.PurgeUserEventsRequest, opts ...gax.CallOption) (*PurgeUserEventsOperation, error) {
	return c.internalClient.PurgeUserEvents(ctx, req, opts...)
}

// PurgeUserEventsOperation returns a new PurgeUserEventsOperation from a given name.
// The name must be that of a previously created PurgeUserEventsOperation, possibly from a different process.
func (c *UserEventClient) PurgeUserEventsOperation(name string) *PurgeUserEventsOperation {
	return c.internalClient.PurgeUserEventsOperation(name)
}

// ImportUserEvents bulk import of User events. Request processing might be
// synchronous. Events that already exist are skipped.
// Use this method for backfilling historical user events.
//
// Operation.response is of type ImportResponse. Note that it is
// possible for a subset of the items to be successfully inserted.
// Operation.metadata is of type ImportMetadata.
func (c *UserEventClient) ImportUserEvents(ctx context.Context, req *recommendationenginepb.ImportUserEventsRequest, opts ...gax.CallOption) (*ImportUserEventsOperation, error) {
	return c.internalClient.ImportUserEvents(ctx, req, opts...)
}

// ImportUserEventsOperation returns a new ImportUserEventsOperation from a given name.
// The name must be that of a previously created ImportUserEventsOperation, possibly from a different process.
func (c *UserEventClient) ImportUserEventsOperation(name string) *ImportUserEventsOperation {
	return c.internalClient.ImportUserEventsOperation(name)
}

// userEventGRPCClient is a client for interacting with Recommendations AI over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type userEventGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing UserEventClient
	CallOptions **UserEventCallOptions

	// The gRPC API client.
	userEventClient recommendationenginepb.UserEventServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewUserEventClient creates a new user event service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for ingesting end user actions on the customer website.
func NewUserEventClient(ctx context.Context, opts ...option.ClientOption) (*UserEventClient, error) {
	clientOpts := defaultUserEventGRPCClientOptions()
	if newUserEventClientHook != nil {
		hookOpts, err := newUserEventClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := UserEventClient{CallOptions: defaultUserEventCallOptions()}

	c := &userEventGRPCClient{
		connPool:        connPool,
		userEventClient: recommendationenginepb.NewUserEventServiceClient(connPool),
		CallOptions:     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *userEventGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *userEventGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *userEventGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type userEventRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing UserEventClient
	CallOptions **UserEventCallOptions
}

// NewUserEventRESTClient creates a new user event service rest client.
//
// Service for ingesting end user actions on the customer website.
func NewUserEventRESTClient(ctx context.Context, opts ...option.ClientOption) (*UserEventClient, error) {
	clientOpts := append(defaultUserEventRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultUserEventRESTCallOptions()
	c := &userEventRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &UserEventClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultUserEventRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://recommendationengine.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://recommendationengine.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://recommendationengine.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://recommendationengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *userEventRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *userEventRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *userEventRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *userEventGRPCClient) WriteUserEvent(ctx context.Context, req *recommendationenginepb.WriteUserEventRequest, opts ...gax.CallOption) (*recommendationenginepb.UserEvent, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).WriteUserEvent[0:len((*c.CallOptions).WriteUserEvent):len((*c.CallOptions).WriteUserEvent)], opts...)
	var resp *recommendationenginepb.UserEvent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.WriteUserEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *userEventGRPCClient) CollectUserEvent(ctx context.Context, req *recommendationenginepb.CollectUserEventRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CollectUserEvent[0:len((*c.CallOptions).CollectUserEvent):len((*c.CallOptions).CollectUserEvent)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.CollectUserEvent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *userEventGRPCClient) ListUserEvents(ctx context.Context, req *recommendationenginepb.ListUserEventsRequest, opts ...gax.CallOption) *UserEventIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListUserEvents[0:len((*c.CallOptions).ListUserEvents):len((*c.CallOptions).ListUserEvents)], opts...)
	it := &UserEventIterator{}
	req = proto.Clone(req).(*recommendationenginepb.ListUserEventsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommendationenginepb.UserEvent, string, error) {
		resp := &recommendationenginepb.ListUserEventsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.userEventClient.ListUserEvents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetUserEvents(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *userEventGRPCClient) PurgeUserEvents(ctx context.Context, req *recommendationenginepb.PurgeUserEventsRequest, opts ...gax.CallOption) (*PurgeUserEventsOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PurgeUserEvents[0:len((*c.CallOptions).PurgeUserEvents):len((*c.CallOptions).PurgeUserEvents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.PurgeUserEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &PurgeUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *userEventGRPCClient) ImportUserEvents(ctx context.Context, req *recommendationenginepb.ImportUserEventsRequest, opts ...gax.CallOption) (*ImportUserEventsOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ImportUserEvents[0:len((*c.CallOptions).ImportUserEvents):len((*c.CallOptions).ImportUserEvents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userEventClient.ImportUserEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// WriteUserEvent writes a single user event.
func (c *userEventRESTClient) WriteUserEvent(ctx context.Context, req *recommendationenginepb.WriteUserEventRequest, opts ...gax.CallOption) (*recommendationenginepb.UserEvent, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetUserEvent()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/userEvents:write", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).WriteUserEvent[0:len((*c.CallOptions).WriteUserEvent):len((*c.CallOptions).WriteUserEvent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &recommendationenginepb.UserEvent{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// CollectUserEvent writes a single user event from the browser. This uses a GET request to
// due to browser restriction of POST-ing to a 3rd party domain.
//
// This method is used only by the Recommendations AI JavaScript pixel.
// Users should not call this method directly.
func (c *userEventRESTClient) CollectUserEvent(ctx context.Context, req *recommendationenginepb.CollectUserEventRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/userEvents:collect", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetEts() != 0 {
		params.Add("ets", fmt.Sprintf("%v", req.GetEts()))
	}
	if req.GetUri() != "" {
		params.Add("uri", fmt.Sprintf("%v", req.GetUri()))
	}
	params.Add("userEvent", fmt.Sprintf("%v", req.GetUserEvent()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CollectUserEvent[0:len((*c.CallOptions).CollectUserEvent):len((*c.CallOptions).CollectUserEvent)], opts...)
	resp := &httpbodypb.HttpBody{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		resp.Data = buf
		if headers := httpRsp.Header; len(headers["Content-Type"]) > 0 {
			resp.ContentType = headers["Content-Type"][0]
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListUserEvents gets a list of user events within a time range, with potential filtering.
func (c *userEventRESTClient) ListUserEvents(ctx context.Context, req *recommendationenginepb.ListUserEventsRequest, opts ...gax.CallOption) *UserEventIterator {
	it := &UserEventIterator{}
	req = proto.Clone(req).(*recommendationenginepb.ListUserEventsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommendationenginepb.UserEvent, string, error) {
		resp := &recommendationenginepb.ListUserEventsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1beta1/%v/userEvents", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetUserEvents(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// PurgeUserEvents deletes permanently all user events specified by the filter provided.
// Depending on the number of events specified by the filter, this operation
// could take hours or days to complete. To test a filter, use the list
// command first.
func (c *userEventRESTClient) PurgeUserEvents(ctx context.Context, req *recommendationenginepb.PurgeUserEventsRequest, opts ...gax.CallOption) (*PurgeUserEventsOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/userEvents:purge", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1beta1/%s", resp.GetName())
	return &PurgeUserEventsOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// ImportUserEvents bulk import of User events. Request processing might be
// synchronous. Events that already exist are skipped.
// Use this method for backfilling historical user events.
//
// Operation.response is of type ImportResponse. Note that it is
// possible for a subset of the items to be successfully inserted.
// Operation.metadata is of type ImportMetadata.
func (c *userEventRESTClient) ImportUserEvents(ctx context.Context, req *recommendationenginepb.ImportUserEventsRequest, opts ...gax.CallOption) (*ImportUserEventsOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/userEvents:import", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1beta1/%s", resp.GetName())
	return &ImportUserEventsOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// ImportUserEventsOperation returns a new ImportUserEventsOperation from a given name.
// The name must be that of a previously created ImportUserEventsOperation, possibly from a different process.
func (c *userEventGRPCClient) ImportUserEventsOperation(name string) *ImportUserEventsOperation {
	return &ImportUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// ImportUserEventsOperation returns a new ImportUserEventsOperation from a given name.
// The name must be that of a previously created ImportUserEventsOperation, possibly from a different process.
func (c *userEventRESTClient) ImportUserEventsOperation(name string) *ImportUserEventsOperation {
	override := fmt.Sprintf("/v1beta1/%s", name)
	return &ImportUserEventsOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// PurgeUserEventsOperation returns a new PurgeUserEventsOperation from a given name.
// The name must be that of a previously created PurgeUserEventsOperation, possibly from a different process.
func (c *userEventGRPCClient) PurgeUserEventsOperation(name string) *PurgeUserEventsOperation {
	return &PurgeUserEventsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// PurgeUserEventsOperation returns a new PurgeUserEventsOperation from a given name.
// The name must be that of a previously created PurgeUserEventsOperation, possibly from a different process.
func (c *userEventRESTClient) PurgeUserEventsOperation(name string) *PurgeUserEventsOperation {
	override := fmt.Sprintf("/v1beta1/%s", name)
	return &PurgeUserEventsOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
