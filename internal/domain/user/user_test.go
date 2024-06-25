package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	username      = "user test"
	email         = "foo@bar.com"
	password      = "testPassword1@"
	validUserType = "Consumer"
)

func Test_NewUser_CustomerSuccessCase(t *testing.T) {
	assert := assert.New(t)
	userCustomerType := "Consumer"

	user, _ := NewUser(username, email, password, userCustomerType)

	assert.NotEmpty(user.ID)
	assert.Equal(user.Username, "user test")
	assert.Equal(user.Email, "foo@bar.com")
	assert.Equal(Consumer, user.UserType)
}

func Test_NewUser_OwnerSuccessCase(t *testing.T) {
	assert := assert.New(t)
	userOwnerType := "Owner"

	user, _ := NewUser(username, email, password, userOwnerType)

	assert.Equal(Owner, user.UserType)
}

func Test_NewUser_InvalidUserType(t *testing.T) {
	assert := assert.New(t)
	userInvalidType := "InvalidTypeUser"

	_, err := NewUser(username, email, password, userInvalidType)

	assert.Error(err)
	assert.EqualError(err, "invalid user type")
}

func Test_NewUser_InvalidUsername(t *testing.T) {
	assert := assert.New(t)
	invalidUsername := ""

	_, err := NewUser(invalidUsername, email, password, validUserType)

	assert.EqualError(err, "Username is required with min 5")
}

func Test_NewUser_InvalidEmail(t *testing.T) {
	assert := assert.New(t)
	invalidEmail := "invalid_email"

	_, err := NewUser(username, invalidEmail, password, validUserType)

	assert.EqualError(err, "Email is invalid")
}

func Test_NewUser_InvalidPasswordHash(t *testing.T) {
	assert := assert.New(t)
	invalidPasswordHash := ""

	_, err := NewUser(username, email, invalidPasswordHash, validUserType)

	assert.EqualError(err, "PasswordHash is required")
}
