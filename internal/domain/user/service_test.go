package user

import (
	"errors"
	"restaurant-evaluator/internal/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type tokenAdapterMock struct {
	mock.Mock
}

func (t *tokenAdapterMock) GenerateToken(username string, email string) (string, error) {
	args := t.Called(username, email)
	return args.String(0), args.Error(1)
}

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(user *User) error {
	args := r.Called(user)
	return args.Error(0)
}

var (
	repository   = new(repositoryMock)
	tokenAdapter = new(tokenAdapterMock)
	userDto      = dto.NewUser{
		Username:     "valid_name",
		Email:        "valid@email.com",
		PasswordHash: "valid_passwordHash",
		UserType:     "Consumer",
	}
	service = Service{
		Repository:     repository,
		TokenGenerator: tokenAdapter,
	}
)

func Test_Create_User(t *testing.T) {
	assert := assert.New(t)
	tokenAdapter.On("GenerateToken", userDto.Username, userDto.Email).Return("mocked_token", nil)
	repository.On("Save", mock.MatchedBy(func(user *User) bool {
		if user.Username != userDto.Username {
			return false
		} else if user.Email != userDto.Email {
			return false
		} else if user.PasswordHash != userDto.PasswordHash {
			return false
		} else if user.UserType != userAppType(userDto.UserType) {
			return false
		}
		return true
	})).Return(nil)

	token, err := service.Create(userDto)

	assert.Nil(err)
	assert.NotNil(token)
	repository.AssertExpectations(t)
	tokenAdapter.AssertExpectations(t)

}

func Test_Create_User_ValidateDomainsErrors(t *testing.T) {
	assert := assert.New(t)
	userDto.PasswordHash = ""
	tokenAdapter.On("GenerateToken", userDto.Username, userDto.Email).Return("mocked_token", nil)

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.Equal(err.Error(), "PasswordHash is required")
	repository.AssertExpectations(t)
	tokenAdapter.AssertExpectations(t)

}

func Test_Create_User_ValidateRepositoryErrors(t *testing.T) {
	assert := assert.New(t)
	userDto.PasswordHash = "valid_passwordHash"
	tokenAdapter.On("GenerateToken", userDto.Username, userDto.Email).Return("mocked_token", nil)
	repositoryMockErr := new(repositoryMock)
	repositoryMockErr.On("Save", mock.Anything).Return(errors.New("error to persist data"))
	service.Repository = repositoryMockErr

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.Equal(err.Error(), "error to persist data")
	repository.AssertExpectations(t)
	tokenAdapter.AssertExpectations(t)

}

func Test_Create_User_ValidateTokenGenerate(t *testing.T) {
	assert := assert.New(t)

	tokenAdapter.On("GenerateToken", userDto.Username, userDto.Email).Return("mocked_token", nil)
	repository.On("Save", mock.Anything).Return(nil)

	token, err := service.Create(userDto)

	assert.Nil(err)
	assert.Equal("mocked_token", token)
	tokenAdapter.AssertExpectations(t)
	repository.AssertExpectations(t)
}
