package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"iam/src/core/model"
	"iam/src/core/types"
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
		fmt.Println(err)
		return errors.New("SERVER_ERROR")
	}

	return nil
}

func (v *VerificationCodeRepository) CountActiveCodes(userId types.Id) (int, error) {
	var count int

	err := v.db.QueryRow("SELECT COUNT(*) FROM email_verification_codes WHERE user_id = $1 AND expires_at > $2", string(userId), time.Now().UTC()).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return 0, errors.New("SERVER_ERROR")
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
		return model.EmailVerificationCodeModel{}, err
	}

	if userId == "" {
		return model.EmailVerificationCodeModel{}, errors.New("CODE_NOT_FOUND")
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
