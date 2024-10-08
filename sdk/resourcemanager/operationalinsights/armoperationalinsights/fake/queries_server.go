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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// QueriesServer is a fake server for instances of the armoperationalinsights.QueriesClient type.
type QueriesServer struct {
	// Delete is the fake for method QueriesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, queryPackName string, id string, options *armoperationalinsights.QueriesClientDeleteOptions) (resp azfake.Responder[armoperationalinsights.QueriesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method QueriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, queryPackName string, id string, options *armoperationalinsights.QueriesClientGetOptions) (resp azfake.Responder[armoperationalinsights.QueriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method QueriesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, queryPackName string, options *armoperationalinsights.QueriesClientListOptions) (resp azfake.PagerResponder[armoperationalinsights.QueriesClientListResponse])

	// Put is the fake for method QueriesClient.Put
	// HTTP status codes to indicate success: http.StatusOK
	Put func(ctx context.Context, resourceGroupName string, queryPackName string, id string, queryPayload armoperationalinsights.LogAnalyticsQueryPackQuery, options *armoperationalinsights.QueriesClientPutOptions) (resp azfake.Responder[armoperationalinsights.QueriesClientPutResponse], errResp azfake.ErrorResponder)

	// NewSearchPager is the fake for method QueriesClient.NewSearchPager
	// HTTP status codes to indicate success: http.StatusOK
	NewSearchPager func(resourceGroupName string, queryPackName string, querySearchProperties armoperationalinsights.LogAnalyticsQueryPackQuerySearchProperties, options *armoperationalinsights.QueriesClientSearchOptions) (resp azfake.PagerResponder[armoperationalinsights.QueriesClientSearchResponse])

	// Update is the fake for method QueriesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, queryPackName string, id string, queryPayload armoperationalinsights.LogAnalyticsQueryPackQuery, options *armoperationalinsights.QueriesClientUpdateOptions) (resp azfake.Responder[armoperationalinsights.QueriesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewQueriesServerTransport creates a new instance of QueriesServerTransport with the provided implementation.
// The returned QueriesServerTransport instance is connected to an instance of armoperationalinsights.QueriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQueriesServerTransport(srv *QueriesServer) *QueriesServerTransport {
	return &QueriesServerTransport{
		srv:            srv,
		newListPager:   newTracker[azfake.PagerResponder[armoperationalinsights.QueriesClientListResponse]](),
		newSearchPager: newTracker[azfake.PagerResponder[armoperationalinsights.QueriesClientSearchResponse]](),
	}
}

// QueriesServerTransport connects instances of armoperationalinsights.QueriesClient to instances of QueriesServer.
// Don't use this type directly, use NewQueriesServerTransport instead.
type QueriesServerTransport struct {
	srv            *QueriesServer
	newListPager   *tracker[azfake.PagerResponder[armoperationalinsights.QueriesClientListResponse]]
	newSearchPager *tracker[azfake.PagerResponder[armoperationalinsights.QueriesClientSearchResponse]]
}

// Do implements the policy.Transporter interface for QueriesServerTransport.
func (q *QueriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "QueriesClient.Delete":
		resp, err = q.dispatchDelete(req)
	case "QueriesClient.Get":
		resp, err = q.dispatchGet(req)
	case "QueriesClient.NewListPager":
		resp, err = q.dispatchNewListPager(req)
	case "QueriesClient.Put":
		resp, err = q.dispatchPut(req)
	case "QueriesClient.NewSearchPager":
		resp, err = q.dispatchNewSearchPager(req)
	case "QueriesClient.Update":
		resp, err = q.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *QueriesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if q.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/queryPacks/(?P<queryPackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	queryPackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryPackName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Delete(req.Context(), resourceGroupNameParam, queryPackNameParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if q.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/queryPacks/(?P<queryPackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	queryPackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryPackName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Get(req.Context(), resourceGroupNameParam, queryPackNameParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LogAnalyticsQueryPackQuery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueriesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if q.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := q.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/queryPacks/(?P<queryPackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries`
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
		queryPackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryPackName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		includeBodyUnescaped, err := url.QueryUnescape(qp.Get("includeBody"))
		if err != nil {
			return nil, err
		}
		includeBodyParam, err := parseOptional(includeBodyUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armoperationalinsights.QueriesClientListOptions
		if topParam != nil || includeBodyParam != nil || skipTokenParam != nil {
			options = &armoperationalinsights.QueriesClientListOptions{
				Top:         topParam,
				IncludeBody: includeBodyParam,
				SkipToken:   skipTokenParam,
			}
		}
		resp := q.srv.NewListPager(resourceGroupNameParam, queryPackNameParam, options)
		newListPager = &resp
		q.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armoperationalinsights.QueriesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		q.newListPager.remove(req)
	}
	return resp, nil
}

func (q *QueriesServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if q.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/queryPacks/(?P<queryPackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.LogAnalyticsQueryPackQuery](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	queryPackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryPackName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Put(req.Context(), resourceGroupNameParam, queryPackNameParam, idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LogAnalyticsQueryPackQuery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueriesServerTransport) dispatchNewSearchPager(req *http.Request) (*http.Response, error) {
	if q.srv.NewSearchPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewSearchPager not implemented")}
	}
	newSearchPager := q.newSearchPager.get(req)
	if newSearchPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/queryPacks/(?P<queryPackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/search`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.LogAnalyticsQueryPackQuerySearchProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		queryPackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryPackName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		includeBodyUnescaped, err := url.QueryUnescape(qp.Get("includeBody"))
		if err != nil {
			return nil, err
		}
		includeBodyParam, err := parseOptional(includeBodyUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armoperationalinsights.QueriesClientSearchOptions
		if topParam != nil || includeBodyParam != nil || skipTokenParam != nil {
			options = &armoperationalinsights.QueriesClientSearchOptions{
				Top:         topParam,
				IncludeBody: includeBodyParam,
				SkipToken:   skipTokenParam,
			}
		}
		resp := q.srv.NewSearchPager(resourceGroupNameParam, queryPackNameParam, body, options)
		newSearchPager = &resp
		q.newSearchPager.add(req, newSearchPager)
		server.PagerResponderInjectNextLinks(newSearchPager, req, func(page *armoperationalinsights.QueriesClientSearchResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newSearchPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newSearchPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newSearchPager) {
		q.newSearchPager.remove(req)
	}
	return resp, nil
}

func (q *QueriesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if q.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/queryPacks/(?P<queryPackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queries/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.LogAnalyticsQueryPackQuery](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	queryPackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("queryPackName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Update(req.Context(), resourceGroupNameParam, queryPackNameParam, idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LogAnalyticsQueryPackQuery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
