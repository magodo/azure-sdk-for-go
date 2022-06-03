//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedInstanceOperationsClient contains the methods for the ManagedInstanceOperations group.
// Don't use this type directly, use NewManagedInstanceOperationsClient() instead.
type ManagedInstanceOperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedInstanceOperationsClient creates a new instance of ManagedInstanceOperationsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedInstanceOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstanceOperationsClient, error) {
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
	client := &ManagedInstanceOperationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Cancel - Cancels the asynchronous operation on the managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstanceOperationsClientCancelOptions contains the optional parameters for the ManagedInstanceOperationsClient.Cancel
// method.
func (client *ManagedInstanceOperationsClient) Cancel(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientCancelOptions) (ManagedInstanceOperationsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, managedInstanceName, operationID, options)
	if err != nil {
		return ManagedInstanceOperationsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceOperationsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceOperationsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return ManagedInstanceOperationsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ManagedInstanceOperationsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations/{operationId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a management operation on a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstanceOperationsClientGetOptions contains the optional parameters for the ManagedInstanceOperationsClient.Get
// method.
func (client *ManagedInstanceOperationsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientGetOptions) (ManagedInstanceOperationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, operationID, options)
	if err != nil {
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceOperationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations/{operationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstanceOperationsClient) getHandleResponse(resp *http.Response) (ManagedInstanceOperationsClientGetResponse, error) {
	result := ManagedInstanceOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceOperation); err != nil {
		return ManagedInstanceOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedInstancePager - Gets a list of operations performed on the managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstanceOperationsClientListByManagedInstanceOptions contains the optional parameters for the ManagedInstanceOperationsClient.ListByManagedInstance
// method.
func (client *ManagedInstanceOperationsClient) NewListByManagedInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstanceOperationsClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstanceOperationsClientListByManagedInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstanceOperationsClientListByManagedInstanceResponse]{
		More: func(page ManagedInstanceOperationsClientListByManagedInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstanceOperationsClientListByManagedInstanceResponse) (ManagedInstanceOperationsClientListByManagedInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstanceOperationsClientListByManagedInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstanceOperationsClientListByManagedInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstanceOperationsClientListByManagedInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagedInstanceHandleResponse(resp)
		},
	})
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceOperationsClientListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceHandleResponse(resp *http.Response) (ManagedInstanceOperationsClientListByManagedInstanceResponse, error) {
	result := ManagedInstanceOperationsClientListByManagedInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceOperationListResult); err != nil {
		return ManagedInstanceOperationsClientListByManagedInstanceResponse{}, err
	}
	return result, nil
}
