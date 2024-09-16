package gateways

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"oshno/config"
)

const telegramAPI = "https://api.telegram.org/bot%s/%s"

type CreateTopicRequest struct {
	ChatID string `json:"chat_id"`
	Name   string `json:"name"`
}

type CreateTopicRes struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	ThreadId  int32  `json:"message_thread_id"`
	Name      string `json:"name"`
	IconColor int32  `json:"icon_color"`
}

type SendMessageRequest struct {
	ChatID          int64  `json:"chat_id"`
	Text            string `json:"text"`
	MessageThreadID int64  `json:"message_thread_id"`
}

func CreateForumTopic(name string) (CreateTopicRes, error) {
	var topicRes CreateTopicRes

	cfg := config.Config()
	url := fmt.Sprintf(telegramAPI, cfg.Telegram.TelegramToken, "createForumTopic")
	chatId := cfg.Telegram.OpertorChatId
	if chatId == 0 {
		chatId = -1002413165254
	}

	topic := CreateTopicRequest{
		ChatID: fmt.Sprintf("%d", chatId),
		Name:   name,
	}

	body, err := json.Marshal(topic)
	if err != nil {
		return CreateTopicRes{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return CreateTopicRes{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CreateTopicRes{}, fmt.Errorf("failed to create topic: %s", resp.Status)
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateTopicRes{}, err
	}
	err = json.Unmarshal(resBody, &topicRes)
	if err != nil {
		return CreateTopicRes{}, err
	}

	fmt.Println("Res body: ", string(resBody))
	fmt.Println("Topic created successfully!")

	return topicRes, nil
}

func SendMessageToTopic(topicID int64, message string) error {
	cfg := config.Config()
	url := fmt.Sprintf(telegramAPI, cfg.Telegram.TelegramToken, "sendMessage")
	chatId := cfg.Telegram.OpertorChatId
	if chatId == 0 {
		chatId = -1002413165254
	}

	msg := SendMessageRequest{
		ChatID:          chatId,
		Text:            message,
		MessageThreadID: topicID,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("topic send message", string(resBody))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}

	fmt.Println("Message sent successfully to topic!")
	return nil
}
