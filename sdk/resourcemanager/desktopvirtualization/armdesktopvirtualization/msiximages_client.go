//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MsixImagesClient contains the methods for the MsixImages group.
// Don't use this type directly, use NewMsixImagesClient() instead.
type MsixImagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMsixImagesClient creates a new instance of MsixImagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMsixImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MsixImagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MsixImagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewExpandPager - Expands and Lists MSIX packages in an Image, given the Image Path.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - msixImageURI - Object containing URI to MSIX Image
//   - options - MsixImagesClientExpandOptions contains the optional parameters for the MsixImagesClient.NewExpandPager method.
func (client *MsixImagesClient) NewExpandPager(resourceGroupName string, hostPoolName string, msixImageURI MSIXImageURI, options *MsixImagesClientExpandOptions) *runtime.Pager[MsixImagesClientExpandResponse] {
	return runtime.NewPager(runtime.PagingHandler[MsixImagesClientExpandResponse]{
		More: func(page MsixImagesClientExpandResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MsixImagesClientExpandResponse) (MsixImagesClientExpandResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MsixImagesClient.NewExpandPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.expandCreateRequest(ctx, resourceGroupName, hostPoolName, msixImageURI, options)
			}, nil)
			if err != nil {
				return MsixImagesClientExpandResponse{}, err
			}
			return client.expandHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// expandCreateRequest creates the Expand request.
func (client *MsixImagesClient) expandCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, msixImageURI MSIXImageURI, options *MsixImagesClientExpandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/expandMsixImage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, msixImageURI); err != nil {
		return nil, err
	}
	return req, nil
}

// expandHandleResponse handles the Expand response.
func (client *MsixImagesClient) expandHandleResponse(resp *http.Response) (MsixImagesClientExpandResponse, error) {
	result := MsixImagesClientExpandResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpandMsixImageList); err != nil {
		return MsixImagesClientExpandResponse{}, err
	}
	return result, nil
}
