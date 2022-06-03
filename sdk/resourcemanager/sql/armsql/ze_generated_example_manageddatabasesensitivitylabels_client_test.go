//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnSensitivityLabelGet.json
func ExampleManagedDatabaseSensitivityLabelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		armsql.SensitivityLabelSourceCurrent,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnSensitivityLabelCreate.json
func ExampleManagedDatabaseSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		armsql.SensitivityLabel{
			Properties: &armsql.SensitivityLabelProperties{
				InformationType:   to.Ptr("PhoneNumber"),
				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
				LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
				LabelName:         to.Ptr("PII"),
				Rank:              to.Ptr(armsql.SensitivityLabelRankHigh),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnSensitivityLabelDelete.json
func ExampleManagedDatabaseSensitivityLabelsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
func ExampleManagedDatabaseSensitivityLabelsClient_DisableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.DisableRecommendation(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelEnable.json
func ExampleManagedDatabaseSensitivityLabelsClient_EnableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.EnableRecommendation(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsListByDatabaseCurrent.json
func ExampleManagedDatabaseSensitivityLabelsClient_NewListCurrentByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListCurrentByDatabasePager("myRG",
		"myManagedInstanceName",
		"myDatabase",
		&armsql.ManagedDatabaseSensitivityLabelsClientListCurrentByDatabaseOptions{SkipToken: nil,
			Count:  nil,
			Filter: nil,
		})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsCurrentUpdate.json
func ExampleManagedDatabaseSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		armsql.SensitivityLabelUpdateList{
			Operations: []*armsql.SensitivityLabelUpdate{
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column1"),
						Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
						SensitivityLabel: &armsql.SensitivityLabel{
							Properties: &armsql.SensitivityLabelProperties{
								InformationType:   to.Ptr("Financial"),
								InformationTypeID: to.Ptr("1D3652D6-422C-4115-82F1-65DAEBC665C8"),
								LabelID:           to.Ptr("3A477B16-9423-432B-AA97-6069B481CEC3"),
								LabelName:         to.Ptr("Highly Confidential"),
							},
						},
						Table: to.Ptr("table1"),
					},
				},
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column2"),
						Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
						SensitivityLabel: &armsql.SensitivityLabel{
							Properties: &armsql.SensitivityLabelProperties{
								InformationType:   to.Ptr("PhoneNumber"),
								InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
								LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
								LabelName:         to.Ptr("PII"),
							},
						},
						Table: to.Ptr("table2"),
					},
				},
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("Column3"),
						Op:     to.Ptr(armsql.SensitivityLabelUpdateKindRemove),
						Table:  to.Ptr("Table1"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsListByDatabaseRecommended.json
func ExampleManagedDatabaseSensitivityLabelsClient_NewListRecommendedByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListRecommendedByDatabasePager("myRG",
		"myManagedInstanceName",
		"myDatabase",
		&armsql.ManagedDatabaseSensitivityLabelsClientListRecommendedByDatabaseOptions{SkipToken: nil,
			IncludeDisabledRecommendations: nil,
			Filter:                         nil,
		})
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
