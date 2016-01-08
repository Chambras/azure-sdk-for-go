package intune

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// IosClient is the microsoft.Intune Resource provider Api features in the
// swagger-2.0 specification
type IosClient struct {
	ManagementClient
}

// NewIosClient creates an instance of the IosClient client.
func NewIosClient(subscriptionID string) IosClient {
	return NewIosClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIosClientWithBaseURI creates an instance of the IosClient client.
func NewIosClientWithBaseURI(baseURI string, subscriptionID string) IosClient {
	return IosClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AddAppForMAMPolicy add app to an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy appName is application unique Name parameters is parameters
// supplied to add an app to an ios policy.
func (client IosClient) AddAppForMAMPolicy(hostName string, policyName string, appName string, parameters MAMPolicyAppIDOrGroupIDPayload) (result autorest.Response, ae error) {
	req, err := client.AddAppForMAMPolicyPreparer(hostName, policyName, appName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddAppForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.AddAppForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddAppForMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.AddAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "AddAppForMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// AddAppForMAMPolicyPreparer prepares the AddAppForMAMPolicy request.
func (client IosClient) AddAppForMAMPolicyPreparer(hostName string, policyName string, appName string, parameters MAMPolicyAppIDOrGroupIDPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":    url.QueryEscape(appName),
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/apps/{appName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// AddAppForMAMPolicySender sends the AddAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) AddAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// AddAppForMAMPolicyResponder handles the response to the AddAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) AddAppForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// AddGroupForMAMPolicy add group to an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy groupID is group Id parameters is parameters supplied to the
// Create or update app to an android policy operation.
func (client IosClient) AddGroupForMAMPolicy(hostName string, policyName string, groupID string, parameters MAMPolicyAppIDOrGroupIDPayload) (result autorest.Response, ae error) {
	req, err := client.AddGroupForMAMPolicyPreparer(hostName, policyName, groupID, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddGroupForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.AddGroupForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "AddGroupForMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.AddGroupForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "AddGroupForMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// AddGroupForMAMPolicyPreparer prepares the AddGroupForMAMPolicy request.
func (client IosClient) AddGroupForMAMPolicyPreparer(hostName string, policyName string, groupID string, parameters MAMPolicyAppIDOrGroupIDPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":    url.QueryEscape(groupID),
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/groups/{groupId}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// AddGroupForMAMPolicySender sends the AddGroupForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) AddGroupForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// AddGroupForMAMPolicyResponder handles the response to the AddGroupForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) AddGroupForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdateMAMPolicy creates or updates iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy parameters is parameters supplied to the Create or update an
// android policy operation.
func (client IosClient) CreateOrUpdateMAMPolicy(hostName string, policyName string, parameters IOSMAMPolicy) (result IOSMAMPolicy, ae error) {
	req, err := client.CreateOrUpdateMAMPolicyPreparer(hostName, policyName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "CreateOrUpdateMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "CreateOrUpdateMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.CreateOrUpdateMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "CreateOrUpdateMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// CreateOrUpdateMAMPolicyPreparer prepares the CreateOrUpdateMAMPolicy request.
func (client IosClient) CreateOrUpdateMAMPolicyPreparer(hostName string, policyName string, parameters IOSMAMPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateMAMPolicySender sends the CreateOrUpdateMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) CreateOrUpdateMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// CreateOrUpdateMAMPolicyResponder handles the response to the CreateOrUpdateMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) CreateOrUpdateMAMPolicyResponder(resp *http.Response) (result IOSMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAppForMAMPolicy delete App for Ios Policy
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy appName is application unique Name
func (client IosClient) DeleteAppForMAMPolicy(hostName string, policyName string, appName string) (result autorest.Response, ae error) {
	req, err := client.DeleteAppForMAMPolicyPreparer(hostName, policyName, appName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteAppForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.DeleteAppForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteAppForMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.DeleteAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "DeleteAppForMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// DeleteAppForMAMPolicyPreparer prepares the DeleteAppForMAMPolicy request.
func (client IosClient) DeleteAppForMAMPolicyPreparer(hostName string, policyName string, appName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appName":    url.QueryEscape(appName),
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/apps/{appName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteAppForMAMPolicySender sends the DeleteAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) DeleteAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteAppForMAMPolicyResponder handles the response to the DeleteAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) DeleteAppForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteGroupForMAMPolicy delete Group for iOS Policy
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy groupID is application unique Name
func (client IosClient) DeleteGroupForMAMPolicy(hostName string, policyName string, groupID string) (result autorest.Response, ae error) {
	req, err := client.DeleteGroupForMAMPolicyPreparer(hostName, policyName, groupID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteGroupForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.DeleteGroupForMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteGroupForMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.DeleteGroupForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "DeleteGroupForMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// DeleteGroupForMAMPolicyPreparer prepares the DeleteGroupForMAMPolicy request.
func (client IosClient) DeleteGroupForMAMPolicyPreparer(hostName string, policyName string, groupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":    url.QueryEscape(groupID),
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/groups/{groupId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteGroupForMAMPolicySender sends the DeleteGroupForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) DeleteGroupForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteGroupForMAMPolicyResponder handles the response to the DeleteGroupForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) DeleteGroupForMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteMAMPolicy delete Ios Policy
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy
func (client IosClient) DeleteMAMPolicy(hostName string, policyName string) (result autorest.Response, ae error) {
	req, err := client.DeleteMAMPolicyPreparer(hostName, policyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.DeleteMAMPolicySender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "DeleteMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.DeleteMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "DeleteMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// DeleteMAMPolicyPreparer prepares the DeleteMAMPolicy request.
func (client IosClient) DeleteMAMPolicyPreparer(hostName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteMAMPolicySender sends the DeleteMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) DeleteMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteMAMPolicyResponder handles the response to the DeleteMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) DeleteMAMPolicyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetAppForMAMPolicy get apps for an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy filter is the filter to apply on the operation. selectParameter
// is select specific fields in entity.
func (client IosClient) GetAppForMAMPolicy(hostName string, policyName string, filter string, top *int, selectParameter string) (result ApplicationCollection, ae error) {
	req, err := client.GetAppForMAMPolicyPreparer(hostName, policyName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetAppForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetAppForMAMPolicyPreparer prepares the GetAppForMAMPolicy request.
func (client IosClient) GetAppForMAMPolicyPreparer(hostName string, policyName string, filter string, top *int, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/apps"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAppForMAMPolicySender sends the GetAppForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetAppForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetAppForMAMPolicyResponder handles the response to the GetAppForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) GetAppForMAMPolicyResponder(resp *http.Response) (result ApplicationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAppForMAMPolicyNextResults retrieves the next set of results, if any.
func (client IosClient) GetAppForMAMPolicyNextResults(lastResults ApplicationCollection) (result ApplicationCollection, ae error) {
	req, err := lastResults.ApplicationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetAppForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", resp.StatusCode, "Failure sending next results request request")
	}

	result, err = client.GetAppForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetAppForMAMPolicy", resp.StatusCode, "Failure responding to next results request request")
	}

	return
}

// GetGroupsForMAMPolicy returns groups for a given iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyName is policy name for
// the tenant
func (client IosClient) GetGroupsForMAMPolicy(hostName string, policyName string) (result GroupsCollection, ae error) {
	req, err := client.GetGroupsForMAMPolicyPreparer(hostName, policyName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetGroupsForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetGroupsForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetGroupsForMAMPolicyPreparer prepares the GetGroupsForMAMPolicy request.
func (client IosClient) GetGroupsForMAMPolicyPreparer(hostName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/groups"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetGroupsForMAMPolicySender sends the GetGroupsForMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetGroupsForMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetGroupsForMAMPolicyResponder handles the response to the GetGroupsForMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) GetGroupsForMAMPolicyResponder(resp *http.Response) (result GroupsCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetGroupsForMAMPolicyNextResults retrieves the next set of results, if any.
func (client IosClient) GetGroupsForMAMPolicyNextResults(lastResults GroupsCollection) (result GroupsCollection, ae error) {
	req, err := lastResults.GroupsCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetGroupsForMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", resp.StatusCode, "Failure sending next results request request")
	}

	result, err = client.GetGroupsForMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetGroupsForMAMPolicy", resp.StatusCode, "Failure responding to next results request request")
	}

	return
}

// GetMAMPolicies returns Intune iOSPolicies.
//
// hostName is location hostName for the tenant filter is the filter to apply
// on the operation. selectParameter is select specific fields in entity.
func (client IosClient) GetMAMPolicies(hostName string, filter string, top *int, selectParameter string) (result IOSMAMPolicyCollection, ae error) {
	req, err := client.GetMAMPoliciesPreparer(hostName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetMAMPoliciesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetMAMPoliciesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetMAMPoliciesPreparer prepares the GetMAMPolicies request.
func (client IosClient) GetMAMPoliciesPreparer(hostName string, filter string, top *int, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMPoliciesSender sends the GetMAMPolicies request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetMAMPoliciesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetMAMPoliciesResponder handles the response to the GetMAMPolicies request. The method always
// closes the http.Response Body.
func (client IosClient) GetMAMPoliciesResponder(resp *http.Response) (result IOSMAMPolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMPoliciesNextResults retrieves the next set of results, if any.
func (client IosClient) GetMAMPoliciesNextResults(lastResults IOSMAMPolicyCollection) (result IOSMAMPolicyCollection, ae error) {
	req, err := lastResults.IOSMAMPolicyCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", autorest.UndefinedStatusCode, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetMAMPoliciesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", resp.StatusCode, "Failure sending next results request request")
	}

	result, err = client.GetMAMPoliciesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicies", resp.StatusCode, "Failure responding to next results request request")
	}

	return
}

// GetMAMPolicyByName returns Intune iOS policies.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy selectParameter is select specific fields in entity.
func (client IosClient) GetMAMPolicyByName(hostName string, policyName string, selectParameter string) (result IOSMAMPolicy, ae error) {
	req, err := client.GetMAMPolicyByNamePreparer(hostName, policyName, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicyByName", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetMAMPolicyByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicyByName", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetMAMPolicyByNameResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "GetMAMPolicyByName", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetMAMPolicyByNamePreparer prepares the GetMAMPolicyByName request.
func (client IosClient) GetMAMPolicyByNamePreparer(hostName string, policyName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMPolicyByNameSender sends the GetMAMPolicyByName request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) GetMAMPolicyByNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetMAMPolicyByNameResponder handles the response to the GetMAMPolicyByName request. The method always
// closes the http.Response Body.
func (client IosClient) GetMAMPolicyByNameResponder(resp *http.Response) (result IOSMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PatchMAMPolicy patch an iOSMAMPolicy.
//
// hostName is location hostName for the tenant policyName is unique name for
// the policy parameters is parameters supplied to the Create or update an
// android policy operation.
func (client IosClient) PatchMAMPolicy(hostName string, policyName string, parameters IOSMAMPolicy) (result IOSMAMPolicy, ae error) {
	req, err := client.PatchMAMPolicyPreparer(hostName, policyName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "PatchMAMPolicy", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.PatchMAMPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/IosClient", "PatchMAMPolicy", resp.StatusCode, "Failure sending request")
	}

	result, err = client.PatchMAMPolicyResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "intune/IosClient", "PatchMAMPolicy", resp.StatusCode, "Failure responding to request")
	}

	return
}

// PatchMAMPolicyPreparer prepares the PatchMAMPolicy request.
func (client IosClient) PatchMAMPolicyPreparer(hostName string, policyName string, parameters IOSMAMPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName":   url.QueryEscape(hostName),
		"policyName": url.QueryEscape(policyName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// PatchMAMPolicySender sends the PatchMAMPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client IosClient) PatchMAMPolicySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// PatchMAMPolicyResponder handles the response to the PatchMAMPolicy request. The method always
// closes the http.Response Body.
func (client IosClient) PatchMAMPolicyResponder(resp *http.Response) (result IOSMAMPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
