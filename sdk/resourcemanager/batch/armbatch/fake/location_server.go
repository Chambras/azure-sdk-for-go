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
	"strconv"
)

// LocationServer is a fake server for instances of the armbatch.LocationClient type.
type LocationServer struct {
	// CheckNameAvailability is the fake for method LocationClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, locationName string, parameters armbatch.CheckNameAvailabilityParameters, options *armbatch.LocationClientCheckNameAvailabilityOptions) (resp azfake.Responder[armbatch.LocationClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// GetQuotas is the fake for method LocationClient.GetQuotas
	// HTTP status codes to indicate success: http.StatusOK
	GetQuotas func(ctx context.Context, locationName string, options *armbatch.LocationClientGetQuotasOptions) (resp azfake.Responder[armbatch.LocationClientGetQuotasResponse], errResp azfake.ErrorResponder)

	// NewListSupportedVirtualMachineSKUsPager is the fake for method LocationClient.NewListSupportedVirtualMachineSKUsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSupportedVirtualMachineSKUsPager func(locationName string, options *armbatch.LocationClientListSupportedVirtualMachineSKUsOptions) (resp azfake.PagerResponder[armbatch.LocationClientListSupportedVirtualMachineSKUsResponse])
}

// NewLocationServerTransport creates a new instance of LocationServerTransport with the provided implementation.
// The returned LocationServerTransport instance is connected to an instance of armbatch.LocationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationServerTransport(srv *LocationServer) *LocationServerTransport {
	return &LocationServerTransport{
		srv:                                     srv,
		newListSupportedVirtualMachineSKUsPager: newTracker[azfake.PagerResponder[armbatch.LocationClientListSupportedVirtualMachineSKUsResponse]](),
	}
}

// LocationServerTransport connects instances of armbatch.LocationClient to instances of LocationServer.
// Don't use this type directly, use NewLocationServerTransport instead.
type LocationServerTransport struct {
	srv                                     *LocationServer
	newListSupportedVirtualMachineSKUsPager *tracker[azfake.PagerResponder[armbatch.LocationClientListSupportedVirtualMachineSKUsResponse]]
}

// Do implements the policy.Transporter interface for LocationServerTransport.
func (l *LocationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocationClient.CheckNameAvailability":
		resp, err = l.dispatchCheckNameAvailability(req)
	case "LocationClient.GetQuotas":
		resp, err = l.dispatchGetQuotas(req)
	case "LocationClient.NewListSupportedVirtualMachineSKUsPager":
		resp, err = l.dispatchNewListSupportedVirtualMachineSKUsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocationServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if l.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbatch.CheckNameAvailabilityParameters](req)
	if err != nil {
		return nil, err
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.CheckNameAvailability(req.Context(), locationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocationServerTransport) dispatchGetQuotas(req *http.Request) (*http.Response, error) {
	if l.srv.GetQuotas == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetQuotas not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.GetQuotas(req.Context(), locationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LocationQuota, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocationServerTransport) dispatchNewListSupportedVirtualMachineSKUsPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListSupportedVirtualMachineSKUsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSupportedVirtualMachineSKUsPager not implemented")}
	}
	newListSupportedVirtualMachineSKUsPager := l.newListSupportedVirtualMachineSKUsPager.get(req)
	if newListSupportedVirtualMachineSKUsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualMachineSkus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		maxresultsUnescaped, err := url.QueryUnescape(qp.Get("maxresults"))
		if err != nil {
			return nil, err
		}
		maxresultsParam, err := parseOptional(maxresultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armbatch.LocationClientListSupportedVirtualMachineSKUsOptions
		if maxresultsParam != nil || filterParam != nil {
			options = &armbatch.LocationClientListSupportedVirtualMachineSKUsOptions{
				Maxresults: maxresultsParam,
				Filter:     filterParam,
			}
		}
		resp := l.srv.NewListSupportedVirtualMachineSKUsPager(locationNameParam, options)
		newListSupportedVirtualMachineSKUsPager = &resp
		l.newListSupportedVirtualMachineSKUsPager.add(req, newListSupportedVirtualMachineSKUsPager)
		server.PagerResponderInjectNextLinks(newListSupportedVirtualMachineSKUsPager, req, func(page *armbatch.LocationClientListSupportedVirtualMachineSKUsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSupportedVirtualMachineSKUsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListSupportedVirtualMachineSKUsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSupportedVirtualMachineSKUsPager) {
		l.newListSupportedVirtualMachineSKUsPager.remove(req)
	}
	return resp, nil
}
