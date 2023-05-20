package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"time"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) SaveUser(user model.UserModel, authMethod model.UsersAuthMethodsModel) error {
	var txErr error
	tx, txErr := r.db.Begin()
	if txErr != nil {
		return errors.New("SERVER_ERROR")
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

func (r *UsersRepository) GetUserById(id types.Id) (model.UserModel, error) {
	var user model.UserModel

	var (
		createdAt           time.Time
		username            string
		usernameFingerprint string
		email               string
		emailVerified       bool
		emailVerifiedAt     sql.NullTime
		blocked             bool
		deleted             bool
		deletedAt           sql.NullTime
	)

	txErr := r.db.QueryRow("SELECT created_at, username, username_fingerprint, email, email_verified, email_verified_at, blocked, deleted, deleted_at FROM users WHERE id = $1", id).Scan(&createdAt, &username, &usernameFingerprint, &email, &emailVerified, &emailVerifiedAt, &blocked, &deleted, &deletedAt)

	if txErr != nil {
		fmt.Println(txErr)
		return user, errors.New("USER_NOT_FOUND")
	}

	var parsedEmailValidatedAt time.Time

	if emailVerifiedAt.Valid {
		parsedEmailValidatedAt = emailVerifiedAt.Time
	}

	var parsedDeletedAt time.Time

	if deletedAt.Valid {
		parsedDeletedAt = deletedAt.Time
	}

	user.Hydrate(string(id), createdAt, username, usernameFingerprint, email, emailVerified, parsedEmailValidatedAt, blocked, deleted, parsedDeletedAt)

	return user, nil
}

func (r *UsersRepository) ValidateEmail(userId types.Id) error {
	_, txErr := r.db.Exec("UPDATE users SET email_verified = true, email_verified_at = $1 WHERE id = $2", time.Now().UTC(), userId)

	if txErr != nil {
		return txErr
	}

	return nil
}
