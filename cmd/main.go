package main

import (
	"log"
	"oshno/config"
	"oshno/db"
	"oshno/handler"
	"oshno/migration"
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
	if conf.Telegram.TelegramToken != "" {
		token = "6785340587:AAEB2tcCgrd3o196hP5_1PE5a_AHj-hzfvc"
	} else {
		token = "5446910492:AAFFFOQzEYMRMhTzDbxg5fsTgN-3aNLJgYw"
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
	handler.Start(h)
}
