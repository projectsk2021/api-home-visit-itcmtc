package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type VisitRepository interface {
	GetAllVisit(*uuid.UUID) ([]model.Visit, error)
	GetDetailByIdVisit(*uuid.UUID, *uuid.UUID) (model.Visit, error)
	CreateVisit(domain.RequestVisit) (model.Visit, error)
	UpdateVisit(req domain.RequestUpdateVisitForm) error
	DeleteVisit(*uuid.UUID) error
}
