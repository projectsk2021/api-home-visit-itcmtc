package services

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type StudentService interface {
	NewStudent(req domain.StudentForm) (domain.ResponseCreateStudentFromExcel, error)
	NewStudentFromExcel(*multipart.FileHeader, uuid.UUID) ([]domain.ResponseCreateStudentFromExcel, error)
	ListInfoStudent(*domain.RequestGetAll) ([]model.StudentCustom, error)
	ListInfoStudentById(domain.RequestGetOneStudent) (model.StudentCustom, error)
	RemoveStudentById([]string) error
}
