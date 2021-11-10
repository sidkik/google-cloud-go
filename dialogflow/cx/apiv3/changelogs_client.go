// Copyright 2021 Google LLC
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

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newChangelogsClientHook clientHook

// ChangelogsCallOptions contains the retry settings for each method of ChangelogsClient.
type ChangelogsCallOptions struct {
	ListChangelogs []gax.CallOption
	GetChangelog   []gax.CallOption
}

func defaultChangelogsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultChangelogsCallOptions() *ChangelogsCallOptions {
	return &ChangelogsCallOptions{
		ListChangelogs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetChangelog: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalChangelogsClient is an interface that defines the methods availaible from Dialogflow API.
type internalChangelogsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListChangelogs(context.Context, *cxpb.ListChangelogsRequest, ...gax.CallOption) *ChangelogIterator
	GetChangelog(context.Context, *cxpb.GetChangelogRequest, ...gax.CallOption) (*cxpb.Changelog, error)
}

// ChangelogsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Changelogs.
type ChangelogsClient struct {
	// The internal transport-dependent client.
	internalClient internalChangelogsClient

	// The call options for this service.
	CallOptions *ChangelogsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ChangelogsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ChangelogsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ChangelogsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListChangelogs returns the list of Changelogs.
func (c *ChangelogsClient) ListChangelogs(ctx context.Context, req *cxpb.ListChangelogsRequest, opts ...gax.CallOption) *ChangelogIterator {
	return c.internalClient.ListChangelogs(ctx, req, opts...)
}

// GetChangelog retrieves the specified Changelog.
func (c *ChangelogsClient) GetChangelog(ctx context.Context, req *cxpb.GetChangelogRequest, opts ...gax.CallOption) (*cxpb.Changelog, error) {
	return c.internalClient.GetChangelog(ctx, req, opts...)
}

// changelogsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type changelogsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ChangelogsClient
	CallOptions **ChangelogsCallOptions

	// The gRPC API client.
	changelogsClient cxpb.ChangelogsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewChangelogsClient creates a new changelogs client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Changelogs.
func NewChangelogsClient(ctx context.Context, opts ...option.ClientOption) (*ChangelogsClient, error) {
	clientOpts := defaultChangelogsGRPCClientOptions()
	if newChangelogsClientHook != nil {
		hookOpts, err := newChangelogsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ChangelogsClient{CallOptions: defaultChangelogsCallOptions()}

	c := &changelogsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		changelogsClient: cxpb.NewChangelogsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *changelogsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *changelogsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *changelogsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *changelogsGRPCClient) ListChangelogs(ctx context.Context, req *cxpb.ListChangelogsRequest, opts ...gax.CallOption) *ChangelogIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListChangelogs[0:len((*c.CallOptions).ListChangelogs):len((*c.CallOptions).ListChangelogs)], opts...)
	it := &ChangelogIterator{}
	req = proto.Clone(req).(*cxpb.ListChangelogsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.Changelog, string, error) {
		resp := &cxpb.ListChangelogsResponse{}
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
			resp, err = c.changelogsClient.ListChangelogs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetChangelogs(), resp.GetNextPageToken(), nil
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

func (c *changelogsGRPCClient) GetChangelog(ctx context.Context, req *cxpb.GetChangelogRequest, opts ...gax.CallOption) (*cxpb.Changelog, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetChangelog[0:len((*c.CallOptions).GetChangelog):len((*c.CallOptions).GetChangelog)], opts...)
	var resp *cxpb.Changelog
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.changelogsClient.GetChangelog(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChangelogIterator manages a stream of *cxpb.Changelog.
type ChangelogIterator struct {
	items    []*cxpb.Changelog
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.Changelog, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ChangelogIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ChangelogIterator) Next() (*cxpb.Changelog, error) {
	var item *cxpb.Changelog
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ChangelogIterator) bufLen() int {
	return len(it.items)
}

func (it *ChangelogIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
