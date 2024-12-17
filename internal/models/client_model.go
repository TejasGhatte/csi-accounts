package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Client struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name         string         `gorm:"type:text;not null" json:"name"`
	Description  string         `gorm:"type:text" json:"description"`
	RedirectURIs pq.StringArray `gorm:"type:text[]" json:"redirectURIs"`
	ClientScopes []ClientScope  `gorm:"foreignKey:ClientID" json:"clientScopes"`
	ClientID     string         `gorm:"type:text;unique;not null" json:"clientID"`
	ClientSecret string         `gorm:"type:text;unique;not null" json:"clientSecret"`
	CreatedAt    time.Time      `gorm:"default:current_timestamp" json:"createdAt"`
}
