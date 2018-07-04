package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/aswinda/notifyme/interfaces"
	"github.com/go-chi/chi"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) GetUserDetail(response http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(req, "user_id")

	detail, err := controller.GetUserDetail(userId)
	if err := nil {
		json.NewEncoder(res).Encode("Something went wrong!!")
	}

	json.NewEncoder(res).Encode(detail)
}