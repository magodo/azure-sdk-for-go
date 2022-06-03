//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

const (
	moduleName    = "armservicebus"
	moduleVersion = "v1.0.0"
)

type AccessRights string

const (
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
	AccessRightsListen AccessRights = "Listen"
)

// PossibleAccessRightsValues returns the possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsManage,
		AccessRightsSend,
		AccessRightsListen,
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

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// EndPointProvisioningState - Provisioning state of the Private Endpoint Connection.
type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

// PossibleEndPointProvisioningStateValues returns the possible values for the EndPointProvisioningState const type.
func PossibleEndPointProvisioningStateValues() []EndPointProvisioningState {
	return []EndPointProvisioningState{
		EndPointProvisioningStateCanceled,
		EndPointProvisioningStateCreating,
		EndPointProvisioningStateDeleting,
		EndPointProvisioningStateFailed,
		EndPointProvisioningStateSucceeded,
		EndPointProvisioningStateUpdating,
	}
}

// EntityStatus - Entity status.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns the possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusDisabled,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusRenaming,
		EntityStatusUnknown,
	}
}

// FilterType - Rule filter types
type FilterType string

const (
	FilterTypeSQLFilter         FilterType = "SqlFilter"
	FilterTypeCorrelationFilter FilterType = "CorrelationFilter"
)

// PossibleFilterTypeValues returns the possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{
		FilterTypeSQLFilter,
		FilterTypeCorrelationFilter,
	}
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

// ManagedServiceIdentityType - Type of managed service identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeUserAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeNone,
	}
}

type MigrationConfigurationName string

const (
	MigrationConfigurationNameDefault MigrationConfigurationName = "$default"
)

// PossibleMigrationConfigurationNameValues returns the possible values for the MigrationConfigurationName const type.
func PossibleMigrationConfigurationNameValues() []MigrationConfigurationName {
	return []MigrationConfigurationName{
		MigrationConfigurationNameDefault,
	}
}

// NetworkRuleIPAction - The IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns the possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{
		NetworkRuleIPActionAllow,
	}
}

// PrivateLinkConnectionStatus - Status of the connection.
type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

// PossiblePrivateLinkConnectionStatusValues returns the possible values for the PrivateLinkConnectionStatus const type.
func PossiblePrivateLinkConnectionStatusValues() []PrivateLinkConnectionStatus {
	return []PrivateLinkConnectionStatus{
		PrivateLinkConnectionStatusApproved,
		PrivateLinkConnectionStatusDisconnected,
		PrivateLinkConnectionStatusPending,
		PrivateLinkConnectionStatusRejected,
	}
}

// ProvisioningStateDR - Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or
// 'Succeeded' or 'Failed'
type ProvisioningStateDR string

const (
	ProvisioningStateDRAccepted  ProvisioningStateDR = "Accepted"
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
	ProvisioningStateDRFailed    ProvisioningStateDR = "Failed"
)

// PossibleProvisioningStateDRValues returns the possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{
		ProvisioningStateDRAccepted,
		ProvisioningStateDRSucceeded,
		ProvisioningStateDRFailed,
	}
}

// PublicNetworkAccessFlag - This determines if traffic is allowed over public network. By default it is enabled.
type PublicNetworkAccessFlag string

const (
	PublicNetworkAccessFlagDisabled PublicNetworkAccessFlag = "Disabled"
	PublicNetworkAccessFlagEnabled  PublicNetworkAccessFlag = "Enabled"
)

// PossiblePublicNetworkAccessFlagValues returns the possible values for the PublicNetworkAccessFlag const type.
func PossiblePublicNetworkAccessFlagValues() []PublicNetworkAccessFlag {
	return []PublicNetworkAccessFlag{
		PublicNetworkAccessFlagDisabled,
		PublicNetworkAccessFlagEnabled,
	}
}

// RoleDisasterRecovery - role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
type RoleDisasterRecovery string

const (
	RoleDisasterRecoveryPrimary               RoleDisasterRecovery = "Primary"
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	RoleDisasterRecoverySecondary             RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns the possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{
		RoleDisasterRecoveryPrimary,
		RoleDisasterRecoveryPrimaryNotReplicating,
		RoleDisasterRecoverySecondary,
	}
}

// SKUName - Name of this SKU.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNameStandard SKUName = "Standard"
	SKUNamePremium  SKUName = "Premium"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameStandard,
		SKUNamePremium,
	}
}

// SKUTier - The billing tier of this particular SKU.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns the possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}
