package handler

import (
	"fmt"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/pkg/gateways"
	"oshno/pkg/utils"
	"oshno/pkg/validation"
	"strings"
	"time"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

// var (
// 	workers map[uint]interface{}
// )

func (h BotHandler) Start(c tele.Context) error {
	h.logger.Info("bot started")
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return c.Send("Здравствуйте, для регистрации пожалуйста, поделитесь своим номером телефона", models.PhoneMarkup)
	}
	err = h.storage.UpdatePhase(user.ID, 0)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	return c.Send("Здравствуйте, для регистрации пожалуйста, поделитесь своим номером телефона", models.PhoneMarkup)
}

func (h BotHandler) Contact(c tele.Context) error {
	h.logger.Info("share contact started ", zap.String("Phone Number", c.Message().Contact.PhoneNumber))
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		if err.Error() != constants.UserNotRegist {
			return c.Send(constants.ConstMessages[constants.ErrorReport], models.StartMarkup)
		}
	}
	phone := c.Message().Contact.PhoneNumber
	// check if shared contact already has "+", because sometimes telegram shares phone number without "+"
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}
	if !validation.IsPhoneValid(phone) {
		return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPhone], models.PhoneMarkup)
	}

	if user.UserPhase == 2 {
		request, err := h.storage.GetLastRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		request.PhoneNumber = phone
		err = h.storage.UpdateRequest(request.ID, request)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		err = h.storage.UpdatePhase(user.ID, 3)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		if request.Service == constants.ServiceConnectProviderRu {
			h.storage.UpdatePhase(user.ID, 4)
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress], models.DisplayMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress], models.DisplayMarkupRu)
			}
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPersonalAccount], models.DisplayMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPersonalAccount], models.DisplayMarkupRu)
		}
	}

	if phone == user.PhoneNumber {
		return c.Send(constants.ConstMessages[constants.Russian][constants.LanguageChanged], models.MenuMarkupRu)
	}
	err = h.storage.CreateUser(models.User{
		PhoneNumber:    phone,
		TelegramUserId: c.Sender().ID,
		TelegramChatId: c.Chat().ID,
		Nickname:       c.Chat().Username,
		Language:       constants.Russian,
		FullName:       c.Message().Contact.FirstName + " " + c.Message().Contact.LastName,
	})
	if err != nil {
		c.Send(constants.ConstMessages[constants.ErrorReport], models.StartMarkup)
		return err
	}
	h.logger.Info("share contact finished", zap.Int64("userId", int64(user.ID)))
	return c.Send(constants.ConstMessages[constants.Russian][constants.LanguageChanged], models.MenuMarkupRu)
}

func (h BotHandler) Location(c tele.Context) error {
	return nil
}

func (h BotHandler) LanguageButton(languageCode string) func(c tele.Context) error {
	return func(c tele.Context) error {
		h.logger.Info("Language change started", zap.String("language", languageCode))
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		if user.UserPhase == 20 {
			err = h.storage.UploadMedia(models.Advertisement{
				Name:     constants.MediaTableName,
				Language: languageCode,
			})
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			switch user.Language {
			case constants.Tajik:
				return c.Send("Send post!", models.MenuMarkupTg)
			default:
				return c.Send("Отпавьте пост!", models.MenuMarkupRu)
			}
		}

		err = h.storage.UpdateLanguage(user.ID, languageCode)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		h.logger.Info("Language change finished", zap.String("language", languageCode), zap.Uint("userID", user.ID))
		switch languageCode {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.LanguageChanged], models.MenuMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.LanguageChanged], models.MenuMarkupRu)
		}
	}
}

func (h BotHandler) Text(languageCode string) func(c tele.Context) error {
	return func(c tele.Context) error {
		if c.Message().TopicMessage {
			err := h.OperatorMessages(c)
			return err
		}

		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if user.ID == 0 {
			return nil
		}
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		message := strings.TrimSpace(c.Message().Text)
		if len(message) > 150 && user.UserPhase != 30 && user.UserPhase != int(constants.PhaseOperatorSupport) {
			switch user.Language {
			case constants.Tajik:
				return c.Send("Паёми шумо хеле дароз аст! Бори дигар кӯшиш кунед, танҳо онро каме содда кунед")
			default:
				return c.Send("Ваше сообщение слишком длиная! Попробуйте снова только немного упростив")
			}
		}

		h.logger.Info("onText started", zap.Any("user", user), zap.String("user message", message), zap.String("workers", h.workers[user.ID]))

		switch user.UserPhase {
		case 1:
			isValid := utils.ValidateString(message)

			if !isValid {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ReEnterFirstName])
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ReEnterFirstName])
				}
			}
			err := h.updateRequest(message, 1, user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPhoneNumber], models.PhoneFillMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPhoneNumber], models.PhoneFillMarkup)
			}
		case 2:
			if !strings.HasPrefix(message, "+") {
				message = "+" + message
			}

			if !validation.IsPhoneValid(message) {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ReEnterPhoneNumber])
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ReEnterPhoneNumber])
				}
			}

			err := h.updateRequest(message, 2, user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			if request.Service == constants.ServiceConnectProviderRu {
				h.storage.UpdatePhase(user.ID, 4)
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress], models.DisplayMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress], models.DisplayMarkupRu)
				}
			}
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPersonalAccount], models.DisplayMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPersonalAccount], models.DisplayMarkupRu)
			}
		case 3:

			err := h.updateRequest(message, 3, user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			if request.Service == constants.ServiceChangeTariffRu {
				err = h.storage.UpdatePhase(user.ID, 5)
				if err != nil {
					return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
				}

				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPreferPlanTariff], models.PlanStatusMarkupToj)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPreferPlanTariff], models.PlanStatusMarkupRu)
				}
			}

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
			}

		case 4:
			err := h.updateRequest(message, 4, user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPreferPlanTariff], models.PlanStatusMarkupToj)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPreferPlanTariff], models.PlanStatusMarkupRu)
			}
		case 9:
			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			request.PersonalAccount = message
			err = h.storage.UpdateRequest(request.ID, request)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			err = h.storage.UpdatePhase(user.ID, 10)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPreferPlanTariff], models.PlanStatusMarkupToj)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPreferPlanTariff], models.PlanStatusMarkupRu)
			}
		case 11:
			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			request.Address = message
			err = h.storage.UpdateRequest(request.ID, request)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			err = h.storage.UpdatePhase(user.ID, 12)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				err = c.Send("Шумо тасдиқ мекунед, ки шумо тарифро илова кардан мехоҳед " + request.Plan + "\nСуроға: " + message)
				if err != nil {
					return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
				}
				return c.Send("шумо ин маълумотро тасдиқ мекунед?", models.ConfirmRequestMarkupTg)
			default:
				err = c.Send("Вы подтверждаете, что хотите подключить дополнительный тариф на " + request.Plan + "\nАдрес: " + message)
				if err != nil {
					return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
				}
				return c.Send("Вы потдверждаете эти данные?", models.ConfirmRequestMarkupRu)
			}

		case 30:
			data, err := gateways.SendMessage(user.AIChatId, message)
			if err != nil {
				h.logger.Error(err.Error())
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			now := time.Now()
			err = h.storage.UpdateUser(user.ID, models.User{AIConfirmSended: false, AILastMessageTime: &now})
			if err != nil {

				h.logger.Error("error in update user: ", zap.Error(err))
			}

			c.Send(data.Data)
		case int(constants.PhaseOperatorSupport):
			// send to topic
			err := gateways.SendMessageToTopic(int64(user.ActiveTopic), message)
			if err != nil {
				h.logger.Error("error in send message to operator group: ", zap.Error(err))
			}

			topic, _ := h.storage.GetTopicByThreadId(user.ActiveTopic)
			h.storage.CreateTopicMessage(models.OperatorChat{Message: message, TopicId: int32(topic.ID), Operator: ""})
			return err
		}

		return nil
	}
}

func (h BotHandler) Photo(c tele.Context) error {
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	h.logger.Info("share photo started", zap.Uint("userId", user.ID))

	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	// only for admins
	if user.UserPhase == 20 {
		err := h.uploadMedia(*c.Message())
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		return c.Send("Подтверждаете ли вы отправку текста объявления пользователям?", models.ConfirmAdMarkup)
	}

	return nil
}

func (h BotHandler) updateRequest(message string, phase, userId uint) error {
	request, err := h.storage.GetLastRequestUser(userId)
	if err != nil {
		return err
	}
	switch phase {
	case 1:
		request.FullName = message
	case 2:
		request.PhoneNumber = message
	case 3:
		request.PersonalAccount = message
	case 4:
		request.Address = message
	}

	err = h.storage.UpdateRequest(request.ID, request)
	if err != nil {
		return err
	}

	err = h.storage.UpdatePhase(userId, int(phase)+1)
	if err != nil {
		return err
	}

	return nil
}

// func (h BotHandler) backgroundFunc(c tele.Context, user models.User, num uint) {
// 	startTime := time.Now()
// 	startTime = startTime.Add(time.Minute * 10)

// 	for {
// 		if utils.IsLastWorker(h.workers[user.ID], num) {
// 			afterDelWorker := utils.DeleteWorker(h.workers[user.ID], num)
// 			h.workers[user.ID] = afterDelWorker
// 			return
// 		}
// 		elapsedTime := time.Since(startTime)
// 		if elapsedTime >= time.Minute {
// 			switch user.Language {
// 			case constants.Tajik:
// 				c.Send("Оё маълумот муфид буд?", models.AIConfirmMarkupTg)
// 			default:
// 				c.Send("Была ли информация полезной?", models.AIConfirmMarkupRu)
// 			}
// 			afterDelWorker := utils.DeleteWorker(h.workers[user.ID], num)
// 			h.workers[user.ID] = afterDelWorker

// 			return
// 		}

// 	}
// }

func confirmMessageTg(payload models.Request) string {
	service := payload.Service

	switch payload.Service {
	case constants.ServiceConnectProviderRu:
		service = constants.ServiceConnectProviderTg

		return fmt.Sprintf("Провайдер: %s\nХизмат: %s\nНом ва насаб: %s\nРақами мобилӣ: %s\nСуроға: %s\nПлан: %s\n",
			payload.Provider,
			service,
			payload.FullName,
			payload.PhoneNumber,
			payload.Address,
			payload.Plan,
		)

	case constants.ServiceChangeTariffRu:
		service = constants.ServiceChangeTariffTg
		return fmt.Sprintf("Провайдер: %s\nХизмат: %s\nНом ва насаб: %s\nРақами мобилӣ: %s\nПлан: %s\nҲисоби шахсӣ: %s\n",
			payload.Provider,
			service,
			payload.FullName,
			payload.PhoneNumber,
			payload.Plan,
			payload.PersonalAccount,
		)

	case constants.ServiceConnectAdditionalTariffRu:
		service = constants.ServiceConnectAdditionalTariffTg
		return fmt.Sprintf("Провайдер: %s\nХизмат: %s\nНом ва насаб: %s\nРақами мобилӣ: %s\nСуроға: %s\nПлан: %s\nҲисоби шахсӣ: %s\n",
			payload.Provider,
			service,
			payload.FullName,
			payload.PhoneNumber,
			payload.Address,
			payload.Plan,
			payload.PersonalAccount,
		)

	}

	return fmt.Sprintf("Провайдер: %s\nХизмат: %s\nНом ва насаб: %s\nРақами мобилӣ: %s\nСуроға: %s\nПлан: %s\n",
		payload.Provider,
		service,
		payload.FullName,
		payload.PhoneNumber,
		payload.Address,
		payload.Plan,
	)
}

func confirmMessageRu(payload models.Request) string {
	switch payload.Service {
	case constants.ServiceConnectProviderRu:
		return fmt.Sprintf("Провайдер: %s\nУслуга: %s\nИмя и фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\n",
			payload.Provider,
			payload.Service,
			payload.FullName,
			payload.PhoneNumber,
			payload.Address,
			payload.Plan,
		)
	case constants.ServiceChangeTariffRu:
		return fmt.Sprintf("Провайдер: %s\nУслуга: %s\nИмя и фамилия: %s\nНомер телефона: %s\nТариф: %s\nЛицевой счёт: %s\n",
			payload.Provider,
			payload.Service,
			payload.FullName,
			payload.PhoneNumber,
			payload.Plan,
			payload.PersonalAccount,
		)

	case constants.ServiceConnectAdditionalTariffRu:
		return fmt.Sprintf("Провайдер: %s\nУслуга: %s\nИмя и фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\nЛицевой счёт: %s\n",
			payload.Provider,
			payload.Service,
			payload.FullName,
			payload.PhoneNumber,
			payload.Address,
			payload.Plan,
			payload.PersonalAccount,
		)
	}

	return fmt.Sprintf("Провайдер: %s\nУслуга: %s\nИмя и фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\n",
		payload.Provider,
		payload.Service,
		payload.FullName,
		payload.PhoneNumber,
		payload.Address,
		payload.Plan,
	)
}

func newRequestMessageToGroup(rq models.Request) string {
	var plan string
	if rq.Provider == constants.RequestOshoProvider {
		plan = constants.OshnoPlans[int(rq.PlanNumber)]
	} else {
		plan = constants.TojNetPlans[int(rq.PlanNumber)]
	}
	switch rq.Service {
	case constants.ServiceConnectProviderRu:
		return fmt.Sprintf("Новая заявка\n\nПровайдер: %s\nУслуга: %s\nНомер заявки: %d\nИмя и фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\n",
			rq.Provider,
			rq.Service,
			rq.ID,
			rq.FullName,
			rq.PhoneNumber,
			rq.Address,
			plan,
		)
	case constants.ServiceChangeTariffRu:
		return fmt.Sprintf("Новая заявка\n\nПровайдер: %s\nУслуга: %s\nНомер заявки: %d\nИмя и фамилия: %s\nНомер телефона: %s\nТариф: %s\nЛицевой счёт: %s\n",
			rq.Provider,
			rq.Service,
			rq.ID,
			rq.FullName,
			rq.PhoneNumber,
			plan,
			rq.PersonalAccount,
		)

	case constants.ServiceConnectAdditionalTariffRu:
		return fmt.Sprintf("Новая заявка\n\nПровайдер: %s\nУслуга: %s\nНомер заявки: %d\nИмя и фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\nЛицевой счёт: %s\n",
			rq.Provider,
			rq.Service,
			rq.ID,
			rq.FullName,
			rq.PhoneNumber,
			rq.Address,
			plan,
			rq.PersonalAccount,
		)
	}
	return fmt.Sprintf("Новая заявка\n\nПровайдер: %s\nУслуга: %s\nНомер заявки: %d\nИмя и фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\n",
		rq.Provider,
		rq.Service,
		rq.ID,
		rq.FullName,
		rq.PhoneNumber,
		rq.Address,
		plan,
	)

}

func (h BotHandler) OperatorMessages(c tele.Context) error {
	message := c.Message().Text
	topicId := c.Message().ThreadID

	h.logger.Info("operator answered message: ", zap.String("message", message))

	if message == "/stop" {
		h.Stop(c)
		gateways.SendMessageToTopic(int64(topicId), "Вы закрыли чат с килентом")
		return nil
	}

	user, err := h.storage.GetUserInActiveTopic(topicId)
	if err != nil {
		return nil
	}

	if user.ID == 0 {

		gateways.SendMessageToTopic(int64(topicId), "Клиент уже закрыл чат и не получил ваш ответ!")
		return nil
	}

	topic, _ := h.storage.GetTopicByThreadId(int32(topicId))
	h.storage.CreateTopicMessage(models.OperatorChat{Message: message, TopicId: int32(topic.ID), Operator: c.Message().Sender.Username})

	now := time.Now()
	err = h.storage.UpdateUser(user.ID, models.User{OpertorLastMessageTime: &now})
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	c.Bot().Send(&tele.User{ID: user.TelegramUserId}, message)
	return nil
}
