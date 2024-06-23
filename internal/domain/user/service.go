package user

import (
	"restaurant-evaluator/internal/dto"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newUser dto.NewUser) (string, error) {
	user, err := NewUser(newUser.Username, newUser.Email, newUser.PasswordHash, newUser.UserType, newUser.Token)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
