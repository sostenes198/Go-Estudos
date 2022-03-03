package controllers

import (
	"devbook/src/controllers/view_models"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/repositories/contracts"
	"devbook/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// List Lista todos usuários
func List(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	repository := repositories.NewUserRepository()
	users, err := repository.List(contracts.UserListFilter{NameOrNick: nameOrNick})
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if len(users) == 0 {
		responses.JSON(w, http.StatusNotFound, nil)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetById obtem usuário por Id
func GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}

	repository := repositories.NewUserRepository()
	user, err := repository.GetById(id)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if user == nil {
		responses.JSON(w, http.StatusNotFound, nil)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// Create Cria um usuário
func Create(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userViewModel view_models.CreateUser
	if err = json.Unmarshal(bodyRequest, &userViewModel); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	user := models.NewUser(models.ParamsUser{
		Name:     userViewModel.Name,
		Nick:     userViewModel.Nick,
		Email:    userViewModel.Email,
		Password: userViewModel.Password,
		CreateAt: userViewModel.CreateAt,
	})
	errs := user.Validate()
	if len(errs) > 0 {
		responses.Erros(w, http.StatusBadRequest, errs)
		return
	}

	repository := repositories.NewUserRepository()
	if err := repository.Create(user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// Update Atualiza um usuário
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}

// Delete excluir um usuário
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
