package auth

import (
	"errors"
	"fmt"
	"github.com/dami-pie/napi/src/config"
	"github.com/dgrijalva/jwt-go/v4"
	"net/http"
	"strings"
	"time"
)

// Gera o token do usuário baseado no seu email e hora atual
func GenerateToken(userEmail string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["auth"] = true
	permissions["expiresAt"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["email"] = userEmail

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.JWTKey))
}

func ValidateToken(req *http.Request) error {
	tokenString := extractToken(req)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return fmt.Errorf("ValidateToken: %w", err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// Extrai o email do usuário presente dentro do token
func ExtractUserEmail(req *http.Request) (string, error) {
	tokenString := extractToken(req)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return "", fmt.Errorf("ExtractUserEmail: %w", err)
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail := permissoes["email"]
		if str, ok := userEmail.(string); ok {
			return str, nil
		} else {
			return "", errors.New("email não reconhecido")
		}
	}

	return "", errors.New("não foi possível extrair o email do usuário do token JWT")
}

// Extrai o token da requisição "Bearer 'token'"
func extractToken(req *http.Request) string {
	token := req.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("returnVerificationKey: método de assinatura inesperado: %v", token.Header["alg"])
	}
	return config.JWTKey, nil
}
