//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_Get.json
func ExampleTopicsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().Get(ctx, "examplerg", "exampletopic2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Topic = armeventgrid.Topic{
	// 	Name: to.Ptr("exampletopic2"),
	// 	Type: to.Ptr("Microsoft.EventGrid/topics"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armeventgrid.TopicProperties{
	// 		Endpoint: to.Ptr("https://exampletopic2.westcentralus-1.eventgrid.azure.net/api/events"),
	// 		ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_CreateOrUpdate.json
func ExampleTopicsClient_BeginCreateOrUpdate_topicsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.Topic{
		Location: to.Ptr("westus2"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		Properties: &armeventgrid.TopicProperties{
			InboundIPRules: []*armeventgrid.InboundIPRule{
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.30.15"),
				},
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.176.1"),
				}},
			PublicNetworkAccess: to.Ptr(armeventgrid.PublicNetworkAccessEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_CreateOrUpdateForAzureArc.json
func ExampleTopicsClient_BeginCreateOrUpdate_topicsCreateOrUpdateForAzureArc() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.Topic{
		Location: to.Ptr("westus2"),
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
		ExtendedLocation: &armeventgrid.ExtendedLocation{
			Name: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourcegroups/examplerg/providers/Microsoft.ExtendedLocation/CustomLocations/exampleCustomLocation"),
			Type: to.Ptr("CustomLocation"),
		},
		Kind: to.Ptr(armeventgrid.ResourceKindAzureArc),
		Properties: &armeventgrid.TopicProperties{
			InputSchema: to.Ptr(armeventgrid.InputSchemaCloudEventSchemaV10),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_Delete.json
func ExampleTopicsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginDelete(ctx, "examplerg1", "exampletopic1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_Update.json
func ExampleTopicsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.TopicUpdateParameters{
		Properties: &armeventgrid.TopicUpdateParameterProperties{
			InboundIPRules: []*armeventgrid.InboundIPRule{
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.30.15"),
				},
				{
					Action: to.Ptr(armeventgrid.IPActionTypeAllow),
					IPMask: to.Ptr("12.18.176.1"),
				}},
			PublicNetworkAccess: to.Ptr(armeventgrid.PublicNetworkAccessEnabled),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_ListBySubscription.json
func ExampleTopicsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListBySubscriptionPager(&armeventgrid.TopicsClientListBySubscriptionOptions{Filter: nil,
		Top: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TopicsListResult = armeventgrid.TopicsListResult{
		// 	Value: []*armeventgrid.Topic{
		// 		{
		// 			Name: to.Ptr("exampletopic1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1"),
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armeventgrid.TopicProperties{
		// 				Endpoint: to.Ptr("https://exampletopic1.westus2-1.eventgrid.azure.net/api/events"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampletopic2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armeventgrid.TopicProperties{
		// 				Endpoint: to.Ptr("https://exampletopic2.westcentralus-1.eventgrid.azure.net/api/events"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_ListByResourceGroup.json
func ExampleTopicsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListByResourceGroupPager("examplerg", &armeventgrid.TopicsClientListByResourceGroupOptions{Filter: nil,
		Top: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TopicsListResult = armeventgrid.TopicsListResult{
		// 	Value: []*armeventgrid.Topic{
		// 		{
		// 			Name: to.Ptr("exampletopic1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1"),
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armeventgrid.TopicProperties{
		// 				Endpoint: to.Ptr("https://exampletopic1.westus2-1.eventgrid.azure.net/api/events"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampletopic2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armeventgrid.TopicProperties{
		// 				Endpoint: to.Ptr("https://exampletopic2.westcentralus-1.eventgrid.azure.net/api/events"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_ListSharedAccessKeys.json
func ExampleTopicsClient_ListSharedAccessKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().ListSharedAccessKeys(ctx, "examplerg", "exampletopic2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TopicSharedAccessKeys = armeventgrid.TopicSharedAccessKeys{
	// 	Key1: to.Ptr("testKey1Value"),
	// 	Key2: to.Ptr("testKey2Value"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_RegenerateKey.json
func ExampleTopicsClient_BeginRegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginRegenerateKey(ctx, "examplerg", "exampletopic2", armeventgrid.TopicRegenerateKeyRequest{
		KeyName: to.Ptr("key1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TopicSharedAccessKeys = armeventgrid.TopicSharedAccessKeys{
	// 	Key1: to.Ptr("testKey1Value"),
	// 	Key2: to.Ptr("testKey2Value"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_ListEventTypes.json
func ExampleTopicsClient_NewListEventTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListEventTypesPager("examplerg", "Microsoft.Storage", "storageAccounts", "ExampleStorageAccount", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.EventTypesListResult = armeventgrid.EventTypesListResult{
		// 	Value: []*armeventgrid.EventType{
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.BlobCreated"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes/eventTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts/eventTypes/Microsoft.Storage.BlobCreated"),
		// 			Properties: &armeventgrid.EventTypeProperties{
		// 				Description: to.Ptr("Raised when a blob is created."),
		// 				DisplayName: to.Ptr("Blob Created"),
		// 				SchemaURL: to.Ptr("tbd"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.BlobDeleted"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes/eventTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts/eventTypes/Microsoft.Storage.BlobDeleted"),
		// 			Properties: &armeventgrid.EventTypeProperties{
		// 				Description: to.Ptr("Raised when a blob is deleted."),
		// 				DisplayName: to.Ptr("Blob Deleted"),
		// 				SchemaURL: to.Ptr("tbd"),
		// 			},
		// 	}},
		// }
	}
}
