package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	Email   string `json:"email"`
	GroupID uint64 `json:"groupID"`
}

func (user *User) ValidateUser() error {
	i := strings.LastIndexByte(user.Email, '@')
	host := user.Email[i+1:]

	if i == -1 {
		return errors.New("o email não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o email inserido é inválido")
	}

	if host != "ecomp.poli.br" {
		return errors.New("domínio inválido")
	}

	user.Email = strings.TrimSpace(user.Email)
	return nil
}
