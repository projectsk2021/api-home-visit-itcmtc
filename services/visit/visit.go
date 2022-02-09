package services

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type VisitService interface {
	ListAllVisit() ([]model.Visit, error)
	ListDetailById(*uuid.UUID) (model.Visit, error)
	NewVisit(domain.RequestVisit) (model.Visit, error)
	UpdateVisit(domain.RequestUpdateVisitForm) error
	DeleteVisit(*uuid.UUID) error
}
