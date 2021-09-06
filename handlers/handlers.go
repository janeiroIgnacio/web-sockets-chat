package handlers

import (
	service "LefkasChat/services"
	"net/http"
)

type Handlers interface {
	SetupRoutes()
}

type handlers struct {
	userService service.UserService
}

func (this handlers) SetupRoutes() {

	//hub := models.NewPool()
	//go this.hubService.Start(hub)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		this.userService.GetWebSocket(w,r)
	})

}

var handle *handlers

func GetInstance() Handlers{
	if handle != nil {
		return handle
	}
	return &handlers{
		userService: service.GetUserService(),
	}
}
