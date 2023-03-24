//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise

// ClientCreateResponse contains the response from method Client.BeginCreate.
type ClientCreateResponse struct {
	Cluster
}

// ClientDeleteResponse contains the response from method Client.BeginDelete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	Cluster
}

// ClientListByResourceGroupResponse contains the response from method Client.NewListByResourceGroupPager.
type ClientListByResourceGroupResponse struct {
	ClusterList
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	ClusterList
}

// ClientUpdateResponse contains the response from method Client.BeginUpdate.
type ClientUpdateResponse struct {
	Cluster
}

// DatabasesClientCreateResponse contains the response from method DatabasesClient.BeginCreate.
type DatabasesClientCreateResponse struct {
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.BeginDelete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientExportResponse contains the response from method DatabasesClient.BeginExport.
type DatabasesClientExportResponse struct {
	// placeholder for future response values
}

// DatabasesClientFlushResponse contains the response from method DatabasesClient.BeginFlush.
type DatabasesClientFlushResponse struct {
	// placeholder for future response values
}

// DatabasesClientForceUnlinkResponse contains the response from method DatabasesClient.BeginForceUnlink.
type DatabasesClientForceUnlinkResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	Database
}

// DatabasesClientImportResponse contains the response from method DatabasesClient.BeginImport.
type DatabasesClientImportResponse struct {
	// placeholder for future response values
}

// DatabasesClientListByClusterResponse contains the response from method DatabasesClient.NewListByClusterPager.
type DatabasesClientListByClusterResponse struct {
	DatabaseList
}

// DatabasesClientListKeysResponse contains the response from method DatabasesClient.ListKeys.
type DatabasesClientListKeysResponse struct {
	AccessKeys
}

// DatabasesClientRegenerateKeyResponse contains the response from method DatabasesClient.BeginRegenerateKey.
type DatabasesClientRegenerateKeyResponse struct {
	AccessKeys
}

// DatabasesClientUpdateResponse contains the response from method DatabasesClient.BeginUpdate.
type DatabasesClientUpdateResponse struct {
	Database
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// OperationsStatusClientGetResponse contains the response from method OperationsStatusClient.Get.
type OperationsStatusClientGetResponse struct {
	OperationStatus
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientPutResponse contains the response from method PrivateEndpointConnectionsClient.BeginPut.
type PrivateEndpointConnectionsClientPutResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListByClusterResponse contains the response from method PrivateLinkResourcesClient.NewListByClusterPager.
type PrivateLinkResourcesClientListByClusterResponse struct {
	PrivateLinkResourceListResult
}

// SKUsClientListResponse contains the response from method SKUsClient.NewListPager.
type SKUsClientListResponse struct {
	RegionSKUDetails
}