package jwt

import (
	"fmt"
	"github.com/fatemehabsavaran/user-authentication.git/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTConfig struct {
	Secret string
	Expire time.Duration
}
type JWTService struct {
	config JWTConfig
}

func NewFromConfig(config JWTConfig) *JWTService {
	return &JWTService{
		config: config,
	}
}

func (j *JWTService) GenerateToken(user models.User) (models.Token, error) {
	ExpiredAt := time.Now().Add(j.config.Expire)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"userId":    user.ID,
		"expiredAt": ExpiredAt,
	})

	tokenStr, err := jwtToken.SignedString(j.config.Secret)
	if err != nil {
		return models.Token{}, err
	}
	return models.Token{
		Token:    tokenStr,
		ExpireAt: ExpiredAt,
	}, nil

}
func (j *JWTService) ValidateToken(token string) (models.User, error) {

	parsedJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return j.config.Secret, nil
	})
	if err != nil {
		return models.User{}, err
	}

	dataClaim, ok := parsedJwt.Claims.(jwt.MapClaims)
	if !ok {
		return models.User{}, fmt.Errorf("invalid sign method")
	}

	return models.User{
		ID: dataClaim["userId"].(uint),
	}, nil
}
