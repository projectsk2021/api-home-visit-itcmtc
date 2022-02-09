package repository

import (
	"github.com/kamchai-n/api-student-home-visit/model"
	"gorm.io/gorm"
)

type staticRepository struct {
	db *gorm.DB
}

func NewStaticRepository(db *gorm.DB) StaticRepository {
	return staticRepository{db: db}
}

func (r staticRepository) GetDegress() (degree []model.Degree, err error) {
	if err = r.db.Model(&model.Degree{}).Find(&degree).Error; err != nil {
		return degree, err
	}
	return degree, nil
}

func (r staticRepository) GetMajors() (major []model.Major, err error) {
	if err = r.db.Model(&model.Major{}).Find(&major).Error; err != nil {
		return major, err
	}
	return major, nil
}

func (r staticRepository) GetFaculty() (faculty []model.Faculty, err error) {
	if err = r.db.Model(&model.Faculty{}).Find(&faculty).Error; err != nil {
		return faculty, err
	}
	return faculty, nil
}
