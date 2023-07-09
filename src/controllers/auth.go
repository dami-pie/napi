package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/dami-pie/napi/models"
	"github.com/dami-pie/napi/src/responses"
	auth "github.com/dami-pie/napi/src/services"
)

func Login(res http.ResponseWriter, req *http.Request) {
	var usuario models.User
	err := responses.DecodeJSON(res, req, usuario)
	if err != nil {
		var mr *responses.MalformedRequest
		log.Printf("[%s %s %s] -> %s\n", req.Method, req.RequestURI, req.Host, err.Error())
		if errors.As(err, &mr) {
			responses.Err(res, mr.Status, err)
		} else {
			responses.Err(res, http.StatusInternalServerError, err)
		}
		return
	}

	if err = usuario.ValidateUser(); err != nil {
		responses.Err(res, http.StatusBadRequest, err)
		return
	}

	userToken, err := auth.GenerateToken(usuario.Email)
	if err != nil {
		responses.Err(res, http.StatusInternalServerError, err)
		return
	}

	responses.EncodeJSON(res, http.StatusAccepted, struct {
		Token string `json:"token"`
	}{
		Token: userToken,
	})
}

// TODO: Precisamos guardar o email e o tempo da entrada do usuário após a abertura da porta. Isso vai ser implementado junto com o banco.
func ValidateOTP(res http.ResponseWriter, req *http.Request) {
	_, err := auth.ExtractUserEmail(req)
	if err != nil {
		responses.Err(res, http.StatusUnauthorized, err)
	}

	var otp models.OTP
	err = responses.DecodeJSON(res, req, otp)
	if err != nil {
		var mr *responses.MalformedRequest
		log.Printf("[%s %s %s] -> %s\n", req.Method, req.RequestURI, req.Host, err.Error())
		if errors.As(err, &mr) {
			responses.Err(res, mr.Status, err)
		} else {
			responses.Err(res, http.StatusInternalServerError, err)
		}
		return
	}

	if isValid, err := otp.ValidateKey(); !isValid {
		responses.Err(res, http.StatusUnauthorized, err)
		return
	} else {
		responses.SendOpenCommand(req)
		responses.EncodeJSON(res, http.StatusOK, err)
		return
	}
}
