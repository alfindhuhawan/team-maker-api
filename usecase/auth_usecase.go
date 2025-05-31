package usecase

import (
	"errors"
	"team-maker-api/domain"
)

type AuthUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(userRepo domain.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}

func (uc *AuthUsecase) Authenticate(username, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, errors.New("authentication failed")
	}
	return user, nil
}
