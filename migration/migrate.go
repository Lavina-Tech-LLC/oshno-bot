package migration

import (
	"oshno/db"
	"oshno/models"
	"oshno/pkg/logger"

	"go.uber.org/zap"
)

func Migrate() {
	err := db.ConnectDB().AutoMigrate(
		&models.User{},
		&models.Advertisement{},
		&models.Request{},
		&models.ChangePlanRequest{},
		&models.AddPlanRequest{},
	)

	if err != nil {
		logger.Logger().Error("error to auto migrate", zap.Error(err))
		panic(err)
	}

	logger.Logger().Info("auto migration is successfully done!")
}
