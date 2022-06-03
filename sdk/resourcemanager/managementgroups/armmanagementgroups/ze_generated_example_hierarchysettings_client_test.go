//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/ListHierarchySettings.json
func ExampleHierarchySettingsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewHierarchySettingsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"root",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetHierarchySettings.json
func ExampleHierarchySettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewHierarchySettingsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"root",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutHierarchySettings.json
func ExampleHierarchySettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewHierarchySettingsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"root",
		armmanagementgroups.CreateOrUpdateSettingsRequest{
			Properties: &armmanagementgroups.CreateOrUpdateSettingsProperties{
				DefaultManagementGroup:               to.Ptr("/providers/Microsoft.Management/managementGroups/DefaultGroup"),
				RequireAuthorizationForGroupCreation: to.Ptr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PatchHierarchySettings.json
func ExampleHierarchySettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewHierarchySettingsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"root",
		armmanagementgroups.CreateOrUpdateSettingsRequest{
			Properties: &armmanagementgroups.CreateOrUpdateSettingsProperties{
				DefaultManagementGroup:               to.Ptr("/providers/Microsoft.Management/managementGroups/DefaultGroup"),
				RequireAuthorizationForGroupCreation: to.Ptr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/DeleteHierarchySettings.json
func ExampleHierarchySettingsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewHierarchySettingsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"root",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
