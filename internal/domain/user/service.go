package user

import (
	"restaurant-evaluator/internal/adapters"
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"
)

type Service struct {
	Repository     Repository
	TokenGenerator adapters.TokenAdapter
}

func (s *Service) Create(newUser dto.NewUser) (string, error) {

	user, err := NewUser(newUser.Username, newUser.Email, newUser.PasswordHash, newUser.UserType)
	if err != nil {
		return "", err
	}

	token, err := s.TokenGenerator.GenerateToken(user.Username, user.Email)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	err = s.Repository.Save(user)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return token, nil
}
