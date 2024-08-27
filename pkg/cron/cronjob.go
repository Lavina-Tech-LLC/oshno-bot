package cron

import (
	"fmt"
	"oshno/config"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/storage"
	"time"

	"github.com/robfig/cron"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Cronjob struct {
	storage storage.Storage
	log     zap.Logger
	bot     *tele.Bot
}

func NewCronjob(storage *storage.Storage, log *zap.Logger, bot *tele.Bot) Cronjob {
	return Cronjob{log: *log, storage: *storage, bot: bot}
}

func StartCronjob(storage *storage.Storage, log *zap.Logger, bot *tele.Bot) {
	cronjob := NewCronjob(storage, log, bot)

	fmt.Println("start cron")

	c := cron.New()
	c.AddFunc("@every 0h1m", cronjob.SendMessageAIConfirm)

	go c.Start()
}

func (cj Cronjob) SendMessageAIConfirm() {
	cfg := config.Config()
	timeDuration := cfg.Cron.Duration
	if timeDuration == 0 {
		timeDuration = 10
	}

	// get all users and check last message send or not
	users, err := cj.storage.GetAllUsersAINotConfirmed()
	if err != nil {

		cj.log.Error("error in update user: ", zap.Error(err))
		return
	}

	for _, user := range users {
		now := time.Now()
		lastTime := user.AILastMessageTime.Add(time.Duration(timeDuration) * time.Minute)
		fmt.Println("now: ", now, "last time: ", lastTime, lastTime.Before(now))
		if lastTime.Before(now) {

			switch user.Language {
			case constants.Tajik:

				cj.bot.Send(&tele.User{ID: user.TelegramUserId}, "Оё маълумот муфид буд?", models.AIConfirmMarkupTg)
			default:

				cj.bot.Send(&tele.User{ID: user.TelegramUserId}, "Была ли информация полезной?", models.AIConfirmMarkupRu)
			}

			user.AIConfirmSended = true
			err := cj.storage.UpdateUser(user.ID, user)
			if err != nil {

				cj.log.Error("error in update user: ", zap.Error(err))
				continue
			}
		}
	}
}
