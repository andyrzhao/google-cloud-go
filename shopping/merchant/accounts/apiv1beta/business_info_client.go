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

package accounts

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	accountspb "cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newBusinessInfoClientHook clientHook

// BusinessInfoCallOptions contains the retry settings for each method of BusinessInfoClient.
type BusinessInfoCallOptions struct {
	GetBusinessInfo    []gax.CallOption
	UpdateBusinessInfo []gax.CallOption
}

func defaultBusinessInfoGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("merchantapi.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("merchantapi.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("merchantapi.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBusinessInfoCallOptions() *BusinessInfoCallOptions {
	return &BusinessInfoCallOptions{
		GetBusinessInfo:    []gax.CallOption{},
		UpdateBusinessInfo: []gax.CallOption{},
	}
}

func defaultBusinessInfoRESTCallOptions() *BusinessInfoCallOptions {
	return &BusinessInfoCallOptions{
		GetBusinessInfo:    []gax.CallOption{},
		UpdateBusinessInfo: []gax.CallOption{},
	}
}

// internalBusinessInfoClient is an interface that defines the methods available from Merchant API.
type internalBusinessInfoClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetBusinessInfo(context.Context, *accountspb.GetBusinessInfoRequest, ...gax.CallOption) (*accountspb.BusinessInfo, error)
	UpdateBusinessInfo(context.Context, *accountspb.UpdateBusinessInfoRequest, ...gax.CallOption) (*accountspb.BusinessInfo, error)
}

// BusinessInfoClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to support business info API.
type BusinessInfoClient struct {
	// The internal transport-dependent client.
	internalClient internalBusinessInfoClient

	// The call options for this service.
	CallOptions *BusinessInfoCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BusinessInfoClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BusinessInfoClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BusinessInfoClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetBusinessInfo retrieves the business info of an account.
func (c *BusinessInfoClient) GetBusinessInfo(ctx context.Context, req *accountspb.GetBusinessInfoRequest, opts ...gax.CallOption) (*accountspb.BusinessInfo, error) {
	return c.internalClient.GetBusinessInfo(ctx, req, opts...)
}

// UpdateBusinessInfo updates the business info of an account. Executing this method requires
// admin access.
func (c *BusinessInfoClient) UpdateBusinessInfo(ctx context.Context, req *accountspb.UpdateBusinessInfoRequest, opts ...gax.CallOption) (*accountspb.BusinessInfo, error) {
	return c.internalClient.UpdateBusinessInfo(ctx, req, opts...)
}

// businessInfoGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type businessInfoGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BusinessInfoClient
	CallOptions **BusinessInfoCallOptions

	// The gRPC API client.
	businessInfoClient accountspb.BusinessInfoServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewBusinessInfoClient creates a new business info service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to support business info API.
func NewBusinessInfoClient(ctx context.Context, opts ...option.ClientOption) (*BusinessInfoClient, error) {
	clientOpts := defaultBusinessInfoGRPCClientOptions()
	if newBusinessInfoClientHook != nil {
		hookOpts, err := newBusinessInfoClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BusinessInfoClient{CallOptions: defaultBusinessInfoCallOptions()}

	c := &businessInfoGRPCClient{
		connPool:           connPool,
		businessInfoClient: accountspb.NewBusinessInfoServiceClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *businessInfoGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *businessInfoGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *businessInfoGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type businessInfoRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing BusinessInfoClient
	CallOptions **BusinessInfoCallOptions
}

// NewBusinessInfoRESTClient creates a new business info service rest client.
//
// Service to support business info API.
func NewBusinessInfoRESTClient(ctx context.Context, opts ...option.ClientOption) (*BusinessInfoClient, error) {
	clientOpts := append(defaultBusinessInfoRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultBusinessInfoRESTCallOptions()
	c := &businessInfoRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &BusinessInfoClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultBusinessInfoRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://merchantapi.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://merchantapi.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://merchantapi.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *businessInfoRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *businessInfoRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *businessInfoRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *businessInfoGRPCClient) GetBusinessInfo(ctx context.Context, req *accountspb.GetBusinessInfoRequest, opts ...gax.CallOption) (*accountspb.BusinessInfo, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetBusinessInfo[0:len((*c.CallOptions).GetBusinessInfo):len((*c.CallOptions).GetBusinessInfo)], opts...)
	var resp *accountspb.BusinessInfo
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.businessInfoClient.GetBusinessInfo(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *businessInfoGRPCClient) UpdateBusinessInfo(ctx context.Context, req *accountspb.UpdateBusinessInfoRequest, opts ...gax.CallOption) (*accountspb.BusinessInfo, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "business_info.name", url.QueryEscape(req.GetBusinessInfo().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateBusinessInfo[0:len((*c.CallOptions).UpdateBusinessInfo):len((*c.CallOptions).UpdateBusinessInfo)], opts...)
	var resp *accountspb.BusinessInfo
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.businessInfoClient.UpdateBusinessInfo(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBusinessInfo retrieves the business info of an account.
func (c *businessInfoRESTClient) GetBusinessInfo(ctx context.Context, req *accountspb.GetBusinessInfoRequest, opts ...gax.CallOption) (*accountspb.BusinessInfo, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetBusinessInfo[0:len((*c.CallOptions).GetBusinessInfo):len((*c.CallOptions).GetBusinessInfo)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.BusinessInfo{}
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

// UpdateBusinessInfo updates the business info of an account. Executing this method requires
// admin access.
func (c *businessInfoRESTClient) UpdateBusinessInfo(ctx context.Context, req *accountspb.UpdateBusinessInfoRequest, opts ...gax.CallOption) (*accountspb.BusinessInfo, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetBusinessInfo()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetBusinessInfo().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "business_info.name", url.QueryEscape(req.GetBusinessInfo().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateBusinessInfo[0:len((*c.CallOptions).UpdateBusinessInfo):len((*c.CallOptions).UpdateBusinessInfo)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.BusinessInfo{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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