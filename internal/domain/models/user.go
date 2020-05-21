package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	Username  string    `gorm:"type:text;not null"`
	Password  string    `gorm:"type:text;not null"`
}

func (p *User) TableName() string {
	return "users"
}

func NewUser(id uuid.UUID, ca time.Time, ua time.Time, uname string, pass string) *User {
	return &User{
		Id:        id,
		CreatedAt: ca,
		UpdatedAt: ua,
		Username:  uname,
		Password:  pass,
	}
}
