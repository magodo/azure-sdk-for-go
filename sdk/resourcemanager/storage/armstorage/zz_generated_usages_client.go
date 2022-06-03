//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsagesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByLocationPager - Gets the current usage count and the limit for the resources of the location under the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01
// location - The location of the Azure Storage resource.
// options - UsagesClientListByLocationOptions contains the optional parameters for the UsagesClient.ListByLocation method.
func (client *UsagesClient) NewListByLocationPager(location string, options *UsagesClientListByLocationOptions) *runtime.Pager[UsagesClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsagesClientListByLocationResponse]{
		More: func(page UsagesClientListByLocationResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UsagesClientListByLocationResponse) (UsagesClientListByLocationResponse, error) {
			req, err := client.listByLocationCreateRequest(ctx, location, options)
			if err != nil {
				return UsagesClientListByLocationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return UsagesClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsagesClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *UsagesClient) listByLocationCreateRequest(ctx context.Context, location string, options *UsagesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *UsagesClient) listByLocationHandleResponse(resp *http.Response) (UsagesClientListByLocationResponse, error) {
	result := UsagesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return UsagesClientListByLocationResponse{}, err
	}
	return result, nil
}
