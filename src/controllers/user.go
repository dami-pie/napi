package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dami-pie/napi/models"
	"github.com/dami-pie/napi/src/responses"
	"github.com/dami-pie/napi/src/services"
)

// TODO(will): as status codes respondidas deveriam variar de acordo com os erros

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		responses.Err(res, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.User
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		responses.Err(res, http.StatusBadRequest, erro)
		return
	}

	userID, erro := services.CreateUser(usuario)
	if erro != nil {
		responses.Err(res, http.StatusBadRequest, erro)
		return
	}

	responses.EncodeJSON(res, http.StatusCreated, struct {
		UserID uint64 `json:"userID"`
	}{
		UserID: userID,
	})
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	untreatedId := mux.Vars(req)["id"]
	id, err := strconv.Atoi(untreatedId)

	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	//Ele existe?
	_, err = services.UserExists(uint64(id))

	if err != nil {
		http.NotFound(res, req)
		return
	}

	err = services.DeleteUser(uint64(id))

	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	res.WriteHeader(http.StatusNoContent)
	return
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	untreatedId := mux.Vars(req)["id"]
	id, err := strconv.Atoi(untreatedId)

	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	var newUser models.User
	err = responses.DecodeJSON(res, req, newUser)

	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	//Existe algum usuário usando o email novo?
	_, err = services.GetUserByEmail(newUser.Email)

	if errors.Is(err, sql.ErrNoRows) {
		services.UpdateUser(uint64(id), newUser)
		res.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(res, "update: já existe um usuário usando o mesmo email", http.StatusConflict)
	return
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	untreatedId := mux.Vars(req)["id"]
	id, err := strconv.Atoi(untreatedId)

	if err != nil {
		responses.Err(res, http.StatusUnprocessableEntity, err)
		return
	}

	user, err := services.GetUser(uint64(id))
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	responses.EncodeJSON(res, http.StatusOK, user)
	return
}
