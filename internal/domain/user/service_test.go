package user

import (
	"errors"
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

var (
	repository = new(repositoryMock)
	userDto    = dto.NewUser{
		Username:     "valid_name",
		Email:        "valid@email.com",
		PasswordHash: "valid_passwordHash",
		UserType:     "Consumer",
		Token:        "valid_token",
	}
	service = Service{
		Repository: repository,
	}
)

func Test_Create_User(t *testing.T) {
	assert := assert.New(t)
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

	id, err := service.Create(userDto)

	assert.Nil(err)
	assert.NotNil(id)
	repository.AssertExpectations(t)
}

func Test_Create_User_ValidateDomainsErrors(t *testing.T) {
	assert := assert.New(t)
	userDto.Username = ""

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.Equal(err.Error(), "Username is required with min 5")
	repository.AssertExpectations(t)
}

func Test_Create_User_ValidateRepositoryErrors(t *testing.T) {
	assert := assert.New(t)
	userDto.Username = "valid_username"
	repositoryMockErr := new(repositoryMock)
	repositoryMockErr.On("Save", mock.Anything).Return(errors.New("error to persist data"))
	service.Repository = repositoryMockErr

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.Equal(err.Error(), "error to persist data")
	repository.AssertExpectations(t)
}
