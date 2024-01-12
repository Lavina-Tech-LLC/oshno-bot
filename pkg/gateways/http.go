package gateways

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oshno/models"
)

type RespBody struct {
	Data    []Data `json:"data"`
	Message string `json:"message"`
	IsOk    bool   `json:"isOk"`
}

type Data struct {
	Role      string `json:"role"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

func GetHistoryChat(chatId string) (*RespBody, error) {
	url := "https://chatly-back.lavina.tech/chats/history/" + fmt.Sprintf("%s", chatId)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Origin", "https://oshno.lavina.tech")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: tr,
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}

	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Print the response body as a string
	var history RespBody
	fmt.Println(string(body))
	err = json.Unmarshal(body, &history)

	if err != nil {
		fmt.Println("error in parse json", err)
		return nil, err
	}

	return &history, nil
}

func (r RespBody) GenerateMessage(user models.User) string {
	msgTg := fmt.Sprintf("Чат айди: %s\nПолное имя: %s\nНикнейм: @%s\nТел номер: %s\n\n", user.AIChatId, user.FullName, user.Nickname, user.PhoneNumber)

	for _, message := range r.Data {
		msgUser := fmt.Sprintf("Роль %s:\n    %s\n\n", message.Role, message.Content)
		msgTg = msgTg + msgUser
	}

	return msgTg
}
