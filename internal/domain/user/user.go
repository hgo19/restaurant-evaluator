package user

import (
	"errors"
	internalerrors "restaurant-evaluator/internal/internal-errors"

	"github.com/rs/xid"
)

type userAppType string

const (
	Consumer userAppType = "Consumer"
	Owner    userAppType = "Owner"
)

type User struct {
	ID           string      `validate:"required"`
	Username     string      `validate:"min=5,max=24"`
	Email        string      `validate:"email"`
	PasswordHash string      `validate:"required"`
	UserType     userAppType `validate:"required"`
	Token        string      `validate:"required"`
}

func NewUser(username string, email string, password string, userType string, token string) (*User, error) {
	ut, errUserType := parseUserAppType(userType)
	if errUserType != nil {
		return nil, errUserType
	}

	user := &User{
		ID:           xid.New().String(),
		Username:     username,
		Email:        email,
		PasswordHash: password,
		UserType:     ut,
		Token:        token,
	}

	err := internalerrors.ValidateStruct(user)

	if err == nil {
		return user, nil
	}

	return nil, err
}

func parseUserAppType(userType string) (userAppType, error) {
	switch userType {
	case "Consumer":
		return Consumer, nil
	case "Owner":
		return Owner, nil
	default:
		return "", errors.New("invalid user type")
	}
}
