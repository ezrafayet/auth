package dbrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"iam/pkg/apperrors"
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

func (r *UsersRepository) SaveUser(
	user model.UserModel,
	authType model.UserAuthTypeModel,
	role types.Role,
	termsAndConditions model.UserTermsAndConditionsModel,
	marketingPreferences model.UserMarketingPreferencesModel) error {
	var txErr error
	tx, txErr := r.db.Begin()
	if txErr != nil {
		return txErr
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

	_, txErr = tx.Exec("INSERT INTO user_account(id, created_at, username, username_fingerprint, email) VALUES ($1, $2, $3, $4, $5)", user.Id, time.Time(user.CreatedAt), user.Username, user.UsernameFingerprint, user.Email)

	if txErr != nil {
		if pgErr, ok := txErr.(*pq.Error); ok {
			if pgErr.Code == "23505" && pgErr.Constraint == "users_unique_fingerprint" {
				return errors.New(apperrors.UsernameAlreadyExists)
			} else if pgErr.Code == "23505" && pgErr.Constraint == "users_unique_email" {
				return errors.New(apperrors.EmailAlreadyExists)
			}
		}
		return txErr
	}

	_, txErr = tx.Exec("INSERT INTO user_auth_type(user_id, auth_type_id, created_at) VALUES ($1, $2, $3)", string(authType.UserId), authType.AuthTypeId, time.Time(authType.CreatedAt))

	if txErr != nil {
		return txErr
	}

	_, txErr = tx.Exec("INSERT INTO user_role(user_id, role_id, created_at) VALUES ($1, $2, $3)", string(user.Id), role, time.Time(user.CreatedAt))

	if txErr != nil {
		return txErr
	}

	_, txErr = tx.Exec("INSERT INTO user_terms_and_conditions(user_id, accepted, accepted_at, terms_version, user_data) VALUES ($1, $2, $3, $4, $5)", string(termsAndConditions.UserId), termsAndConditions.Accepted, time.Time(termsAndConditions.AcceptedAt), termsAndConditions.TermsVersion, termsAndConditions.UserData)

	if txErr != nil {
		return txErr
	}

	_, txErr = tx.Exec("INSERT INTO user_marketing_preferences(user_id, accepted_marketing, updated_marketing_at, accepted_newsletter, updated_newsletter_at) VALUES ($1, $2, $3, $4, $5)", string(marketingPreferences.UserId), marketingPreferences.AcceptedMarketing, time.Time(marketingPreferences.UpdatedMarketingAt), marketingPreferences.AcceptedNewsletter, time.Time(marketingPreferences.UpdatedNewsletterAt))

	if txErr != nil {
		return txErr
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

	txErr := r.db.QueryRow("SELECT created_at, username, username_fingerprint, email, email_verified, email_verified_at, blocked, deleted, deleted_at FROM user_account WHERE id = $1", id).Scan(&createdAt, &username, &usernameFingerprint, &email, &emailVerified, &emailVerifiedAt, &blocked, &deleted, &deletedAt)

	if txErr != nil {
		fmt.Println(txErr)
		// todo: handle error better
		return user, errors.New(apperrors.UserNotFound)
	}

	var parsedEmailValidatedAt time.Time

	if emailVerifiedAt.Valid {
		parsedEmailValidatedAt = emailVerifiedAt.Time
	}

	var parsedDeletedAt time.Time

	if deletedAt.Valid {
		parsedDeletedAt = deletedAt.Time
	}

	err := user.Hydrate(string(id), createdAt, username, usernameFingerprint, email, emailVerified, parsedEmailValidatedAt, blocked, deleted, parsedDeletedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UsersRepository) GetUserByEmail(email types.Email) (model.UserModel, error) {
	var user model.UserModel

	var (
		id                  types.Id
		createdAt           time.Time
		username            string
		usernameFingerprint string
		emailVerified       bool
		emailVerifiedAt     sql.NullTime
		blocked             bool
		deleted             bool
		deletedAt           sql.NullTime
	)

	txErr := r.db.QueryRow("SELECT id, created_at, username, username_fingerprint, email_verified, email_verified_at, blocked, deleted, deleted_at FROM user_account WHERE email = $1", string(email)).Scan(&id, &createdAt, &username, &usernameFingerprint, &emailVerified, &emailVerifiedAt, &blocked, &deleted, &deletedAt)

	if txErr != nil {
		fmt.Println(txErr)
		// todo: handle error better
		return user, errors.New(apperrors.UserNotFound)
	}

	var parsedEmailValidatedAt time.Time

	if emailVerifiedAt.Valid {
		parsedEmailValidatedAt = emailVerifiedAt.Time
	}

	var parsedDeletedAt time.Time

	if deletedAt.Valid {
		parsedDeletedAt = deletedAt.Time
	}

	err := user.Hydrate(string(id), createdAt, username, usernameFingerprint, string(email), emailVerified, parsedEmailValidatedAt, blocked, deleted, parsedDeletedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UsersRepository) ValidateEmail(userId types.Id) error {
	_, txErr := r.db.Exec("UPDATE user_account SET email_verified = true, email_verified_at = $1 WHERE id = $2", time.Now().UTC(), userId)

	if txErr != nil {
		// todo: handle error better
		return txErr
	}

	return nil
}
