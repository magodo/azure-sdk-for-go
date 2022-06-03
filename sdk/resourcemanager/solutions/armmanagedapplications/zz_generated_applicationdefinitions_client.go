//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

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

// ApplicationDefinitionsClient contains the methods for the ApplicationDefinitions group.
// Don't use this type directly, use NewApplicationDefinitionsClient() instead.
type ApplicationDefinitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationDefinitionsClient creates a new instance of ApplicationDefinitionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationDefinitionsClient, error) {
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
	client := &ApplicationDefinitionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationDefinitionName - The name of the managed application definition.
// parameters - Parameters supplied to the create or update an managed application definition.
// options - ApplicationDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationDefinitionsClient.BeginCreateOrUpdate
// method.
func (client *ApplicationDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ApplicationDefinitionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationDefinitionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a new managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
func (client *ApplicationDefinitionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
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
func (client *ApplicationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationDefinitionName - The name of the managed application definition.
// options - ApplicationDefinitionsClientBeginDeleteOptions contains the optional parameters for the ApplicationDefinitionsClient.BeginDelete
// method.
func (client *ApplicationDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteOptions) (*runtime.Poller[ApplicationDefinitionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, applicationDefinitionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ApplicationDefinitionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationDefinitionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
func (client *ApplicationDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
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
func (client *ApplicationDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationDefinitionName - The name of the managed application definition.
// options - ApplicationDefinitionsClientGetOptions contains the optional parameters for the ApplicationDefinitionsClient.Get
// method.
func (client *ApplicationDefinitionsClient) Get(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetOptions) (ApplicationDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationDefinitionsClient) getHandleResponse(resp *http.Response) (ApplicationDefinitionsClientGetResponse, error) {
	result := ApplicationDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the managed application definitions in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ApplicationDefinitionsClientListByResourceGroupOptions contains the optional parameters for the ApplicationDefinitionsClient.ListByResourceGroup
// method.
func (client *ApplicationDefinitionsClient) NewListByResourceGroupPager(resourceGroupName string, options *ApplicationDefinitionsClientListByResourceGroupOptions) *runtime.Pager[ApplicationDefinitionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationDefinitionsClientListByResourceGroupResponse]{
		More: func(page ApplicationDefinitionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationDefinitionsClientListByResourceGroupResponse) (ApplicationDefinitionsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationDefinitionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationDefinitionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationDefinitionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationDefinitionsClientListByResourceGroupResponse, error) {
	result := ApplicationDefinitionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinitionListResult); err != nil {
		return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all the application definitions within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
// options - ApplicationDefinitionsClientListBySubscriptionOptions contains the optional parameters for the ApplicationDefinitionsClient.ListBySubscription
// method.
func (client *ApplicationDefinitionsClient) NewListBySubscriptionPager(options *ApplicationDefinitionsClientListBySubscriptionOptions) *runtime.Pager[ApplicationDefinitionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationDefinitionsClientListBySubscriptionResponse]{
		More: func(page ApplicationDefinitionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationDefinitionsClientListBySubscriptionResponse) (ApplicationDefinitionsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationDefinitionsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationDefinitionsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationDefinitionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationDefinitionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationDefinitionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/applicationDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationDefinitionsClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationDefinitionsClientListBySubscriptionResponse, error) {
	result := ApplicationDefinitionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinitionListResult); err != nil {
		return ApplicationDefinitionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// applicationDefinitionName - The name of the managed application definition.
// parameters - Parameters supplied to the update a managed application definition.
// options - ApplicationDefinitionsClientUpdateOptions contains the optional parameters for the ApplicationDefinitionsClient.Update
// method.
func (client *ApplicationDefinitionsClient) Update(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinitionPatchable, options *ApplicationDefinitionsClientUpdateOptions) (ApplicationDefinitionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationDefinitionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationDefinitionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinitionPatchable, options *ApplicationDefinitionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ApplicationDefinitionsClient) updateHandleResponse(resp *http.Response) (ApplicationDefinitionsClientUpdateResponse, error) {
	result := ApplicationDefinitionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	return result, nil
}
