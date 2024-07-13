package user

import (
	"restaurant-evaluator/internal/adapters"
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"
)

type Service struct {
	Repository Repository
	Encrypter  adapters.Encrypter
	TokenAuth  adapters.TokenAuth
}

func (s *Service) Create(newUser dto.User) (string, error) {

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

func (s *Service) GenerateToken(email string) (string, error) {
	user, err := s.Repository.FindByEmail(email)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	if user == nil {
		return "", internalerrors.NotFound
	}

	token, err := s.TokenAuth.Generate(user.Email)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return token, nil
}
