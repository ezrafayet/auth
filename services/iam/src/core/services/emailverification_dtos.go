package services

type SendVerificationCodeArgs struct {
	UserId string `json:"userId"`
}

type SendVerificationCodeAnswer struct{}

type ConfirmVerificationCodeArgs struct {
	VerificationCode string `json:"verificationCode"`
}

type ConfirmVerificationCodeAnswer struct {
	AuthorizationCode string `json:"authorizationCode"`
}
