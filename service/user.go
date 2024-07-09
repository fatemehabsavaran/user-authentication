package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/fatemehabsavaran/user-authentication.git/database"
	"github.com/fatemehabsavaran/user-authentication.git/models"
)

type UserService struct {
	Db database.UserDbProvider
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

	if err = s.Db.CreateUser(&models.User{
		Email: email,
		Name:  name,
		Pass:  hashedPass,
	}); err != nil {
		return models.User{}, err
	}
	return models.User{}, err

}
