package database

import "github.com/fatemehabsavaran/user-authentication.git/models"

type UserDbProvider interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (models.User, error)
	AddToken(userID uint, token *models.Token) error
	RemoveToken(token string) error
}
