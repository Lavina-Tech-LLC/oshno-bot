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
	Name     string   `json:"name"`
	Plan     string   `json:"plan"`
	Phone    string   `json:"phoneNumber"`
	Address  string   `json:"address"`
	Messages []AIData `json:"messages"`
}

func (r AIRespBody) GenerateMessage(user User) string {
	msgTg := fmt.Sprintf("Чат айди: %s\nПолное имя: %s\nНикнейм: @%s\nТел номер: %s\n\n", user.AIChatId, user.FullName, user.Nickname, user.PhoneNumber)

	for _, message := range r.Data {
		msgUser := fmt.Sprintf("Роль %s:\n    %s\n\n", message.Role, message.Content)
		msgTg = msgTg + msgUser
	}

	return msgTg
}