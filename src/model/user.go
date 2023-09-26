package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Username  string     `gorm:"type:varchar(100)" json:"username"`
	Password  string     `gorm:"type:varchar(100)" json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type AuthUser struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
