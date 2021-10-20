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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRegionTargetHttpProxiesClientHook clientHook

// RegionTargetHttpProxiesCallOptions contains the retry settings for each method of RegionTargetHttpProxiesClient.
type RegionTargetHttpProxiesCallOptions struct {
	Delete    []gax.CallOption
	Get       []gax.CallOption
	Insert    []gax.CallOption
	List      []gax.CallOption
	SetUrlMap []gax.CallOption
}

// internalRegionTargetHttpProxiesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalRegionTargetHttpProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteRegionTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetRegionTargetHttpProxyRequest, ...gax.CallOption) (*computepb.TargetHttpProxy, error)
	Insert(context.Context, *computepb.InsertRegionTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListRegionTargetHttpProxiesRequest, ...gax.CallOption) *TargetHttpProxyIterator
	SetUrlMap(context.Context, *computepb.SetUrlMapRegionTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
}

// RegionTargetHttpProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionTargetHttpProxies API.
type RegionTargetHttpProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionTargetHttpProxiesClient

	// The call options for this service.
	CallOptions *RegionTargetHttpProxiesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionTargetHttpProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionTargetHttpProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *RegionTargetHttpProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified TargetHttpProxy resource.
func (c *RegionTargetHttpProxiesClient) Delete(ctx context.Context, req *computepb.DeleteRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetHttpProxy resource in the specified region. Gets a list of available target HTTP proxies by making a list() request.
func (c *RegionTargetHttpProxiesClient) Get(ctx context.Context, req *computepb.GetRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetHttpProxy resource in the specified project and region using the data included in the request.
func (c *RegionTargetHttpProxiesClient) Insert(ctx context.Context, req *computepb.InsertRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetHttpProxy resources available to the specified project in the specified region.
func (c *RegionTargetHttpProxiesClient) List(ctx context.Context, req *computepb.ListRegionTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetUrlMap changes the URL map for TargetHttpProxy.
func (c *RegionTargetHttpProxiesClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetUrlMap(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionTargetHttpProxiesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewRegionTargetHttpProxiesRESTClient creates a new region target http proxies rest client.
//
// The RegionTargetHttpProxies API.
func NewRegionTargetHttpProxiesRESTClient(ctx context.Context, opts ...option.ClientOption) (*RegionTargetHttpProxiesClient, error) {
	clientOpts := append(defaultRegionTargetHttpProxiesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &regionTargetHttpProxiesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &RegionTargetHttpProxiesClient{internalClient: c, CallOptions: &RegionTargetHttpProxiesCallOptions{}}, nil
}

func defaultRegionTargetHttpProxiesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionTargetHttpProxiesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionTargetHttpProxiesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *regionTargetHttpProxiesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified TargetHttpProxy resource.
func (c *regionTargetHttpProxiesRESTClient) Delete(ctx context.Context, req *computepb.DeleteRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetHttpProxies/%v", req.GetProject(), req.GetRegion(), req.GetTargetHttpProxy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified TargetHttpProxy resource in the specified region. Gets a list of available target HTTP proxies by making a list() request.
func (c *regionTargetHttpProxiesRESTClient) Get(ctx context.Context, req *computepb.GetRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpProxy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetHttpProxies/%v", req.GetProject(), req.GetRegion(), req.GetTargetHttpProxy())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.TargetHttpProxy{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a TargetHttpProxy resource in the specified project and region using the data included in the request.
func (c *regionTargetHttpProxiesRESTClient) Insert(ctx context.Context, req *computepb.InsertRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetHttpProxyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetHttpProxies", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves the list of TargetHttpProxy resources available to the specified project in the specified region.
func (c *regionTargetHttpProxiesRESTClient) List(ctx context.Context, req *computepb.ListRegionTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxyIterator {
	it := &TargetHttpProxyIterator{}
	req = proto.Clone(req).(*computepb.ListRegionTargetHttpProxiesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetHttpProxy, string, error) {
		resp := &computepb.TargetHttpProxyList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetHttpProxies", req.GetProject(), req.GetRegion())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// SetUrlMap changes the URL map for TargetHttpProxy.
func (c *regionTargetHttpProxiesRESTClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapRegionTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetUrlMapReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetHttpProxies/%v/setUrlMap", req.GetProject(), req.GetRegion(), req.GetTargetHttpProxy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, err
	}
	op := &Operation{proto: rsp}
	return op, err
}

// TargetHttpProxyIterator manages a stream of *computepb.TargetHttpProxy.
type TargetHttpProxyIterator struct {
	items    []*computepb.TargetHttpProxy
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
	InternalFetch func(pageSize int, pageToken string) (results []*computepb.TargetHttpProxy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TargetHttpProxyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TargetHttpProxyIterator) Next() (*computepb.TargetHttpProxy, error) {
	var item *computepb.TargetHttpProxy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TargetHttpProxyIterator) bufLen() int {
	return len(it.items)
}

func (it *TargetHttpProxyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
