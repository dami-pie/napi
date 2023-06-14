package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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

func init() {
	//durante o modo de testes, os testes não procuram pela .env na raiz do projeto, mas sim na pasta onde o teste é executado
	//https://github.com/joho/godotenv/issues/43
	err := godotenv.Load(os.Getenv("TEST_WD") + "\\.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	JWTKey = []byte(os.Getenv("JWT_SECRET"))
	OTPKey = os.Getenv("OTP_SECRET")
	Device = os.Getenv("DEVICE")
	CertFile = os.Getenv("CERTIFICATE")
	KeyFile = os.Getenv("PRIVATE_KEY")
}
