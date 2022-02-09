package services

import (
	"github.com/kamchai-n/api-student-home-visit/domain"
)

type StaticService interface {
	ListEducation() (domain.ResponseGetEducation, error)
}
