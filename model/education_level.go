package model

type EducationLevel struct {
	EnducationLevelId    uint   `gorm:"primarykey" json:"enducation_level_id"`
	EnducationLevelLabel string `json:"enducation_level_label"`
}
