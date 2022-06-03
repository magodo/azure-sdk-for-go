//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armelastic

const (
	moduleName    = "armelastic"
	moduleVersion = "v0.5.0"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ElasticDeploymentStatus - Flag specifying if the Elastic deployment status is healthy or not.
type ElasticDeploymentStatus string

const (
	ElasticDeploymentStatusHealthy   ElasticDeploymentStatus = "Healthy"
	ElasticDeploymentStatusUnhealthy ElasticDeploymentStatus = "Unhealthy"
)

// PossibleElasticDeploymentStatusValues returns the possible values for the ElasticDeploymentStatus const type.
func PossibleElasticDeploymentStatusValues() []ElasticDeploymentStatus {
	return []ElasticDeploymentStatus{
		ElasticDeploymentStatusHealthy,
		ElasticDeploymentStatusUnhealthy,
	}
}

type LiftrResourceCategories string

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = "MonitorLogs"
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = "Unknown"
)

// PossibleLiftrResourceCategoriesValues returns the possible values for the LiftrResourceCategories const type.
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return []LiftrResourceCategories{
		LiftrResourceCategoriesMonitorLogs,
		LiftrResourceCategoriesUnknown,
	}
}

// ManagedIdentityTypes - Managed Identity types.
type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = "SystemAssigned"
)

// PossibleManagedIdentityTypesValues returns the possible values for the ManagedIdentityTypes const type.
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return []ManagedIdentityTypes{
		ManagedIdentityTypesSystemAssigned,
	}
}

// MonitoringStatus - Flag specifying if the resource monitoring is enabled or disabled.
type MonitoringStatus string

const (
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
)

// PossibleMonitoringStatusValues returns the possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{
		MonitoringStatusDisabled,
		MonitoringStatusEnabled,
	}
}

// OperationName - Operation to be performed on the given vm resource id.
type OperationName string

const (
	OperationNameAdd    OperationName = "Add"
	OperationNameDelete OperationName = "Delete"
)

// PossibleOperationNameValues returns the possible values for the OperationName const type.
func PossibleOperationNameValues() []OperationName {
	return []OperationName{
		OperationNameAdd,
		OperationNameDelete,
	}
}

// ProvisioningState - Provisioning state of Elastic resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SendingLogs - Flag indicating the status of the resource for sending logs operation to Elastic.
type SendingLogs string

const (
	SendingLogsFalse SendingLogs = "False"
	SendingLogsTrue  SendingLogs = "True"
)

// PossibleSendingLogsValues returns the possible values for the SendingLogs const type.
func PossibleSendingLogsValues() []SendingLogs {
	return []SendingLogs{
		SendingLogsFalse,
		SendingLogsTrue,
	}
}

// TagAction - Valid actions for a filtering tag. Exclusion takes priority over inclusion.
type TagAction string

const (
	TagActionExclude TagAction = "Exclude"
	TagActionInclude TagAction = "Include"
)

// PossibleTagActionValues returns the possible values for the TagAction const type.
func PossibleTagActionValues() []TagAction {
	return []TagAction{
		TagActionExclude,
		TagActionInclude,
	}
}
