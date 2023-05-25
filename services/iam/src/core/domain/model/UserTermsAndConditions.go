package model

import (
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/types"
)

type UserTermsAndConditionsModel struct {
	UserId       types.Id
	Accepted     bool
	AcceptedAt   types.Timestamp
	TermsVersion string
	UserData     string
}

func NewUserTermsAndConditionsModel(userId types.Id) UserTermsAndConditionsModel {
	return UserTermsAndConditionsModel{
		UserId:   userId,
		Accepted: false,
	}
}

func (v *UserTermsAndConditionsModel) Accept(acceptedTerms bool, termsVersion, UserData string) error {
	if !acceptedTerms {
		return errors.New(apperrors.RefusedTerms)
	}

	v.Accepted = true
	v.AcceptedAt = types.NewTimestamp()
	v.TermsVersion = termsVersion
	v.UserData = UserData // todo: move to creation phase

	return nil
}
