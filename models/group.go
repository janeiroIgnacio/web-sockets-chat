package models

type Group struct {
	Group	map[int][]*User
}

func NewGroup() *Group {
	return &Group{
		Group:	make(map[int][]*User),
	}
}
