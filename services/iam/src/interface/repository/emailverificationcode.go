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

func (v *VerificationCodeRepository) SaveVerificationCode(code model.EmailVerificationCodeModel) error {
	_, err := v.db.Exec("INSERT INTO email_verification_codes (user_id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)", string(code.UserId), string(code.Code), time.Time(code.CreatedAt), time.Time(code.ExpiresAt))

	if err != nil {
		fmt.Println(err)
		return errors.New("SERVER_ERROR")
	}

	return nil
}

func (v *VerificationCodeRepository) ConfirmVerificationCode(code types.Code) error {
	return nil
}
