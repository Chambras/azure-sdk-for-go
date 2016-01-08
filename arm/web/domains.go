package web

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

// DomainsClient is the use these APIs to manage Azure Websites resources
// through the Azure Resource Manager. All task operations conform to the
// HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type DomainsClient struct {
	ManagementClient
}

// NewDomainsClient creates an instance of the DomainsClient client.
func NewDomainsClient(subscriptionID string) DomainsClient {
	return NewDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDomainsClientWithBaseURI creates an instance of the DomainsClient client.
func NewDomainsClientWithBaseURI(baseURI string, subscriptionID string) DomainsClient {
	return DomainsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateDomain sends the create or update domain request.
//
// resourceGroupName is &gt;Name of the resource group domainName is name of
// the domain domain is domain registration information
func (client DomainsClient) CreateOrUpdateDomain(resourceGroupName string, domainName string, domain Domain) (result Domain, ae error) {
	req, err := client.CreateOrUpdateDomainPreparer(resourceGroupName, domainName, domain)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "CreateOrUpdateDomain", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "CreateOrUpdateDomain", resp.StatusCode, "Failure sending request")
	}

	result, err = client.CreateOrUpdateDomainResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/DomainsClient", "CreateOrUpdateDomain", resp.StatusCode, "Failure responding to request")
	}

	return
}

// CreateOrUpdateDomainPreparer prepares the CreateOrUpdateDomain request.
func (client DomainsClient) CreateOrUpdateDomainPreparer(resourceGroupName string, domainName string, domain Domain) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        url.QueryEscape(domainName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}"),
		autorest.WithJSON(domain),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateDomainSender sends the CreateOrUpdateDomain request. The method will close the
// http.Response Body if it receives an error.
func (client DomainsClient) CreateOrUpdateDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusAccepted, http.StatusOK)
}

// CreateOrUpdateDomainResponder handles the response to the CreateOrUpdateDomain request. The method always
// closes the http.Response Body.
func (client DomainsClient) CreateOrUpdateDomainResponder(resp *http.Response) (result Domain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteDomain sends the delete domain request.
//
// resourceGroupName is name of the resource group domainName is name of the
// domain forceHardDeleteDomain is if true then the domain will be deleted
// immediately instead of after 24 hours
func (client DomainsClient) DeleteDomain(resourceGroupName string, domainName string, forceHardDeleteDomain *bool) (result ObjectSet, ae error) {
	req, err := client.DeleteDomainPreparer(resourceGroupName, domainName, forceHardDeleteDomain)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "DeleteDomain", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.DeleteDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "DeleteDomain", resp.StatusCode, "Failure sending request")
	}

	result, err = client.DeleteDomainResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/DomainsClient", "DeleteDomain", resp.StatusCode, "Failure responding to request")
	}

	return
}

// DeleteDomainPreparer prepares the DeleteDomain request.
func (client DomainsClient) DeleteDomainPreparer(resourceGroupName string, domainName string, forceHardDeleteDomain *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        url.QueryEscape(domainName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if forceHardDeleteDomain != nil {
		queryParameters["forceHardDeleteDomain"] = forceHardDeleteDomain
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteDomainSender sends the DeleteDomain request. The method will close the
// http.Response Body if it receives an error.
func (client DomainsClient) DeleteDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteDomainResponder handles the response to the DeleteDomain request. The method always
// closes the http.Response Body.
func (client DomainsClient) DeleteDomainResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDomain sends the get domain request.
//
// resourceGroupName is name of the resource group domainName is name of the
// domain
func (client DomainsClient) GetDomain(resourceGroupName string, domainName string) (result Domain, ae error) {
	req, err := client.GetDomainPreparer(resourceGroupName, domainName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomain", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomain", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetDomainResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomain", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetDomainPreparer prepares the GetDomain request.
func (client DomainsClient) GetDomainPreparer(resourceGroupName string, domainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        url.QueryEscape(domainName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetDomainSender sends the GetDomain request. The method will close the
// http.Response Body if it receives an error.
func (client DomainsClient) GetDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetDomainResponder handles the response to the GetDomain request. The method always
// closes the http.Response Body.
func (client DomainsClient) GetDomainResponder(resp *http.Response) (result Domain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDomainOperation sends the get domain operation request.
//
// resourceGroupName is name of the resource group domainName is name of the
// domain operationID is domain purchase operation Id
func (client DomainsClient) GetDomainOperation(resourceGroupName string, domainName string, operationID string) (result Domain, ae error) {
	req, err := client.GetDomainOperationPreparer(resourceGroupName, domainName, operationID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomainOperation", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetDomainOperationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomainOperation", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetDomainOperationResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomainOperation", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetDomainOperationPreparer prepares the GetDomainOperation request.
func (client DomainsClient) GetDomainOperationPreparer(resourceGroupName string, domainName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        url.QueryEscape(domainName),
		"operationId":       url.QueryEscape(operationID),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/operationresults/{operationId}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetDomainOperationSender sends the GetDomainOperation request. The method will close the
// http.Response Body if it receives an error.
func (client DomainsClient) GetDomainOperationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusAccepted, http.StatusOK, http.StatusInternalServerError)
}

// GetDomainOperationResponder handles the response to the GetDomainOperation request. The method always
// closes the http.Response Body.
func (client DomainsClient) GetDomainOperationResponder(resp *http.Response) (result Domain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDomains sends the get domains request.
//
// resourceGroupName is name of the resource group
func (client DomainsClient) GetDomains(resourceGroupName string) (result DomainCollection, ae error) {
	req, err := client.GetDomainsPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomains", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.GetDomainsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomains", resp.StatusCode, "Failure sending request")
	}

	result, err = client.GetDomainsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/DomainsClient", "GetDomains", resp.StatusCode, "Failure responding to request")
	}

	return
}

// GetDomainsPreparer prepares the GetDomains request.
func (client DomainsClient) GetDomainsPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetDomainsSender sends the GetDomains request. The method will close the
// http.Response Body if it receives an error.
func (client DomainsClient) GetDomainsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetDomainsResponder handles the response to the GetDomains request. The method always
// closes the http.Response Body.
func (client DomainsClient) GetDomainsResponder(resp *http.Response) (result DomainCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateDomain sends the update domain request.
//
// resourceGroupName is &gt;Name of the resource group domainName is name of
// the domain domain is domain registration information
func (client DomainsClient) UpdateDomain(resourceGroupName string, domainName string, domain Domain) (result Domain, ae error) {
	req, err := client.UpdateDomainPreparer(resourceGroupName, domainName, domain)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "UpdateDomain", autorest.UndefinedStatusCode, "Failure preparing request")
	}

	resp, err := client.UpdateDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/DomainsClient", "UpdateDomain", resp.StatusCode, "Failure sending request")
	}

	result, err = client.UpdateDomainResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "web/DomainsClient", "UpdateDomain", resp.StatusCode, "Failure responding to request")
	}

	return
}

// UpdateDomainPreparer prepares the UpdateDomain request.
func (client DomainsClient) UpdateDomainPreparer(resourceGroupName string, domainName string, domain Domain) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        url.QueryEscape(domainName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}"),
		autorest.WithJSON(domain),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// UpdateDomainSender sends the UpdateDomain request. The method will close the
// http.Response Body if it receives an error.
func (client DomainsClient) UpdateDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusAccepted, http.StatusOK)
}

// UpdateDomainResponder handles the response to the UpdateDomain request. The method always
// closes the http.Response Body.
func (client DomainsClient) UpdateDomainResponder(resp *http.Response) (result Domain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
