package http

import (
	"net/http"
	"strconv"

	"github.com/andisis/gocrud/src/api/user"
	"github.com/andisis/gocrud/src/helper"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	repo user.Repository
}

// UserHandler represent the httphandler for user
type UserHandler struct {
	UsrUsecase user.Usecase
}

// NewUserHandler will initialize the /users resources endpoint
func NewUserHandler(r *mux.Router, us user.Usecase) {
	handler := &UserHandler{
		UsrUsecase: us,
	}

	r.HandleFunc("/users", handler.FetchUser).Methods("GET")
	r.HandleFunc("/users/{userID}", handler.GetByID).Methods("GET")
}

// FetchUser handler
func (usr *UserHandler) FetchUser(w http.ResponseWriter, r *http.Request) {
	listUser, err := usr.UsrUsecase.Fetch(r.Context())
	if err != nil {
		dataErr := helper.ResponseError{
			Error: helper.ErrPayload{
				Code:    helper.GetStatusCode(err),
				Message: err.Error(),
			},
		}

		helper.JSONResponse(w, helper.GetStatusCode(err), dataErr)
		return
	}

	data := helper.ResponseSuccess{
		Data: listUser,
	}

	helper.JSONResponse(w, http.StatusOK, data)
	return
}

// GetByID handler
func (usr *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userID"])
	if err != nil {
		dataErr := helper.ResponseError{
			Error: helper.ErrPayload{
				Code:    helper.GetStatusCode(err),
				Message: err.Error(),
			},
		}

		helper.JSONResponse(w, helper.GetStatusCode(err), dataErr)
		return
	}

	user, err := usr.UsrUsecase.GetByID(r.Context(), int(id))
	if err != nil {
		dataErr := helper.ResponseError{
			Error: helper.ErrPayload{
				Code:    helper.GetStatusCode(err),
				Message: err.Error(),
			},
		}

		helper.JSONResponse(w, helper.GetStatusCode(err), dataErr)
		return
	}

	data := helper.ResponseSuccess{
		Data: user,
	}

	helper.JSONResponse(w, http.StatusOK, data)
	return
}
