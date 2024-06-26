//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

// AlertRuleClassification provides polymorphic access to related types.
// Call the interface's GetAlertRule() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AlertRule, *FusionAlertRule, *MLBehaviorAnalyticsAlertRule, *MicrosoftSecurityIncidentCreationAlertRule, *NrtAlertRule,
// - *ScheduledAlertRule, *ThreatIntelligenceAlertRule
type AlertRuleClassification interface {
	// GetAlertRule returns the AlertRule content of the underlying type.
	GetAlertRule() *AlertRule
}

// AlertRuleTemplateClassification provides polymorphic access to related types.
// Call the interface's GetAlertRuleTemplate() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AlertRuleTemplate, *FusionAlertRuleTemplate, *MLBehaviorAnalyticsAlertRuleTemplate, *MicrosoftSecurityIncidentCreationAlertRuleTemplate,
// - *NrtAlertRuleTemplate, *ScheduledAlertRuleTemplate, *ThreatIntelligenceAlertRuleTemplate
type AlertRuleTemplateClassification interface {
	// GetAlertRuleTemplate returns the AlertRuleTemplate content of the underlying type.
	GetAlertRuleTemplate() *AlertRuleTemplate
}

// AutomationRuleActionClassification provides polymorphic access to related types.
// Call the interface's GetAutomationRuleAction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutomationRuleAction, *AutomationRuleModifyPropertiesAction, *AutomationRuleRunPlaybookAction
type AutomationRuleActionClassification interface {
	// GetAutomationRuleAction returns the AutomationRuleAction content of the underlying type.
	GetAutomationRuleAction() *AutomationRuleAction
}

// AutomationRuleConditionClassification provides polymorphic access to related types.
// Call the interface's GetAutomationRuleCondition() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutomationRuleCondition, *BooleanConditionProperties, *PropertyArrayChangedConditionProperties, *PropertyArrayConditionProperties,
// - *PropertyChangedConditionProperties, *PropertyConditionProperties
type AutomationRuleConditionClassification interface {
	// GetAutomationRuleCondition returns the AutomationRuleCondition content of the underlying type.
	GetAutomationRuleCondition() *AutomationRuleCondition
}

// CustomEntityQueryClassification provides polymorphic access to related types.
// Call the interface's GetCustomEntityQuery() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ActivityCustomEntityQuery, *CustomEntityQuery
type CustomEntityQueryClassification interface {
	// GetCustomEntityQuery returns the CustomEntityQuery content of the underlying type.
	GetCustomEntityQuery() *CustomEntityQuery
}

// DataConnectorClassification provides polymorphic access to related types.
// Call the interface's GetDataConnector() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AADDataConnector, *AATPDataConnector, *ASCDataConnector, *AwsCloudTrailDataConnector, *AwsS3DataConnector, *CodelessAPIPollingDataConnector,
// - *CodelessUIDataConnector, *DataConnector, *Dynamics365DataConnector, *IoTDataConnector, *MCASDataConnector, *MDATPDataConnector,
// - *MSTIDataConnector, *MTPDataConnector, *Office365ProjectDataConnector, *OfficeATPDataConnector, *OfficeDataConnector,
// - *OfficeIRMDataConnector, *OfficePowerBIDataConnector, *TIDataConnector, *TiTaxiiDataConnector
type DataConnectorClassification interface {
	// GetDataConnector returns the DataConnector content of the underlying type.
	GetDataConnector() *DataConnector
}

// DataConnectorsCheckRequirementsClassification provides polymorphic access to related types.
// Call the interface's GetDataConnectorsCheckRequirements() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AADCheckRequirements, *AATPCheckRequirements, *ASCCheckRequirements, *AwsCloudTrailCheckRequirements, *AwsS3CheckRequirements,
// - *DataConnectorsCheckRequirements, *Dynamics365CheckRequirements, *IoTCheckRequirements, *MCASCheckRequirements, *MDATPCheckRequirements,
// - *MSTICheckRequirements, *MtpCheckRequirements, *Office365ProjectCheckRequirements, *OfficeATPCheckRequirements, *OfficeIRMCheckRequirements,
// - *OfficePowerBICheckRequirements, *TICheckRequirements, *TiTaxiiCheckRequirements
type DataConnectorsCheckRequirementsClassification interface {
	// GetDataConnectorsCheckRequirements returns the DataConnectorsCheckRequirements content of the underlying type.
	GetDataConnectorsCheckRequirements() *DataConnectorsCheckRequirements
}

// EntityClassification provides polymorphic access to related types.
// Call the interface's GetEntity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccountEntity, *AzureResourceEntity, *CloudApplicationEntity, *DNSEntity, *Entity, *FileEntity, *FileHashEntity, *HostEntity,
// - *HuntingBookmark, *IPEntity, *IoTDeviceEntity, *MailClusterEntity, *MailMessageEntity, *MailboxEntity, *MalwareEntity,
// - *NicEntity, *ProcessEntity, *RegistryKeyEntity, *RegistryValueEntity, *SecurityAlert, *SecurityGroupEntity, *SubmissionMailEntity,
// - *URLEntity
type EntityClassification interface {
	// GetEntity returns the Entity content of the underlying type.
	GetEntity() *Entity
}

// EntityQueryClassification provides polymorphic access to related types.
// Call the interface's GetEntityQuery() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ActivityEntityQuery, *EntityQuery, *ExpansionEntityQuery
type EntityQueryClassification interface {
	// GetEntityQuery returns the EntityQuery content of the underlying type.
	GetEntityQuery() *EntityQuery
}

// EntityQueryItemClassification provides polymorphic access to related types.
// Call the interface's GetEntityQueryItem() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *EntityQueryItem, *InsightQueryItem
type EntityQueryItemClassification interface {
	// GetEntityQueryItem returns the EntityQueryItem content of the underlying type.
	GetEntityQueryItem() *EntityQueryItem
}

// EntityQueryTemplateClassification provides polymorphic access to related types.
// Call the interface's GetEntityQueryTemplate() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ActivityEntityQueryTemplate, *EntityQueryTemplate
type EntityQueryTemplateClassification interface {
	// GetEntityQueryTemplate returns the EntityQueryTemplate content of the underlying type.
	GetEntityQueryTemplate() *EntityQueryTemplate
}

// EntityTimelineItemClassification provides polymorphic access to related types.
// Call the interface's GetEntityTimelineItem() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ActivityTimelineItem, *AnomalyTimelineItem, *BookmarkTimelineItem, *EntityTimelineItem, *SecurityAlertTimelineItem
type EntityTimelineItemClassification interface {
	// GetEntityTimelineItem returns the EntityTimelineItem content of the underlying type.
	GetEntityTimelineItem() *EntityTimelineItem
}

// SecurityMLAnalyticsSettingClassification provides polymorphic access to related types.
// Call the interface's GetSecurityMLAnalyticsSetting() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AnomalySecurityMLAnalyticsSettings, *SecurityMLAnalyticsSetting
type SecurityMLAnalyticsSettingClassification interface {
	// GetSecurityMLAnalyticsSetting returns the SecurityMLAnalyticsSetting content of the underlying type.
	GetSecurityMLAnalyticsSetting() *SecurityMLAnalyticsSetting
}

// SettingsClassification provides polymorphic access to related types.
// Call the interface's GetSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Anomalies, *EntityAnalytics, *EyesOn, *Settings, *Ueba
type SettingsClassification interface {
	// GetSettings returns the Settings content of the underlying type.
	GetSettings() *Settings
}

// ThreatIntelligenceInformationClassification provides polymorphic access to related types.
// Call the interface's GetThreatIntelligenceInformation() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ThreatIntelligenceIndicatorModel, *ThreatIntelligenceInformation
type ThreatIntelligenceInformationClassification interface {
	// GetThreatIntelligenceInformation returns the ThreatIntelligenceInformation content of the underlying type.
	GetThreatIntelligenceInformation() *ThreatIntelligenceInformation
}
