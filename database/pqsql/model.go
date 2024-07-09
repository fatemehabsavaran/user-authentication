package pqsql

import "time"

type User struct {
	Email     string `gorm:"unique"`
	Name      string
	Pass      string
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Tokens    []Token
}

type Token struct {
	Token    string
	ExpireAt time.Time
	User     User
	UserID   uint
	ID       uint `gorm:"primaryKey"`
}
