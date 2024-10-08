//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkSecurityPerimeterServer is a fake server for instances of the armbatch.NetworkSecurityPerimeterClient type.
type NetworkSecurityPerimeterServer struct {
	// GetConfiguration is the fake for method NetworkSecurityPerimeterClient.GetConfiguration
	// HTTP status codes to indicate success: http.StatusOK
	GetConfiguration func(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *armbatch.NetworkSecurityPerimeterClientGetConfigurationOptions) (resp azfake.Responder[armbatch.NetworkSecurityPerimeterClientGetConfigurationResponse], errResp azfake.ErrorResponder)

	// NewListConfigurationsPager is the fake for method NetworkSecurityPerimeterClient.NewListConfigurationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListConfigurationsPager func(resourceGroupName string, accountName string, options *armbatch.NetworkSecurityPerimeterClientListConfigurationsOptions) (resp azfake.PagerResponder[armbatch.NetworkSecurityPerimeterClientListConfigurationsResponse])

	// BeginReconcileConfiguration is the fake for method NetworkSecurityPerimeterClient.BeginReconcileConfiguration
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginReconcileConfiguration func(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *armbatch.NetworkSecurityPerimeterClientBeginReconcileConfigurationOptions) (resp azfake.PollerResponder[armbatch.NetworkSecurityPerimeterClientReconcileConfigurationResponse], errResp azfake.ErrorResponder)
}

// NewNetworkSecurityPerimeterServerTransport creates a new instance of NetworkSecurityPerimeterServerTransport with the provided implementation.
// The returned NetworkSecurityPerimeterServerTransport instance is connected to an instance of armbatch.NetworkSecurityPerimeterClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkSecurityPerimeterServerTransport(srv *NetworkSecurityPerimeterServer) *NetworkSecurityPerimeterServerTransport {
	return &NetworkSecurityPerimeterServerTransport{
		srv:                         srv,
		newListConfigurationsPager:  newTracker[azfake.PagerResponder[armbatch.NetworkSecurityPerimeterClientListConfigurationsResponse]](),
		beginReconcileConfiguration: newTracker[azfake.PollerResponder[armbatch.NetworkSecurityPerimeterClientReconcileConfigurationResponse]](),
	}
}

// NetworkSecurityPerimeterServerTransport connects instances of armbatch.NetworkSecurityPerimeterClient to instances of NetworkSecurityPerimeterServer.
// Don't use this type directly, use NewNetworkSecurityPerimeterServerTransport instead.
type NetworkSecurityPerimeterServerTransport struct {
	srv                         *NetworkSecurityPerimeterServer
	newListConfigurationsPager  *tracker[azfake.PagerResponder[armbatch.NetworkSecurityPerimeterClientListConfigurationsResponse]]
	beginReconcileConfiguration *tracker[azfake.PollerResponder[armbatch.NetworkSecurityPerimeterClientReconcileConfigurationResponse]]
}

// Do implements the policy.Transporter interface for NetworkSecurityPerimeterServerTransport.
func (n *NetworkSecurityPerimeterServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkSecurityPerimeterClient.GetConfiguration":
		resp, err = n.dispatchGetConfiguration(req)
	case "NetworkSecurityPerimeterClient.NewListConfigurationsPager":
		resp, err = n.dispatchNewListConfigurationsPager(req)
	case "NetworkSecurityPerimeterClient.BeginReconcileConfiguration":
		resp, err = n.dispatchBeginReconcileConfiguration(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkSecurityPerimeterServerTransport) dispatchGetConfiguration(req *http.Request) (*http.Response, error) {
	if n.srv.GetConfiguration == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetConfiguration not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<networkSecurityPerimeterConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.GetConfiguration(req.Context(), resourceGroupNameParam, accountNameParam, networkSecurityPerimeterConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSecurityPerimeterConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterServerTransport) dispatchNewListConfigurationsPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListConfigurationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListConfigurationsPager not implemented")}
	}
	newListConfigurationsPager := n.newListConfigurationsPager.get(req)
	if newListConfigurationsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListConfigurationsPager(resourceGroupNameParam, accountNameParam, nil)
		newListConfigurationsPager = &resp
		n.newListConfigurationsPager.add(req, newListConfigurationsPager)
		server.PagerResponderInjectNextLinks(newListConfigurationsPager, req, func(page *armbatch.NetworkSecurityPerimeterClientListConfigurationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListConfigurationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListConfigurationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListConfigurationsPager) {
		n.newListConfigurationsPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterServerTransport) dispatchBeginReconcileConfiguration(req *http.Request) (*http.Response, error) {
	if n.srv.BeginReconcileConfiguration == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReconcileConfiguration not implemented")}
	}
	beginReconcileConfiguration := n.beginReconcileConfiguration.get(req)
	if beginReconcileConfiguration == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<networkSecurityPerimeterConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reconcile`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		networkSecurityPerimeterConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginReconcileConfiguration(req.Context(), resourceGroupNameParam, accountNameParam, networkSecurityPerimeterConfigurationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReconcileConfiguration = &respr
		n.beginReconcileConfiguration.add(req, beginReconcileConfiguration)
	}

	resp, err := server.PollerResponderNext(beginReconcileConfiguration, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		n.beginReconcileConfiguration.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReconcileConfiguration) {
		n.beginReconcileConfiguration.remove(req)
	}

	return resp, nil
}
