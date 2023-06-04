package services

import (
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
)

type UsersService struct {
	usersRepository secondaryports.UsersRepository
	emailRepository secondaryports.EmailRepository
}

// NewUsersService initializes a new instance of UsersService with the provided repositories
func NewUsersService(usersRepository secondaryports.UsersRepository, emailRepository secondaryports.EmailRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
		emailRepository: emailRepository,
	}
}

// Register registers a new user
// /!\ It currently handles magic link registration only, but it can be extended
func (s *UsersService) Register(args primaryports.RegisterArgs) (primaryports.RegisterAnswer, error) {
	authType, err := types.ParseAndValidateAuthType(args.AuthType)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	email, err := types.ParseAndValidateEmail(args.Email)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	username, err := types.ParseAndValidateUsername(args.Username)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	if !args.HasAcceptedTerms || args.AcceptedTermsVersion == "" {
		return primaryports.RegisterAnswer{}, errors.New(apperrors.RefusedTerms)
	}

	user := model.NewUserModel(username, email, types.NewTimestamp())

	userAuthMethod := model.NewUserMagicLinkAuth(user.Id, authType)

	termsAndConditions := model.NewUserTermsAndConditionsModel(user.Id)

	err = termsAndConditions.Accept(args.HasAcceptedTerms, args.AcceptedTermsVersion)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	marketingPreferences := model.NewUserMarketingPreferencesModel(user.Id)

	if args.HasAcceptedNewsletter {
		marketingPreferences.AcceptNewsletter()
	}

	if args.HasAcceptedMarketing {
		marketingPreferences.AcceptMarketing()
	}

	// quite a few things are happening here, we should consider breaking it down to multiple operations while
	// keeping the benefits of the transaction - not urgent
	err = s.usersRepository.SaveUser(user, userAuthMethod, types.RoleUser, termsAndConditions, marketingPreferences)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	go func() {
		// this email could later be handled by a worker
		_ = s.emailRepository.WelcomeNewUser(user.Email, user.Username)
	}()

	return primaryports.RegisterAnswer{
		UserId: string(user.Id),
	}, nil
}
