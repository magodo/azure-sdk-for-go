package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DiscoverTenantsClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type DiscoverTenantsClient struct {
	BaseClient
}

// NewDiscoverTenantsClient creates an instance of the DiscoverTenantsClient client.
func NewDiscoverTenantsClient(subscriptionID string) DiscoverTenantsClient {
	return NewDiscoverTenantsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDiscoverTenantsClientWithBaseURI creates an instance of the DiscoverTenantsClient client.
func NewDiscoverTenantsClientWithBaseURI(baseURI string, subscriptionID string) DiscoverTenantsClient {
	return DiscoverTenantsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a Tenant Properties.
// Parameters:
// billingAccountID - billing Account Id.
// billingProfileID - billing Profile Id.
func (client DiscoverTenantsClient) Get(ctx context.Context, billingAccountID string, billingProfileID string) (result DiscoverTenant, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DiscoverTenantsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountID, billingProfileID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.DiscoverTenantsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.DiscoverTenantsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.DiscoverTenantsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DiscoverTenantsClient) GetPreparer(ctx context.Context, billingAccountID string, billingProfileID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
		"billingProfileId": autorest.Encode("path", billingProfileID),
	}

	const APIVersion = "2018-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/discoverTenants", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DiscoverTenantsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DiscoverTenantsClient) GetResponder(resp *http.Response) (result DiscoverTenant, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
