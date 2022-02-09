package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type SettingRepository interface {
	GetSettingById(uuid.UUID) (model.Setting, error)
	CreateSetting(uuid.UUID) error
	UpdatedIsFromSdcSetting(domain.RequestIsFromSdcUpdateSetting) error
}
