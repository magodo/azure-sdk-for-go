//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolOperationResultsClient contains the methods for the SQLPoolOperationResults group.
// Don't use this type directly, use NewSQLPoolOperationResultsClient() instead.
type SQLPoolOperationResultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolOperationResultsClient creates a new instance of SQLPoolOperationResultsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLPoolOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolOperationResultsClient, error) {
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
	client := &SQLPoolOperationResultsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetLocationHeaderResult - Get the status of a SQL pool operation
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// operationID - Operation ID
// options - SQLPoolOperationResultsClientGetLocationHeaderResultOptions contains the optional parameters for the SQLPoolOperationResultsClient.GetLocationHeaderResult
// method.
func (client *SQLPoolOperationResultsClient) GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, operationID string, options *SQLPoolOperationResultsClientGetLocationHeaderResultOptions) (SQLPoolOperationResultsClientGetLocationHeaderResultResponse, error) {
	req, err := client.getLocationHeaderResultCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, operationID, options)
	if err != nil {
		return SQLPoolOperationResultsClientGetLocationHeaderResultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolOperationResultsClientGetLocationHeaderResultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return SQLPoolOperationResultsClientGetLocationHeaderResultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLocationHeaderResultHandleResponse(resp)
}

// getLocationHeaderResultCreateRequest creates the GetLocationHeaderResult request.
func (client *SQLPoolOperationResultsClient) getLocationHeaderResultCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, operationID string, options *SQLPoolOperationResultsClientGetLocationHeaderResultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLocationHeaderResultHandleResponse handles the GetLocationHeaderResult response.
func (client *SQLPoolOperationResultsClient) getLocationHeaderResultHandleResponse(resp *http.Response) (SQLPoolOperationResultsClientGetLocationHeaderResultResponse, error) {
	result := SQLPoolOperationResultsClientGetLocationHeaderResultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return SQLPoolOperationResultsClientGetLocationHeaderResultResponse{}, err
	}
	return result, nil
}
