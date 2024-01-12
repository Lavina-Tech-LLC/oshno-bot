package storage

import (
	"errors"
	"oshno/models"
	"oshno/pkg/constants"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Storage struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewStorage(db *gorm.DB, log *zap.Logger) *Storage {
	return &Storage{
		db:  db,
		log: log,
	}
}

func (s *Storage) CreateUser(payload models.User) error {
	err := s.db.Create(&payload).Error
	if err != nil {
		s.log.Error("error in create user", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) GetUserByTgId(id int64) (models.User, error) {
	var user models.User
	var count int64
	err := s.db.Where("telegram_user_id = ?", id).Model(&models.User{}).Count(&count).Error
	if err != nil {
		s.log.Error("error in get user failed", zap.Error(err))
		return user, err
	}

	if count < 1 {
		return user, errors.New(constants.UserNotRegist)
	}

	err = s.db.Where("telegram_user_id = ?", id).Find(&user).Error
	if err != nil {
		s.log.Error("get user by telegram id failed", zap.Error(err))
		return user, err
	}
	return user, nil
}

func (s *Storage) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	if err != nil {
		s.log.Error("get all users failed", zap.Error(err))
		return users, err
	}
	return users, nil
}

func (s *Storage) UpdateLanguage(id uint, language string) error {
	err := s.db.Model(&models.User{}).Where("id = ?", id).Update("language", language).Error
	if err != nil {
		s.log.Error("update language failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) UpdatePhase(id uint, phase int) error {
	err := s.db.Model(&models.User{}).Where("id = ?", id).Update("user_phase", phase).Error
	if err != nil {
		s.log.Error("update language failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) UpdateAIChatId(id uint, chatId string) error {
	err := s.db.Model(&models.User{}).Where("id = ?", id).Update("ai_chat_id", chatId).Error
	if err != nil {
		s.log.Error("update ai chatId failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) UploadMedia(payload models.Advertisement) error {
	var count int64

	err := s.db.Model(&models.Advertisement{}).Where("name = ?", constants.MediaTableName).Count(&count).Error

	if err != nil {
		s.log.Error("error in upload media", zap.Error(err))
		return err
	}

	payload.Name = constants.MediaTableName
	if count == 0 {
		err = s.db.Create(&payload).Error
		if err != nil {
			s.log.Error("error in upload media", zap.Error(err))
			return err
		}
	} else {
		var media models.Advertisement
		err = s.db.Where("name = ?", constants.MediaTableName).First(&media).Error

		if err != nil {
			s.log.Error("error in upload media", zap.Error(err))
			return err
		}

		media.Body = payload.Body
		if payload.Language != "" {
			media.Language = payload.Language
		}
		err = s.db.Where("name = ?", constants.MediaTableName).Updates(&media).Error
		if err != nil {
			s.log.Error("error in upload media", zap.Error(err))
			return err
		}
	}

	return nil
}

func (s *Storage) GetMedia() (models.Advertisement, error) {
	var media models.Advertisement
	err := s.db.Where("name = ?", constants.MediaTableName).First(&media).Error

	if err != nil {
		s.log.Error("error in upload media", zap.Error(err))
		return media, err
	}
	return media, nil
}

func (s *Storage) CreateRequest(payload models.Request) error {
	err := s.db.Create(&payload).Error
	if err != nil {
		s.log.Error("error in create request", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) UpdateRequest(id uint, payload *models.Request) error {
	err := s.db.Model(&models.Request{}).Where("id = ?", id).Updates(payload).Error
	if err != nil {
		s.log.Error("update payload niyozbek", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) GetLastRequestUser(userId uint) (*models.Request, error) {
	var request models.Request
	err := s.db.Where("user_id = ?", userId).Order("created_at DESC").Find(&request).Error
	if err != nil {
		s.log.Error("update payload niyozbek", zap.Error(err))
		return nil, err
	}
	return &request, nil
}

func (s *Storage) GetLastFilledRequestUser(userId uint) (*models.Request, error) {
	var request models.Request
	err := s.db.Where("user_id = ? and is_filled = true", userId).Order("created_at DESC").Find(&request).Error
	if err != nil {
		s.log.Error("update payload niyozbek", zap.Error(err))
		return nil, err
	}
	return &request, nil
}
