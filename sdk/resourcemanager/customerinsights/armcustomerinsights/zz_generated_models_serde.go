//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AssignmentPrincipal.
func (a AssignmentPrincipal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", a.PrincipalID)
	populate(objectMap, "principalMetadata", a.PrincipalMetadata)
	populate(objectMap, "principalType", a.PrincipalType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AuthorizationPolicy.
func (a AuthorizationPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "permissions", a.Permissions)
	populate(objectMap, "policyName", a.PolicyName)
	populate(objectMap, "primaryKey", a.PrimaryKey)
	populate(objectMap, "secondaryKey", a.SecondaryKey)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Connector.
func (c Connector) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectorId", c.ConnectorID)
	populate(objectMap, "connectorName", c.ConnectorName)
	populate(objectMap, "connectorProperties", c.ConnectorProperties)
	populate(objectMap, "connectorType", c.ConnectorType)
	populateTimeRFC3339(objectMap, "created", c.Created)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "displayName", c.DisplayName)
	populate(objectMap, "isInternal", c.IsInternal)
	populateTimeRFC3339(objectMap, "lastModified", c.LastModified)
	populate(objectMap, "state", c.State)
	populate(objectMap, "tenantId", c.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Connector.
func (c *Connector) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectorId":
			err = unpopulate(val, "ConnectorID", &c.ConnectorID)
			delete(rawMsg, key)
		case "connectorName":
			err = unpopulate(val, "ConnectorName", &c.ConnectorName)
			delete(rawMsg, key)
		case "connectorProperties":
			err = unpopulate(val, "ConnectorProperties", &c.ConnectorProperties)
			delete(rawMsg, key)
		case "connectorType":
			err = unpopulate(val, "ConnectorType", &c.ConnectorType)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC3339(val, "Created", &c.Created)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &c.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &c.DisplayName)
			delete(rawMsg, key)
		case "isInternal":
			err = unpopulate(val, "IsInternal", &c.IsInternal)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, "LastModified", &c.LastModified)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &c.State)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &c.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectorMapping.
func (c ConnectorMapping) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectorMappingName", c.ConnectorMappingName)
	populate(objectMap, "connectorName", c.ConnectorName)
	populate(objectMap, "connectorType", c.ConnectorType)
	populateTimeRFC3339(objectMap, "created", c.Created)
	populate(objectMap, "dataFormatId", c.DataFormatID)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "displayName", c.DisplayName)
	populate(objectMap, "entityType", c.EntityType)
	populate(objectMap, "entityTypeName", c.EntityTypeName)
	populateTimeRFC3339(objectMap, "lastModified", c.LastModified)
	populate(objectMap, "mappingProperties", c.MappingProperties)
	populateTimeRFC3339(objectMap, "nextRunTime", c.NextRunTime)
	populate(objectMap, "runId", c.RunID)
	populate(objectMap, "state", c.State)
	populate(objectMap, "tenantId", c.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectorMapping.
func (c *ConnectorMapping) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectorMappingName":
			err = unpopulate(val, "ConnectorMappingName", &c.ConnectorMappingName)
			delete(rawMsg, key)
		case "connectorName":
			err = unpopulate(val, "ConnectorName", &c.ConnectorName)
			delete(rawMsg, key)
		case "connectorType":
			err = unpopulate(val, "ConnectorType", &c.ConnectorType)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC3339(val, "Created", &c.Created)
			delete(rawMsg, key)
		case "dataFormatId":
			err = unpopulate(val, "DataFormatID", &c.DataFormatID)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &c.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &c.DisplayName)
			delete(rawMsg, key)
		case "entityType":
			err = unpopulate(val, "EntityType", &c.EntityType)
			delete(rawMsg, key)
		case "entityTypeName":
			err = unpopulate(val, "EntityTypeName", &c.EntityTypeName)
			delete(rawMsg, key)
		case "lastModified":
			err = unpopulateTimeRFC3339(val, "LastModified", &c.LastModified)
			delete(rawMsg, key)
		case "mappingProperties":
			err = unpopulate(val, "MappingProperties", &c.MappingProperties)
			delete(rawMsg, key)
		case "nextRunTime":
			err = unpopulateTimeRFC3339(val, "NextRunTime", &c.NextRunTime)
			delete(rawMsg, key)
		case "runId":
			err = unpopulate(val, "RunID", &c.RunID)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &c.State)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &c.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectorMappingProperties.
func (c ConnectorMappingProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availability", c.Availability)
	populate(objectMap, "completeOperation", c.CompleteOperation)
	populate(objectMap, "errorManagement", c.ErrorManagement)
	populate(objectMap, "fileFilter", c.FileFilter)
	populate(objectMap, "folderPath", c.FolderPath)
	populate(objectMap, "format", c.Format)
	populate(objectMap, "hasHeader", c.HasHeader)
	populate(objectMap, "structure", c.Structure)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EnrichingKpi.
func (e EnrichingKpi) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aliases", e.Aliases)
	populate(objectMap, "calculationWindow", e.CalculationWindow)
	populate(objectMap, "calculationWindowFieldName", e.CalculationWindowFieldName)
	populate(objectMap, "description", e.Description)
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "entityType", e.EntityType)
	populate(objectMap, "entityTypeName", e.EntityTypeName)
	populate(objectMap, "expression", e.Expression)
	populate(objectMap, "extracts", e.Extracts)
	populate(objectMap, "filter", e.Filter)
	populate(objectMap, "function", e.Function)
	populate(objectMap, "groupBy", e.GroupBy)
	populate(objectMap, "groupByMetadata", e.GroupByMetadata)
	populate(objectMap, "kpiName", e.KpiName)
	populate(objectMap, "participantProfilesMetadata", e.ParticipantProfilesMetadata)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	populate(objectMap, "tenantId", e.TenantID)
	populate(objectMap, "thresHolds", e.ThresHolds)
	populate(objectMap, "unit", e.Unit)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EntityTypeDefinition.
func (e EntityTypeDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "apiEntitySetName", e.APIEntitySetName)
	populate(objectMap, "attributes", e.Attributes)
	populate(objectMap, "description", e.Description)
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "entityType", e.EntityType)
	populate(objectMap, "fields", e.Fields)
	populate(objectMap, "instancesCount", e.InstancesCount)
	populate(objectMap, "largeImage", e.LargeImage)
	populateTimeRFC3339(objectMap, "lastChangedUtc", e.LastChangedUTC)
	populate(objectMap, "localizedAttributes", e.LocalizedAttributes)
	populate(objectMap, "mediumImage", e.MediumImage)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	populate(objectMap, "schemaItemTypeLink", e.SchemaItemTypeLink)
	populate(objectMap, "smallImage", e.SmallImage)
	populate(objectMap, "tenantId", e.TenantID)
	populate(objectMap, "timestampFieldName", e.TimestampFieldName)
	populate(objectMap, "typeName", e.TypeName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EntityTypeDefinition.
func (e *EntityTypeDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "apiEntitySetName":
			err = unpopulate(val, "APIEntitySetName", &e.APIEntitySetName)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "Attributes", &e.Attributes)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &e.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &e.DisplayName)
			delete(rawMsg, key)
		case "entityType":
			err = unpopulate(val, "EntityType", &e.EntityType)
			delete(rawMsg, key)
		case "fields":
			err = unpopulate(val, "Fields", &e.Fields)
			delete(rawMsg, key)
		case "instancesCount":
			err = unpopulate(val, "InstancesCount", &e.InstancesCount)
			delete(rawMsg, key)
		case "largeImage":
			err = unpopulate(val, "LargeImage", &e.LargeImage)
			delete(rawMsg, key)
		case "lastChangedUtc":
			err = unpopulateTimeRFC3339(val, "LastChangedUTC", &e.LastChangedUTC)
			delete(rawMsg, key)
		case "localizedAttributes":
			err = unpopulate(val, "LocalizedAttributes", &e.LocalizedAttributes)
			delete(rawMsg, key)
		case "mediumImage":
			err = unpopulate(val, "MediumImage", &e.MediumImage)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &e.ProvisioningState)
			delete(rawMsg, key)
		case "schemaItemTypeLink":
			err = unpopulate(val, "SchemaItemTypeLink", &e.SchemaItemTypeLink)
			delete(rawMsg, key)
		case "smallImage":
			err = unpopulate(val, "SmallImage", &e.SmallImage)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &e.TenantID)
			delete(rawMsg, key)
		case "timestampFieldName":
			err = unpopulate(val, "TimestampFieldName", &e.TimestampFieldName)
			delete(rawMsg, key)
		case "typeName":
			err = unpopulate(val, "TypeName", &e.TypeName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Hub.
func (h Hub) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", h.ID)
	populate(objectMap, "location", h.Location)
	populate(objectMap, "name", h.Name)
	populate(objectMap, "properties", h.Properties)
	populate(objectMap, "tags", h.Tags)
	populate(objectMap, "type", h.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InteractionTypeDefinition.
func (i InteractionTypeDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "apiEntitySetName", i.APIEntitySetName)
	populate(objectMap, "attributes", i.Attributes)
	populate(objectMap, "dataSourcePrecedenceRules", i.DataSourcePrecedenceRules)
	populate(objectMap, "defaultDataSource", i.DefaultDataSource)
	populate(objectMap, "description", i.Description)
	populate(objectMap, "displayName", i.DisplayName)
	populate(objectMap, "entityType", i.EntityType)
	populate(objectMap, "fields", i.Fields)
	populate(objectMap, "idPropertyNames", i.IDPropertyNames)
	populate(objectMap, "instancesCount", i.InstancesCount)
	populate(objectMap, "isActivity", i.IsActivity)
	populate(objectMap, "largeImage", i.LargeImage)
	populateTimeRFC3339(objectMap, "lastChangedUtc", i.LastChangedUTC)
	populate(objectMap, "localizedAttributes", i.LocalizedAttributes)
	populate(objectMap, "mediumImage", i.MediumImage)
	populate(objectMap, "participantProfiles", i.ParticipantProfiles)
	populate(objectMap, "primaryParticipantProfilePropertyName", i.PrimaryParticipantProfilePropertyName)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "schemaItemTypeLink", i.SchemaItemTypeLink)
	populate(objectMap, "smallImage", i.SmallImage)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "timestampFieldName", i.TimestampFieldName)
	populate(objectMap, "typeName", i.TypeName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InteractionTypeDefinition.
func (i *InteractionTypeDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "apiEntitySetName":
			err = unpopulate(val, "APIEntitySetName", &i.APIEntitySetName)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "Attributes", &i.Attributes)
			delete(rawMsg, key)
		case "dataSourcePrecedenceRules":
			err = unpopulate(val, "DataSourcePrecedenceRules", &i.DataSourcePrecedenceRules)
			delete(rawMsg, key)
		case "defaultDataSource":
			err = unpopulate(val, "DefaultDataSource", &i.DefaultDataSource)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &i.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &i.DisplayName)
			delete(rawMsg, key)
		case "entityType":
			err = unpopulate(val, "EntityType", &i.EntityType)
			delete(rawMsg, key)
		case "fields":
			err = unpopulate(val, "Fields", &i.Fields)
			delete(rawMsg, key)
		case "idPropertyNames":
			err = unpopulate(val, "IDPropertyNames", &i.IDPropertyNames)
			delete(rawMsg, key)
		case "instancesCount":
			err = unpopulate(val, "InstancesCount", &i.InstancesCount)
			delete(rawMsg, key)
		case "isActivity":
			err = unpopulate(val, "IsActivity", &i.IsActivity)
			delete(rawMsg, key)
		case "largeImage":
			err = unpopulate(val, "LargeImage", &i.LargeImage)
			delete(rawMsg, key)
		case "lastChangedUtc":
			err = unpopulateTimeRFC3339(val, "LastChangedUTC", &i.LastChangedUTC)
			delete(rawMsg, key)
		case "localizedAttributes":
			err = unpopulate(val, "LocalizedAttributes", &i.LocalizedAttributes)
			delete(rawMsg, key)
		case "mediumImage":
			err = unpopulate(val, "MediumImage", &i.MediumImage)
			delete(rawMsg, key)
		case "participantProfiles":
			err = unpopulate(val, "ParticipantProfiles", &i.ParticipantProfiles)
			delete(rawMsg, key)
		case "primaryParticipantProfilePropertyName":
			err = unpopulate(val, "PrimaryParticipantProfilePropertyName", &i.PrimaryParticipantProfilePropertyName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &i.ProvisioningState)
			delete(rawMsg, key)
		case "schemaItemTypeLink":
			err = unpopulate(val, "SchemaItemTypeLink", &i.SchemaItemTypeLink)
			delete(rawMsg, key)
		case "smallImage":
			err = unpopulate(val, "SmallImage", &i.SmallImage)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &i.TenantID)
			delete(rawMsg, key)
		case "timestampFieldName":
			err = unpopulate(val, "TimestampFieldName", &i.TimestampFieldName)
			delete(rawMsg, key)
		case "typeName":
			err = unpopulate(val, "TypeName", &i.TypeName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type KpiDefinition.
func (k KpiDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aliases", k.Aliases)
	populate(objectMap, "calculationWindow", k.CalculationWindow)
	populate(objectMap, "calculationWindowFieldName", k.CalculationWindowFieldName)
	populate(objectMap, "description", k.Description)
	populate(objectMap, "displayName", k.DisplayName)
	populate(objectMap, "entityType", k.EntityType)
	populate(objectMap, "entityTypeName", k.EntityTypeName)
	populate(objectMap, "expression", k.Expression)
	populate(objectMap, "extracts", k.Extracts)
	populate(objectMap, "filter", k.Filter)
	populate(objectMap, "function", k.Function)
	populate(objectMap, "groupBy", k.GroupBy)
	populate(objectMap, "groupByMetadata", k.GroupByMetadata)
	populate(objectMap, "kpiName", k.KpiName)
	populate(objectMap, "participantProfilesMetadata", k.ParticipantProfilesMetadata)
	populate(objectMap, "provisioningState", k.ProvisioningState)
	populate(objectMap, "tenantId", k.TenantID)
	populate(objectMap, "thresHolds", k.ThresHolds)
	populate(objectMap, "unit", k.Unit)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KpiGroupByMetadata.
func (k KpiGroupByMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", k.DisplayName)
	populate(objectMap, "fieldName", k.FieldName)
	populate(objectMap, "fieldType", k.FieldType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LinkDefinition.
func (l LinkDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", l.Description)
	populate(objectMap, "displayName", l.DisplayName)
	populate(objectMap, "linkName", l.LinkName)
	populate(objectMap, "mappings", l.Mappings)
	populate(objectMap, "operationType", l.OperationType)
	populate(objectMap, "participantPropertyReferences", l.ParticipantPropertyReferences)
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "referenceOnly", l.ReferenceOnly)
	populate(objectMap, "sourceEntityType", l.SourceEntityType)
	populate(objectMap, "sourceEntityTypeName", l.SourceEntityTypeName)
	populate(objectMap, "targetEntityType", l.TargetEntityType)
	populate(objectMap, "targetEntityTypeName", l.TargetEntityTypeName)
	populate(objectMap, "tenantId", l.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetadataDefinitionBase.
func (m MetadataDefinitionBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "attributes", m.Attributes)
	populate(objectMap, "description", m.Description)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "largeImage", m.LargeImage)
	populate(objectMap, "localizedAttributes", m.LocalizedAttributes)
	populate(objectMap, "mediumImage", m.MediumImage)
	populate(objectMap, "smallImage", m.SmallImage)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Participant.
func (p Participant) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", p.Description)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "participantName", p.ParticipantName)
	populate(objectMap, "participantPropertyReferences", p.ParticipantPropertyReferences)
	populate(objectMap, "profileTypeName", p.ProfileTypeName)
	populate(objectMap, "role", p.Role)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Prediction.
func (p Prediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoAnalyze", p.AutoAnalyze)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "grades", p.Grades)
	populate(objectMap, "involvedInteractionTypes", p.InvolvedInteractionTypes)
	populate(objectMap, "involvedKpiTypes", p.InvolvedKpiTypes)
	populate(objectMap, "involvedRelationships", p.InvolvedRelationships)
	populate(objectMap, "mappings", p.Mappings)
	populate(objectMap, "negativeOutcomeExpression", p.NegativeOutcomeExpression)
	populate(objectMap, "positiveOutcomeExpression", p.PositiveOutcomeExpression)
	populate(objectMap, "predictionName", p.PredictionName)
	populate(objectMap, "primaryProfileType", p.PrimaryProfileType)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "scopeExpression", p.ScopeExpression)
	populate(objectMap, "scoreLabel", p.ScoreLabel)
	populate(objectMap, "systemGeneratedEntities", p.SystemGeneratedEntities)
	populate(objectMap, "tenantId", p.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PredictionSystemGeneratedEntities.
func (p PredictionSystemGeneratedEntities) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "generatedInteractionTypes", p.GeneratedInteractionTypes)
	populate(objectMap, "generatedKpis", p.GeneratedKpis)
	populate(objectMap, "generatedLinks", p.GeneratedLinks)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProfileEnumValidValuesFormat.
func (p ProfileEnumValidValuesFormat) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "localizedValueNames", p.LocalizedValueNames)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProfileTypeDefinition.
func (p ProfileTypeDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "apiEntitySetName", p.APIEntitySetName)
	populate(objectMap, "attributes", p.Attributes)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "entityType", p.EntityType)
	populate(objectMap, "fields", p.Fields)
	populate(objectMap, "instancesCount", p.InstancesCount)
	populate(objectMap, "largeImage", p.LargeImage)
	populateTimeRFC3339(objectMap, "lastChangedUtc", p.LastChangedUTC)
	populate(objectMap, "localizedAttributes", p.LocalizedAttributes)
	populate(objectMap, "mediumImage", p.MediumImage)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "schemaItemTypeLink", p.SchemaItemTypeLink)
	populate(objectMap, "smallImage", p.SmallImage)
	populate(objectMap, "strongIds", p.StrongIDs)
	populate(objectMap, "tenantId", p.TenantID)
	populate(objectMap, "timestampFieldName", p.TimestampFieldName)
	populate(objectMap, "typeName", p.TypeName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProfileTypeDefinition.
func (p *ProfileTypeDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "apiEntitySetName":
			err = unpopulate(val, "APIEntitySetName", &p.APIEntitySetName)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "Attributes", &p.Attributes)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &p.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "entityType":
			err = unpopulate(val, "EntityType", &p.EntityType)
			delete(rawMsg, key)
		case "fields":
			err = unpopulate(val, "Fields", &p.Fields)
			delete(rawMsg, key)
		case "instancesCount":
			err = unpopulate(val, "InstancesCount", &p.InstancesCount)
			delete(rawMsg, key)
		case "largeImage":
			err = unpopulate(val, "LargeImage", &p.LargeImage)
			delete(rawMsg, key)
		case "lastChangedUtc":
			err = unpopulateTimeRFC3339(val, "LastChangedUTC", &p.LastChangedUTC)
			delete(rawMsg, key)
		case "localizedAttributes":
			err = unpopulate(val, "LocalizedAttributes", &p.LocalizedAttributes)
			delete(rawMsg, key)
		case "mediumImage":
			err = unpopulate(val, "MediumImage", &p.MediumImage)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		case "schemaItemTypeLink":
			err = unpopulate(val, "SchemaItemTypeLink", &p.SchemaItemTypeLink)
			delete(rawMsg, key)
		case "smallImage":
			err = unpopulate(val, "SmallImage", &p.SmallImage)
			delete(rawMsg, key)
		case "strongIds":
			err = unpopulate(val, "StrongIDs", &p.StrongIDs)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &p.TenantID)
			delete(rawMsg, key)
		case "timestampFieldName":
			err = unpopulate(val, "TimestampFieldName", &p.TimestampFieldName)
			delete(rawMsg, key)
		case "typeName":
			err = unpopulate(val, "TypeName", &p.TypeName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PropertyDefinition.
func (p PropertyDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "arrayValueSeparator", p.ArrayValueSeparator)
	populate(objectMap, "dataSourcePrecedenceRules", p.DataSourcePrecedenceRules)
	populate(objectMap, "enumValidValues", p.EnumValidValues)
	populate(objectMap, "fieldName", p.FieldName)
	populate(objectMap, "fieldType", p.FieldType)
	populate(objectMap, "isArray", p.IsArray)
	populate(objectMap, "isAvailableInGraph", p.IsAvailableInGraph)
	populate(objectMap, "isEnum", p.IsEnum)
	populate(objectMap, "isFlagEnum", p.IsFlagEnum)
	populate(objectMap, "isImage", p.IsImage)
	populate(objectMap, "isLocalizedString", p.IsLocalizedString)
	populate(objectMap, "isName", p.IsName)
	populate(objectMap, "isRequired", p.IsRequired)
	populate(objectMap, "maxLength", p.MaxLength)
	populate(objectMap, "propertyId", p.PropertyID)
	populate(objectMap, "schemaItemPropLink", p.SchemaItemPropLink)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RelationshipDefinition.
func (r RelationshipDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "cardinality", r.Cardinality)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "displayName", r.DisplayName)
	populateTimeRFC3339(objectMap, "expiryDateTimeUtc", r.ExpiryDateTimeUTC)
	populate(objectMap, "fields", r.Fields)
	populate(objectMap, "lookupMappings", r.LookupMappings)
	populate(objectMap, "profileType", r.ProfileType)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "relatedProfileType", r.RelatedProfileType)
	populate(objectMap, "relationshipGuidId", r.RelationshipGUIDID)
	populate(objectMap, "relationshipName", r.RelationshipName)
	populate(objectMap, "tenantId", r.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RelationshipDefinition.
func (r *RelationshipDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cardinality":
			err = unpopulate(val, "Cardinality", &r.Cardinality)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &r.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &r.DisplayName)
			delete(rawMsg, key)
		case "expiryDateTimeUtc":
			err = unpopulateTimeRFC3339(val, "ExpiryDateTimeUTC", &r.ExpiryDateTimeUTC)
			delete(rawMsg, key)
		case "fields":
			err = unpopulate(val, "Fields", &r.Fields)
			delete(rawMsg, key)
		case "lookupMappings":
			err = unpopulate(val, "LookupMappings", &r.LookupMappings)
			delete(rawMsg, key)
		case "profileType":
			err = unpopulate(val, "ProfileType", &r.ProfileType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "relatedProfileType":
			err = unpopulate(val, "RelatedProfileType", &r.RelatedProfileType)
			delete(rawMsg, key)
		case "relationshipGuidId":
			err = unpopulate(val, "RelationshipGUIDID", &r.RelationshipGUIDID)
			delete(rawMsg, key)
		case "relationshipName":
			err = unpopulate(val, "RelationshipName", &r.RelationshipName)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &r.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RelationshipLinkDefinition.
func (r RelationshipLinkDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", r.Description)
	populate(objectMap, "displayName", r.DisplayName)
	populate(objectMap, "interactionType", r.InteractionType)
	populate(objectMap, "linkName", r.LinkName)
	populate(objectMap, "mappings", r.Mappings)
	populate(objectMap, "profilePropertyReferences", r.ProfilePropertyReferences)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "relatedProfilePropertyReferences", r.RelatedProfilePropertyReferences)
	populate(objectMap, "relationshipGuidId", r.RelationshipGUIDID)
	populate(objectMap, "relationshipName", r.RelationshipName)
	populate(objectMap, "tenantId", r.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RelationshipTypeMapping.
func (r RelationshipTypeMapping) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fieldMappings", r.FieldMappings)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSetDescription.
func (r ResourceSetDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "elements", r.Elements)
	populate(objectMap, "exceptions", r.Exceptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RoleAssignment.
func (r RoleAssignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignmentName", r.AssignmentName)
	populate(objectMap, "conflationPolicies", r.ConflationPolicies)
	populate(objectMap, "connectors", r.Connectors)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "displayName", r.DisplayName)
	populate(objectMap, "interactions", r.Interactions)
	populate(objectMap, "kpis", r.Kpis)
	populate(objectMap, "links", r.Links)
	populate(objectMap, "principals", r.Principals)
	populate(objectMap, "profiles", r.Profiles)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "relationshipLinks", r.RelationshipLinks)
	populate(objectMap, "relationships", r.Relationships)
	populate(objectMap, "role", r.Role)
	populate(objectMap, "roleAssignments", r.RoleAssignments)
	populate(objectMap, "sasPolicies", r.SasPolicies)
	populate(objectMap, "segments", r.Segments)
	populate(objectMap, "tenantId", r.TenantID)
	populate(objectMap, "views", r.Views)
	populate(objectMap, "widgetTypes", r.WidgetTypes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type StrongID.
func (s StrongID) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "keyPropertyNames", s.KeyPropertyNames)
	populate(objectMap, "strongIdName", s.StrongIDName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type View.
func (v View) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "changed", v.Changed)
	populateTimeRFC3339(objectMap, "created", v.Created)
	populate(objectMap, "definition", v.Definition)
	populate(objectMap, "displayName", v.DisplayName)
	populate(objectMap, "tenantId", v.TenantID)
	populate(objectMap, "userId", v.UserID)
	populate(objectMap, "viewName", v.ViewName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type View.
func (v *View) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "changed":
			err = unpopulateTimeRFC3339(val, "Changed", &v.Changed)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC3339(val, "Created", &v.Created)
			delete(rawMsg, key)
		case "definition":
			err = unpopulate(val, "Definition", &v.Definition)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &v.DisplayName)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &v.TenantID)
			delete(rawMsg, key)
		case "userId":
			err = unpopulate(val, "UserID", &v.UserID)
			delete(rawMsg, key)
		case "viewName":
			err = unpopulate(val, "ViewName", &v.ViewName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WidgetType.
func (w WidgetType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "changed", w.Changed)
	populateTimeRFC3339(objectMap, "created", w.Created)
	populate(objectMap, "definition", w.Definition)
	populate(objectMap, "description", w.Description)
	populate(objectMap, "displayName", w.DisplayName)
	populate(objectMap, "imageUrl", w.ImageURL)
	populate(objectMap, "tenantId", w.TenantID)
	populate(objectMap, "widgetTypeName", w.WidgetTypeName)
	populate(objectMap, "widgetVersion", w.WidgetVersion)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WidgetType.
func (w *WidgetType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "changed":
			err = unpopulateTimeRFC3339(val, "Changed", &w.Changed)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC3339(val, "Created", &w.Created)
			delete(rawMsg, key)
		case "definition":
			err = unpopulate(val, "Definition", &w.Definition)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &w.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &w.DisplayName)
			delete(rawMsg, key)
		case "imageUrl":
			err = unpopulate(val, "ImageURL", &w.ImageURL)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &w.TenantID)
			delete(rawMsg, key)
		case "widgetTypeName":
			err = unpopulate(val, "WidgetTypeName", &w.WidgetTypeName)
			delete(rawMsg, key)
		case "widgetVersion":
			err = unpopulate(val, "WidgetVersion", &w.WidgetVersion)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}