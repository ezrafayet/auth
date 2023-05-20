package registry

import (
	"database/sql"
	"iam/src/core/services"
	"iam/src/infra/emailprovider"
	"iam/src/interface/emailrepository"
	"iam/src/interface/handlers"
	"iam/src/interface/repository"
)

type Registry struct {
	UsersHandler            *handlers.UsersHandler
	VerificationCodeHandler *handlers.EmailVerificationHandler
}

func NewRegistry(db *sql.DB, emailProvider *emailprovider.Provider) *Registry {
	// Repositories
	var usersRepository = repository.NewUsersRepository(db)
	var verificationCodeRepository = repository.NewVerificationCodeRepository(db)
	var authorizationCodeRepository = repository.NewAuthorizationCodeRepository(db)

	// Email repository
	var emailRepository = emailrepository.NewEmailRepository(emailProvider)

	// Services
	var usersService = services.NewUserService(usersRepository, emailRepository)
	var verificationCodeService = services.NewEmailVerificationService(usersRepository, verificationCodeRepository, authorizationCodeRepository, emailRepository)

	// Handlers
	var usersHandler = handlers.NewUsersHandler(usersService)
	var verificationCodeHandler = handlers.NewEmailVerificationHandler(verificationCodeService)

	// Registry
	return &Registry{
		UsersHandler:            usersHandler,
		VerificationCodeHandler: verificationCodeHandler,
	}
}
