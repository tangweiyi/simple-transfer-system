package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int64
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
