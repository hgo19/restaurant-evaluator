package user

import (
	"errors"
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type encrypterMock struct {
	mock.Mock
}

func (r *encrypterMock) HashPassword(password string) (string, error) {
	args := r.Called(password)
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
	repository = new(repositoryMock)
	encrypter  = new(encrypterMock)
	userDto    = dto.NewUser{
		Username: "valid_name",
		Email:    "valid@email.com",
		Password: "valid_passwordHash",
		UserType: "Consumer",
	}
	service = Service{
		Repository: repository,
		Encrypter:  encrypter,
	}
)

func Test_Create_User(t *testing.T) {
	assert := assert.New(t)
	encrypter.On("HashPassword", userDto.Password).Return("hashed_password", nil)
	repository.On("Save", mock.MatchedBy(func(user *User) bool {
		if user.Username != userDto.Username {
			return false
		} else if user.Email != userDto.Email {
			return false
		} else if user.PasswordHash != "hashed_password" {
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
	encrypter.AssertExpectations(t)

}

func Test_Create_User_ValidateDomainsErrors(t *testing.T) {
	assert := assert.New(t)
	encrypter.On("HashPassword", userDto.Password).Return("hashed_password", nil)
	userDto.Email = ""

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.Equal(err.Error(), "Email is invalid")
	encrypter.AssertExpectations(t)
}

func Test_Create_User_ValidateRepositoryErrors(t *testing.T) {
	assert := assert.New(t)
	userDto.Email = "valid@email.com"
	encrypter.On("HashPassword", userDto.Password).Return("hashed_password", nil)
	repositoryMockErr := new(repositoryMock)
	repositoryMockErr.On("Save", mock.Anything).Return(errors.New("error to persist data"))
	service.Repository = repositoryMockErr

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.Equal(err, internalerrors.ErrInternal)
	encrypter.AssertExpectations(t)

}

func Test_Create_User_ValidateEncrypterErrors(t *testing.T) {
	assert := assert.New(t)
	userDto.Email = "valid@email.com"
	encrypterErr := new(encrypterMock)
	encrypterErr.On("HashPassword", userDto.Password).Return("", errors.New("error to encrypt password"))
	service.Encrypter = encrypterErr

	_, err := service.Create(userDto)

	assert.NotNil(err)
	assert.EqualError(err, internalerrors.ErrInternal.Error())
	encrypterErr.AssertExpectations(t)
}
