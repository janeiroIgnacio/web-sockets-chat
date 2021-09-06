package repository

import (
	"LefkasChat/config"
	"LefkasChat/models"
	"github.com/jinzhu/gorm"
)

type MessageRepository interface {
	SaveMessage(message models.Message)
	GetAllMessage()
	GetMessagesById(moduleId int) []*models.Message
}

type messageRepository struct {
	db *gorm.DB
}

func (this *messageRepository) GetAllMessage() {

}

func (this *messageRepository) GetMessagesById(messageId int) []*models.Message {
	var messageRemain []*models.Message
	_ = this.db.Where(&models.Message{ReceiverId: messageId}).Find(&messageRemain)
	return messageRemain
}

func (this *messageRepository) SaveMessage(message models.Message){
}

var messageRepo *messageRepository

func GetMessageRepository() MessageRepository {
	if messageRepo != nil{
		return messageRepo
	}
	messageRepo = &messageRepository{
		db: config.GetDBInstance(),
	}
	return messageRepo
}
