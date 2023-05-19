//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azsecrets

// BackupSecretResponse contains the response from method Client.BackupSecret.
type BackupSecretResponse struct {
	BackupSecretResult
}

// DeleteSecretResponse contains the response from method Client.DeleteSecret.
type DeleteSecretResponse struct {
	DeletedSecret
}

// GetDeletedSecretResponse contains the response from method Client.GetDeletedSecret.
type GetDeletedSecretResponse struct {
	DeletedSecret
}

// GetSecretResponse contains the response from method Client.GetSecret.
type GetSecretResponse struct {
	Secret
}

// ListDeletedSecretPropertiesResponse contains the response from method Client.NewListDeletedSecretPropertiesPager.
type ListDeletedSecretPropertiesResponse struct {
	DeletedSecretPropertiesListResult
}

// ListSecretPropertiesResponse contains the response from method Client.NewListSecretPropertiesPager.
type ListSecretPropertiesResponse struct {
	SecretPropertiesListResult
}

// ListSecretPropertiesVersionsResponse contains the response from method Client.NewListSecretPropertiesVersionsPager.
type ListSecretPropertiesVersionsResponse struct {
	SecretPropertiesListResult
}

// PurgeDeletedSecretResponse contains the response from method Client.PurgeDeletedSecret.
type PurgeDeletedSecretResponse struct {
	// placeholder for future response values
}

// RecoverDeletedSecretResponse contains the response from method Client.RecoverDeletedSecret.
type RecoverDeletedSecretResponse struct {
	Secret
}

// RestoreSecretResponse contains the response from method Client.RestoreSecret.
type RestoreSecretResponse struct {
	Secret
}

// SetSecretResponse contains the response from method Client.SetSecret.
type SetSecretResponse struct {
	Secret
}

// UpdateSecretPropertiesResponse contains the response from method Client.UpdateSecretProperties.
type UpdateSecretPropertiesResponse struct {
	Secret
}
