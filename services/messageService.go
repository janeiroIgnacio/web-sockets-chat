package services

import (
	"LefkasChat/dtos"
	"LefkasChat/mapper"
	"LefkasChat/models"
	"LefkasChat/repository"
	"fmt"
	"github.com/gorilla/websocket"
)

type messageService struct {
	repository 	repository.MessageRepository
	mapper 		mapper.MessageMapper
}

func (this messageService) SendMessage(conn *websocket.Conn, message dtos.MessageDto){
	if err := conn.WriteJSON(message);  err != nil{
		fmt.Println(err)
		return
	}
}

func (this messageService) SendMessagesInFile(user *models.User) {
	for _ , message := range this.repository.GetMessagesById(user.ID){
		this.SendMessage(user.Conn, *this.mapper.MessageToDto(*message))
	}
}

var mService *messageService
func GetMessageService() ISendMessage {
	if mService != nil{
		return mService
	}
	mService = &messageService{
		repository: repository.GetMessageRepository(),
		mapper: mapper.MessageMapper{},
	}
	return mService
}