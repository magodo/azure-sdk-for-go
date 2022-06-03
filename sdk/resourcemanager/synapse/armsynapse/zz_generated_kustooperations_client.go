//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// KustoOperationsClient contains the methods for the KustoOperations group.
// Don't use this type directly, use NewKustoOperationsClient() instead.
type KustoOperationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewKustoOperationsClient creates a new instance of KustoOperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewKustoOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*KustoOperationsClient, error) {
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
	client := &KustoOperationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists available operations for the Kusto sub-resources inside Microsoft.Synapse provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01-preview
// options - KustoOperationsClientListOptions contains the optional parameters for the KustoOperationsClient.List method.
func (client *KustoOperationsClient) NewListPager(options *KustoOperationsClientListOptions) *runtime.Pager[KustoOperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[KustoOperationsClientListResponse]{
		More: func(page KustoOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KustoOperationsClientListResponse) (KustoOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return KustoOperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return KustoOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KustoOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *KustoOperationsClient) listCreateRequest(ctx context.Context, options *KustoOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Synapse/kustooperations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *KustoOperationsClient) listHandleResponse(resp *http.Response) (KustoOperationsClientListResponse, error) {
	result := KustoOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return KustoOperationsClientListResponse{}, err
	}
	return result, nil
}
