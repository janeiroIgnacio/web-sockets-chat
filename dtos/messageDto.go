package dtos

type MessageDto struct {
	Type		int		`json:"type"`
	SenderId 	int 	`json:"send_from"`
	ReceiverId  int     `json:"send_to"`
	Message		string	`json:"message"`
}
