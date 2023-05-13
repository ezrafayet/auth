package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"iam/src/core/model"
	"time"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) CreateUser(user model.UserModel, authMethod model.UsersAuthMethodsModel) error {
	var txErr error
	tx, txErr := r.db.Begin()
	if txErr != nil {
		return errors.New("ERROR")
	}
	defer func(tx *sql.Tx) {
		switch txErr {
		case nil:
			txErr = tx.Commit()
		default:
			err := tx.Rollback()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}(tx)

	_, txErr = tx.Exec("INSERT INTO users(id, created_at, username, username_fingerprint, email) VALUES ($1, $2, $3, $4, $5)", user.Id, time.Time(user.CreatedAt), user.Username, user.UsernameFingerprint, user.Email)

	if txErr != nil {
		return txErr
	}

	_, txErr = tx.Exec("INSERT INTO users_auth_methods(id, user_id, auth_method) VALUES ($1, $2, $3)", authMethod.Id, authMethod.UserId, authMethod.AuthMethod)

	if txErr != nil {
		return txErr
	}

	return nil
}
