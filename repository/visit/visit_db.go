package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return visitRepository{db: db}
}

func (s visitRepository) CreateVisit(req domain.RequestVisit) (visit model.Visit, err error) {
	visitForm := model.VisitForm{}
	tx := s.db.Begin()

	if err := tx.Model(&model.VisitForm{}).Create(&visitForm).Error; err != nil {
		tx.Rollback()
		return visit, err
	}

	visit = model.Visit{
		StudentCustomID: req.StudentCustomID,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
		IsInArea:        req.IsInArea,
		VisitorFormID:   visitForm.VisitorFormID,
		CreatedBy:       req.CreatedBy,
	}

	if err := tx.Model(&model.Visit{}).Create(&visit).Error; err != nil {
		tx.Rollback()
		return visit, err
	}

	tx.Commit()

	return visit, nil
}

func (s visitRepository) UpdateVisit(req domain.RequestUpdateVisitForm) error {
	visitForm := model.VisitForm{
		DegreeID:           req.DegreeID,
		FacultyID:          req.FacultyID,
		MajorID:            req.MajorID,
		Address:            req.Address,
		PhoneNumber:        req.PhoneNumber,
		ParentName:         req.ParentName,
		ParentAddress:      req.ParentAddress,
		ParentPhoneNumber:  req.ParentPhoneNumber,
		ParentCareer:       req.ParentCareer,
		ParentIncome:       req.ParentIncome,
		TypeHomeID:         req.TypeHomeID,
		TypeHomeRemark:     req.TypeHomeRemark,
		HomeAddress:        req.HomeAddress,
		HomePhoneNumber:    req.HomePhoneNumber,
		HomeConditions:     req.HomeConditions,
		AboutFamily:        req.AboutFamily,
		RoleInFamily:       req.RoleInFamily,
		IssueInFamily:      req.IssueInFamily,
		Comment:            req.Comment,
		CommentOfAdviser:   req.CommentOfAdviser,
		SignatureOfStudent: req.SignatureOfStudent,
		SignatureOfAdviser: req.SignatureOfAdviser,
		ImagesOne:          req.ImagesOne,
		ImagesTwo:          req.ImagesTwo,
	}

	if err := s.db.Model(&model.VisitForm{}).Where("visitor_form_id = ?", req.VisitorFormID).Updates(&visitForm).Error; err != nil {
		return err
	}

	return nil
}

func (s visitRepository) GetAllVisit(user_id *uuid.UUID) (visit []model.Visit, err error) {
	if err := s.db.Debug().Preload("VisitForm.Degree").Preload("VisitForm.Faculty").Preload("VisitForm.Major").Preload(clause.Associations).Where("created_by = ?", user_id).Order("created_at DESC").Find(&visit).Error; err != nil {
		return visit, err
	}
	return visit, nil
}

func (s visitRepository) GetDetailByIdVisit(user_id *uuid.UUID, visit_id *uuid.UUID) (visit model.Visit, err error) {
	if err := s.db.Preload(clause.Associations).Where("created_by = ? AND visit_id = ?", user_id, visit_id).Order("created_at DESC").Find(&visit).Error; err != nil {
		return visit, err
	}
	return visit, nil
}

func (s visitRepository) DeleteVisit(visit_id *uuid.UUID) error {
	if err := s.db.Where("visit_id = ?", visit_id).Delete(&model.Visit{}).Error; err != nil {
		return err
	}
	return nil
}
