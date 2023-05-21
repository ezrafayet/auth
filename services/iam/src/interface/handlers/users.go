package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/apperrors"
	"iam/pkg/httphelpers"
	"iam/src/core/ports/primaryports"
	"net/http"
)

type UsersHandler struct {
	usersService primaryports.UsersService
}

func NewUsersHandler(usersService primaryports.UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

func (h *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	var args primaryports.RegisterArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		return
	}

	answer, err := h.usersService.Register(args)

	if err != nil {
		switch err.Error() {
		case apperrors.InvalidAuthType:
			fallthrough
		case apperrors.InvalidUsername:
			fallthrough
		case apperrors.InvalidEmail:
			httphelpers.WriteError(http.StatusBadRequest, "error", err.Error())(w, r)
		case apperrors.UsernameAlreadyExists:
			fallthrough
		case apperrors.EmailAlreadyExists:
			httphelpers.WriteError(http.StatusConflict, "error", err.Error())(w, r)
		default:
			fmt.Println(err)
			httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		}
		return
	}

	httphelpers.WriteSuccess(http.StatusCreated, "User created successfully", answer)(w, r)
}

func (h *UsersHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	httphelpers.WriteSuccess(http.StatusOK, "Foobar", struct{}{})(w, r)
}
