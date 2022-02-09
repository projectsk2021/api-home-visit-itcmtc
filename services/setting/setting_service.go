package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
	repoSetting "github.com/kamchai-n/api-student-home-visit/repository/setting"
)

type settingService struct {
	settingRepo repoSetting.SettingRepository
}

func NewSettingService(settingRepo repoSetting.SettingRepository) SettingService {
	return settingService{settingRepo: settingRepo}
}

func (s settingService) ListSetting(user_id uuid.UUID) (model.Setting, error) {
	setting, err := s.settingRepo.GetSettingById(user_id)
	if err != nil {
		return setting, fmt.Errorf("1317")
	}
	return setting, nil
}

func (s settingService) UpdateIsFromSdcSetting(req domain.RequestIsFromSdcUpdateSetting) error {

	if err := s.settingRepo.UpdatedIsFromSdcSetting(req); err != nil {
		return fmt.Errorf("1317")
	}
	return nil
}
