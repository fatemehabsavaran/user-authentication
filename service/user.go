package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/fatemehabsavaran/user-authentication.git/auth"
	"github.com/fatemehabsavaran/user-authentication.git/database"
	"github.com/fatemehabsavaran/user-authentication.git/models"
)

type UserService struct {
	Db   database.UserDbProvider
	Auth auth.AuthProvider
}

func (s *UserService) SignUp(email string, name string, pass string) (models.User, error) {
	checkUser, err := s.Db.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if checkUser.Email == email {
		return models.User{}, fmt.Errorf("user exists in database")
	}

	hasher := sha256.New()
	hasher.Write([]byte(pass))
	hashedPass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	user := models.User{
		Email: email,
		Name:  name,
		Pass:  hashedPass,
	}
	if err = s.Db.CreateUser(&user); err != nil {
		return models.User{}, err
	}
	token, err := s.GenerateToken(user)
	if err != nil {
		return models.User{}, err
	}
	user.Tokens = []models.Token{token}
	return user, err

}

func (s *UserService) GenerateToken(user models.User) (models.Token, error) {
	token, err := s.Auth.GenerateToken(user)
	if err != nil {
		return models.Token{}, err
	}
	return token, nil
}
