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

package aiplatform

import (
	"context"
	"fmt"
	"math"
	"net/url"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newIndexEndpointClientHook clientHook

// IndexEndpointCallOptions contains the retry settings for each method of IndexEndpointClient.
type IndexEndpointCallOptions struct {
	CreateIndexEndpoint []gax.CallOption
	GetIndexEndpoint    []gax.CallOption
	ListIndexEndpoints  []gax.CallOption
	UpdateIndexEndpoint []gax.CallOption
	DeleteIndexEndpoint []gax.CallOption
	DeployIndex         []gax.CallOption
	UndeployIndex       []gax.CallOption
	MutateDeployedIndex []gax.CallOption
	GetLocation         []gax.CallOption
	ListLocations       []gax.CallOption
	GetIamPolicy        []gax.CallOption
	SetIamPolicy        []gax.CallOption
	TestIamPermissions  []gax.CallOption
	CancelOperation     []gax.CallOption
	DeleteOperation     []gax.CallOption
	GetOperation        []gax.CallOption
	ListOperations      []gax.CallOption
	WaitOperation       []gax.CallOption
}

func defaultIndexEndpointGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("aiplatform.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIndexEndpointCallOptions() *IndexEndpointCallOptions {
	return &IndexEndpointCallOptions{
		CreateIndexEndpoint: []gax.CallOption{},
		GetIndexEndpoint:    []gax.CallOption{},
		ListIndexEndpoints:  []gax.CallOption{},
		UpdateIndexEndpoint: []gax.CallOption{},
		DeleteIndexEndpoint: []gax.CallOption{},
		DeployIndex:         []gax.CallOption{},
		UndeployIndex:       []gax.CallOption{},
		MutateDeployedIndex: []gax.CallOption{},
		GetLocation:         []gax.CallOption{},
		ListLocations:       []gax.CallOption{},
		GetIamPolicy:        []gax.CallOption{},
		SetIamPolicy:        []gax.CallOption{},
		TestIamPermissions:  []gax.CallOption{},
		CancelOperation:     []gax.CallOption{},
		DeleteOperation:     []gax.CallOption{},
		GetOperation:        []gax.CallOption{},
		ListOperations:      []gax.CallOption{},
		WaitOperation:       []gax.CallOption{},
	}
}

// internalIndexEndpointClient is an interface that defines the methods available from Vertex AI API.
type internalIndexEndpointClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateIndexEndpoint(context.Context, *aiplatformpb.CreateIndexEndpointRequest, ...gax.CallOption) (*CreateIndexEndpointOperation, error)
	CreateIndexEndpointOperation(name string) *CreateIndexEndpointOperation
	GetIndexEndpoint(context.Context, *aiplatformpb.GetIndexEndpointRequest, ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error)
	ListIndexEndpoints(context.Context, *aiplatformpb.ListIndexEndpointsRequest, ...gax.CallOption) *IndexEndpointIterator
	UpdateIndexEndpoint(context.Context, *aiplatformpb.UpdateIndexEndpointRequest, ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error)
	DeleteIndexEndpoint(context.Context, *aiplatformpb.DeleteIndexEndpointRequest, ...gax.CallOption) (*DeleteIndexEndpointOperation, error)
	DeleteIndexEndpointOperation(name string) *DeleteIndexEndpointOperation
	DeployIndex(context.Context, *aiplatformpb.DeployIndexRequest, ...gax.CallOption) (*DeployIndexOperation, error)
	DeployIndexOperation(name string) *DeployIndexOperation
	UndeployIndex(context.Context, *aiplatformpb.UndeployIndexRequest, ...gax.CallOption) (*UndeployIndexOperation, error)
	UndeployIndexOperation(name string) *UndeployIndexOperation
	MutateDeployedIndex(context.Context, *aiplatformpb.MutateDeployedIndexRequest, ...gax.CallOption) (*MutateDeployedIndexOperation, error)
	MutateDeployedIndexOperation(name string) *MutateDeployedIndexOperation
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
	WaitOperation(context.Context, *longrunningpb.WaitOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// IndexEndpointClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing Vertex AI’s IndexEndpoints.
type IndexEndpointClient struct {
	// The internal transport-dependent client.
	internalClient internalIndexEndpointClient

	// The call options for this service.
	CallOptions *IndexEndpointCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IndexEndpointClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IndexEndpointClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *IndexEndpointClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateIndexEndpoint creates an IndexEndpoint.
func (c *IndexEndpointClient) CreateIndexEndpoint(ctx context.Context, req *aiplatformpb.CreateIndexEndpointRequest, opts ...gax.CallOption) (*CreateIndexEndpointOperation, error) {
	return c.internalClient.CreateIndexEndpoint(ctx, req, opts...)
}

// CreateIndexEndpointOperation returns a new CreateIndexEndpointOperation from a given name.
// The name must be that of a previously created CreateIndexEndpointOperation, possibly from a different process.
func (c *IndexEndpointClient) CreateIndexEndpointOperation(name string) *CreateIndexEndpointOperation {
	return c.internalClient.CreateIndexEndpointOperation(name)
}

// GetIndexEndpoint gets an IndexEndpoint.
func (c *IndexEndpointClient) GetIndexEndpoint(ctx context.Context, req *aiplatformpb.GetIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	return c.internalClient.GetIndexEndpoint(ctx, req, opts...)
}

// ListIndexEndpoints lists IndexEndpoints in a Location.
func (c *IndexEndpointClient) ListIndexEndpoints(ctx context.Context, req *aiplatformpb.ListIndexEndpointsRequest, opts ...gax.CallOption) *IndexEndpointIterator {
	return c.internalClient.ListIndexEndpoints(ctx, req, opts...)
}

// UpdateIndexEndpoint updates an IndexEndpoint.
func (c *IndexEndpointClient) UpdateIndexEndpoint(ctx context.Context, req *aiplatformpb.UpdateIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	return c.internalClient.UpdateIndexEndpoint(ctx, req, opts...)
}

// DeleteIndexEndpoint deletes an IndexEndpoint.
func (c *IndexEndpointClient) DeleteIndexEndpoint(ctx context.Context, req *aiplatformpb.DeleteIndexEndpointRequest, opts ...gax.CallOption) (*DeleteIndexEndpointOperation, error) {
	return c.internalClient.DeleteIndexEndpoint(ctx, req, opts...)
}

// DeleteIndexEndpointOperation returns a new DeleteIndexEndpointOperation from a given name.
// The name must be that of a previously created DeleteIndexEndpointOperation, possibly from a different process.
func (c *IndexEndpointClient) DeleteIndexEndpointOperation(name string) *DeleteIndexEndpointOperation {
	return c.internalClient.DeleteIndexEndpointOperation(name)
}

// DeployIndex deploys an Index into this IndexEndpoint, creating a DeployedIndex within
// it.
// Only non-empty Indexes can be deployed.
func (c *IndexEndpointClient) DeployIndex(ctx context.Context, req *aiplatformpb.DeployIndexRequest, opts ...gax.CallOption) (*DeployIndexOperation, error) {
	return c.internalClient.DeployIndex(ctx, req, opts...)
}

// DeployIndexOperation returns a new DeployIndexOperation from a given name.
// The name must be that of a previously created DeployIndexOperation, possibly from a different process.
func (c *IndexEndpointClient) DeployIndexOperation(name string) *DeployIndexOperation {
	return c.internalClient.DeployIndexOperation(name)
}

// UndeployIndex undeploys an Index from an IndexEndpoint, removing a DeployedIndex from it,
// and freeing all resources it’s using.
func (c *IndexEndpointClient) UndeployIndex(ctx context.Context, req *aiplatformpb.UndeployIndexRequest, opts ...gax.CallOption) (*UndeployIndexOperation, error) {
	return c.internalClient.UndeployIndex(ctx, req, opts...)
}

// UndeployIndexOperation returns a new UndeployIndexOperation from a given name.
// The name must be that of a previously created UndeployIndexOperation, possibly from a different process.
func (c *IndexEndpointClient) UndeployIndexOperation(name string) *UndeployIndexOperation {
	return c.internalClient.UndeployIndexOperation(name)
}

// MutateDeployedIndex update an existing DeployedIndex under an IndexEndpoint.
func (c *IndexEndpointClient) MutateDeployedIndex(ctx context.Context, req *aiplatformpb.MutateDeployedIndexRequest, opts ...gax.CallOption) (*MutateDeployedIndexOperation, error) {
	return c.internalClient.MutateDeployedIndex(ctx, req, opts...)
}

// MutateDeployedIndexOperation returns a new MutateDeployedIndexOperation from a given name.
// The name must be that of a previously created MutateDeployedIndexOperation, possibly from a different process.
func (c *IndexEndpointClient) MutateDeployedIndexOperation(name string) *MutateDeployedIndexOperation {
	return c.internalClient.MutateDeployedIndexOperation(name)
}

// GetLocation gets information about a location.
func (c *IndexEndpointClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *IndexEndpointClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *IndexEndpointClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *IndexEndpointClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *IndexEndpointClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *IndexEndpointClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *IndexEndpointClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *IndexEndpointClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *IndexEndpointClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// WaitOperation is a utility method from google.longrunning.Operations.
func (c *IndexEndpointClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.WaitOperation(ctx, req, opts...)
}

// indexEndpointGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type indexEndpointGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing IndexEndpointClient
	CallOptions **IndexEndpointCallOptions

	// The gRPC API client.
	indexEndpointClient aiplatformpb.IndexEndpointServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewIndexEndpointClient creates a new index endpoint service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing Vertex AI’s IndexEndpoints.
func NewIndexEndpointClient(ctx context.Context, opts ...option.ClientOption) (*IndexEndpointClient, error) {
	clientOpts := defaultIndexEndpointGRPCClientOptions()
	if newIndexEndpointClientHook != nil {
		hookOpts, err := newIndexEndpointClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IndexEndpointClient{CallOptions: defaultIndexEndpointCallOptions()}

	c := &indexEndpointGRPCClient{
		connPool:            connPool,
		indexEndpointClient: aiplatformpb.NewIndexEndpointServiceClient(connPool),
		CallOptions:         &client.CallOptions,
		operationsClient:    longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:     iampb.NewIAMPolicyClient(connPool),
		locationsClient:     locationpb.NewLocationsClient(connPool),
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
func (c *indexEndpointGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *indexEndpointGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *indexEndpointGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *indexEndpointGRPCClient) CreateIndexEndpoint(ctx context.Context, req *aiplatformpb.CreateIndexEndpointRequest, opts ...gax.CallOption) (*CreateIndexEndpointOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateIndexEndpoint[0:len((*c.CallOptions).CreateIndexEndpoint):len((*c.CallOptions).CreateIndexEndpoint)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.CreateIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) GetIndexEndpoint(ctx context.Context, req *aiplatformpb.GetIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIndexEndpoint[0:len((*c.CallOptions).GetIndexEndpoint):len((*c.CallOptions).GetIndexEndpoint)], opts...)
	var resp *aiplatformpb.IndexEndpoint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.GetIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) ListIndexEndpoints(ctx context.Context, req *aiplatformpb.ListIndexEndpointsRequest, opts ...gax.CallOption) *IndexEndpointIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListIndexEndpoints[0:len((*c.CallOptions).ListIndexEndpoints):len((*c.CallOptions).ListIndexEndpoints)], opts...)
	it := &IndexEndpointIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListIndexEndpointsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.IndexEndpoint, string, error) {
		resp := &aiplatformpb.ListIndexEndpointsResponse{}
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
			resp, err = c.indexEndpointClient.ListIndexEndpoints(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIndexEndpoints(), resp.GetNextPageToken(), nil
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

func (c *indexEndpointGRPCClient) UpdateIndexEndpoint(ctx context.Context, req *aiplatformpb.UpdateIndexEndpointRequest, opts ...gax.CallOption) (*aiplatformpb.IndexEndpoint, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint.name", url.QueryEscape(req.GetIndexEndpoint().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateIndexEndpoint[0:len((*c.CallOptions).UpdateIndexEndpoint):len((*c.CallOptions).UpdateIndexEndpoint)], opts...)
	var resp *aiplatformpb.IndexEndpoint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.UpdateIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) DeleteIndexEndpoint(ctx context.Context, req *aiplatformpb.DeleteIndexEndpointRequest, opts ...gax.CallOption) (*DeleteIndexEndpointOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteIndexEndpoint[0:len((*c.CallOptions).DeleteIndexEndpoint):len((*c.CallOptions).DeleteIndexEndpoint)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.DeleteIndexEndpoint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) DeployIndex(ctx context.Context, req *aiplatformpb.DeployIndexRequest, opts ...gax.CallOption) (*DeployIndexOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint", url.QueryEscape(req.GetIndexEndpoint()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeployIndex[0:len((*c.CallOptions).DeployIndex):len((*c.CallOptions).DeployIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.DeployIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) UndeployIndex(ctx context.Context, req *aiplatformpb.UndeployIndexRequest, opts ...gax.CallOption) (*UndeployIndexOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint", url.QueryEscape(req.GetIndexEndpoint()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UndeployIndex[0:len((*c.CallOptions).UndeployIndex):len((*c.CallOptions).UndeployIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.UndeployIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UndeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) MutateDeployedIndex(ctx context.Context, req *aiplatformpb.MutateDeployedIndexRequest, opts ...gax.CallOption) (*MutateDeployedIndexOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "index_endpoint", url.QueryEscape(req.GetIndexEndpoint()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateDeployedIndex[0:len((*c.CallOptions).MutateDeployedIndex):len((*c.CallOptions).MutateDeployedIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.indexEndpointClient.MutateDeployedIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &MutateDeployedIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *indexEndpointGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

func (c *indexEndpointGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *indexEndpointGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *indexEndpointGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *indexEndpointGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

func (c *indexEndpointGRPCClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).WaitOperation[0:len((*c.CallOptions).WaitOperation):len((*c.CallOptions).WaitOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.WaitOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateIndexEndpointOperation returns a new CreateIndexEndpointOperation from a given name.
// The name must be that of a previously created CreateIndexEndpointOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) CreateIndexEndpointOperation(name string) *CreateIndexEndpointOperation {
	return &CreateIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteIndexEndpointOperation returns a new DeleteIndexEndpointOperation from a given name.
// The name must be that of a previously created DeleteIndexEndpointOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) DeleteIndexEndpointOperation(name string) *DeleteIndexEndpointOperation {
	return &DeleteIndexEndpointOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeployIndexOperation returns a new DeployIndexOperation from a given name.
// The name must be that of a previously created DeployIndexOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) DeployIndexOperation(name string) *DeployIndexOperation {
	return &DeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// MutateDeployedIndexOperation returns a new MutateDeployedIndexOperation from a given name.
// The name must be that of a previously created MutateDeployedIndexOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) MutateDeployedIndexOperation(name string) *MutateDeployedIndexOperation {
	return &MutateDeployedIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UndeployIndexOperation returns a new UndeployIndexOperation from a given name.
// The name must be that of a previously created UndeployIndexOperation, possibly from a different process.
func (c *indexEndpointGRPCClient) UndeployIndexOperation(name string) *UndeployIndexOperation {
	return &UndeployIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
