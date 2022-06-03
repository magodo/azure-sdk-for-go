//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// AuthorizationsClient contains the methods for the Authorizations group.
// Don't use this type directly, use NewAuthorizationsClient() instead.
type AuthorizationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAuthorizationsClient creates a new instance of AuthorizationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAuthorizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationsClient, error) {
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
	client := &AuthorizationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update an ExpressRoute Circuit Authorization in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - The name of the private cloud.
// authorizationName - Name of the ExpressRoute Circuit Authorization in the private cloud
// authorization - An ExpressRoute Circuit Authorization
// options - AuthorizationsClientBeginCreateOrUpdateOptions contains the optional parameters for the AuthorizationsClient.BeginCreateOrUpdate
// method.
func (client *AuthorizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization ExpressRouteAuthorization, options *AuthorizationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AuthorizationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, authorizationName, authorization, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AuthorizationsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AuthorizationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update an ExpressRoute Circuit Authorization in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
func (client *AuthorizationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization ExpressRouteAuthorization, options *AuthorizationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, authorizationName, authorization, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AuthorizationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, authorization ExpressRouteAuthorization, options *AuthorizationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, authorization)
}

// BeginDelete - Delete an ExpressRoute Circuit Authorization in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// authorizationName - Name of the ExpressRoute Circuit Authorization in the private cloud
// options - AuthorizationsClientBeginDeleteOptions contains the optional parameters for the AuthorizationsClient.BeginDelete
// method.
func (client *AuthorizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientBeginDeleteOptions) (*runtime.Poller[AuthorizationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, authorizationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AuthorizationsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AuthorizationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an ExpressRoute Circuit Authorization in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
func (client *AuthorizationsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, authorizationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AuthorizationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an ExpressRoute Circuit Authorization by name in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// authorizationName - Name of the ExpressRoute Circuit Authorization in the private cloud
// options - AuthorizationsClientGetOptions contains the optional parameters for the AuthorizationsClient.Get method.
func (client *AuthorizationsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientGetOptions) (AuthorizationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, authorizationName, options)
	if err != nil {
		return AuthorizationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AuthorizationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, authorizationName string, options *AuthorizationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations/{authorizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if authorizationName == "" {
		return nil, errors.New("parameter authorizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationName}", url.PathEscape(authorizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AuthorizationsClient) getHandleResponse(resp *http.Response) (AuthorizationsClientGetResponse, error) {
	result := AuthorizationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteAuthorization); err != nil {
		return AuthorizationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List ExpressRoute Circuit Authorizations in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - AuthorizationsClientListOptions contains the optional parameters for the AuthorizationsClient.List method.
func (client *AuthorizationsClient) NewListPager(resourceGroupName string, privateCloudName string, options *AuthorizationsClientListOptions) *runtime.Pager[AuthorizationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthorizationsClientListResponse]{
		More: func(page AuthorizationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthorizationsClientListResponse) (AuthorizationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AuthorizationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AuthorizationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AuthorizationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AuthorizationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *AuthorizationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/authorizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AuthorizationsClient) listHandleResponse(resp *http.Response) (AuthorizationsClientListResponse, error) {
	result := AuthorizationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteAuthorizationList); err != nil {
		return AuthorizationsClientListResponse{}, err
	}
	return result, nil
}
