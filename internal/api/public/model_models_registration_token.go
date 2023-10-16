/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsRegistrationToken struct for ModelsRegistrationToken
type ModelsRegistrationToken struct {
	// BearerToken is the token the client should use to authenticate the device registration request.
	BearerToken    string `json:"bearer_token,omitempty"`
	Description    string `json:"description,omitempty"`
	DeviceId       string `json:"device_id,omitempty"`
	Expiration     string `json:"expiration,omitempty"`
	Id             string `json:"id,omitempty"`
	OrganizationId string `json:"organization_id,omitempty"`
	UserId         string `json:"user_id,omitempty"`
}
