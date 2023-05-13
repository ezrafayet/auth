package repository

import (
	"database/sql"
	"iam/src/core/model"
	"iam/src/core/types"
	"time"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) CreateUser(user model.UserModel, authMethod types.AuthMethod) error {
	_, err := r.db.Exec("INSERT INTO users(id, created_at, username, username_fingerprint, email) VALUES ($1, $2, $3, $4, $5)", user.Id, time.Time(user.CreatedAt), user.Username, user.UsernameFingerprint, user.Email)
	if err != nil {
		return err
	}
	return nil
}
