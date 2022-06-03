//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

const (
	moduleName    = "armservicefabricmesh"
	moduleVersion = "v0.5.0"
)

// ApplicationScopedVolumeKind - Specifies the application-scoped volume kind.
type ApplicationScopedVolumeKind string

const (
	// ApplicationScopedVolumeKindServiceFabricVolumeDisk - Provides Service Fabric High Availability Volume Disk
	ApplicationScopedVolumeKindServiceFabricVolumeDisk ApplicationScopedVolumeKind = "ServiceFabricVolumeDisk"
)

// PossibleApplicationScopedVolumeKindValues returns the possible values for the ApplicationScopedVolumeKind const type.
func PossibleApplicationScopedVolumeKindValues() []ApplicationScopedVolumeKind {
	return []ApplicationScopedVolumeKind{
		ApplicationScopedVolumeKindServiceFabricVolumeDisk,
	}
}

// AutoScalingMechanismKind - Enumerates the mechanisms for auto scaling.
type AutoScalingMechanismKind string

const (
	// AutoScalingMechanismKindAddRemoveReplica - Indicates that scaling should be performed by adding or removing replicas.
	AutoScalingMechanismKindAddRemoveReplica AutoScalingMechanismKind = "AddRemoveReplica"
)

// PossibleAutoScalingMechanismKindValues returns the possible values for the AutoScalingMechanismKind const type.
func PossibleAutoScalingMechanismKindValues() []AutoScalingMechanismKind {
	return []AutoScalingMechanismKind{
		AutoScalingMechanismKindAddRemoveReplica,
	}
}

// AutoScalingMetricKind - Enumerates the metrics that are used for triggering auto scaling.
type AutoScalingMetricKind string

const (
	// AutoScalingMetricKindResource - Indicates that the metric is one of resources, like cpu or memory.
	AutoScalingMetricKindResource AutoScalingMetricKind = "Resource"
)

// PossibleAutoScalingMetricKindValues returns the possible values for the AutoScalingMetricKind const type.
func PossibleAutoScalingMetricKindValues() []AutoScalingMetricKind {
	return []AutoScalingMetricKind{
		AutoScalingMetricKindResource,
	}
}

// AutoScalingResourceMetricName - Enumerates the resources that are used for triggering auto scaling.
type AutoScalingResourceMetricName string

const (
	// AutoScalingResourceMetricNameCPU - Indicates that the resource is CPU cores.
	AutoScalingResourceMetricNameCPU AutoScalingResourceMetricName = "cpu"
	// AutoScalingResourceMetricNameMemoryInGB - Indicates that the resource is memory in GB.
	AutoScalingResourceMetricNameMemoryInGB AutoScalingResourceMetricName = "memoryInGB"
)

// PossibleAutoScalingResourceMetricNameValues returns the possible values for the AutoScalingResourceMetricName const type.
func PossibleAutoScalingResourceMetricNameValues() []AutoScalingResourceMetricName {
	return []AutoScalingResourceMetricName{
		AutoScalingResourceMetricNameCPU,
		AutoScalingResourceMetricNameMemoryInGB,
	}
}

// AutoScalingTriggerKind - Enumerates the triggers for auto scaling.
type AutoScalingTriggerKind string

const (
	// AutoScalingTriggerKindAverageLoad - Indicates that scaling should be performed based on average load of all replicas in
	// the service.
	AutoScalingTriggerKindAverageLoad AutoScalingTriggerKind = "AverageLoad"
)

// PossibleAutoScalingTriggerKindValues returns the possible values for the AutoScalingTriggerKind const type.
func PossibleAutoScalingTriggerKindValues() []AutoScalingTriggerKind {
	return []AutoScalingTriggerKind{
		AutoScalingTriggerKindAverageLoad,
	}
}

// DiagnosticsSinkKind - The kind of DiagnosticsSink.
type DiagnosticsSinkKind string

const (
	// DiagnosticsSinkKindAzureInternalMonitoringPipeline - Diagnostics settings for Geneva.
	DiagnosticsSinkKindAzureInternalMonitoringPipeline DiagnosticsSinkKind = "AzureInternalMonitoringPipeline"
	// DiagnosticsSinkKindInvalid - Indicates an invalid sink kind. All Service Fabric enumerations have the invalid type.
	DiagnosticsSinkKindInvalid DiagnosticsSinkKind = "Invalid"
)

// PossibleDiagnosticsSinkKindValues returns the possible values for the DiagnosticsSinkKind const type.
func PossibleDiagnosticsSinkKindValues() []DiagnosticsSinkKind {
	return []DiagnosticsSinkKind{
		DiagnosticsSinkKindAzureInternalMonitoringPipeline,
		DiagnosticsSinkKindInvalid,
	}
}

// HeaderMatchType - how to match header value
type HeaderMatchType string

const (
	HeaderMatchTypeExact HeaderMatchType = "exact"
)

// PossibleHeaderMatchTypeValues returns the possible values for the HeaderMatchType const type.
func PossibleHeaderMatchTypeValues() []HeaderMatchType {
	return []HeaderMatchType{
		HeaderMatchTypeExact,
	}
}

// HealthState - The health state of a Service Fabric entity such as Cluster, Node, Application, Service, Partition, Replica
// etc.
type HealthState string

const (
	// HealthStateError - Indicates the health state is at an error level. Error health state should be investigated, as they
	// can impact the correct functionality of the cluster. The value is 3.
	HealthStateError HealthState = "Error"
	// HealthStateInvalid - Indicates an invalid health state. All Service Fabric enumerations have the invalid type. The value
	// is zero.
	HealthStateInvalid HealthState = "Invalid"
	// HealthStateOk - Indicates the health state is okay. The value is 1.
	HealthStateOk HealthState = "Ok"
	// HealthStateUnknown - Indicates an unknown health status. The value is 65535.
	HealthStateUnknown HealthState = "Unknown"
	// HealthStateWarning - Indicates the health state is at a warning level. The value is 2.
	HealthStateWarning HealthState = "Warning"
)

// PossibleHealthStateValues returns the possible values for the HealthState const type.
func PossibleHealthStateValues() []HealthState {
	return []HealthState{
		HealthStateError,
		HealthStateInvalid,
		HealthStateOk,
		HealthStateUnknown,
		HealthStateWarning,
	}
}

// NetworkKind - The type of a Service Fabric container network.
type NetworkKind string

const (
	// NetworkKindLocal - Indicates a container network local to a single Service Fabric cluster. The value is 1.
	NetworkKindLocal NetworkKind = "Local"
)

// PossibleNetworkKindValues returns the possible values for the NetworkKind const type.
func PossibleNetworkKindValues() []NetworkKind {
	return []NetworkKind{
		NetworkKindLocal,
	}
}

// OperatingSystemType - The operation system required by the code in service.
type OperatingSystemType string

const (
	// OperatingSystemTypeLinux - The required operating system is Linux.
	OperatingSystemTypeLinux OperatingSystemType = "Linux"
	// OperatingSystemTypeWindows - The required operating system is Windows.
	OperatingSystemTypeWindows OperatingSystemType = "Windows"
)

// PossibleOperatingSystemTypeValues returns the possible values for the OperatingSystemType const type.
func PossibleOperatingSystemTypeValues() []OperatingSystemType {
	return []OperatingSystemType{
		OperatingSystemTypeLinux,
		OperatingSystemTypeWindows,
	}
}

// PathMatchType - how to match value in the Uri
type PathMatchType string

const (
	PathMatchTypePrefix PathMatchType = "prefix"
)

// PossiblePathMatchTypeValues returns the possible values for the PathMatchType const type.
func PossiblePathMatchTypeValues() []PathMatchType {
	return []PathMatchType{
		PathMatchTypePrefix,
	}
}

// ResourceStatus - Status of the resource.
type ResourceStatus string

const (
	// ResourceStatusCreating - Indicates the resource is being created. The value is 3.
	ResourceStatusCreating ResourceStatus = "Creating"
	// ResourceStatusDeleting - Indicates the resource is being deleted. The value is 4.
	ResourceStatusDeleting ResourceStatus = "Deleting"
	// ResourceStatusFailed - Indicates the resource is not functional due to persistent failures. See statusDetails property
	// for more details. The value is 5.
	ResourceStatusFailed ResourceStatus = "Failed"
	// ResourceStatusReady - Indicates the resource is ready. The value is 1.
	ResourceStatusReady ResourceStatus = "Ready"
	// ResourceStatusUnknown - Indicates the resource status is unknown. The value is zero.
	ResourceStatusUnknown ResourceStatus = "Unknown"
	// ResourceStatusUpgrading - Indicates the resource is upgrading. The value is 2.
	ResourceStatusUpgrading ResourceStatus = "Upgrading"
)

// PossibleResourceStatusValues returns the possible values for the ResourceStatus const type.
func PossibleResourceStatusValues() []ResourceStatus {
	return []ResourceStatus{
		ResourceStatusCreating,
		ResourceStatusDeleting,
		ResourceStatusFailed,
		ResourceStatusReady,
		ResourceStatusUnknown,
		ResourceStatusUpgrading,
	}
}

// SecretKind - Describes the kind of secret.
type SecretKind string

const (
	// SecretKindInlinedValue - A simple secret resource whose plaintext value is provided by the user.
	SecretKindInlinedValue SecretKind = "inlinedValue"
)

// PossibleSecretKindValues returns the possible values for the SecretKind const type.
func PossibleSecretKindValues() []SecretKind {
	return []SecretKind{
		SecretKindInlinedValue,
	}
}

// SizeTypes - Volume size
type SizeTypes string

const (
	SizeTypesLarge  SizeTypes = "Large"
	SizeTypesMedium SizeTypes = "Medium"
	SizeTypesSmall  SizeTypes = "Small"
)

// PossibleSizeTypesValues returns the possible values for the SizeTypes const type.
func PossibleSizeTypesValues() []SizeTypes {
	return []SizeTypes{
		SizeTypesLarge,
		SizeTypesMedium,
		SizeTypesSmall,
	}
}

// VolumeProvider - Describes the provider of the volume resource.
type VolumeProvider string

const (
	// VolumeProviderSFAzureFile - Provides volumes that are backed by Azure Files.
	VolumeProviderSFAzureFile VolumeProvider = "SFAzureFile"
)

// PossibleVolumeProviderValues returns the possible values for the VolumeProvider const type.
func PossibleVolumeProviderValues() []VolumeProvider {
	return []VolumeProvider{
		VolumeProviderSFAzureFile,
	}
}
