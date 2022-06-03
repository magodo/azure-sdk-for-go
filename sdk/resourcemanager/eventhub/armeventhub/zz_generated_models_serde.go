//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ApplicationGroupProperties.
func (a ApplicationGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "clientAppGroupIdentifier", a.ClientAppGroupIdentifier)
	populate(objectMap, "isEnabled", a.IsEnabled)
	populate(objectMap, "policies", a.Policies)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApplicationGroupProperties.
func (a *ApplicationGroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clientAppGroupIdentifier":
			err = unpopulate(val, "ClientAppGroupIdentifier", &a.ClientAppGroupIdentifier)
			delete(rawMsg, key)
		case "isEnabled":
			err = unpopulate(val, "IsEnabled", &a.IsEnabled)
			delete(rawMsg, key)
		case "policies":
			a.Policies, err = unmarshalApplicationGroupPolicyClassificationArray(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AuthorizationRuleProperties.
func (a AuthorizationRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "rights", a.Rights)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Cluster.
func (c Cluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ClusterQuotaConfigurationProperties.
func (c ClusterQuotaConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "settings", c.Settings)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConsumerGroupProperties.
func (c ConsumerGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", c.CreatedAt)
	populateTimeRFC3339(objectMap, "updatedAt", c.UpdatedAt)
	populate(objectMap, "userMetadata", c.UserMetadata)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConsumerGroupProperties.
func (c *ConsumerGroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &c.CreatedAt)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, "UpdatedAt", &c.UpdatedAt)
			delete(rawMsg, key)
		case "userMetadata":
			err = unpopulate(val, "UserMetadata", &c.UserMetadata)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EHNamespace.
func (e EHNamespace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "identity", e.Identity)
	populate(objectMap, "location", e.Location)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "sku", e.SKU)
	populate(objectMap, "systemData", e.SystemData)
	populate(objectMap, "tags", e.Tags)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EHNamespaceProperties.
func (e EHNamespaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "alternateName", e.AlternateName)
	populate(objectMap, "clusterArmId", e.ClusterArmID)
	populateTimeRFC3339(objectMap, "createdAt", e.CreatedAt)
	populate(objectMap, "disableLocalAuth", e.DisableLocalAuth)
	populate(objectMap, "encryption", e.Encryption)
	populate(objectMap, "isAutoInflateEnabled", e.IsAutoInflateEnabled)
	populate(objectMap, "kafkaEnabled", e.KafkaEnabled)
	populate(objectMap, "maximumThroughputUnits", e.MaximumThroughputUnits)
	populate(objectMap, "metricId", e.MetricID)
	populate(objectMap, "minimumTlsVersion", e.MinimumTLSVersion)
	populate(objectMap, "privateEndpointConnections", e.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", e.PublicNetworkAccess)
	populate(objectMap, "serviceBusEndpoint", e.ServiceBusEndpoint)
	populate(objectMap, "status", e.Status)
	populateTimeRFC3339(objectMap, "updatedAt", e.UpdatedAt)
	populate(objectMap, "zoneRedundant", e.ZoneRedundant)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EHNamespaceProperties.
func (e *EHNamespaceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "alternateName":
			err = unpopulate(val, "AlternateName", &e.AlternateName)
			delete(rawMsg, key)
		case "clusterArmId":
			err = unpopulate(val, "ClusterArmID", &e.ClusterArmID)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &e.CreatedAt)
			delete(rawMsg, key)
		case "disableLocalAuth":
			err = unpopulate(val, "DisableLocalAuth", &e.DisableLocalAuth)
			delete(rawMsg, key)
		case "encryption":
			err = unpopulate(val, "Encryption", &e.Encryption)
			delete(rawMsg, key)
		case "isAutoInflateEnabled":
			err = unpopulate(val, "IsAutoInflateEnabled", &e.IsAutoInflateEnabled)
			delete(rawMsg, key)
		case "kafkaEnabled":
			err = unpopulate(val, "KafkaEnabled", &e.KafkaEnabled)
			delete(rawMsg, key)
		case "maximumThroughputUnits":
			err = unpopulate(val, "MaximumThroughputUnits", &e.MaximumThroughputUnits)
			delete(rawMsg, key)
		case "metricId":
			err = unpopulate(val, "MetricID", &e.MetricID)
			delete(rawMsg, key)
		case "minimumTlsVersion":
			err = unpopulate(val, "MinimumTLSVersion", &e.MinimumTLSVersion)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, "PrivateEndpointConnections", &e.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &e.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, "PublicNetworkAccess", &e.PublicNetworkAccess)
			delete(rawMsg, key)
		case "serviceBusEndpoint":
			err = unpopulate(val, "ServiceBusEndpoint", &e.ServiceBusEndpoint)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &e.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, "UpdatedAt", &e.UpdatedAt)
			delete(rawMsg, key)
		case "zoneRedundant":
			err = unpopulate(val, "ZoneRedundant", &e.ZoneRedundant)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Encryption.
func (e Encryption) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "keySource", e.KeySource)
	populate(objectMap, "keyVaultProperties", e.KeyVaultProperties)
	populate(objectMap, "requireInfrastructureEncryption", e.RequireInfrastructureEncryption)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Identity.
func (i Identity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	populate(objectMap, "userAssignedIdentities", i.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSetProperties.
func (n NetworkRuleSetProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultAction", n.DefaultAction)
	populate(objectMap, "ipRules", n.IPRules)
	populate(objectMap, "publicNetworkAccess", n.PublicNetworkAccess)
	populate(objectMap, "trustedServiceAccessEnabled", n.TrustedServiceAccessEnabled)
	populate(objectMap, "virtualNetworkRules", n.VirtualNetworkRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkSecurityPerimeterConfiguration.
func (n NetworkSecurityPerimeterConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", n.ID)
	populate(objectMap, "location", n.Location)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "properties", n.Properties)
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "type", n.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkSecurityPerimeterConfigurationProperties.
func (n NetworkSecurityPerimeterConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "networkSecurityPerimeter", n.NetworkSecurityPerimeter)
	populate(objectMap, "profile", n.Profile)
	populate(objectMap, "provisioningIssues", n.ProvisioningIssues)
	populate(objectMap, "provisioningState", n.ProvisioningState)
	populate(objectMap, "resourceAssociation", n.ResourceAssociation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkSecurityPerimeterConfigurationPropertiesProfile.
func (n NetworkSecurityPerimeterConfigurationPropertiesProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessRules", n.AccessRules)
	populate(objectMap, "accessRulesVersion", n.AccessRulesVersion)
	populate(objectMap, "name", n.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NspAccessRuleProperties.
func (n NspAccessRuleProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "addressPrefixes", n.AddressPrefixes)
	populate(objectMap, "direction", n.Direction)
	populate(objectMap, "fullyQualifiedDomainNames", n.FullyQualifiedDomainNames)
	populate(objectMap, "networkSecurityPerimeters", n.NetworkSecurityPerimeters)
	populate(objectMap, "subscriptions", n.Subscriptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "captureDescription", p.CaptureDescription)
	populateTimeRFC3339(objectMap, "createdAt", p.CreatedAt)
	populate(objectMap, "messageRetentionInDays", p.MessageRetentionInDays)
	populate(objectMap, "partitionCount", p.PartitionCount)
	populate(objectMap, "partitionIds", p.PartitionIDs)
	populate(objectMap, "status", p.Status)
	populateTimeRFC3339(objectMap, "updatedAt", p.UpdatedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "captureDescription":
			err = unpopulate(val, "CaptureDescription", &p.CaptureDescription)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &p.CreatedAt)
			delete(rawMsg, key)
		case "messageRetentionInDays":
			err = unpopulate(val, "MessageRetentionInDays", &p.MessageRetentionInDays)
			delete(rawMsg, key)
		case "partitionCount":
			err = unpopulate(val, "PartitionCount", &p.PartitionCount)
			delete(rawMsg, key)
		case "partitionIds":
			err = unpopulate(val, "PartitionIDs", &p.PartitionIDs)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &p.Status)
			delete(rawMsg, key)
		case "updatedAt":
			err = unpopulateTimeRFC3339(val, "UpdatedAt", &p.UpdatedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SchemaGroupProperties.
func (s SchemaGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAtUtc", s.CreatedAtUTC)
	populate(objectMap, "eTag", s.ETag)
	populate(objectMap, "groupProperties", s.GroupProperties)
	populate(objectMap, "schemaCompatibility", s.SchemaCompatibility)
	populate(objectMap, "schemaType", s.SchemaType)
	populateTimeRFC3339(objectMap, "updatedAtUtc", s.UpdatedAtUTC)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SchemaGroupProperties.
func (s *SchemaGroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAtUtc":
			err = unpopulateTimeRFC3339(val, "CreatedAtUTC", &s.CreatedAtUTC)
			delete(rawMsg, key)
		case "eTag":
			err = unpopulate(val, "ETag", &s.ETag)
			delete(rawMsg, key)
		case "groupProperties":
			err = unpopulate(val, "GroupProperties", &s.GroupProperties)
			delete(rawMsg, key)
		case "schemaCompatibility":
			err = unpopulate(val, "SchemaCompatibility", &s.SchemaCompatibility)
			delete(rawMsg, key)
		case "schemaType":
			err = unpopulate(val, "SchemaType", &s.SchemaType)
			delete(rawMsg, key)
		case "updatedAtUtc":
			err = unpopulateTimeRFC3339(val, "UpdatedAtUTC", &s.UpdatedAtUTC)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ThrottlingPolicy.
func (t ThrottlingPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "metricId", t.MetricID)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "rateLimitThreshold", t.RateLimitThreshold)
	objectMap["type"] = ApplicationGroupPolicyTypeThrottlingPolicy
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThrottlingPolicy.
func (t *ThrottlingPolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "metricId":
			err = unpopulate(val, "MetricID", &t.MetricID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "rateLimitThreshold":
			err = unpopulate(val, "RateLimitThreshold", &t.RateLimitThreshold)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &t.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
