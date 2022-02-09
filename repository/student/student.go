package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type StudentRepository interface {
	CreateStudent(domain.StudentForm) (domain.ResponseCreateStudentFromExcel, error)
	CreateStudentFromExcel([]domain.StudentForm) ([]domain.ResponseCreateStudentFromExcel, error)
	GetInfoStudent(*domain.RequestGetAll, uuid.UUID) ([]model.StudentCustom, error)
	GetInfoStudentById(domain.RequestGetOneStudent) (model.StudentCustom, error)
	DeleteStudentById([]string) error
}
