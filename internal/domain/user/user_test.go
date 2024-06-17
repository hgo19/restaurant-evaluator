package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUser_SuccessCase(t *testing.T) {
	assert := assert.New(t)
	username := "user test"
	email := "foo@bar.com"
	password := "testPassword1@"
	userType := "Consumer"
	token := "token_for_test"

	user, _ := NewUser(username, email, password, userType, token)

	assert.Equal(user.ID, "id_test")
	assert.Equal(user.Username, "user test")
	assert.Equal(user.Email, "foo@bar.com")
	assert.Equal(Consumer, user.UserType)
}

func Test_NewUser_InvalidUserType(t *testing.T) {
	assert := assert.New(t)
	username := "user test"
	email := "foo@bar.com"
	password := "testPassword1@"
	userType := "InvalidType"
	token := "token_for_test"

	_, err := NewUser(username, email, password, userType, token)

	assert.Error(err)
	assert.EqualError(err, "invalid user type")
}
