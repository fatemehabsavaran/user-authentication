package models

import "time"

type User struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Pass      string    `json:"pass"`
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Tokens    []Token
}

type Token struct {
	Token    string
	ExpireAt time.Time
	User     User
}
