//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armsecurityinsights.ClientFactory type.
type ServerFactory struct {
	ActionsServer                            ActionsServer
	AlertRuleTemplatesServer                 AlertRuleTemplatesServer
	AlertRulesServer                         AlertRulesServer
	AutomationRulesServer                    AutomationRulesServer
	BookmarkServer                           BookmarkServer
	BookmarkRelationsServer                  BookmarkRelationsServer
	BookmarksServer                          BookmarksServer
	DataConnectorsCheckRequirementsServer    DataConnectorsCheckRequirementsServer
	DataConnectorsServer                     DataConnectorsServer
	DomainWhoisServer                        DomainWhoisServer
	EntitiesServer                           EntitiesServer
	EntitiesGetTimelineServer                EntitiesGetTimelineServer
	EntitiesRelationsServer                  EntitiesRelationsServer
	EntityQueriesServer                      EntityQueriesServer
	EntityQueryTemplatesServer               EntityQueryTemplatesServer
	EntityRelationsServer                    EntityRelationsServer
	FileImportsServer                        FileImportsServer
	IPGeodataServer                          IPGeodataServer
	IncidentCommentsServer                   IncidentCommentsServer
	IncidentRelationsServer                  IncidentRelationsServer
	IncidentsServer                          IncidentsServer
	MetadataServer                           MetadataServer
	OfficeConsentsServer                     OfficeConsentsServer
	OperationsServer                         OperationsServer
	ProductSettingsServer                    ProductSettingsServer
	SecurityMLAnalyticsSettingsServer        SecurityMLAnalyticsSettingsServer
	SentinelOnboardingStatesServer           SentinelOnboardingStatesServer
	SourceControlServer                      SourceControlServer
	SourceControlsServer                     SourceControlsServer
	ThreatIntelligenceIndicatorServer        ThreatIntelligenceIndicatorServer
	ThreatIntelligenceIndicatorMetricsServer ThreatIntelligenceIndicatorMetricsServer
	ThreatIntelligenceIndicatorsServer       ThreatIntelligenceIndicatorsServer
	WatchlistItemsServer                     WatchlistItemsServer
	WatchlistsServer                         WatchlistsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armsecurityinsights.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armsecurityinsights.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                        *ServerFactory
	trMu                                       sync.Mutex
	trActionsServer                            *ActionsServerTransport
	trAlertRuleTemplatesServer                 *AlertRuleTemplatesServerTransport
	trAlertRulesServer                         *AlertRulesServerTransport
	trAutomationRulesServer                    *AutomationRulesServerTransport
	trBookmarkServer                           *BookmarkServerTransport
	trBookmarkRelationsServer                  *BookmarkRelationsServerTransport
	trBookmarksServer                          *BookmarksServerTransport
	trDataConnectorsCheckRequirementsServer    *DataConnectorsCheckRequirementsServerTransport
	trDataConnectorsServer                     *DataConnectorsServerTransport
	trDomainWhoisServer                        *DomainWhoisServerTransport
	trEntitiesServer                           *EntitiesServerTransport
	trEntitiesGetTimelineServer                *EntitiesGetTimelineServerTransport
	trEntitiesRelationsServer                  *EntitiesRelationsServerTransport
	trEntityQueriesServer                      *EntityQueriesServerTransport
	trEntityQueryTemplatesServer               *EntityQueryTemplatesServerTransport
	trEntityRelationsServer                    *EntityRelationsServerTransport
	trFileImportsServer                        *FileImportsServerTransport
	trIPGeodataServer                          *IPGeodataServerTransport
	trIncidentCommentsServer                   *IncidentCommentsServerTransport
	trIncidentRelationsServer                  *IncidentRelationsServerTransport
	trIncidentsServer                          *IncidentsServerTransport
	trMetadataServer                           *MetadataServerTransport
	trOfficeConsentsServer                     *OfficeConsentsServerTransport
	trOperationsServer                         *OperationsServerTransport
	trProductSettingsServer                    *ProductSettingsServerTransport
	trSecurityMLAnalyticsSettingsServer        *SecurityMLAnalyticsSettingsServerTransport
	trSentinelOnboardingStatesServer           *SentinelOnboardingStatesServerTransport
	trSourceControlServer                      *SourceControlServerTransport
	trSourceControlsServer                     *SourceControlsServerTransport
	trThreatIntelligenceIndicatorServer        *ThreatIntelligenceIndicatorServerTransport
	trThreatIntelligenceIndicatorMetricsServer *ThreatIntelligenceIndicatorMetricsServerTransport
	trThreatIntelligenceIndicatorsServer       *ThreatIntelligenceIndicatorsServerTransport
	trWatchlistItemsServer                     *WatchlistItemsServerTransport
	trWatchlistsServer                         *WatchlistsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ActionsClient":
		initServer(s, &s.trActionsServer, func() *ActionsServerTransport { return NewActionsServerTransport(&s.srv.ActionsServer) })
		resp, err = s.trActionsServer.Do(req)
	case "AlertRuleTemplatesClient":
		initServer(s, &s.trAlertRuleTemplatesServer, func() *AlertRuleTemplatesServerTransport {
			return NewAlertRuleTemplatesServerTransport(&s.srv.AlertRuleTemplatesServer)
		})
		resp, err = s.trAlertRuleTemplatesServer.Do(req)
	case "AlertRulesClient":
		initServer(s, &s.trAlertRulesServer, func() *AlertRulesServerTransport { return NewAlertRulesServerTransport(&s.srv.AlertRulesServer) })
		resp, err = s.trAlertRulesServer.Do(req)
	case "AutomationRulesClient":
		initServer(s, &s.trAutomationRulesServer, func() *AutomationRulesServerTransport {
			return NewAutomationRulesServerTransport(&s.srv.AutomationRulesServer)
		})
		resp, err = s.trAutomationRulesServer.Do(req)
	case "BookmarkClient":
		initServer(s, &s.trBookmarkServer, func() *BookmarkServerTransport { return NewBookmarkServerTransport(&s.srv.BookmarkServer) })
		resp, err = s.trBookmarkServer.Do(req)
	case "BookmarkRelationsClient":
		initServer(s, &s.trBookmarkRelationsServer, func() *BookmarkRelationsServerTransport {
			return NewBookmarkRelationsServerTransport(&s.srv.BookmarkRelationsServer)
		})
		resp, err = s.trBookmarkRelationsServer.Do(req)
	case "BookmarksClient":
		initServer(s, &s.trBookmarksServer, func() *BookmarksServerTransport { return NewBookmarksServerTransport(&s.srv.BookmarksServer) })
		resp, err = s.trBookmarksServer.Do(req)
	case "DataConnectorsCheckRequirementsClient":
		initServer(s, &s.trDataConnectorsCheckRequirementsServer, func() *DataConnectorsCheckRequirementsServerTransport {
			return NewDataConnectorsCheckRequirementsServerTransport(&s.srv.DataConnectorsCheckRequirementsServer)
		})
		resp, err = s.trDataConnectorsCheckRequirementsServer.Do(req)
	case "DataConnectorsClient":
		initServer(s, &s.trDataConnectorsServer, func() *DataConnectorsServerTransport {
			return NewDataConnectorsServerTransport(&s.srv.DataConnectorsServer)
		})
		resp, err = s.trDataConnectorsServer.Do(req)
	case "DomainWhoisClient":
		initServer(s, &s.trDomainWhoisServer, func() *DomainWhoisServerTransport { return NewDomainWhoisServerTransport(&s.srv.DomainWhoisServer) })
		resp, err = s.trDomainWhoisServer.Do(req)
	case "EntitiesClient":
		initServer(s, &s.trEntitiesServer, func() *EntitiesServerTransport { return NewEntitiesServerTransport(&s.srv.EntitiesServer) })
		resp, err = s.trEntitiesServer.Do(req)
	case "EntitiesGetTimelineClient":
		initServer(s, &s.trEntitiesGetTimelineServer, func() *EntitiesGetTimelineServerTransport {
			return NewEntitiesGetTimelineServerTransport(&s.srv.EntitiesGetTimelineServer)
		})
		resp, err = s.trEntitiesGetTimelineServer.Do(req)
	case "EntitiesRelationsClient":
		initServer(s, &s.trEntitiesRelationsServer, func() *EntitiesRelationsServerTransport {
			return NewEntitiesRelationsServerTransport(&s.srv.EntitiesRelationsServer)
		})
		resp, err = s.trEntitiesRelationsServer.Do(req)
	case "EntityQueriesClient":
		initServer(s, &s.trEntityQueriesServer, func() *EntityQueriesServerTransport {
			return NewEntityQueriesServerTransport(&s.srv.EntityQueriesServer)
		})
		resp, err = s.trEntityQueriesServer.Do(req)
	case "EntityQueryTemplatesClient":
		initServer(s, &s.trEntityQueryTemplatesServer, func() *EntityQueryTemplatesServerTransport {
			return NewEntityQueryTemplatesServerTransport(&s.srv.EntityQueryTemplatesServer)
		})
		resp, err = s.trEntityQueryTemplatesServer.Do(req)
	case "EntityRelationsClient":
		initServer(s, &s.trEntityRelationsServer, func() *EntityRelationsServerTransport {
			return NewEntityRelationsServerTransport(&s.srv.EntityRelationsServer)
		})
		resp, err = s.trEntityRelationsServer.Do(req)
	case "FileImportsClient":
		initServer(s, &s.trFileImportsServer, func() *FileImportsServerTransport { return NewFileImportsServerTransport(&s.srv.FileImportsServer) })
		resp, err = s.trFileImportsServer.Do(req)
	case "IPGeodataClient":
		initServer(s, &s.trIPGeodataServer, func() *IPGeodataServerTransport { return NewIPGeodataServerTransport(&s.srv.IPGeodataServer) })
		resp, err = s.trIPGeodataServer.Do(req)
	case "IncidentCommentsClient":
		initServer(s, &s.trIncidentCommentsServer, func() *IncidentCommentsServerTransport {
			return NewIncidentCommentsServerTransport(&s.srv.IncidentCommentsServer)
		})
		resp, err = s.trIncidentCommentsServer.Do(req)
	case "IncidentRelationsClient":
		initServer(s, &s.trIncidentRelationsServer, func() *IncidentRelationsServerTransport {
			return NewIncidentRelationsServerTransport(&s.srv.IncidentRelationsServer)
		})
		resp, err = s.trIncidentRelationsServer.Do(req)
	case "IncidentsClient":
		initServer(s, &s.trIncidentsServer, func() *IncidentsServerTransport { return NewIncidentsServerTransport(&s.srv.IncidentsServer) })
		resp, err = s.trIncidentsServer.Do(req)
	case "MetadataClient":
		initServer(s, &s.trMetadataServer, func() *MetadataServerTransport { return NewMetadataServerTransport(&s.srv.MetadataServer) })
		resp, err = s.trMetadataServer.Do(req)
	case "OfficeConsentsClient":
		initServer(s, &s.trOfficeConsentsServer, func() *OfficeConsentsServerTransport {
			return NewOfficeConsentsServerTransport(&s.srv.OfficeConsentsServer)
		})
		resp, err = s.trOfficeConsentsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ProductSettingsClient":
		initServer(s, &s.trProductSettingsServer, func() *ProductSettingsServerTransport {
			return NewProductSettingsServerTransport(&s.srv.ProductSettingsServer)
		})
		resp, err = s.trProductSettingsServer.Do(req)
	case "SecurityMLAnalyticsSettingsClient":
		initServer(s, &s.trSecurityMLAnalyticsSettingsServer, func() *SecurityMLAnalyticsSettingsServerTransport {
			return NewSecurityMLAnalyticsSettingsServerTransport(&s.srv.SecurityMLAnalyticsSettingsServer)
		})
		resp, err = s.trSecurityMLAnalyticsSettingsServer.Do(req)
	case "SentinelOnboardingStatesClient":
		initServer(s, &s.trSentinelOnboardingStatesServer, func() *SentinelOnboardingStatesServerTransport {
			return NewSentinelOnboardingStatesServerTransport(&s.srv.SentinelOnboardingStatesServer)
		})
		resp, err = s.trSentinelOnboardingStatesServer.Do(req)
	case "SourceControlClient":
		initServer(s, &s.trSourceControlServer, func() *SourceControlServerTransport {
			return NewSourceControlServerTransport(&s.srv.SourceControlServer)
		})
		resp, err = s.trSourceControlServer.Do(req)
	case "SourceControlsClient":
		initServer(s, &s.trSourceControlsServer, func() *SourceControlsServerTransport {
			return NewSourceControlsServerTransport(&s.srv.SourceControlsServer)
		})
		resp, err = s.trSourceControlsServer.Do(req)
	case "ThreatIntelligenceIndicatorClient":
		initServer(s, &s.trThreatIntelligenceIndicatorServer, func() *ThreatIntelligenceIndicatorServerTransport {
			return NewThreatIntelligenceIndicatorServerTransport(&s.srv.ThreatIntelligenceIndicatorServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorServer.Do(req)
	case "ThreatIntelligenceIndicatorMetricsClient":
		initServer(s, &s.trThreatIntelligenceIndicatorMetricsServer, func() *ThreatIntelligenceIndicatorMetricsServerTransport {
			return NewThreatIntelligenceIndicatorMetricsServerTransport(&s.srv.ThreatIntelligenceIndicatorMetricsServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorMetricsServer.Do(req)
	case "ThreatIntelligenceIndicatorsClient":
		initServer(s, &s.trThreatIntelligenceIndicatorsServer, func() *ThreatIntelligenceIndicatorsServerTransport {
			return NewThreatIntelligenceIndicatorsServerTransport(&s.srv.ThreatIntelligenceIndicatorsServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorsServer.Do(req)
	case "WatchlistItemsClient":
		initServer(s, &s.trWatchlistItemsServer, func() *WatchlistItemsServerTransport {
			return NewWatchlistItemsServerTransport(&s.srv.WatchlistItemsServer)
		})
		resp, err = s.trWatchlistItemsServer.Do(req)
	case "WatchlistsClient":
		initServer(s, &s.trWatchlistsServer, func() *WatchlistsServerTransport { return NewWatchlistsServerTransport(&s.srv.WatchlistsServer) })
		resp, err = s.trWatchlistsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
