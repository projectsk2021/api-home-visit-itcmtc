package domain

import "github.com/kamchai-n/api-student-home-visit/model"

type ResponseGetEducation struct {
	Faculty []model.Faculty `json:"faculty"`
	Degree  []model.Degree  `json:"degree"`
	Major   []model.Major   `json:"major"`
}
