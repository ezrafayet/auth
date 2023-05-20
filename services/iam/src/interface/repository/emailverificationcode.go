package repository

import (
	"database/sql"
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"time"
)

type VerificationCodeRepository struct {
	db *sql.DB
}

func NewVerificationCodeRepository(db *sql.DB) *VerificationCodeRepository {
	return &VerificationCodeRepository{
		db: db,
	}
}

func (v *VerificationCodeRepository) SaveCode(code model.EmailVerificationCodeModel) error {
	_, err := v.db.Exec("INSERT INTO email_verification_codes (user_id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)", string(code.UserId), string(code.Code), time.Time(code.CreatedAt), time.Time(code.ExpiresAt))

	if err != nil {
		// todo: handle error
		return err
	}

	return nil
}

func (v *VerificationCodeRepository) CountActiveCodes(userId types.Id) (int, error) {
	var count int

	err := v.db.QueryRow("SELECT COUNT(*) FROM email_verification_codes WHERE user_id = $1 AND expires_at > $2", string(userId), time.Now().UTC()).Scan(&count)

	if err != nil {
		// todo: handle error
		return count, err
	}

	return count, nil
}

func (v *VerificationCodeRepository) GetCode(code types.Code) (model.EmailVerificationCodeModel, error) {
	var verificationCode model.EmailVerificationCodeModel
	var userId string
	var createdAt time.Time
	var expiresAt time.Time

	err := v.db.QueryRow("SELECT user_id, created_at, expires_at FROM email_verification_codes WHERE code = $1", string(code)).Scan(&userId, &createdAt, &expiresAt)

	if err != nil {
		// todo: handle error
		return model.EmailVerificationCodeModel{}, err
	}

	if userId == "" {
		return model.EmailVerificationCodeModel{}, errors.New(apperrors.VerificationCodeNotFound)
	}

	verificationCode.Hydrate(userId, string(code), createdAt, expiresAt)

	return verificationCode, nil
}

func (v *VerificationCodeRepository) DeleteCode(code types.Code) error {
	_, err := v.db.Exec("DELETE FROM email_verification_codes WHERE code = $1", string(code))

	if err != nil {
		return err
	}

	return nil
}
