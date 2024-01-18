package models

import "fmt"

type AIRespBody struct {
	Data    []AIData `json:"data"`
	Message string   `json:"message"`
	IsOk    bool     `json:"isOk"`
}

type AIData struct {
	Role      string `json:"role"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type AIBidRequest struct {
	ChatId  string `json:"chat_id"`
	Subject string `json:"subject"`
}

func (r AIRespBody) GenerateMessage(user User, subject string) string {
	msgTg := fmt.Sprintf("Чат айди: %s\nТема: %s\nПолное имя: %s\nНикнейм: @%s\nТел номер: %s\n\n", user.AIChatId, subject, user.FullName, user.Nickname, user.PhoneNumber)

	for _, message := range r.Data {
		msgUser := fmt.Sprintf("Роль %s:\n    %s\n\n", message.Role, message.Content)
		msgTg = msgTg + msgUser
	}

	return msgTg
}
