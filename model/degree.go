package model

type Degree struct {
	DegreeID          uint   `gorm:"primarykey" json:"degree_id"`
	DegreeName        string `json:"degree_name"`
	EnducationLevelId uint   `json:"enducation_level_id"`

	EducationLevel *EducationLevel `gorm:"foreignKey:EnducationLevelId;references:EnducationLevelId" json:"education_level,omitempty"`
}
