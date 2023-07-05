package models

import (
	"errors"
	"github.com/dami-pie/napi/src/config"
	"time"

	"github.com/pquerna/otp/totp"
)

type OTP struct {
	Tempo time.Time `json:"time" validate:"required"`
	Key   string    `json:"key" validate:"required"`
}

func (otp *OTP) ValidateKey() (bool, error) {
	if totp.Validate(otp.Key, config.OTPKey) {
		return true, nil
	} else {
		return false, errors.New("OTP inválida")
	}
}
