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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CodeVersionsServer is a fake server for instances of the armmachinelearning.CodeVersionsClient type.
type CodeVersionsServer struct {
	// CreateOrGetStartPendingUpload is the fake for method CodeVersionsClient.CreateOrGetStartPendingUpload
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrGetStartPendingUpload func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body armmachinelearning.PendingUploadRequestDto, options *armmachinelearning.CodeVersionsClientCreateOrGetStartPendingUploadOptions) (resp azfake.Responder[armmachinelearning.CodeVersionsClientCreateOrGetStartPendingUploadResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method CodeVersionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body armmachinelearning.CodeVersion, options *armmachinelearning.CodeVersionsClientCreateOrUpdateOptions) (resp azfake.Responder[armmachinelearning.CodeVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method CodeVersionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *armmachinelearning.CodeVersionsClientDeleteOptions) (resp azfake.Responder[armmachinelearning.CodeVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CodeVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *armmachinelearning.CodeVersionsClientGetOptions) (resp azfake.Responder[armmachinelearning.CodeVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CodeVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, name string, options *armmachinelearning.CodeVersionsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.CodeVersionsClientListResponse])

	// BeginPublish is the fake for method CodeVersionsClient.BeginPublish
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPublish func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body armmachinelearning.DestinationAsset, options *armmachinelearning.CodeVersionsClientBeginPublishOptions) (resp azfake.PollerResponder[armmachinelearning.CodeVersionsClientPublishResponse], errResp azfake.ErrorResponder)
}

// NewCodeVersionsServerTransport creates a new instance of CodeVersionsServerTransport with the provided implementation.
// The returned CodeVersionsServerTransport instance is connected to an instance of armmachinelearning.CodeVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCodeVersionsServerTransport(srv *CodeVersionsServer) *CodeVersionsServerTransport {
	return &CodeVersionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.CodeVersionsClientListResponse]](),
		beginPublish: newTracker[azfake.PollerResponder[armmachinelearning.CodeVersionsClientPublishResponse]](),
	}
}

// CodeVersionsServerTransport connects instances of armmachinelearning.CodeVersionsClient to instances of CodeVersionsServer.
// Don't use this type directly, use NewCodeVersionsServerTransport instead.
type CodeVersionsServerTransport struct {
	srv          *CodeVersionsServer
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.CodeVersionsClientListResponse]]
	beginPublish *tracker[azfake.PollerResponder[armmachinelearning.CodeVersionsClientPublishResponse]]
}

// Do implements the policy.Transporter interface for CodeVersionsServerTransport.
func (c *CodeVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CodeVersionsClient.CreateOrGetStartPendingUpload":
		resp, err = c.dispatchCreateOrGetStartPendingUpload(req)
	case "CodeVersionsClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "CodeVersionsClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "CodeVersionsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CodeVersionsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "CodeVersionsClient.BeginPublish":
		resp, err = c.dispatchBeginPublish(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CodeVersionsServerTransport) dispatchCreateOrGetStartPendingUpload(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrGetStartPendingUpload == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrGetStartPendingUpload not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codes/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startPendingUpload`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.PendingUploadRequestDto](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrGetStartPendingUpload(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PendingUploadResponseDto, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CodeVersionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codes/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.CodeVersion](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CodeVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CodeVersionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codes/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, nil)
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

func (c *CodeVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codes/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CodeVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CodeVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codes/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		hashUnescaped, err := url.QueryUnescape(qp.Get("hash"))
		if err != nil {
			return nil, err
		}
		hashParam := getOptional(hashUnescaped)
		hashVersionUnescaped, err := url.QueryUnescape(qp.Get("hashVersion"))
		if err != nil {
			return nil, err
		}
		hashVersionParam := getOptional(hashVersionUnescaped)
		var options *armmachinelearning.CodeVersionsClientListOptions
		if orderByParam != nil || topParam != nil || skipParam != nil || hashParam != nil || hashVersionParam != nil {
			options = &armmachinelearning.CodeVersionsClientListOptions{
				OrderBy:     orderByParam,
				Top:         topParam,
				Skip:        skipParam,
				Hash:        hashParam,
				HashVersion: hashVersionParam,
			}
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nameParam, options)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.CodeVersionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *CodeVersionsServerTransport) dispatchBeginPublish(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPublish == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPublish not implemented")}
	}
	beginPublish := c.beginPublish.get(req)
	if beginPublish == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codes/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publish`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.DestinationAsset](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginPublish(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPublish = &respr
		c.beginPublish.add(req, beginPublish)
	}

	resp, err := server.PollerResponderNext(beginPublish, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginPublish.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPublish) {
		c.beginPublish.remove(req)
	}

	return resp, nil
}
