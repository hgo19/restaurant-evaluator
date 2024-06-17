package user

import "errors"

type userAppType string

const (
	Consumer userAppType = "Consumer"
	Owner    userAppType = "Owner"
)

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	UserType     userAppType
	Token        string
}

func NewUser(username string, email string, password string, userType string) (*User, error) {
	ut, err := parseUserAppType(userType)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           "id_test",
		Username:     username,
		Email:        email,
		PasswordHash: password,
		UserType:     ut,
		Token:        "token_test",
	}, nil
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
