package main

import (
	"log"
	"oshno/config"
	"oshno/db"
	"oshno/handler"
	"oshno/migration"
	"oshno/pkg/cron"
	"oshno/pkg/logger"
	"oshno/storage"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	logger := logger.Logger()
	conf := config.Config()
	migration.Migrate()
	db := db.ConnectDB()

	var token string
	if conf.Telegram.TelegramToken == "" {
		token = "5446910492:AAFFFOQzEYMRMhTzDbxg5fsTgN-3aNLJgYw"
	} else {
		token = conf.Telegram.TelegramToken
	}
	storage := storage.NewStorage(db, logger)
	pref := tele.Settings{
		Token:       token,
		Poller:      &tele.LongPoller{Timeout: 10 * time.Second},
		Synchronous: false,
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// socket.Connection()

	h := handler.NewBotHandler(bot, logger, storage)

	cron.StartCronjob(storage, logger, bot)
	handler.Start(h)
}
