package user

import (
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newUser dto.NewUser) (string, error) {

	user, err := NewUser(newUser.Username, newUser.Email, newUser.PasswordHash, newUser.UserType)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(user)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return user.ID, nil
}
