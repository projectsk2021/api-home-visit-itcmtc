package domain

import "github.com/google/uuid"

type RequestVisit struct {
	StudentCustomID uuid.UUID `json:"student_custom_id"`
	Latitude        string    `json:"latitude"`
	Longitude       string    `json:"longitude"`
	IsInArea        bool      `json:"is_in_area"`
	CreatedBy       uuid.UUID `json:"created_by"`
}

type RequestUpdateVisitForm struct {
	VisitorFormID      uint    `json:"visitor_form_id"`
	DegreeID           *uint   `json:"degree_id"`            // id ของระดับชั้น
	FacultyID          *uint   `json:"faculty_id"`           // แผนกวิชา
	MajorID            *uint   `json:"major_id"`             // สาขา
	Address            *string `json:"address"`              // ที่อยู่
	PhoneNumber        *string `json:"phone_number"`         // เบอร์โทร
	ParentName         *string `json:"parent_name"`          // ชื่อผู้ปกครอง
	ParentAddress      *string `json:"parent_address"`       // ที่อยู่ผู้ปกครอง
	ParentPhoneNumber  *string `json:"parent_phone_number"`  // เบอร์โทรผู้ปกครอง
	ParentCareer       *string `json:"parent_career"`        // อาชีพผู้ปกครอง
	ParentIncome       *uint   `json:"parent_income"`        // รายได้ผู้ปกครอง
	TypeHomeID         *uint   `json:"type_home_id"`         // ชนิดบ้าน
	TypeHomeRemark     *string `json:"type_home_remark"`     // หมายเหตุ อื่น ๆ
	HomeAddress        *string `json:"home_address"`         // ที่อยู่บ้าน
	HomePhoneNumber    *string `json:"home_phone_number"`    // เบอร์โทรบ้าน
	HomeConditions     *string `json:"home_conditions"`      // สภาพบ้าน
	AboutFamily        *string `json:"about_family"`         // ข้อมูลเกี่ยวกับครอบครัว
	RoleInFamily       *string `json:"role_in_family"`       // หน้าที่ในครอบครัว
	IssueInFamily      *string `json:"issue_in_family"`      // ปัญหาในครอบครัว
	Comment            *string `json:"comment"`              // ความคิดเห็น
	CommentOfAdviser   *string `json:"comment_of_adviser"`   // ความคิดเห็นของที่ปรึกษา
	SignatureOfStudent *string `json:"signature_of_student"` // ลายเซ็นของนักศึกษา
	SignatureOfAdviser *string `json:"signature_of_adviser"` // ลายเซ็นของที่ปรึกษา
	ImagesOne          *string `json:"images_one"`           // รูปภาพที่ 1
	ImagesTwo          *string `json:"images_two"`           // รูปภาพที่ 2
}
