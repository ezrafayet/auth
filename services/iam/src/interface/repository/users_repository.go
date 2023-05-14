package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
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
		if pgErr, ok := txErr.(*pq.Error); ok {
			if pgErr.Code == "23505" && pgErr.Constraint == "users_unique_fingerprint" {
				return errors.New("USERNAME_NOT_AVAILABLE")
			} else if pgErr.Code == "23505" && pgErr.Constraint == "users_unique_email" {
				return errors.New("EMAIL_NOT_AVAILABLE")
			}
		}
		fmt.Println(txErr)
		return errors.New("SERVER_ERROR")
	}

	_, txErr = tx.Exec("INSERT INTO users_auth_methods(id, user_id, auth_method) VALUES ($1, $2, $3)", authMethod.Id, authMethod.UserId, authMethod.AuthMethod)

	if txErr != nil {
		fmt.Println(txErr)
		return errors.New("SERVER_ERROR")
	}

	return nil
}
