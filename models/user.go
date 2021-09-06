package models

import (
	"github.com/gorilla/websocket"
)

type Tabler interface {
	TableName() string
}

type User struct {
	ID   int
	Name string
	Email string
	Messages []Message 		`gorm:"foreignKey:UserRefer"`
	Conn *websocket.Conn 	`gorm:"-"`
}

func (User) TableName() string {
	return "usuario"
}
