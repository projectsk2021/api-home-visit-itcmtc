package model

type Faculty struct {
	FacultyId         uint   `gorm:"primarykey" json:"faculty_id"`
	FacultyName       string `json:"faculty_name"`
	EnducationLevelId uint   `json:"enducation_level_id"`

	EducationLevel *EducationLevel `gorm:"foreignKey:EnducationLevelId;references:EnducationLevelId" json:"education_level,omitempty"`
}
