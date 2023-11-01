/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsAddVPC struct for ModelsAddVPC
type ModelsAddVPC struct {
	Cidr           string `json:"cidr,omitempty"`
	CidrV6         string `json:"cidr_v6,omitempty"`
	Description    string `json:"description,omitempty"`
	HubZone        bool   `json:"hub_zone,omitempty"`
	OrganizationId string `json:"organization_id,omitempty"`
	PrivateCidr    bool   `json:"private_cidr,omitempty"`
}
