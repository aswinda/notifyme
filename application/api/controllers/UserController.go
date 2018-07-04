package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aswinda/notifyme/application/api/interfaces"
	"github.com/go-chi/chi"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) GetUserDetailAction(response http.ResponseWriter, request *http.Request) {
	userIdParam := chi.URLParam(request, "user_id")
	userId, err := strconv.Atoi(userIdParam)

	detail, err := controller.GetUserDetail(userId)
	if err == nil {
		json.NewEncoder(response).Encode("Something went wrong!!")
	}

	json.NewEncoder(response).Encode(detail)
}
