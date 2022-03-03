package controllers

import (
	"devbook/src/authentication"
	"devbook/src/controllers/view_models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"devbook/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var loginViewModel view_models.LoginVw
	if err = json.Unmarshal(bodyRequest, &loginViewModel); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	repository := repositories.NewUserRepository()
	user, err := repository.GetUserToLogin(loginViewModel.Email)
	if err != nil{
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.ValidatePassword(user.Password, loginViewModel.Password); err != nil{
		responses.Erro(w, http.StatusUnauthorized, nil)
		return
	}

	token, err := authentication.CreateToken(user.Id)
	if err != nil{
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, struct {
		Token string `json:"token,omitempty"`
	}{
		Token: token,
	})
}
