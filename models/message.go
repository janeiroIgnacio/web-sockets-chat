package models

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	SenderId 	int 	`json:"send_to"`
	ReceiverId 	int 	`json:"from"`
	Message		string	`json:"message"`
	Type		int     `json:"type"`
}

func (Message) TableName() string {
	return "mensajes"
}
