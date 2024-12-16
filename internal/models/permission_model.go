package models

import "github.com/google/uuid"

type Permission struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name string `gorm:"type:text;not null" json:"name"`
}