//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsList_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("resRg",
		"myDigitalTwinsService",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsGet_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"resRg",
		"myDigitalTwinsService",
		"myConnection",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsPut_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resRg",
		"myDigitalTwinsService",
		"myConnection",
		armdigitaltwins.TimeSeriesDatabaseConnection{
			Properties: &armdigitaltwins.AzureDataExplorerConnectionProperties{
				ConnectionType:              to.Ptr(armdigitaltwins.ConnectionTypeAzureDataExplorer),
				AdxDatabaseName:             to.Ptr("myDatabase"),
				AdxEndpointURI:              to.Ptr("https://mycluster.kusto.windows.net"),
				AdxResourceID:               to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.Kusto/clusters/mycluster"),
				AdxTableName:                to.Ptr("myTable"),
				EventHubEndpointURI:         to.Ptr("sb://myeh.servicebus.windows.net/"),
				EventHubEntityPath:          to.Ptr("myeh"),
				EventHubNamespaceResourceID: to.Ptr("/subscriptions/c493073e-2460-45ba-a403-f3e0df1e9feg/resourceGroups/testrg/providers/Microsoft.EventHub/namespaces/myeh"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsDelete_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"resRg",
		"myDigitalTwinsService",
		"myConnection",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
