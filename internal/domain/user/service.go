package user

import (
	"errors"
	"restaurant-evaluator/internal/adapters"
	"restaurant-evaluator/internal/dto"
)

type Service struct {
	Repository     Repository
	TokenGenerator adapters.TokenAdapter
}

func (s *Service) Create(newUser dto.NewUser) (string, error) {
	token, err := s.TokenGenerator.GenerateToken(newUser.Username, newUser.Email)

	if err != nil {
		return "", errors.New("token not generated")
	}

	user, err := NewUser(newUser.Username, newUser.Email, newUser.PasswordHash, newUser.UserType, token)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(user)
	if err != nil {
		return "", err
	}

	return user.Token, nil
}
