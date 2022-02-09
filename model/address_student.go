package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressStudent struct {
	ID                 uint      `gorm:"primarykey"`
	AddressNo          *string   `gorm:"null" json:"address_no"`
	AddressLane        *string   `gorm:"null" json:"address_lane"`
	AddressRoad        *string   `gorm:"null" json:"address_road"`
	AddressSubDistrict string    `gorm:"not null" json:"address_sub_district"`
	AddressDistrict    string    `gorm:"not null" json:"address_district"`
	AddressProvince    string    `gorm:"not null" json:"address_province"`
	Latitude           string    `gorm:"null" json:"latitude"`
	Longitude          string    `gorm:"null" json:"longitude"`
	StudentCustomID    uuid.UUID `gorm:"not null" json:"student_custom_id"`

	StudentCustom *StudentCustom `gorm:"foreignKey:StudentCustomID" json:"student_custom,omitempty"`

	CreatedAt *time.Time     `gorm:"default:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"default:now()" json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
