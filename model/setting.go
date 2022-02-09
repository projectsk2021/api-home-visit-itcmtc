package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Setting struct {
	SettingID uint      `gorm:"primaryKey; autoIncrement;" json:"setting_id"`
	UserID    uuid.UUID `gorm:"not null" json:"user_id"`
	IsFromSdc bool      `gorm:"default:false" json:"is_from_sdc"`

	User *User `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`

	CreatedAt *time.Time     `gorm:"default:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"default:now()" json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
