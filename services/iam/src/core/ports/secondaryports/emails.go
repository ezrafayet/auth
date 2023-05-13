package secondaryports

type Email interface {
	NewUser() error
	VerifyEmail() error
}
