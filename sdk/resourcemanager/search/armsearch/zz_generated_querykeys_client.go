//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// QueryKeysClient contains the methods for the QueryKeys group.
// Don't use this type directly, use NewQueryKeysClient() instead.
type QueryKeysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewQueryKeysClient creates a new instance of QueryKeysClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
// Manager API or the portal.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewQueryKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QueryKeysClient, error) {
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
	client := &QueryKeysClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Generates a new query key for the specified search service. You can create up to 50 query keys per service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
// Azure Resource Manager API or the portal.
// searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
// name - The name of the new query API key.
// SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
// method.
// options - QueryKeysClientCreateOptions contains the optional parameters for the QueryKeysClient.Create method.
func (client *QueryKeysClient) Create(ctx context.Context, resourceGroupName string, searchServiceName string, name string, searchManagementRequestOptions *SearchManagementRequestOptions, options *QueryKeysClientCreateOptions) (QueryKeysClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, searchServiceName, name, searchManagementRequestOptions, options)
	if err != nil {
		return QueryKeysClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueryKeysClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueryKeysClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *QueryKeysClient) createCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, name string, searchManagementRequestOptions *SearchManagementRequestOptions, options *QueryKeysClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/createQueryKey/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *QueryKeysClient) createHandleResponse(resp *http.Response) (QueryKeysClientCreateResponse, error) {
	result := QueryKeysClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryKey); err != nil {
		return QueryKeysClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified query key. Unlike admin keys, query keys are not regenerated. The process for regenerating
// a query key is to delete and then recreate it.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
// Azure Resource Manager API or the portal.
// searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
// key - The query key to be deleted. Query keys are identified by value, not by name.
// SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
// method.
// options - QueryKeysClientDeleteOptions contains the optional parameters for the QueryKeysClient.Delete method.
func (client *QueryKeysClient) Delete(ctx context.Context, resourceGroupName string, searchServiceName string, key string, searchManagementRequestOptions *SearchManagementRequestOptions, options *QueryKeysClientDeleteOptions) (QueryKeysClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, searchServiceName, key, searchManagementRequestOptions, options)
	if err != nil {
		return QueryKeysClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueryKeysClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent, http.StatusNotFound) {
		return QueryKeysClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return QueryKeysClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *QueryKeysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, key string, searchManagementRequestOptions *SearchManagementRequestOptions, options *QueryKeysClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/deleteQueryKey/{key}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if key == "" {
		return nil, errors.New("parameter key cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{key}", url.PathEscape(key))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListBySearchServicePager - Returns the list of query API keys for the given Azure Cognitive Search service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
// Azure Resource Manager API or the portal.
// searchServiceName - The name of the Azure Cognitive Search service associated with the specified resource group.
// SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
// method.
// options - QueryKeysClientListBySearchServiceOptions contains the optional parameters for the QueryKeysClient.ListBySearchService
// method.
func (client *QueryKeysClient) NewListBySearchServicePager(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *QueryKeysClientListBySearchServiceOptions) *runtime.Pager[QueryKeysClientListBySearchServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueryKeysClientListBySearchServiceResponse]{
		More: func(page QueryKeysClientListBySearchServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueryKeysClientListBySearchServiceResponse) (QueryKeysClientListBySearchServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySearchServiceCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return QueryKeysClientListBySearchServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return QueryKeysClientListBySearchServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return QueryKeysClientListBySearchServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySearchServiceHandleResponse(resp)
		},
	})
}

// listBySearchServiceCreateRequest creates the ListBySearchService request.
func (client *QueryKeysClient) listBySearchServiceCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *QueryKeysClientListBySearchServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listQueryKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySearchServiceHandleResponse handles the ListBySearchService response.
func (client *QueryKeysClient) listBySearchServiceHandleResponse(resp *http.Response) (QueryKeysClientListBySearchServiceResponse, error) {
	result := QueryKeysClientListBySearchServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListQueryKeysResult); err != nil {
		return QueryKeysClientListBySearchServiceResponse{}, err
	}
	return result, nil
}
