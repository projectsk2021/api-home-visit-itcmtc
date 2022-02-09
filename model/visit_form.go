package model

import (
	"time"
)

type VisitForm struct {
	VisitorFormID uint    `gorm:"primaryKey; autoIncrement;" json:"visitor_form_id"`
	DegreeID      *uint   `json:"degree_id"`    // id ของระดับชั้น
	FacultyID     *uint   `json:"faculty_id"`   // แผนกวิชา
	MajorID       *uint   `json:"major_id"`     // สาขา
	Address       *string `json:"address"`      // ที่อยู่
	PhoneNumber   *string `json:"phone_number"` // เบอร์โทร
	// ผู้ปกครอง
	ParentName        *string `json:"parent_name"`         // ชื่อผู้ปกครอง
	ParentAddress     *string `json:"parent_address"`      // ที่อยู่ผู้ปกครอง
	ParentPhoneNumber *string `json:"parent_phone_number"` // เบอร์โทรผู้ปกครอง
	ParentCareer      *string `json:"parent_career"`       // อาชีพผู้ปกครอง
	ParentIncome      *uint   `json:"parent_income"`       // รายได้ผู้ปกครอง
	// สถานที่อยู่
	TypeHomeID      *uint   `json:"type_home_id"`      // ชนิดบ้าน
	TypeHomeRemark  *string `json:"type_home_remark"`  // หมายเหตุ อื่น ๆ
	HomeAddress     *string `json:"home_address"`      // ที่อยู่บ้าน
	HomePhoneNumber *string `json:"home_phone_number"` // เบอร์โทรบ้าน
	HomeConditions  *string `json:"home_conditions"`   // สภาพบ้าน
	// ข้อมูลเกี่ยวกับครอบครัว
	AboutFamily   *string `json:"about_family"`    // ข้อมูลเกี่ยวกับครอบครัว
	RoleInFamily  *string `json:"role_in_family"`  // หน้าที่ในครอบครัว
	IssueInFamily *string `json:"issue_in_family"` // ปัญหาในครอบครัว
	// ความคิดเห็น
	Comment          *string `json:"comment"`            // ความคิดเห็น
	CommentOfAdviser *string `json:"comment_of_adviser"` // ความคิดเห็นของที่ปรึกษา
	// ลายเซ็น
	SignatureOfStudent *string `json:"signature_of_student"` // ลายเซ็นของนักศึกษา
	SignatureOfAdviser *string `json:"signature_of_adviser"` // ลายเซ็นของที่ปรึกษา
	// รูปภาพ
	ImagesOne *string `json:"images_one"` // รูปภาพที่ 1
	ImagesTwo *string `json:"images_two"` // รูปภาพที่ 2

	Degree  *Degree  `gorm:"foreignKey:DegreeID" json:"degree"`
	Faculty *Faculty `gorm:"foreignKey:FacultyID" json:"faculty"`
	Major   *Major   `gorm:"foreignKey:MajorID" json:"major"`

	CreatedAt *time.Time `gorm:"default:now()" json:"-"`
	UpdatedAt *time.Time `gorm:"default:now()" json:"-"`
}
