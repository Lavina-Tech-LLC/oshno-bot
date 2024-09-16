package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		PhoneNumber            string
		TelegramUserId         int64
		TelegramChatId         int64
		Nickname               string
		Language               string
		FullName               string
		UserPhase              int
		Role                   string
		AIChatId               string
		AILastMessageTime      *time.Time
		AIConfirmSended        bool `gorm:"default:true"`
		Topics                 []Topic
		ActiveTopic            int32
		OpertorLastMessageTime *time.Time
		gorm.Model
	}

	Advertisement struct {
		Body     []byte
		Name     string
		Language string
	}

	Request struct {
		Service         string
		Provider        string
		UserId          int
		FullName        string
		PhoneNumber     string
		Address         string
		Plan            string
		PlanNumber      uint
		PersonalAccount string
		IsUnlimit       bool `gorm:"default:null"`
		IsFilled        bool `gorm:"default:false"`
		gorm.Model
	}

	ChangePlanRequest struct {
		PersonalAccount string
		NewPlan         string
		RequestId       uint
		Request         Request
		gorm.Model
	}

	AddPlanRequest struct {
		Plan            string
		Address         string
		PersonalAccount string
		RequestId       uint
		Request         Request
		gorm.Model
	}

	OperatorChat struct {
		Message  string
		TopicId  int32
		Topic    Topic
		Operator string
		gorm.Model
	}

	Topic struct {
		Name     string
		ThreadId int32
		UserId   uint
		User     User
		gorm.Model
	}
)
