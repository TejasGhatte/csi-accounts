package schemas

import "time"

type Did struct {
	Scheme           string `json:"scheme" validate:"required"`
	Method           string `json:"method" validate:"required"`
	MethodSpecificID string `json:"methodSpecificId" validate:"required"`
}

type DID string

type PublicKey struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	PublicKey string `json:"publicKey"`
}

type Authentication struct {
	Type      string    `json:"type"`
	PublicKey PublicKey `json:"publicKey"`
}

type ServiceEndpoint struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	ServiceURL string `json:"serviceEndpoint"`
}

type DIDDocument struct {
	DID            DID               `json:"id"`             // The DID unique id
	PublicKeys     PublicKey         `json:"publicKeys"`     // List of public keys
	Authentication Authentication    `json:"authentication"` // Authentication methods
	Services       []ServiceEndpoint `json:"services"`       // Service endpoints
	CreatedAt      time.Time         `json:"createdAt"`      // Date when DID was created
	UpdatedAt      time.Time         `json:"updatedAt"`      // Date when DID was last updated
}
