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
	//durante o modo de testes, os testes não procuram pela .env na raiz do projeto, mas sim na pasta onde o teste é executado
	//https://github.com/joho/godotenv/issues/43
	err := godotenv.Load(os.Getenv("TEST_WD") + "\\.env")

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
