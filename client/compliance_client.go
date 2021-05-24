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

package client

import (
	"context"
	"fmt"
	"math"
	"net/url"

	"github.com/golang/protobuf/proto"
	genprotopb "github.com/googleapis/gapic-showcase/server/genproto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newComplianceClientHook clientHook

// ComplianceCallOptions contains the retry settings for each method of ComplianceClient.
type ComplianceCallOptions struct {
	RepeatDataBody                 []gax.CallOption
	RepeatDataBodyInfo             []gax.CallOption
	RepeatDataQuery                []gax.CallOption
	RepeatDataSimplePath           []gax.CallOption
	RepeatDataPathResource         []gax.CallOption
	RepeatDataPathTrailingResource []gax.CallOption
	ListLocations                  []gax.CallOption
	GetLocation                    []gax.CallOption
	SetIamPolicy                   []gax.CallOption
	GetIamPolicy                   []gax.CallOption
	TestIamPermissions             []gax.CallOption
	ListOperations                 []gax.CallOption
	GetOperation                   []gax.CallOption
	DeleteOperation                []gax.CallOption
	CancelOperation                []gax.CallOption
}

func defaultComplianceGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("localhost:7469"),
		internaloption.WithDefaultMTLSEndpoint("localhost:7469"),
		internaloption.WithDefaultAudience("https://localhost/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultComplianceCallOptions() *ComplianceCallOptions {
	return &ComplianceCallOptions{
		RepeatDataBody:                 []gax.CallOption{},
		RepeatDataBodyInfo:             []gax.CallOption{},
		RepeatDataQuery:                []gax.CallOption{},
		RepeatDataSimplePath:           []gax.CallOption{},
		RepeatDataPathResource:         []gax.CallOption{},
		RepeatDataPathTrailingResource: []gax.CallOption{},
		ListLocations:                  []gax.CallOption{},
		GetLocation:                    []gax.CallOption{},
		SetIamPolicy:                   []gax.CallOption{},
		GetIamPolicy:                   []gax.CallOption{},
		TestIamPermissions:             []gax.CallOption{},
		ListOperations:                 []gax.CallOption{},
		GetOperation:                   []gax.CallOption{},
		DeleteOperation:                []gax.CallOption{},
		CancelOperation:                []gax.CallOption{},
	}
}

// internalComplianceClient is an interface that defines the methods availaible from Client Libraries Showcase API.
type internalComplianceClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	RepeatDataBody(context.Context, *genprotopb.RepeatRequest, ...gax.CallOption) (*genprotopb.RepeatResponse, error)
	RepeatDataBodyInfo(context.Context, *genprotopb.RepeatRequest, ...gax.CallOption) (*genprotopb.RepeatResponse, error)
	RepeatDataQuery(context.Context, *genprotopb.RepeatRequest, ...gax.CallOption) (*genprotopb.RepeatResponse, error)
	RepeatDataSimplePath(context.Context, *genprotopb.RepeatRequest, ...gax.CallOption) (*genprotopb.RepeatResponse, error)
	RepeatDataPathResource(context.Context, *genprotopb.RepeatRequest, ...gax.CallOption) (*genprotopb.RepeatResponse, error)
	RepeatDataPathTrailingResource(context.Context, *genprotopb.RepeatRequest, ...gax.CallOption) (*genprotopb.RepeatResponse, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
}

// ComplianceClient is a client for interacting with Client Libraries Showcase API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service is used to test that GAPICs can transcode proto3 requests to
// REST format correctly for various types of HTTP annotations.
type ComplianceClient struct {
	// The internal transport-dependent client.
	internalClient internalComplianceClient

	// The call options for this service.
	CallOptions *ComplianceCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ComplianceClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ComplianceClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ComplianceClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// RepeatDataBody this method echoes the ComplianceData request. This method exercises
// sending the entire request object in the REST body.
func (c *ComplianceClient) RepeatDataBody(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	return c.internalClient.RepeatDataBody(ctx, req, opts...)
}

// RepeatDataBodyInfo this method echoes the ComplianceData request. This method exercises
// sending the a message-type field in the REST body. Per AIP-127, only
// top-level, non-repeated fields can be sent this way.
func (c *ComplianceClient) RepeatDataBodyInfo(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	return c.internalClient.RepeatDataBodyInfo(ctx, req, opts...)
}

// RepeatDataQuery this method echoes the ComplianceData request. This method exercises
// sending all request fields as query parameters.
func (c *ComplianceClient) RepeatDataQuery(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	return c.internalClient.RepeatDataQuery(ctx, req, opts...)
}

// RepeatDataSimplePath this method echoes the ComplianceData request. This method exercises
// sending some parameters as “simple” path variables (i.e., of the form
// “/bar/{foo}” rather than “/{foo=bar/*}”), and the rest as query parameters.
func (c *ComplianceClient) RepeatDataSimplePath(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	return c.internalClient.RepeatDataSimplePath(ctx, req, opts...)
}

// RepeatDataPathResource same as RepeatDataSimplePath, but with a path resource.
func (c *ComplianceClient) RepeatDataPathResource(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	return c.internalClient.RepeatDataPathResource(ctx, req, opts...)
}

// RepeatDataPathTrailingResource same as RepeatDataSimplePath, but with a trailing resource.
func (c *ComplianceClient) RepeatDataPathTrailingResource(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	return c.internalClient.RepeatDataPathTrailingResource(ctx, req, opts...)
}

// ListLocations is a utility method from google.cloud.location.Locations.
func (c *ComplianceClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetLocation is a utility method from google.cloud.location.Locations.
func (c *ComplianceClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// SetIamPolicy is a utility method from google.iam.v1.IAMPolicy.
func (c *ComplianceClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// GetIamPolicy is a utility method from google.iam.v1.IAMPolicy.
func (c *ComplianceClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions is a utility method from google.iam.v1.IAMPolicy.
func (c *ComplianceClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *ComplianceClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *ComplianceClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *ComplianceClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *ComplianceClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// complianceGRPCClient is a client for interacting with Client Libraries Showcase API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type complianceGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ComplianceClient
	CallOptions **ComplianceCallOptions

	// The gRPC API client.
	complianceClient genprotopb.ComplianceClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewComplianceClient creates a new compliance client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service is used to test that GAPICs can transcode proto3 requests to
// REST format correctly for various types of HTTP annotations.
func NewComplianceClient(ctx context.Context, opts ...option.ClientOption) (*ComplianceClient, error) {
	clientOpts := defaultComplianceGRPCClientOptions()
	if newComplianceClientHook != nil {
		hookOpts, err := newComplianceClientHook(ctx, clientHookParams{})
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
	client := ComplianceClient{CallOptions: defaultComplianceCallOptions()}

	c := &complianceGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		complianceClient: genprotopb.NewComplianceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:  iampb.NewIAMPolicyClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *complianceGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *complianceGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *complianceGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *complianceGRPCClient) RepeatDataBody(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).RepeatDataBody[0:len((*c.CallOptions).RepeatDataBody):len((*c.CallOptions).RepeatDataBody)], opts...)
	var resp *genprotopb.RepeatResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.complianceClient.RepeatDataBody(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *complianceGRPCClient) RepeatDataBodyInfo(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).RepeatDataBodyInfo[0:len((*c.CallOptions).RepeatDataBodyInfo):len((*c.CallOptions).RepeatDataBodyInfo)], opts...)
	var resp *genprotopb.RepeatResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.complianceClient.RepeatDataBodyInfo(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *complianceGRPCClient) RepeatDataQuery(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).RepeatDataQuery[0:len((*c.CallOptions).RepeatDataQuery):len((*c.CallOptions).RepeatDataQuery)], opts...)
	var resp *genprotopb.RepeatResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.complianceClient.RepeatDataQuery(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *complianceGRPCClient) RepeatDataSimplePath(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v&%s=%v", "info.f_string", url.QueryEscape(req.GetInfo().GetFString()), "info.f_int32", req.GetInfo().GetFInt32(), "info.f_double", url.QueryEscape(fmt.Sprintf("%g", req.GetInfo().GetFDouble())), "info.f_bool", req.GetInfo().GetFBool(), "info.f_kingdom", req.GetInfo().GetFKingdom()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RepeatDataSimplePath[0:len((*c.CallOptions).RepeatDataSimplePath):len((*c.CallOptions).RepeatDataSimplePath)], opts...)
	var resp *genprotopb.RepeatResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.complianceClient.RepeatDataSimplePath(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *complianceGRPCClient) RepeatDataPathResource(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "info.f_string", url.QueryEscape(req.GetInfo().GetFString()), "info.f_child.f_string", url.QueryEscape(req.GetInfo().GetFChild().GetFString()), "info.f_bool", req.GetInfo().GetFBool()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RepeatDataPathResource[0:len((*c.CallOptions).RepeatDataPathResource):len((*c.CallOptions).RepeatDataPathResource)], opts...)
	var resp *genprotopb.RepeatResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.complianceClient.RepeatDataPathResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *complianceGRPCClient) RepeatDataPathTrailingResource(ctx context.Context, req *genprotopb.RepeatRequest, opts ...gax.CallOption) (*genprotopb.RepeatResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "info.f_string", url.QueryEscape(req.GetInfo().GetFString()), "info.f_child.f_string", url.QueryEscape(req.GetInfo().GetFChild().GetFString())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RepeatDataPathTrailingResource[0:len((*c.CallOptions).RepeatDataPathTrailingResource):len((*c.CallOptions).RepeatDataPathTrailingResource)], opts...)
	var resp *genprotopb.RepeatResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.complianceClient.RepeatDataPathTrailingResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *complianceGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		var resp *locationpb.ListLocationsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
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

func (c *complianceGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
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

func (c *complianceGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
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

func (c *complianceGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
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

func (c *complianceGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
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

func (c *complianceGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		var resp *longrunningpb.ListOperationsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
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

func (c *complianceGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
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

func (c *complianceGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *complianceGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// LocationIterator manages a stream of *locationpb.Location.
type LocationIterator struct {
	items    []*locationpb.Location
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
	InternalFetch func(pageSize int, pageToken string) (results []*locationpb.Location, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *LocationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *LocationIterator) Next() (*locationpb.Location, error) {
	var item *locationpb.Location
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *LocationIterator) bufLen() int {
	return len(it.items)
}

func (it *LocationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// OperationIterator manages a stream of *longrunningpb.Operation.
type OperationIterator struct {
	items    []*longrunningpb.Operation
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
	InternalFetch func(pageSize int, pageToken string) (results []*longrunningpb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*longrunningpb.Operation, error) {
	var item *longrunningpb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
