package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andisis/gocrud/src/api/user/mocks"
	"github.com/andisis/gocrud/src/api/user/usecase"
	"github.com/andisis/gocrud/src/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	mockUser := &model.User{
		ID:        1,
		Fullname:  "Andi Siswanto",
		Email:     "andisis92@gmail.com",
		Username:  "andisis",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockListUser := make([]*model.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything).Return(mockListUser, nil).Once()
		usr := usecase.NewUserUsecase(mockUserRepo, time.Second*2)
		list, err := usr.Fetch(context.TODO())

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything).Return(nil, errors.New("unexpexted error")).Once()
		usr := usecase.NewUserUsecase(mockUserRepo, time.Second*2)
		list, err := usr.Fetch(context.TODO())

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	mockUser := model.User{
		ID:        1,
		Fullname:  "Andi Siswanto",
		Email:     "andisis92@gmail.com",
		Username:  "andisis",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(&mockUser, nil).Once()
		usr := usecase.NewUserUsecase(mockUserRepo, time.Second*2)
		user, err := usr.GetByID(context.TODO(), mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("unexpected")).Once()
		usr := usecase.NewUserUsecase(mockUserRepo, time.Second*2)
		user, err := usr.GetByID(context.TODO(), mockUser.ID)

		assert.Error(t, err)
		assert.Nil(t, user)
		mockUserRepo.AssertExpectations(t)
	})
}
