//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

const (
	moduleName    = "armmachinelearningservices"
	moduleVersion = "v1.0.1"
)

// AllocationState - Allocation state of the compute. Possible values are: steady - Indicates that the compute is not resizing.
// There are no changes to the number of compute nodes in the compute in progress. A compute
// enters this state when it is created and when no operations are being performed on the compute to change the number of
// compute nodes. resizing - Indicates that the compute is resizing; that is,
// compute nodes are being added to or removed from the compute.
type AllocationState string

const (
	AllocationStateResizing AllocationState = "Resizing"
	AllocationStateSteady   AllocationState = "Steady"
)

// PossibleAllocationStateValues returns the possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{
		AllocationStateResizing,
		AllocationStateSteady,
	}
}

// ApplicationSharingPolicy - Policy for sharing applications on this compute instance among users of parent workspace. If
// Personal, only the creator can access applications on this compute instance. When Shared, any workspace
// user can access applications on this instance depending on his/her assigned role.
type ApplicationSharingPolicy string

const (
	ApplicationSharingPolicyPersonal ApplicationSharingPolicy = "Personal"
	ApplicationSharingPolicyShared   ApplicationSharingPolicy = "Shared"
)

// PossibleApplicationSharingPolicyValues returns the possible values for the ApplicationSharingPolicy const type.
func PossibleApplicationSharingPolicyValues() []ApplicationSharingPolicy {
	return []ApplicationSharingPolicy{
		ApplicationSharingPolicyPersonal,
		ApplicationSharingPolicyShared,
	}
}

// BillingCurrency - Three lettered code specifying the currency of the VM price. Example: USD
type BillingCurrency string

const (
	BillingCurrencyUSD BillingCurrency = "USD"
)

// PossibleBillingCurrencyValues returns the possible values for the BillingCurrency const type.
func PossibleBillingCurrencyValues() []BillingCurrency {
	return []BillingCurrency{
		BillingCurrencyUSD,
	}
}

// ClusterPurpose - Intended usage of the cluster
type ClusterPurpose string

const (
	ClusterPurposeDenseProd ClusterPurpose = "DenseProd"
	ClusterPurposeDevTest   ClusterPurpose = "DevTest"
	ClusterPurposeFastProd  ClusterPurpose = "FastProd"
)

// PossibleClusterPurposeValues returns the possible values for the ClusterPurpose const type.
func PossibleClusterPurposeValues() []ClusterPurpose {
	return []ClusterPurpose{
		ClusterPurposeDenseProd,
		ClusterPurposeDevTest,
		ClusterPurposeFastProd,
	}
}

// ComputeInstanceAuthorizationType - The Compute Instance Authorization type. Available values are personal (default).
type ComputeInstanceAuthorizationType string

const (
	ComputeInstanceAuthorizationTypePersonal ComputeInstanceAuthorizationType = "personal"
)

// PossibleComputeInstanceAuthorizationTypeValues returns the possible values for the ComputeInstanceAuthorizationType const type.
func PossibleComputeInstanceAuthorizationTypeValues() []ComputeInstanceAuthorizationType {
	return []ComputeInstanceAuthorizationType{
		ComputeInstanceAuthorizationTypePersonal,
	}
}

// ComputeInstanceState - Current state of an ComputeInstance.
type ComputeInstanceState string

const (
	ComputeInstanceStateCreateFailed    ComputeInstanceState = "CreateFailed"
	ComputeInstanceStateCreating        ComputeInstanceState = "Creating"
	ComputeInstanceStateDeleting        ComputeInstanceState = "Deleting"
	ComputeInstanceStateJobRunning      ComputeInstanceState = "JobRunning"
	ComputeInstanceStateRestarting      ComputeInstanceState = "Restarting"
	ComputeInstanceStateRunning         ComputeInstanceState = "Running"
	ComputeInstanceStateSettingUp       ComputeInstanceState = "SettingUp"
	ComputeInstanceStateSetupFailed     ComputeInstanceState = "SetupFailed"
	ComputeInstanceStateStarting        ComputeInstanceState = "Starting"
	ComputeInstanceStateStopped         ComputeInstanceState = "Stopped"
	ComputeInstanceStateStopping        ComputeInstanceState = "Stopping"
	ComputeInstanceStateUnknown         ComputeInstanceState = "Unknown"
	ComputeInstanceStateUnusable        ComputeInstanceState = "Unusable"
	ComputeInstanceStateUserSettingUp   ComputeInstanceState = "UserSettingUp"
	ComputeInstanceStateUserSetupFailed ComputeInstanceState = "UserSetupFailed"
)

// PossibleComputeInstanceStateValues returns the possible values for the ComputeInstanceState const type.
func PossibleComputeInstanceStateValues() []ComputeInstanceState {
	return []ComputeInstanceState{
		ComputeInstanceStateCreateFailed,
		ComputeInstanceStateCreating,
		ComputeInstanceStateDeleting,
		ComputeInstanceStateJobRunning,
		ComputeInstanceStateRestarting,
		ComputeInstanceStateRunning,
		ComputeInstanceStateSettingUp,
		ComputeInstanceStateSetupFailed,
		ComputeInstanceStateStarting,
		ComputeInstanceStateStopped,
		ComputeInstanceStateStopping,
		ComputeInstanceStateUnknown,
		ComputeInstanceStateUnusable,
		ComputeInstanceStateUserSettingUp,
		ComputeInstanceStateUserSetupFailed,
	}
}

// ComputeType - The type of compute
type ComputeType string

const (
	ComputeTypeAKS               ComputeType = "AKS"
	ComputeTypeAmlCompute        ComputeType = "AmlCompute"
	ComputeTypeComputeInstance   ComputeType = "ComputeInstance"
	ComputeTypeDataFactory       ComputeType = "DataFactory"
	ComputeTypeDataLakeAnalytics ComputeType = "DataLakeAnalytics"
	ComputeTypeDatabricks        ComputeType = "Databricks"
	ComputeTypeHDInsight         ComputeType = "HDInsight"
	ComputeTypeKubernetes        ComputeType = "Kubernetes"
	ComputeTypeSynapseSpark      ComputeType = "SynapseSpark"
	ComputeTypeVirtualMachine    ComputeType = "VirtualMachine"
)

// PossibleComputeTypeValues returns the possible values for the ComputeType const type.
func PossibleComputeTypeValues() []ComputeType {
	return []ComputeType{
		ComputeTypeAKS,
		ComputeTypeAmlCompute,
		ComputeTypeComputeInstance,
		ComputeTypeDataFactory,
		ComputeTypeDataLakeAnalytics,
		ComputeTypeDatabricks,
		ComputeTypeHDInsight,
		ComputeTypeKubernetes,
		ComputeTypeSynapseSpark,
		ComputeTypeVirtualMachine,
	}
}

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

// DiagnoseResultLevel - Level of workspace setup error
type DiagnoseResultLevel string

const (
	DiagnoseResultLevelError       DiagnoseResultLevel = "Error"
	DiagnoseResultLevelInformation DiagnoseResultLevel = "Information"
	DiagnoseResultLevelWarning     DiagnoseResultLevel = "Warning"
)

// PossibleDiagnoseResultLevelValues returns the possible values for the DiagnoseResultLevel const type.
func PossibleDiagnoseResultLevelValues() []DiagnoseResultLevel {
	return []DiagnoseResultLevel{
		DiagnoseResultLevelError,
		DiagnoseResultLevelInformation,
		DiagnoseResultLevelWarning,
	}
}

// EncryptionStatus - Indicates whether or not the encryption is enabled for the workspace.
type EncryptionStatus string

const (
	EncryptionStatusDisabled EncryptionStatus = "Disabled"
	EncryptionStatusEnabled  EncryptionStatus = "Enabled"
)

// PossibleEncryptionStatusValues returns the possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{
		EncryptionStatusDisabled,
		EncryptionStatusEnabled,
	}
}

// LoadBalancerType - Load Balancer Type
type LoadBalancerType string

const (
	LoadBalancerTypeInternalLoadBalancer LoadBalancerType = "InternalLoadBalancer"
	LoadBalancerTypePublicIP             LoadBalancerType = "PublicIp"
)

// PossibleLoadBalancerTypeValues returns the possible values for the LoadBalancerType const type.
func PossibleLoadBalancerTypeValues() []LoadBalancerType {
	return []LoadBalancerType{
		LoadBalancerTypeInternalLoadBalancer,
		LoadBalancerTypePublicIP,
	}
}

// NodeState - State of the compute node. Values are idle, running, preparing, unusable, leaving and preempted.
type NodeState string

const (
	NodeStateIdle      NodeState = "idle"
	NodeStateLeaving   NodeState = "leaving"
	NodeStatePreempted NodeState = "preempted"
	NodeStatePreparing NodeState = "preparing"
	NodeStateRunning   NodeState = "running"
	NodeStateUnusable  NodeState = "unusable"
)

// PossibleNodeStateValues returns the possible values for the NodeState const type.
func PossibleNodeStateValues() []NodeState {
	return []NodeState{
		NodeStateIdle,
		NodeStateLeaving,
		NodeStatePreempted,
		NodeStatePreparing,
		NodeStateRunning,
		NodeStateUnusable,
	}
}

// OperationName - Name of the last operation.
type OperationName string

const (
	OperationNameCreate  OperationName = "Create"
	OperationNameDelete  OperationName = "Delete"
	OperationNameReimage OperationName = "Reimage"
	OperationNameRestart OperationName = "Restart"
	OperationNameStart   OperationName = "Start"
	OperationNameStop    OperationName = "Stop"
)

// PossibleOperationNameValues returns the possible values for the OperationName const type.
func PossibleOperationNameValues() []OperationName {
	return []OperationName{
		OperationNameCreate,
		OperationNameDelete,
		OperationNameReimage,
		OperationNameRestart,
		OperationNameStart,
		OperationNameStop,
	}
}

// OperationStatus - Operation status.
type OperationStatus string

const (
	OperationStatusCreateFailed  OperationStatus = "CreateFailed"
	OperationStatusDeleteFailed  OperationStatus = "DeleteFailed"
	OperationStatusInProgress    OperationStatus = "InProgress"
	OperationStatusReimageFailed OperationStatus = "ReimageFailed"
	OperationStatusRestartFailed OperationStatus = "RestartFailed"
	OperationStatusStartFailed   OperationStatus = "StartFailed"
	OperationStatusStopFailed    OperationStatus = "StopFailed"
	OperationStatusSucceeded     OperationStatus = "Succeeded"
)

// PossibleOperationStatusValues returns the possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{
		OperationStatusCreateFailed,
		OperationStatusDeleteFailed,
		OperationStatusInProgress,
		OperationStatusReimageFailed,
		OperationStatusRestartFailed,
		OperationStatusStartFailed,
		OperationStatusStopFailed,
		OperationStatusSucceeded,
	}
}

// OsType - Compute OS Type
type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeLinux,
		OsTypeWindows,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
	PrivateEndpointServiceConnectionStatusTimeout      PrivateEndpointServiceConnectionStatus = "Timeout"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusDisconnected,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
		PrivateEndpointServiceConnectionStatusTimeout,
	}
}

// ProvisioningState - The current deployment state of workspace resource. The provisioningState is to indicate states for
// resource provisioning.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether requests from Public Network are allowed.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// QuotaUnit - An enum describing the unit of quota measurement.
type QuotaUnit string

const (
	QuotaUnitCount QuotaUnit = "Count"
)

// PossibleQuotaUnitValues returns the possible values for the QuotaUnit const type.
func PossibleQuotaUnitValues() []QuotaUnit {
	return []QuotaUnit{
		QuotaUnitCount,
	}
}

// ReasonCode - The reason for the restriction.
type ReasonCode string

const (
	ReasonCodeNotAvailableForRegion       ReasonCode = "NotAvailableForRegion"
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeNotSpecified                ReasonCode = "NotSpecified"
)

// PossibleReasonCodeValues returns the possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{
		ReasonCodeNotAvailableForRegion,
		ReasonCodeNotAvailableForSubscription,
		ReasonCodeNotSpecified,
	}
}

// RemoteLoginPortPublicAccess - State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh
// port is closed on all nodes of the cluster. Enabled - Indicates that the public ssh port is open on all
// nodes of the cluster. NotSpecified - Indicates that the public ssh port is closed on all nodes of the cluster if VNet is
// defined, else is open all public nodes. It can be default only during cluster
// creation time, after creation it will be either enabled or disabled.
type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessDisabled     RemoteLoginPortPublicAccess = "Disabled"
	RemoteLoginPortPublicAccessEnabled      RemoteLoginPortPublicAccess = "Enabled"
	RemoteLoginPortPublicAccessNotSpecified RemoteLoginPortPublicAccess = "NotSpecified"
)

// PossibleRemoteLoginPortPublicAccessValues returns the possible values for the RemoteLoginPortPublicAccess const type.
func PossibleRemoteLoginPortPublicAccessValues() []RemoteLoginPortPublicAccess {
	return []RemoteLoginPortPublicAccess{
		RemoteLoginPortPublicAccessDisabled,
		RemoteLoginPortPublicAccessEnabled,
		RemoteLoginPortPublicAccessNotSpecified,
	}
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// SSHPublicAccess - State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh port is closed
// on this instance. Enabled - Indicates that the public ssh port is open and accessible
// according to the VNet/subnet policy if applicable.
type SSHPublicAccess string

const (
	SSHPublicAccessDisabled SSHPublicAccess = "Disabled"
	SSHPublicAccessEnabled  SSHPublicAccess = "Enabled"
)

// PossibleSSHPublicAccessValues returns the possible values for the SSHPublicAccess const type.
func PossibleSSHPublicAccessValues() []SSHPublicAccess {
	return []SSHPublicAccess{
		SSHPublicAccessDisabled,
		SSHPublicAccessEnabled,
	}
}

// SSLConfigurationStatus - Enable or disable ssl for scoring
type SSLConfigurationStatus string

const (
	SSLConfigurationStatusAuto     SSLConfigurationStatus = "Auto"
	SSLConfigurationStatusDisabled SSLConfigurationStatus = "Disabled"
	SSLConfigurationStatusEnabled  SSLConfigurationStatus = "Enabled"
)

// PossibleSSLConfigurationStatusValues returns the possible values for the SSLConfigurationStatus const type.
func PossibleSSLConfigurationStatusValues() []SSLConfigurationStatus {
	return []SSLConfigurationStatus{
		SSLConfigurationStatusAuto,
		SSLConfigurationStatusDisabled,
		SSLConfigurationStatusEnabled,
	}
}

// Status - Status of update workspace quota.
type Status string

const (
	StatusFailure                              Status = "Failure"
	StatusInvalidQuotaBelowClusterMinimum      Status = "InvalidQuotaBelowClusterMinimum"
	StatusInvalidQuotaExceedsSubscriptionLimit Status = "InvalidQuotaExceedsSubscriptionLimit"
	StatusInvalidVMFamilyName                  Status = "InvalidVMFamilyName"
	StatusOperationNotEnabledForRegion         Status = "OperationNotEnabledForRegion"
	StatusOperationNotSupportedForSKU          Status = "OperationNotSupportedForSku"
	StatusSuccess                              Status = "Success"
	StatusUndefined                            Status = "Undefined"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusFailure,
		StatusInvalidQuotaBelowClusterMinimum,
		StatusInvalidQuotaExceedsSubscriptionLimit,
		StatusInvalidVMFamilyName,
		StatusOperationNotEnabledForRegion,
		StatusOperationNotSupportedForSKU,
		StatusSuccess,
		StatusUndefined,
	}
}

type UnderlyingResourceAction string

const (
	UnderlyingResourceActionDelete UnderlyingResourceAction = "Delete"
	UnderlyingResourceActionDetach UnderlyingResourceAction = "Detach"
)

// PossibleUnderlyingResourceActionValues returns the possible values for the UnderlyingResourceAction const type.
func PossibleUnderlyingResourceActionValues() []UnderlyingResourceAction {
	return []UnderlyingResourceAction{
		UnderlyingResourceActionDelete,
		UnderlyingResourceActionDetach,
	}
}

// UnitOfMeasure - The unit of time measurement for the specified VM price. Example: OneHour
type UnitOfMeasure string

const (
	UnitOfMeasureOneHour UnitOfMeasure = "OneHour"
)

// PossibleUnitOfMeasureValues returns the possible values for the UnitOfMeasure const type.
func PossibleUnitOfMeasureValues() []UnitOfMeasure {
	return []UnitOfMeasure{
		UnitOfMeasureOneHour,
	}
}

// UsageUnit - An enum describing the unit of usage measurement.
type UsageUnit string

const (
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
	}
}

// VMPriceOSType - Operating system type used by the VM.
type VMPriceOSType string

const (
	VMPriceOSTypeLinux   VMPriceOSType = "Linux"
	VMPriceOSTypeWindows VMPriceOSType = "Windows"
)

// PossibleVMPriceOSTypeValues returns the possible values for the VMPriceOSType const type.
func PossibleVMPriceOSTypeValues() []VMPriceOSType {
	return []VMPriceOSType{
		VMPriceOSTypeLinux,
		VMPriceOSTypeWindows,
	}
}

// VMPriority - Virtual Machine priority
type VMPriority string

const (
	VMPriorityDedicated   VMPriority = "Dedicated"
	VMPriorityLowPriority VMPriority = "LowPriority"
)

// PossibleVMPriorityValues returns the possible values for the VMPriority const type.
func PossibleVMPriorityValues() []VMPriority {
	return []VMPriority{
		VMPriorityDedicated,
		VMPriorityLowPriority,
	}
}

// VMTier - The type of the VM.
type VMTier string

const (
	VMTierLowPriority VMTier = "LowPriority"
	VMTierSpot        VMTier = "Spot"
	VMTierStandard    VMTier = "Standard"
)

// PossibleVMTierValues returns the possible values for the VMTier const type.
func PossibleVMTierValues() []VMTier {
	return []VMTier{
		VMTierLowPriority,
		VMTierSpot,
		VMTierStandard,
	}
}

// ValueFormat - format for the workspace connection value
type ValueFormat string

const (
	ValueFormatJSON ValueFormat = "JSON"
)

// PossibleValueFormatValues returns the possible values for the ValueFormat const type.
func PossibleValueFormatValues() []ValueFormat {
	return []ValueFormat{
		ValueFormatJSON,
	}
}
