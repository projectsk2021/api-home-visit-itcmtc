package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	"github.com/kamchai-n/api-student-home-visit/model"
	repoVisit "github.com/kamchai-n/api-student-home-visit/repository/visit"
)

type visitService struct {
	visitRepo repoVisit.VisitRepository
}

func NewStudentService(visitRepo repoVisit.VisitRepository) VisitService {
	return visitService{visitRepo: visitRepo}
}

func (s visitService) NewVisit(req domain.RequestVisit) (model.Visit, error) {
	visit, err := s.visitRepo.CreateVisit(req)
	if err != nil {
		return visit, fmt.Errorf("1323")
	}
	return visit, nil
}

func (s visitService) UpdateVisit(req domain.RequestUpdateVisitForm) error {
	if err := s.visitRepo.UpdateVisit(req); err != nil {
		return fmt.Errorf("1324")
	}
	return nil
}

func (s visitService) ListAllVisit() ([]model.Visit, error) {
	visit, err := s.visitRepo.GetAllVisit(middlewares.UserClaims.UserId)
	if err != nil {
		return visit, fmt.Errorf("1324")
	}
	return visit, nil
}

func (s visitService) ListDetailById(visit_id *uuid.UUID) (model.Visit, error) {
	visit, err := s.visitRepo.GetDetailByIdVisit(middlewares.UserClaims.UserId, visit_id)
	if err != nil {
		return visit, fmt.Errorf("1328")
	}
	return visit, nil
}

func (s visitService) DeleteVisit(visit_id *uuid.UUID) error {
	if err := s.visitRepo.DeleteVisit(visit_id); err != nil {
		return fmt.Errorf("1329")
	}
	return nil
}
