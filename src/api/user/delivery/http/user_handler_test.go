package http_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	userHandler "github.com/andisis/gocrud/src/api/user/delivery/http"
	"github.com/andisis/gocrud/src/api/user/mocks"
	"github.com/andisis/gocrud/src/api/user/usecase"
	"github.com/andisis/gocrud/src/helper"
	"github.com/andisis/gocrud/src/model"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	var mockUser model.User

	mockUsecase := new(mocks.Usecase)
	mockListUser := make([]*model.User, 0)
	mockListUser = append(mockListUser, &mockUser)
	mockUsecase.On("Fetch", mock.Anything).Return(mockListUser, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := userHandler.UserHandler{
		UsrUsecase: mockUsecase,
	}
	handler.FetchUser(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	mockUsecase.On("Fetch", mock.Anything).Return(nil, helper.ErrInternalServerError)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := userHandler.UserHandler{
		UsrUsecase: mockUsecase,
	}
	handler.FetchUser(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockUser model.User

	mockUsecase := new(mocks.Usecase)
	id := mockUser.ID
	mockUsecase.On("GetByID", mock.Anything, id).Return(&mockUser, nil)

	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	req = mux.SetURLVars(req, map[string]string{"userID": strconv.Itoa(id)})
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := userHandler.UserHandler{
		UsrUsecase: mockUsecase,
	}
	handler.GetByID(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestGetByIDError500(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	id := "id"
	mockUsecase.On("GetByID", mock.Anything, id).Return(nil, helper.ErrInternalServerError)

	req, err := http.NewRequest("GET", "/users/"+id, nil)
	req = mux.SetURLVars(req, map[string]string{"userID": id})
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := userHandler.UserHandler{
		UsrUsecase: mockUsecase,
	}
	handler.GetByID(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetByIDError404(t *testing.T) {
	mockUsecase := new(mocks.Usecase)
	id := 0
	mockUsecase.On("GetByID", mock.Anything, id).Return(nil, helper.ErrNotFound)

	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	req = mux.SetURLVars(req, map[string]string{"userID": strconv.Itoa(id)})
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := userHandler.UserHandler{
		UsrUsecase: mockUsecase,
	}
	handler.GetByID(rec, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestNewUserHandler(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	stdn := usecase.NewUserUsecase(mockUserRepo, time.Second*2)

	r := mux.NewRouter()
	userHandler.NewUserHandler(r, stdn)

	mockUserRepo.AssertExpectations(t)
}
