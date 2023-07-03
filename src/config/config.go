package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Port               = ""
	JWTKey             []byte
	OTPKey             = ""
	Device             = ""
	CertFile           = ""
	KeyFile            = ""
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	JWTKey = []byte(os.Getenv("JWT_SECRET"))
	OTPKey = os.Getenv("OTP_SECRET")
	Device = os.Getenv("DEVICE")
	CertFile = os.Getenv("CERTIFICATE")
	KeyFile = os.Getenv("PRIVATE_KEY")
}
