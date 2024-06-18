package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	username = "user test"
	email    = "foo@bar.com"
	password = "testPassword1@"
	token    = "token_for_test"
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
