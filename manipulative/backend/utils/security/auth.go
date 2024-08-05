package security

import (
	"net/http"
	. "tinyrdm/backend/types"
)

var users = map[string]User{
	"admin": {"admin", "$2a$14$zTmGzUoYbV8S6HzYPX/WJ.bfPnX60J8fHfchH8RXdx/D6GcQ/2laG", "admin"}, // Password: admin
}

var roles = map[string]Role{
	"admin": {"admin", []string{"read", "write", "delete"}},
	"user":  {"user", []string{"read"}},
}

func Login(username, password string) (string, error) {
	user, exists := users[username]
	if !exists || !CheckPasswordHash(password, user.Password) {
		return "", http.ErrNoCookie
	}

	return GenerateJWT(user.Username, user.Role)
}

func AuthenticatedEndpoint(token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	return "Hello, " + claims.Username + "! Your role is: " + claims.Role, nil
}
