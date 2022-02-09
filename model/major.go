package model

type Major struct {
	MajorId   uint   `gorm:"primarykey" json:"major_id"`
	MajorName string `json:"major_name"`

	FacultyId uint     `json:"faculty_id"`
	Faculty   *Faculty `gorm:"foreignKey:FacultyId" json:"faculty,omitempty"`
}
