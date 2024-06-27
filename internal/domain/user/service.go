package user

import (
	"restaurant-evaluator/internal/adapters"
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"
)

type Service struct {
	Repository Repository
	Encrypter  adapters.Encrypter
}

func (s *Service) Create(newUser dto.NewUser) (string, error) {

	passwordHash, err := s.Encrypter.HashPassword(newUser.Password)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	user, err := NewUser(newUser.Username, newUser.Email, passwordHash, newUser.UserType)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(user)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return user.ID, nil
}
