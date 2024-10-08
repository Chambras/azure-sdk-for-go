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

// PoolServer is a fake server for instances of the armbatch.PoolClient type.
type PoolServer struct {
	// Create is the fake for method PoolClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters armbatch.Pool, options *armbatch.PoolClientCreateOptions) (resp azfake.Responder[armbatch.PoolClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PoolClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *armbatch.PoolClientBeginDeleteOptions) (resp azfake.PollerResponder[armbatch.PoolClientDeleteResponse], errResp azfake.ErrorResponder)

	// DisableAutoScale is the fake for method PoolClient.DisableAutoScale
	// HTTP status codes to indicate success: http.StatusOK
	DisableAutoScale func(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *armbatch.PoolClientDisableAutoScaleOptions) (resp azfake.Responder[armbatch.PoolClientDisableAutoScaleResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PoolClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *armbatch.PoolClientGetOptions) (resp azfake.Responder[armbatch.PoolClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByBatchAccountPager is the fake for method PoolClient.NewListByBatchAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBatchAccountPager func(resourceGroupName string, accountName string, options *armbatch.PoolClientListByBatchAccountOptions) (resp azfake.PagerResponder[armbatch.PoolClientListByBatchAccountResponse])

	// StopResize is the fake for method PoolClient.StopResize
	// HTTP status codes to indicate success: http.StatusOK
	StopResize func(ctx context.Context, resourceGroupName string, accountName string, poolName string, options *armbatch.PoolClientStopResizeOptions) (resp azfake.Responder[armbatch.PoolClientStopResizeResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method PoolClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters armbatch.Pool, options *armbatch.PoolClientUpdateOptions) (resp azfake.Responder[armbatch.PoolClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPoolServerTransport creates a new instance of PoolServerTransport with the provided implementation.
// The returned PoolServerTransport instance is connected to an instance of armbatch.PoolClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPoolServerTransport(srv *PoolServer) *PoolServerTransport {
	return &PoolServerTransport{
		srv:                        srv,
		beginDelete:                newTracker[azfake.PollerResponder[armbatch.PoolClientDeleteResponse]](),
		newListByBatchAccountPager: newTracker[azfake.PagerResponder[armbatch.PoolClientListByBatchAccountResponse]](),
	}
}

// PoolServerTransport connects instances of armbatch.PoolClient to instances of PoolServer.
// Don't use this type directly, use NewPoolServerTransport instead.
type PoolServerTransport struct {
	srv                        *PoolServer
	beginDelete                *tracker[azfake.PollerResponder[armbatch.PoolClientDeleteResponse]]
	newListByBatchAccountPager *tracker[azfake.PagerResponder[armbatch.PoolClientListByBatchAccountResponse]]
}

// Do implements the policy.Transporter interface for PoolServerTransport.
func (p *PoolServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PoolClient.Create":
		resp, err = p.dispatchCreate(req)
	case "PoolClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PoolClient.DisableAutoScale":
		resp, err = p.dispatchDisableAutoScale(req)
	case "PoolClient.Get":
		resp, err = p.dispatchGet(req)
	case "PoolClient.NewListByBatchAccountPager":
		resp, err = p.dispatchNewListByBatchAccountPager(req)
	case "PoolClient.StopResize":
		resp, err = p.dispatchStopResize(req)
	case "PoolClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PoolServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if p.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbatch.Pool](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armbatch.PoolClientCreateOptions
	if ifMatchParam != nil || ifNoneMatchParam != nil {
		options = &armbatch.PoolClientCreateOptions{
			IfMatch:     ifMatchParam,
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := p.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pool, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *PoolServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PoolServerTransport) dispatchDisableAutoScale(req *http.Request) (*http.Response, error) {
	if p.srv.DisableAutoScale == nil {
		return nil, &nonRetriableError{errors.New("fake for method DisableAutoScale not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disableAutoScale`
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
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DisableAutoScale(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pool, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *PoolServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pool, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *PoolServerTransport) dispatchNewListByBatchAccountPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByBatchAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBatchAccountPager not implemented")}
	}
	newListByBatchAccountPager := p.newListByBatchAccountPager.get(req)
	if newListByBatchAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
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
		selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
		if err != nil {
			return nil, err
		}
		selectParam := getOptional(selectUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armbatch.PoolClientListByBatchAccountOptions
		if maxresultsParam != nil || selectParam != nil || filterParam != nil {
			options = &armbatch.PoolClientListByBatchAccountOptions{
				Maxresults: maxresultsParam,
				Select:     selectParam,
				Filter:     filterParam,
			}
		}
		resp := p.srv.NewListByBatchAccountPager(resourceGroupNameParam, accountNameParam, options)
		newListByBatchAccountPager = &resp
		p.newListByBatchAccountPager.add(req, newListByBatchAccountPager)
		server.PagerResponderInjectNextLinks(newListByBatchAccountPager, req, func(page *armbatch.PoolClientListByBatchAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBatchAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByBatchAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBatchAccountPager) {
		p.newListByBatchAccountPager.remove(req)
	}
	return resp, nil
}

func (p *PoolServerTransport) dispatchStopResize(req *http.Request) (*http.Response, error) {
	if p.srv.StopResize == nil {
		return nil, &nonRetriableError{errors.New("fake for method StopResize not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stopResize`
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
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.StopResize(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pool, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *PoolServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armbatch.Pool](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armbatch.PoolClientUpdateOptions
	if ifMatchParam != nil {
		options = &armbatch.PoolClientUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameParam, accountNameParam, poolNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pool, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}
