/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsSecurityGroup struct for ModelsSecurityGroup
type ModelsSecurityGroup struct {
	Description    string               `json:"description,omitempty"`
	Id             string               `json:"id,omitempty"`
	InboundRules   []ModelsSecurityRule `json:"inbound_rules,omitempty"`
	OrganizationId string               `json:"organization_id,omitempty"`
	OutboundRules  []ModelsSecurityRule `json:"outbound_rules,omitempty"`
	Revision       int32                `json:"revision,omitempty"`
}
