package model

import (
	"iam/src/core/domain/types"
)

type UserMarketingPreferenceModel struct {
	Id                types.Id
	UserId            types.Id
	AcceptedMarketing bool
	UpdatedAt         types.Timestamp
}

func NewUserMarketingPreferenceModel(userId types.Id) UserMarketingPreferenceModel {
	return UserMarketingPreferenceModel{
		Id:                types.NewId(),
		UserId:            userId,
		AcceptedMarketing: false,
		UpdatedAt:         types.NewTimestamp(),
	}
}

func (v *UserMarketingPreferenceModel) Accept() {
	v.AcceptedMarketing = true
	v.UpdatedAt = types.NewTimestamp()
}

func (v *UserMarketingPreferenceModel) Decline() {
	v.AcceptedMarketing = false
	v.UpdatedAt = types.NewTimestamp()
}
