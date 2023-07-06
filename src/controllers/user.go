package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dami-pie/napi/models"
	"github.com/dami-pie/napi/src/responses"
	"github.com/dami-pie/napi/src/services"
)

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

	responses.EncodeJSON(res, http.StatusAccepted, struct {
		UserID uint64 `json:"userID"`
	}{
		UserID: userID,
	})
}
