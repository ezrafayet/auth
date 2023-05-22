package model

import (
	"iam/src/core/domain/types"
)

type UserMarketingPreferencesModel struct {
	Id                  types.Id
	UserId              types.Id
	AcceptedMarketing   bool
	UpdatedMarketingAt  types.Timestamp
	AcceptedNewsletter  bool
	UpdatedNewsletterAt types.Timestamp
}

func NewUserMarketingPreferencesModel(userId types.Id) UserMarketingPreferencesModel {
	return UserMarketingPreferencesModel{
		Id:                  types.NewId(),
		UserId:              userId,
		AcceptedMarketing:   false,
		UpdatedMarketingAt:  types.NewTimestamp(),
		AcceptedNewsletter:  false,
		UpdatedNewsletterAt: types.NewTimestamp(),
	}
}

func (v *UserMarketingPreferencesModel) AcceptMarketing() {
	v.AcceptedMarketing = true
	v.UpdatedMarketingAt = types.NewTimestamp()
}

func (v *UserMarketingPreferencesModel) DeclineMarketing() {
	v.AcceptedMarketing = false
	v.UpdatedMarketingAt = types.NewTimestamp()
}

func (v *UserMarketingPreferencesModel) AcceptNewsletter() {
	v.AcceptedNewsletter = true
	v.UpdatedNewsletterAt = types.NewTimestamp()
}

func (v *UserMarketingPreferencesModel) DeclineNewsletter() {
	v.AcceptedNewsletter = false
	v.UpdatedNewsletterAt = types.NewTimestamp()
}
