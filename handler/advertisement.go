package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/pkg/utils"
	"sync"
	"time"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func (h BotHandler) Admin(c tele.Context) error {
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	h.logger.Info("admin started", zap.Any("user", user), zap.Any("chat_id", c.Chat()))

	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}
	err = h.storage.UpdatePhase(user.ID, 20)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}
	switch user.Language {
	case constants.Tajik:
		return c.Send(fmt.Sprintf("Забонро интихоб кунед, паёмро ба кадом забон фиристед?"), models.LanguageMarkup)
	default:
		return c.Send(fmt.Sprintf("Выберите язык, на каком языку вы хотите отправить пост?"), models.LanguageMarkup)
	}
}

func (h BotHandler) OnMedia(c tele.Context) error {
	msg := c.Message()
	body, err := json.Marshal(msg)
	if err != nil {
		return c.Send("Sorry. Bot's failure!, Try restarting the bot.", models.StartMarkup)
	}

	err = h.storage.UploadMedia(models.Advertisement{Body: body})
	if err != nil {
		return c.Send("Sorry. Bot's failure!, Try restarting the bot.", models.StartMarkup)
	}
	return c.Send(fmt.Sprintf("Подтверждаете ли вы отправку текста объявления пользователям?"), models.ConfirmAdMarkup)

}

func (h BotHandler) uploadMedia(msg tele.Message) error {

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = h.storage.UploadMedia(models.Advertisement{Body: body})
	if err != nil {
		return err
	}

	return nil
}

func (h BotHandler) SendAd(c tele.Context) error {
	c.Delete()
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	h.logger.Info("sending photo started", zap.Uint("userId", user.ID))

	media, err := h.storage.GetMedia()
	if err != nil {
		return c.Send("Sorry. Bot's failure!, Try restarting the bot.", models.StartMarkup)
	}

	var msg = &tele.Message{}
	err = json.Unmarshal(media.Body, msg)
	if err != nil {
		return c.Send("Sorry. Bot's failure!, Try restarting the bot.", models.StartMarkup)
	}

	users, err := h.storage.GetAllUsers()
	if err != nil {
		return c.Send("Sorry. Bot's failure!, Try restarting the bot.", models.StartMarkup)
	}
	const (
		MaxConcurrentRequests = 30
		RequestsPerMinute     = 100

		DelayBetweenRequests = time.Minute / time.Duration(RequestsPerMinute)
	)

	var wg sync.WaitGroup
	sem := make(chan struct{}, MaxConcurrentRequests)

	for _, u := range users {
		wg.Add(1)
		sem <- struct{}{}

		go func(u models.User) {
			defer wg.Done()
			defer func() { <-sem }()

			if media.Language != u.Language {
				return
			}

			if u.TelegramUserId == c.Chat().ID {
				return
			}

			_, err := c.Bot().Send(&tele.User{ID: u.TelegramUserId}, utils.CreateMessage(msg), msg.ReplyMarkup)
			if err != nil {
				if !errors.Is(err, tele.ErrBlockedByUser) || !errors.Is(err, tele.ErrUserIsDeactivated) || !errors.Is(err, tele.ErrChatNotFound) {
					h.logger.Error("error in recieve media", zap.Error(err))
				}
				return
			}
		}(u)
		time.Sleep(DelayBetweenRequests)
	}

	wg.Wait()

	err = h.storage.UpdatePhase(user.ID, 0)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}
	h.logger.Info("All ads sent successfully!")
	switch user.Language {
	case constants.Tajik:
		return c.Send(fmt.Sprintf("Эълон ба ҳама корбарон фиристода мешавад!"), models.MenuMarkupTg)
	default:
		return c.Send(fmt.Sprintf("Реклама всем юзерам отправлена!"), models.MenuMarkupRu)
	}
}

func (h BotHandler) UnsendAd(c tele.Context) error {
	c.Delete()
	return nil
}
