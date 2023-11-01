package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/google/uuid"
)

// RegistrationToken is used to register devices without an interactive login.
type RegistrationToken struct {
	Base
	OwnerID        string     `json:"owner_id,omitempty"`
	VpcID          uuid.UUID  `json:"vpc_id,omitempty"`
	OrganizationID uuid.UUID  `json:"-"`                      // OrganizationID is denormalized from the VPC record for performance
	BearerToken    string     `json:"bearer_token,omitempty"` // BearerToken is the token the client should use to authenticate the device registration request.
	Description    string     `json:"description,omitempty"`
	DeviceId       *uuid.UUID `json:"device_id,omitempty"`
	Expiration     *time.Time `json:"expiration,omitempty"`
}

type NexodusClaims struct {
	jwt.RegisteredClaims
	Scope    string    `json:"scope,omitempty"`
	VpcID    uuid.UUID `json:"vpc_id,omitempty"`
	DeviceID uuid.UUID `json:"device_id,omitempty"`
}

type AddRegistrationToken struct {
	VpcID       uuid.UUID  `json:"organization_id,omitempty"`
	Description string     `json:"description,omitempty"`
	SingleUse   bool       `json:"single_use,omitempty"` // SingleUse only allows the registration token to be used once.
	Expiration  *time.Time `json:"expiration,omitempty"` // Expiration is optional, if set the registration token is only valid until the Expiration time.
}
