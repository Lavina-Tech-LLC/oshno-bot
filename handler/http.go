package handler

import (
	"encoding/json"
	"net/http"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/pkg/gateways"

	tele "gopkg.in/telebot.v3"
)

func (h BotHandler) Test(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h BotHandler) SendToGroupRequest(w http.ResponseWriter, req *http.Request) {
	var body models.AIBidRequest

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, err := gateways.GetHistoryChat(body.ChatId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.storage.GetUserByChatId(body.ChatId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := messages.GenerateMessage(user)
	_, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	return
}
