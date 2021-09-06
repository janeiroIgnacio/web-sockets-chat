package services

import (
	"LefkasChat/dtos"
	"LefkasChat/mapper"
	"LefkasChat/models"
	"LefkasChat/repository"
	"fmt"
	"github.com/gorilla/websocket"
)

type groupService struct {
	repository 	repository.MessageRepository
	msgService	ISendMessage
	mapper		mapper.MessageMapper
}

func (groupService) SendMessage(conn *websocket.Conn, message dtos.MessageDto){
	if err := conn.WriteJSON(message);  err != nil{
		fmt.Println(err)
		return
	}
}

func (this groupService) SendMessagesInFile(user *models.User) {
	for _ , message := range this.repository.GetMessagesById(user.ID){
		this.SendMessage(user.Conn, *this.mapper.MessageToDto(*message))
	}
}

var gService *groupService
func GetGroupService() ISendMessage {
	if gService != nil{
		return gService
	}
	gService = &groupService{
		repository: repository.GetMessageRepository(),
		msgService: GetMessageService(),
		mapper: mapper.MessageMapper{},
	}
	return gService
}