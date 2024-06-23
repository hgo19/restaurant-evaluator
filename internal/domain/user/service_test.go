package user

import (
	"restaurant-evaluator/internal/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(user *User) error {
	args := r.Called(user)
	return args.Error(0)
}

func Test_Create_User(t *testing.T) {
	assert := assert.New(t)
	userDto := dto.NewUser{
		Username:     "valid_name",
		Email:        "valid@email.com",
		PasswordHash: "valid_passwordHash",
		UserType:     "Consumer",
		Token:        "valid_token",
	}
	repository := new(repositoryMock)
	repository.On("Save", mock.MatchedBy(func(user *User) bool {
		if user.Username != userDto.Username {
			return false
		} else if user.Email != userDto.Email {
			return false
		} else if user.PasswordHash != userDto.PasswordHash {
			return false
		} else if user.UserType != userAppType(userDto.UserType) {
			return false
		} else if user.Token != userDto.Token {
			return false
		}
		return true
	})).Return(nil)

	service := Service{
		Repository: repository,
	}

	id, err := service.Create(userDto)

	assert.Nil(err)
	assert.NotNil(id)

	repository.AssertExpectations(t)
}
