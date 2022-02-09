package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ParentStudent struct {
	ID              uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4();" json:"id"`
	ParentName      string    `gorm:"not null" json:"parent_name"`
	ParentPhone     string    `json:"parent_phone"`
	StudentCustomID uuid.UUID `gorm:"not null" json:"student_custom_id"`

	StudentCustom *StudentCustom `gorm:"foreignKey:StudentCustomID" json:"student_custom,omitempty"`

	CreatedAt *time.Time     `gorm:"default:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"default:now()" json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
