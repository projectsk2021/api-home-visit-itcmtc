package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Visit struct {
	VisitID         uuid.UUID      `gorm:"primaryKey; type:uuid; default:uuid_generate_v4();" json:"visitor_id"`
	StudentCustomID uuid.UUID      `json:"student_custom_id"` // id ของนักศึกษา
	Latitude        string         `json:"latitude"`
	Longitude       string         `json:"longitude"`
	IsInArea        bool           `json:"is_in_area"`
	VisitorFormID   uint           `json:"visitor_form_id"`
	CreatedBy       uuid.UUID      `json:"created_by"`
	CreatedAt       *time.Time     `gorm:"default:now()" json:"created_at"`
	DeleteAt        gorm.DeletedAt `json:"-"`

	StudentCustom *StudentCustom `gorm:"foreignKey:ID;references:StudentCustomID" json:"student_custom,omitempty"`
	VisitForm     *VisitForm     `gorm:"foreignKey:VisitorFormID;" json:"visit_form,omitempty"`
}
