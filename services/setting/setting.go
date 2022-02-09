package services

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type SettingService interface {
	ListSetting(uuid.UUID) (model.Setting, error)
	UpdateIsFromSdcSetting(domain.RequestIsFromSdcUpdateSetting) error
}
