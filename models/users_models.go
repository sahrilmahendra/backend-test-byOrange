package models

import (
	"time"

	"gorm.io/gorm"
)

// struktur data users
type Users struct {
	gorm.Model
	Email     string    `gorm:"unique" json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	LastLogin time.Time `json:"lastlogin" form:"lastlogin"`
	Token     string
	Items     []Items
}

// struktur get user *
type GetUser struct {
	ID    uint
	Email string
}

// struktur get login user
type GetLoginUser struct {
	ID    uint
	Token string
}
