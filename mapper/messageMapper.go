package mapper

import (
	"LefkasChat/dtos"
	"LefkasChat/models"
)

type MessageMapper struct {

}

func (this *MessageMapper) RequestToResponse(message *dtos.MessageDto, userId int){

	message.SenderId = userId

}

func (this *MessageMapper) MessageToDto(message models.Message) *dtos.MessageDto{

	return &dtos.MessageDto{SenderId: message.SenderId, ReceiverId: message.ReceiverId, Message: message.Message}

}