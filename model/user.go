package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    uuid.UUID      `gorm:"primaryKey; type:uuid; default:uuid_generate_v4();" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Username  string         `gorm:"not null,unique" json:"username"`
	Password  *string        `gorm:"not null" json:"password,omitempty"`
	Images    *string        `gorm:"null" json:"images,omitempty"`
	CreatedAt *time.Time     `gorm:"default:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"default:now()" json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
