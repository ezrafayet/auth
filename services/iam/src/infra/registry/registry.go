package registry

import (
	"database/sql"
	"iam/src/core/services"
	"iam/src/infra/emailprovider"
	"iam/src/interface/dbrepository"
	"iam/src/interface/emailrepository"
	"iam/src/interface/handlers"
)

type Registry struct {
	UsersHandler            *handlers.UsersHandler
	VerificationCodeHandler *handlers.EmailVerificationHandler
	AuthenticationHandler   *handlers.AuthenticationHandler
	AuthorizationHandler    *handlers.AuthorizationHandler
}

func NewRegistry(db *sql.DB, emailProvider *emailprovider.Provider) *Registry {
	// Repositories
	var usersRepository = dbrepository.NewUsersRepository(db)
	var verificationCodeRepository = dbrepository.NewVerificationCodeRepository(db)
	var authorizationCodeRepository = dbrepository.NewAuthorizationCodeRepository(db)
	var refreshTokenRepository = dbrepository.NewRefreshTokenRepository(db)

	// Email repository
	var emailRepository = emailrepository.NewEmailRepository(emailProvider)

	// Services
	var usersService = services.NewUsersService(usersRepository, emailRepository)
	var verificationCodeService = services.NewEmailVerificationService(usersRepository, verificationCodeRepository, authorizationCodeRepository, emailRepository)
	var authenticationService = services.NewAuthenticationService(usersRepository, authorizationCodeRepository, emailRepository, refreshTokenRepository)
	var authorizationService = services.NewAuthorizationService()

	// Handlers
	var usersHandler = handlers.NewUsersHandler(usersService)
	var verificationCodeHandler = handlers.NewEmailVerificationHandler(verificationCodeService)
	var authenticationHandler = handlers.NewAuthenticationHandler(authenticationService)
	var authorizationHandler = handlers.NewAuthorizationHandler(authorizationService)

	// Registry
	return &Registry{
		UsersHandler:            usersHandler,
		VerificationCodeHandler: verificationCodeHandler,
		AuthenticationHandler:   authenticationHandler,
		AuthorizationHandler:    authorizationHandler,
	}
}
