package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"userID"`
	Action string    `gorm:"type:text;not null" json:"action"`
	Timestamp time.Time `gorm:"default:current_timestamp" json:"auditLog"`
}