package dbrepository

import (
	"database/sql"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
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

func (r *RefreshTokenRepository) GetAndDeleteByToken(token types.Code) (model.RefreshTokenModel, error) {
	var refreshToken model.RefreshTokenModel

	var (
		userId    string
		createdAt time.Time
		expiresAt time.Time
		revoked   bool
		revokedAt sql.NullTime
	)

	err := r.db.QueryRow("SELECT user_id, created_at, expires_at, revoked, revoked_at FROM refresh_token WHERE token = $1", token).Scan(&userId, &createdAt, &expiresAt, &revoked, &revokedAt)

	if err != nil {
		return model.RefreshTokenModel{}, err
	}

	var parsedRevokedAt time.Time

	if revokedAt.Valid {
		parsedRevokedAt = revokedAt.Time
	}

	err = refreshToken.Hydrate(userId, createdAt, expiresAt, string(token), revoked, parsedRevokedAt)

	if err != nil {
		return model.RefreshTokenModel{}, err
	}

	_, err = r.db.Exec("DELETE FROM refresh_token WHERE token = $1", token)

	if err != nil {
		return model.RefreshTokenModel{}, err
	}

	return refreshToken, nil
}
