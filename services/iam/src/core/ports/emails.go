package ports

type Email interface {
	NewUser() error
	VerifyEmail() error
}
