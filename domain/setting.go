package domain

import "github.com/google/uuid"

type RequestIsFromSdcUpdateSetting struct {
	ID        uuid.UUID `json:"id"`
	IsFromSdc bool      `json:"is_from_sdc"`
}
