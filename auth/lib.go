package auth

import "github.com/fatemehabsavaran/user-authentication.git/models"

type AuthProvider interface {
	GenerateToken(user models.User) (models.Token, error)
	ValidateToken(token string) (models.User, error)
}
