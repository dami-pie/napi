package responses

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dami-pie/napi/models"
	"github.com/dami-pie/napi/src/config"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

func DecodeJSON[K models.Modelable](w http.ResponseWriter, r *http.Request, dst K) error {
	value := r.Header.Get("Content-Type")
	if value != "" {
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("O corpo da request contém JSON mal-formado (na posição %d)", syntaxError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "O corpo da request contém JSON mal-formado"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("O corpo da request contém um valor inválido para o campo %q (na posição %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{Status: http.StatusUnprocessableEntity, Msg: msg}

		case strings.HasPrefix(err.Error(), "responses: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "responses: unknown field ")
			msg := fmt.Sprintf("O corpo da request contém um campo irreconhecido chamado %s", fieldName)
			return &MalformedRequest{Status: http.StatusUnprocessableEntity, Msg: msg}

		case errors.Is(err, io.EOF):
			msg := "O corpo da request não pode estar vazio"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case err.Error() == "http: request body too large":
			msg := "O corpo da request não pode ter mais de 1MB"
			return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "O corpo da request só pode conter um único objeto JSON"
		return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
	}

	//TODO(will): é necessário verificar o tipo do erro de validação
	validate := validator.New()
	err = validate.Struct(dst)
	if err != nil {
		var invalid *validator.InvalidValidationError
		if errors.As(err, &invalid) {
			msg := "O corpo da request não está válido para processamento"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		}
	}

	return nil
}

func EncodeJSON(res http.ResponseWriter, statusCode int, data any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	if err := json.NewEncoder(res).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// TODO(will): essa função deveria retornar a string do erro ou uma própria string formatada? essa função pode ser substituída pelo http.Error()?

func Err(res http.ResponseWriter, statusCode int, err error) {
	EncodeJSON(res, statusCode, struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	})
}

// TODO(will): essa função aqui vai morrer depois que o broker chegar

func SendOpenCommand(r *http.Request) {
	body, _ := json.Marshal(map[string]any{
		"hash": config.OTPKey,
		"open": true,
	})
	payload := bytes.NewBuffer(body)
	http.Post(config.Device, "application/json", payload)
	log.Printf("[%s %s %s] -> %s\n", r.Method, r.RequestURI, r.Host, "Sent open command")
}
