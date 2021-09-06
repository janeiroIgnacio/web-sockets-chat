package models

type Hub struct {
	Register   chan *User
	Unregister chan *User
	Users      map[int]*User
	Broadcast  chan Message
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *User),
		Unregister: make(chan *User),
		Users:    make(map[int]*User),
		Broadcast:  make(chan Message),
	}
}
