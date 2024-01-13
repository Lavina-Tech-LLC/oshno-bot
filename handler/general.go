package handler

import (
	"fmt"
	"log"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/pkg/socket"
	"oshno/pkg/utils"
	"oshno/pkg/validation"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

var (
	workers map[uint]interface{}
)

func (h BotHandler) Start(c tele.Context) error {
	h.logger.Info("bot started")
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return nil
	}
	err = h.storage.UpdatePhase(user.ID, 0)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	return c.Send("Ассалому алайкум, барои сабти ном лутфан рақами телефони худро мубодила кунед\nЗдравствуйте, для регистрации пожалуйста, поделитесь своим номером телефона", models.PhoneMarkup)
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
		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress], models.DisplayMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress], models.DisplayMarkupRu)
		}
	}

	if phone == user.PhoneNumber {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseLanguageStart], models.LanguageMarkup)
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
	return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseLanguageStart], models.LanguageMarkup)
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
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if user.ID == 0 {
			return nil
		}
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		message := strings.TrimSpace(c.Message().Text)
		if len(message) > 150 && user.UserPhase != 30 {
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

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
			}
		case 3:
			err := h.updateRequest(message, 3, user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPreferPlanTariff], models.PlanStatusMarkupToj)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPreferPlanTariff], models.PlanStatusMarkupRu)
			}
			/*
				case 4:
					validated := utils.ValidateDateBirth(message)
					if !validated {
						switch user.Language {
						case constants.Tajik:
							return c.Send(constants.ConstMessages[constants.Tajik][constants.ReEnterDateBirth])
						default:
							return c.Send(constants.ConstMessages[constants.Russian][constants.ReEnterDateBirth])
						}
					}

					err := h.updateRequest(message, 4, user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}

					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPhoneNumber], models.PhoneMarkupTg)
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPhoneNumber], models.PhoneMarkup)
					}
				case 5:
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

					err := h.updateRequest(message, 5, user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}

					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterEmail], models.DisplayMarkupTg)
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterEmail], models.DisplayMarkupRu)
					}
				case 6:
					validated := utils.ValidateEmail(message)
					if !validated {
						switch user.Language {
						case constants.Tajik:
							return c.Send(constants.ConstMessages[constants.Tajik][constants.ReEnterEmail])
						default:
							return c.Send(constants.ConstMessages[constants.Russian][constants.ReEnterEmail])
						}
					}

					err := h.updateRequest(message, 6, user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}

					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
					}
				case 7:
					err := h.updateRequest(message, 7, user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}

					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPassportData])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPassportData])
					}
				case 8:
					err := h.updateRequest(message, 8, user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}

					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPassportFront])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPassportFront])
					}
			*/
		case 6:

			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			request.PersonalAccount = message
			err = h.storage.UpdateRequest(request.ID, request)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			err = h.storage.UpdatePhase(user.ID, 7)
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
			// var aiMessage string
			conn, err := socket.Connection(user.AIChatId, user.Language)
			if err != nil {
				h.logger.Error(err.Error())
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			messageSend := []byte(message)
			err = conn.WriteMessage(websocket.TextMessage, messageSend)
			if err != nil {
				h.logger.Error(err.Error())
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			num, workers := utils.AddWorker(h.workers[user.ID])
			h.workers[user.ID] = workers
			go h.backgroundFunc(c, user, num)

			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return nil
			}
			c.Send(string(message))
			// Закрыть соединение после получения сообщения

			err = conn.Close()
			if err != nil {
				fmt.Println("close error:", err)
				return nil
			}

			return nil
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

	/*
		request, err := h.storage.GetLastRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		PHOTO FRONT phase
		if user.UserPhase == 9 {
			body, err := json.Marshal(c.Message().Photo)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			request.PhotoFront = body
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
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPassportBack])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPassportBack])
			}
		}
		// PHOTO BACK phase
		if user.UserPhase == 10 {
			body, err := json.Marshal(c.Message().Photo)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			request.PhotoBack = body
			err = h.storage.UpdateRequest(request.ID, request)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			err = h.storage.UpdatePhase(user.ID, 11)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPhotoWith])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPhotoWith])
			}
		}
		// PHOTO WITH phase
		if user.UserPhase == 11 {
			body, err := json.Marshal(c.Message().Photo)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			request.PhotoWith = body
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
				if request.Provider == constants.RequestTojNet {
					return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterTariffSelectionTojNet], models.PlanListMarkupToj)
				} else {
					return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterTariffSelectionOshnoService], models.PlanListMarkupOsh)
				}
			default:
				if request.Provider == constants.RequestTojNet {
					return c.Send(constants.ConstMessages[constants.Russian][constants.EnterTariffSelectionTojNet], models.PlanListMarkupToj)
				} else {
					return c.Send(constants.ConstMessages[constants.Russian][constants.EnterTariffSelectionOshnoService], models.PlanListMarkupOsh)
				}
			}
		}
	*/

	// only for admins
	if user.UserPhase == 20 {
		err := h.uploadMedia(*c.Message())
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		return c.Send(fmt.Sprintf("Подтверждаете ли вы отправку текста объявления пользователям?"), models.ConfirmAdMarkup)
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

func (h BotHandler) backgroundFunc(c tele.Context, user models.User, num uint) {
	startTime := time.Now()
	startTime.Add(time.Minute * 1)

	for {
		if utils.IsLastWorker(h.workers[user.ID], num) {
			afterDelWorker := utils.DeleteWorker(h.workers[user.ID], num)
			h.workers[user.ID] = afterDelWorker
			return
		}
		elapsedTime := time.Since(startTime)
		if elapsedTime >= time.Minute {
			switch user.Language {
			case constants.Tajik:
				c.Send("Оё маълумот муфид буд?", models.AIConfirmMarkupTg)
			default:
				c.Send("Была ли информация полезной?", models.AIConfirmMarkupRu)
			}
			afterDelWorker := utils.DeleteWorker(h.workers[user.ID], num)
			h.workers[user.ID] = afterDelWorker

			return
		}

	}
}

func confirmMessageTg(payload models.Request) string {
	service := payload.Service

	switch payload.Service {
	case constants.ServiceConnectProviderRu:
		service = "1"
	case constants.ServiceChangeTariffRu:
		service = "2"
	case constants.ServiceConnectAdditionalTariffRu:
		service = "3"
	}

	if payload.Service != constants.ServiceConnectProviderRu {
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
	if payload.Service != constants.ServiceConnectProviderRu {
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
	return fmt.Sprintf("Провайдер: %s\nУслуга: %s\nИмя м фамилия: %s\nНомер телефона: %s\nАдрес: %s\nТариф: %s\n",
		payload.Provider,
		payload.Service,
		payload.FullName,
		payload.PhoneNumber,
		payload.Address,
		payload.Plan,
	)
}

func newRequestMessageToGroup(rq models.Request) string {
	if rq.Service != constants.ServiceConnectProviderRu {
		return fmt.Sprintf("Новая заявка на провайдер %s услуга %s\nАйди заявки: %d\nИмя и фамилия: %s\nТелефон номер: %s\nАдрес: %s\nТариф: %s\nЛицевой счёт: %s\n",
			rq.Provider,
			rq.Service,
			rq.ID,
			rq.FullName,
			rq.PhoneNumber,
			rq.Address,
			rq.Plan,
			rq.PersonalAccount,
		)
	}
	return fmt.Sprintf("Новая заявка на провайдер %s услуга %s\nАйди заявки: %d\nИмя и фамилия: %s\nТелефон номер: %s\nАдрес: %s\nТариф: %s\n",
		rq.Provider,
		rq.Service,
		rq.ID,
		rq.FullName,
		rq.PhoneNumber,
		rq.Address,
		rq.Plan,
	)

}
