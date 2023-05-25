package dbrepository

import (
	"database/sql"
	"iam/src/core/domain/model"
	"time"
)

type RefreshTokenRepository struct {
	db *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) SaveToken(refreshToken model.RefreshTokenModel) error {
	_, err := r.db.Exec("INSERT INTO refresh_token (user_id, token, created_at, expires_at, revoked) VALUES ($1, $2, $3, $4, $5)", string(refreshToken.UserId), string(refreshToken.Token), time.Time(refreshToken.CreatedAt), time.Time(refreshToken.ExpiresAt), false)

	if err != nil {
		return err
	}

	return nil
}
