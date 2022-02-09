package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
	"gorm.io/gorm"
)

type settingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) SettingRepository {
	return settingRepository{db: db}
}

func (r settingRepository) GetSettingById(user_id uuid.UUID) (model.Setting, error) {
	var setting model.Setting
	if err := r.db.Model(&model.Setting{}).Where("user_id = ?", user_id).First(&setting).Error; err != nil {
		return setting, err
	}
	return setting, nil
}

func (r settingRepository) CreateSetting(user_id uuid.UUID) error {
	if err := r.db.Model(&model.Setting{}).Create(&model.Setting{
		UserID: user_id,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r settingRepository) UpdatedIsFromSdcSetting(req domain.RequestIsFromSdcUpdateSetting) error {
	if err := r.db.Model(&model.Setting{}).Where("user_id = ?", req.ID).Updates(map[string]interface{}{
		"is_from_sdc": req.IsFromSdc,
	}).Error; err != nil {
		return err
	}
	return nil
}
