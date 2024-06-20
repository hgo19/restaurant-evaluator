package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	username      = "user test"
	email         = "foo@bar.com"
	password      = "testPassword1@"
	token         = "token_for_test"
	validUserType = "Consumer"
)

func Test_NewUser_CustomerSuccessCase(t *testing.T) {
	assert := assert.New(t)
	userCustomerType := "Consumer"

	user, _ := NewUser(username, email, password, userCustomerType, token)

	assert.NotEmpty(user.ID)
	assert.Equal(user.Username, "user test")
	assert.Equal(user.Email, "foo@bar.com")
	assert.Equal(Consumer, user.UserType)
}

func Test_NewUser_OwnerSuccessCase(t *testing.T) {
	assert := assert.New(t)
	userOwnerType := "Owner"

	user, _ := NewUser(username, email, password, userOwnerType, token)

	assert.Equal(Owner, user.UserType)
}

func Test_NewUser_InvalidUserType(t *testing.T) {
	assert := assert.New(t)
	userInvalidType := "InvalidTypeUser"

	_, err := NewUser(username, email, password, userInvalidType, token)

	assert.Error(err)
	assert.EqualError(err, "invalid user type")
}

func Test_NewUser_InvalidUsername(t *testing.T) {
	assert := assert.New(t)
	invalidUsername := ""

	_, err := NewUser(invalidUsername, email, password, validUserType, token)

	assert.EqualError(err, "Username is required with min 5")
}

func Test_NewUser_InvalidEmail(t *testing.T) {
	assert := assert.New(t)
	invalidEmail := "invalid_email"

	_, err := NewUser(username, invalidEmail, password, validUserType, token)

	assert.EqualError(err, "Email is invalid")
}

func Test_NewUser_InvalidPasswordHash(t *testing.T) {
	assert := assert.New(t)
	invalidPasswordHash := ""

	_, err := NewUser(username, email, invalidPasswordHash, validUserType, token)

	assert.EqualError(err, "PasswordHash is required")
}

func Test_NewUser_InvalidToken(t *testing.T) {
	assert := assert.New(t)
	invalidToken := ""

	_, err := NewUser(username, email, password, validUserType, invalidToken)

	assert.EqualError(err, "Token is required")
}
