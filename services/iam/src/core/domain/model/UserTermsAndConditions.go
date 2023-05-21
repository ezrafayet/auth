package model

import (
	"iam/src/core/domain/types"
)

type UserTermsAndConditionsModel struct {
	Id           types.Id
	UserId       types.Id
	Accepted     bool
	AcceptedAt   types.Timestamp
	TermsVersion string
}

func NewUserTermsAndConditionsModel(userId types.Id) UserTermsAndConditionsModel {
	return UserTermsAndConditionsModel{
		Id:       types.NewId(),
		UserId:   userId,
		Accepted: false,
	}
}

func (v *UserTermsAndConditionsModel) Accept(termsVersion string) {
	v.Accepted = true
	v.AcceptedAt = types.NewTimestamp()
	v.TermsVersion = termsVersion
}
