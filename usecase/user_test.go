package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/szczynk/MyGram/helpers"
	mocks "github.com/szczynk/MyGram/mocks/models"
	"github.com/szczynk/MyGram/models"
	"github.com/szczynk/MyGram/usecase"
)

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tempNewUser := models.User{
			Age:      18,
			Email:    "test@gmail.com",
			Password: "1234567",
			Username: "test123",
		}

		mockUserRepo := new(mocks.UserRepo)
		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		mockUserRepo.On("Register", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

		err := userUsecase.Register(context.Background(), &tempNewUser)
		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempNewUser)
		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("invalid", func(t *testing.T) {
		tempNewUser := models.User{
			Age:      7,
			Email:    "test@gmail.com",
			Password: "1234567",
			Username: "test123",
		}

		mockUserRepo := new(mocks.UserRepo)
		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		mockUserRepo.On("Register", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

		err := userUsecase.Register(context.Background(), &tempNewUser)
		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempNewUser)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tempRegisteredUser := models.User{
			Email:    "test@gmail.com",
			Password: "1234567",
		}
		mockUserRepo := new(mocks.UserRepo)
		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		mockUserRepo.On("Login", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

		err := userUsecase.Login(context.Background(), &tempRegisteredUser)
		assert.NoError(t, err)

		originalPassword := "$2a$08$PjOHykW6CkN2H4LUnmbZ7.ZIgrGRrQnfvNrzRktnFG0XEOBeksJKC"
		isValid := helpers.CompareHash(originalPassword, tempRegisteredUser.Password)
		assert.True(t, isValid, "invalid password")

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		tempRegisteredUser := models.User{
			Email:    "test@gmail.com",
			Password: "123456",
		}

		mockUserRepo := new(mocks.UserRepo)
		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		mockUserRepo.On("Login", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

		err := userUsecase.Login(context.Background(), &tempRegisteredUser)
		assert.NoError(t, err)

		originalPassword := "$2a$08$PjOHykW6CkN2H4LUnmbZ7.ZIgrGRrQnfvNrzRktnFG0XEOBeksJKC"
		isValid := helpers.CompareHash(originalPassword, tempRegisteredUser.Password)
		assert.False(t, isValid, "valid password")

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tempInputUpdateUser := models.User{
			Username: "test12345",
		}
		tempUpdatedUser := models.User{
			ID:       3,
			Age:      18,
			Email:    "test@gmail.com",
			Username: "test12345",
		}
		mockUserRepo := new(mocks.UserRepo)
		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("models.User"), mock.AnythingOfType("uint")).Return(tempUpdatedUser, nil)

		tempInputUpdateUser.Age = 18
		tempInputUpdateUser.Email = "test@gmail.com"
		tempInputUpdateUser.Password = "1234567"

		user, err := userUsecase.Update(context.Background(), tempInputUpdateUser, 3)
		assert.NoError(t, err)
		assert.Equal(t, tempUpdatedUser, user, "result has to be same")

		_, err = govalidator.ValidateStruct(tempInputUpdateUser)
		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("invalid", func(t *testing.T) {
		tempInputUpdateUser := models.User{
			Username: "",
		}
		tempUpdatedUser := models.User{
			ID:       3,
			Age:      18,
			Email:    "test@gmail.com",
			Username: "test12345",
		}
		mockUserRepo := new(mocks.UserRepo)
		userUsecase := usecase.NewUserUsecase(mockUserRepo)
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("models.User"), mock.AnythingOfType("uint")).Return(models.User{}, errors.New("Username is required"))

		tempInputUpdateUser.Age = 18
		tempInputUpdateUser.Email = "test@gmail.com"
		tempInputUpdateUser.Password = "1234567"

		user, err := userUsecase.Update(context.Background(), tempInputUpdateUser, 3)
		fmt.Println(user, err)
		assert.Error(t, err)
		assert.NotEqual(t, tempUpdatedUser, user, "it should be not updated")

		_, err = govalidator.ValidateStruct(tempInputUpdateUser)
		fmt.Println(err)
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}
