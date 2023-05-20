package repository

import (
	"database/sql"
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"time"
)

type AuthorizationCodeRepository struct {
	db *sql.DB
}

func NewAuthorizationCodeRepository(db *sql.DB) *AuthorizationCodeRepository {
	return &AuthorizationCodeRepository{
		db: db,
	}
}

func (a *AuthorizationCodeRepository) SaveCode(code model.AuthorizationCodeModel) error {
	_, err := a.db.Exec("INSERT INTO authorization_codes (user_id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)", string(code.UserId), string(code.Code), time.Time(code.CreatedAt), time.Time(code.ExpiresAt))

	if err != nil {
		return err
	}

	return nil
}

func (a *AuthorizationCodeRepository) CountCodes(userId types.Id) (int, error) {
	var count int

	err := a.db.QueryRow("SELECT COUNT(*) FROM authorization_codes WHERE user_id = $1 AND expires_at > $2", string(userId), time.Now().UTC()).Scan(&count)

	if err != nil {
		return count, err
	}

	return count, nil
}

func (a *AuthorizationCodeRepository) GetCode(code types.Code) (model.AuthorizationCodeModel, error) {
	var authorizationCode model.AuthorizationCodeModel
	var userId string
	var createdAt time.Time
	var expiresAt time.Time

	err := a.db.QueryRow("SELECT user_id, created_at, expires_at FROM authorization_codes WHERE code = $1", string(code)).Scan(&userId, &createdAt, &expiresAt)

	if err != nil {
		return model.AuthorizationCodeModel{}, err
	}

	if userId == "" {
		return model.AuthorizationCodeModel{}, errors.New(apperrors.AuthorizationCodeNotFound)
	}

	authorizationCode.Hydrate(userId, string(code), createdAt, expiresAt)

	return authorizationCode, nil
}

func (a *AuthorizationCodeRepository) DeleteCode(code types.Code) error {
	_, err := a.db.Exec("DELETE FROM authorization_codes WHERE code = $1", string(code))

	if err != nil {
		return err
	}

	return nil
}
