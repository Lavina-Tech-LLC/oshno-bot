package handler

import (
	"encoding/json"
	"net/http"
	"oshno/models"
	"oshno/pkg/constants"

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

	newRequest := models.Request{
		FullName:    body.Name,
		Plan:        body.Plan,
		PhoneNumber: body.Phone,
		Address:     body.Address,
	}
	// message := messages.GenerateMessage(user)
	// fmt.Println(message)
	// _, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	// }
	message := newRequestMessageToGroup(newRequest)
	_, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)

	w.Header().Set("Content-Type", "text/json")
	w.WriteHeader(http.StatusOK)
	return
}
