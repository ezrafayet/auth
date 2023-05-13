package registry

import (
	"database/sql"
	"iam/src/core/services"
	"iam/src/interface/handlers"
	"iam/src/interface/repository"
)

type Registry struct {
	UsersHandler *handlers.UsersHandler
}

func NewRegistry(db *sql.DB) *Registry {
	// Repositories
	var usersRepository = repository.NewUsersRepository(db)

	// Services
	var usersService = services.NewUserService(usersRepository)

	// Handlers
	var usersHandler = handlers.NewUsersHandler(usersService)

	// Registry
	return &Registry{
		UsersHandler: usersHandler,
	}
}
