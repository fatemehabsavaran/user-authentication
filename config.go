package main

import (
	"github.com/fatemehabsavaran/user-authentication.git/auth/jwt"
	"github.com/fatemehabsavaran/user-authentication.git/database/pgsql"
	"github.com/fatemehabsavaran/user-authentication.git/io/http"
)

type AppConfig struct {
	GinConfig http.GinConfig    `json:"ginConfig"`
	DbConfig  pgsql.PgsqlConfig `json:"dbConfig"`
	JWtConfig jwt.JWTConfig     `json:"JWTConfig"`
}
