package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/httphelpers"
	"iam/src/core/ports/primaryports"
	"iam/src/core/services"
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
	var args services.RegisterArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}

	answer, err := h.usersService.Register(args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusInternalServerError, "error", err.Error())(w, r)
		return
	}

	httphelpers.WriteSuccess(http.StatusOK, "User created successfully", answer)(w, r)
}

func (h *UsersHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	httphelpers.WriteSuccess(http.StatusOK, "Foobar", struct{}{})(w, r)
}