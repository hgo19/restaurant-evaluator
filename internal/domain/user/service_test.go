package user

import (
	"errors"
	"restaurant-evaluator/internal/dto"
	internalerrors "restaurant-evaluator/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type tokenAuthMock struct {
	mock.Mock
}

func (t *tokenAuthMock) Generate(email string) (string, error) {
	args := t.Called(email)
	return args.String(0), args.Error(1)
}

type encrypterMock struct {
	mock.Mock
}

func (e *encrypterMock) HashPassword(password string) (string, error) {
	args := e.Called(password)
	return args.String(0), args.Error(1)
}

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(user *User) error {
	args := r.Called(user)
	return args.Error(0)
}

func (r *repositoryMock) FindByEmail(email string) (*User, error) {
	args := r.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*User), args.Error(1)
	}
	return nil, args.Error(1)
}

var (
	userBaseEntity = User{
		ID:           "valid_id",
		Username:     "valid_username",
		Email:        "valid@email.com",
		PasswordHash: "valid_password",
		UserType:     Consumer,
	}
	repository = new(repositoryMock)
	encrypter  = new(encrypterMock)
	tokenAuth  = new(tokenAuthMock)
	userDto    = dto.User{
		Username: "valid_name",
		Email:    "valid@email.com",
		Password: "valid_passwordHash",
		UserType: "Consumer",
	}
	service = Service{
		Repository: repository,
		Encrypter:  encrypter,
		TokenAuth:  tokenAuth,
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

func Test_GenerateToken_Success(t *testing.T) {
	assert := assert.New(t)
	validToken := "valid@token"

	service.Repository = repository
	service.Encrypter = encrypter
	repository.On("FindByEmail", email).Return(&userBaseEntity, nil)
	tokenAuth.On("Generate", email).Return(validToken, nil)

	token, _ := service.GenerateToken(email)

	assert.NotNil(token)
	assert.Equal(token, validToken)
	repository.AssertExpectations(t)
	tokenAuth.AssertExpectations(t)

}

func Test_GenerateToken_Repository_Error(t *testing.T) {
	assert := assert.New(t)
	email := "invalid@email.com"
	repositoryMockErr := new(repositoryMock)
	service.Repository = repositoryMockErr
	repositoryMockErr.On("FindByEmail", email).Return(nil, errors.New("database found error"))

	_, err := service.GenerateToken(email)

	assert.NotNil(err)
	assert.EqualError(err, internalerrors.ErrInternal.Error())
	repository.AssertExpectations(t)
}

func Test_GenerateToken_Repository_NotFound(t *testing.T) {
	assert := assert.New(t)
	email := "invalid@email.com"
	service.Repository = repository
	repository.On("FindByEmail", email).Return(nil, nil)

	_, err := service.GenerateToken(email)

	assert.NotNil(err)
	assert.EqualError(err, internalerrors.NotFound.Error())
	repository.AssertExpectations(t)
}

func Test_GenerateToken_TokenAuth_Error(t *testing.T) {
	assert := assert.New(t)
	email := "valid@email.com"
	tokenAuthMockErr := new(tokenAuthMock)
	service.TokenAuth = tokenAuthMockErr
	repository.On("FindByEmail", email).Return(&userBaseEntity, nil)
	tokenAuthMockErr.On("Generate", email).Return("", errors.New("Some error to generate token"))

	_, err := service.GenerateToken(email)

	assert.NotNil(err)
	assert.EqualError(err, internalerrors.ErrInternal.Error())
	repository.AssertExpectations(t)
}
