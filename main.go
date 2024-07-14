package main

import (
	"github.com/fatemehabsavaran/user-authentication.git/auth/jwt"
	"github.com/fatemehabsavaran/user-authentication.git/database/pgsql"
	"github.com/fatemehabsavaran/user-authentication.git/io/http"
	"github.com/fatemehabsavaran/user-authentication.git/service"
	"github.com/spf13/viper"
	"strings"
)

var appConfig AppConfig

func init() {
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(err)
	}
}

func main() {
	psqlService, err := pgsql.NewPgsqlFromConfig(appConfig.DbConfig)
	if err != nil {
		panic(err)
	}
	jwtService := jwt.NewFromConfig(appConfig.JWtConfig)
	userService := service.UserService{
		Db:   psqlService,
		Auth: jwtService,
	}
	userController := http.NewUserController(appConfig.GinConfig, userService)
	userController.RegisterRoutes()
	if err := userController.Start(); err != nil {
		panic(err)
	}

}
