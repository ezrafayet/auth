package registry

import (
	"database/sql"
	"iam/src/core/services"
	"iam/src/interface/handlers"
	"iam/src/interface/repository"
)

type Registry struct {
	UsersHandler            *handlers.UsersHandler
	VerificationCodeHandler *handlers.EmailVerificationHandler
}

func NewRegistry(db *sql.DB) *Registry {
	// Repositories
	var usersRepository = repository.NewUsersRepository(db)
	var verificationCodeRepository = repository.NewVerificationCodeRepository(db)
	var authorizationCodeRepository = repository.NewAuthorizationCodeRepository(db)

	// Services
	var usersService = services.NewUserService(usersRepository)
	var verificationCodeService = services.NewEmailVerificationService(usersRepository, verificationCodeRepository, authorizationCodeRepository)

	// Handlers
	var usersHandler = handlers.NewUsersHandler(usersService)
	var verificationCodeHandler = handlers.NewEmailVerificationHandler(verificationCodeService)

	// Registry
	return &Registry{
		UsersHandler:            usersHandler,
		VerificationCodeHandler: verificationCodeHandler,
	}
}
