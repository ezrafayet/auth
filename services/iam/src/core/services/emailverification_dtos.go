package services

type SendVerificationCodeArgs struct {
	UserId string `json:"userId"`
}

type SendVerificationCodeAnswer struct{}

type ConfirmEmailArgs struct {
	VerificationCode string `json:"verificationCode"`
}

type ConfirmEmailAnswer struct {
	AuthorizationCode string `json:"authorizationCode"`
}
