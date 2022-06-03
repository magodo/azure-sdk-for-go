//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
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
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-17-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scopeName - The name of the Azure Monitor PrivateLinkScope resource.
// groupName - The name of the private link resource.
// options - PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get
// method.
func (client *PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, groupName string, options *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, groupName, options)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, groupName string, options *PrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/privateLinkResources/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesClientGetResponse, error) {
	result := PrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPrivateLinkScopePager - Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-17-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scopeName - The name of the Azure Monitor PrivateLinkScope resource.
// options - PrivateLinkResourcesClientListByPrivateLinkScopeOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByPrivateLinkScope
// method.
func (client *PrivateLinkResourcesClient) NewListByPrivateLinkScopePager(resourceGroupName string, scopeName string, options *PrivateLinkResourcesClientListByPrivateLinkScopeOptions) *runtime.Pager[PrivateLinkResourcesClientListByPrivateLinkScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListByPrivateLinkScopeResponse]{
		More: func(page PrivateLinkResourcesClientListByPrivateLinkScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListByPrivateLinkScopeResponse) (PrivateLinkResourcesClientListByPrivateLinkScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkResourcesClientListByPrivateLinkScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkResourcesClientListByPrivateLinkScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkResourcesClientListByPrivateLinkScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPrivateLinkScopeHandleResponse(resp)
		},
	})
}

// listByPrivateLinkScopeCreateRequest creates the ListByPrivateLinkScope request.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkResourcesClientListByPrivateLinkScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPrivateLinkScopeHandleResponse handles the ListByPrivateLinkScope response.
func (client *PrivateLinkResourcesClient) listByPrivateLinkScopeHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByPrivateLinkScopeResponse, error) {
	result := PrivateLinkResourcesClientListByPrivateLinkScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByPrivateLinkScopeResponse{}, err
	}
	return result, nil
}
