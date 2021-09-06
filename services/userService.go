package services

import (
	"LefkasChat/dtos"
	"LefkasChat/mapper"
	"LefkasChat/models"
	"LefkasChat/repository"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ISendMessage interface {
	SendMessagesInFile(user *models.User)
	SendMessage(conn *websocket.Conn, message dtos.MessageDto)
}
type UserService interface {
	Read(user *models.User)
	GetWebSocket(w http.ResponseWriter, r *http.Request)
}

type userService struct {
	repository 	repository.MessageRepository
	hubService 	HubService
	mapper		mapper.MessageMapper
	userRepo	repository.UserRepository
	messagesServices map[int]ISendMessage
}

func (this userService) Read(user *models.User) {
	var message dtos.MessageDto

	defer func() {
		this.hubService.UnRegister(user)
		user.Conn.Close()
	}()
	for {
		_, p, _ := user.Conn.ReadMessage()
		json.Unmarshal(p,&message)
		this.messagesServices[message.Type].SendMessage(this.hubService.GetUserinHub(message.ReceiverId).Conn, message)
	}

}
func (this userService) GetWebSocket(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.FormValue("id"))
	user, _ := this.userRepo.GetUserById(id)
	user.Conn, _ = upgrader.Upgrade(w, r, nil)

	this.hubService.Register(&user)
	this.messagesServices[1].SendMessagesInFile(&user)
	go this.Read(&user)

}


var uService *userService
func GetUserService() UserService {
	if uService != nil{
		return uService
	}
	uService = &userService{
		userRepo: repository.GetUserRepository(),
		repository: repository.GetMessageRepository(),
		hubService: GetHubService(),
		messagesServices: make(map[int]ISendMessage),
	}
	uService.messagesServices[1] = GetMessageService()
	uService.messagesServices[2] = GetGroupService()
	return uService
}