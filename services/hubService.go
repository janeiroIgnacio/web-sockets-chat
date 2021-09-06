package services

import (
	"LefkasChat/models"
)

type HubService interface {
	Register(user *models.User)
	UnRegister(user *models.User)
	GetUserinHub(id int) (user *models.User)
}

type hubService struct {
	userService UserService
	hub *models.Hub
}

func (this *hubService)Register(user *models.User){
	this.hub.Users[user.ID] = user
}

func (this *hubService) UnRegister(user *models.User){
	delete(this.hub.Users, user.ID)
}

func (this *hubService) GetUserinHub(id int) (user *models.User){
	return this.hub.Users[id]
}


var hService *hubService

func GetHubService() HubService {
	if hService != nil {
		return hService
	}
	hService = &hubService{
		hub: models.NewHub(),
	}
	return hService
}

