//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevops

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PipelineTemplateDefinitionsClientListResponse contains the response from method PipelineTemplateDefinitionsClient.NewListPager.
type PipelineTemplateDefinitionsClientListResponse struct {
	PipelineTemplateDefinitionListResult
}

// PipelinesClientCreateOrUpdateResponse contains the response from method PipelinesClient.BeginCreateOrUpdate.
type PipelinesClientCreateOrUpdateResponse struct {
	Pipeline
}

// PipelinesClientDeleteResponse contains the response from method PipelinesClient.Delete.
type PipelinesClientDeleteResponse struct {
	// placeholder for future response values
}

// PipelinesClientGetResponse contains the response from method PipelinesClient.Get.
type PipelinesClientGetResponse struct {
	Pipeline
}

// PipelinesClientListByResourceGroupResponse contains the response from method PipelinesClient.NewListByResourceGroupPager.
type PipelinesClientListByResourceGroupResponse struct {
	PipelineListResult
}

// PipelinesClientListBySubscriptionResponse contains the response from method PipelinesClient.NewListBySubscriptionPager.
type PipelinesClientListBySubscriptionResponse struct {
	PipelineListResult
}

// PipelinesClientUpdateResponse contains the response from method PipelinesClient.Update.
type PipelinesClientUpdateResponse struct {
	Pipeline
}