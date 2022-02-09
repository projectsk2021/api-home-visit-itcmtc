package repository

import (
	"github.com/kamchai-n/api-student-home-visit/model"
)

type StaticRepository interface {
	GetDegress() ([]model.Degree, error)
	GetMajors() ([]model.Major, error)
	GetFaculty() ([]model.Faculty, error)
}
