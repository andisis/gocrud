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
		helper.RespondWithError(w, helper.GetStatusCode(err), err.Error())
		return
	}

	data := helper.Response{
		Status:  int64(helper.GetStatusCode(err)),
		Message: "Success",
		Data:    listUser,
	}
	helper.RespondWithJSON(w, http.StatusOK, data)
	return
}

// GetByID handler
func (usr *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userID"])
	if err != nil {
		helper.RespondWithError(w, helper.GetStatusCode(err), err.Error())
		return
	}

	user, err := usr.UsrUsecase.GetByID(r.Context(), int(id))
	if err != nil {
		helper.RespondWithError(w, helper.GetStatusCode(err), err.Error())
		return
	}

	data := helper.Response{
		Status:  int64(helper.GetStatusCode(err)),
		Message: "Success",
		Data:    user,
	}
	helper.RespondWithJSON(w, http.StatusOK, data)
	return
}
