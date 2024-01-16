package handler

import (
	"fmt"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/pkg/gateways"
	"oshno/pkg/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func (h BotHandler) RequestProvider(c tele.Context) error {
	h.logger.Info("button request provider")
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	err = h.storage.UpdatePhase(user.ID, 0)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	switch user.Language {
	case constants.Tajik:
		return c.Send(constants.ConstMessages[constants.Tajik][constants.RequestProvider], models.RequestProviderMarkupTg)
	default:
		return c.Send(constants.ConstMessages[constants.Russian][constants.RequestProvider], models.RequestProviderMarkupRu)
	}
}

func (h BotHandler) TechnicalSupport(languageCode string) func(c tele.Context) error {
	return func(c tele.Context) error {
		h.logger.Info("button technical support")
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		err = h.storage.UpdatePhase(user.ID, 30)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		AIChatId := uuid.New()
		err = h.storage.UpdateAIChatId(user.ID, AIChatId.String())
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send("Салом, ман ба шумо чӣ гуна кӯмак карда метавонам, лутфан мушкилоти худро шарҳ диҳед?")
		default:
			return c.Send("Здравствуйте, чем могу помочь, опишите пожалуйста вашу проблему?")
		}
	}
}

func (h BotHandler) AIConfirm(q string) func(c tele.Context) error {
	return func(c tele.Context) error {
		h.logger.Info("AI confirm")

		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		err = h.storage.UpdatePhase(user.ID, 0)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		if q == "no" {
			messages, err := gateways.GetHistoryChat(user.AIChatId)
			message := messages.GenerateMessage(user)
			_, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)
			if err != nil {
				fmt.Println(err.Error())
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
		}
		switch user.Language {
		case constants.Tajik:
			if q == "yes" {
				return c.Send("Ташаккур барои интихоб кардани мо", models.MenuMarkupTg)
			} else {
				return c.Send("Ҳама маълумот гирифта шудааст. Ташаккур барои ҷавобҳо. Пас аз чанд вақт кормандони дастгирии мо ба шумо занг мезананд ва барномаи шуморо ҳал мекунанд. Мо аз кор кардан бо Шумо хурсандем!", models.MenuMarkupTg)
			}
		default:
			if q == "yes" {
				return c.Send("Спасибо если ещё будут вопросы обращайтесь", models.MenuMarkupRu)
			} else {
				return c.Send("Вся информация получена. Спасибо за обращения.\nВ скором времени наши сотрудники свяжутся с вами и помогут с решением вашего обращения\nМы рады работать с Вами!", models.MenuMarkupRu)
			}
		}
	}
}

func (h BotHandler) OshnoService(c tele.Context) error {
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	h.logger.Info("onText started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language), zap.String("provider", "oshno service"))

	switch user.Language {
	case constants.Tajik:
		return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.OshnoProviderMarkupTg)
	default:
		return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.OshnoProviderMarkupRu)
	}
}

func (h BotHandler) TojNet(c tele.Context) error {
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}

	h.logger.Info("onText started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language), zap.String("provider", "tojNet"))
	switch user.Language {
	case constants.Tajik:
		return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.TojNetProviderMarkupTg)
	default:
		return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.TojNetProviderMarkupRu)
	}
}

func (h BotHandler) ConnectProvider(provider, service string) func(c tele.Context) error {
	return func(c tele.Context) error {
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		h.logger.Info("onText started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language), zap.String("provider", provider))

		err = h.createRequest(int(user.ID), provider, service)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterFirstName], models.DisplayMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.EnterFirstName], models.DisplayMarkupRu)
		}
	}
}

func (h BotHandler) ChangeTariff(provider, service string) func(c tele.Context) error {
	return func(c tele.Context) error {
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		h.logger.Info("change tariff service started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language), zap.String("provider", provider))

		err = h.createRequest(int(user.ID), provider, service)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterFirstName], models.DisplayMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.EnterFirstName], models.DisplayMarkupRu)
		}
	}
}

func (h BotHandler) ConnectAdditional(provider, service string) func(c tele.Context) error {
	return func(c tele.Context) error {
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		h.logger.Info("connect additional started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language), zap.String("provider", provider))

		err = h.createRequest(int(user.ID), provider, service)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterFirstName], models.DisplayMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.EnterFirstName], models.DisplayMarkupRu)
		}
	}
}

func (h BotHandler) ConfirmRequest(c tele.Context) error {
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}
	h.logger.Info("confirm request started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language))

	// sending request to group chat
	switch user.UserPhase {
	case 5:
		lastRequest, err := h.storage.GetLastRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		message := newRequestMessageToGroup(*lastRequest)
		_, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)
		if err != nil {
			h.logger.Error("error in sending message", zap.Error(err))
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		err = h.storage.UpdatePhase(user.ID, 0)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		lastRequest.IsFilled = true
		err = h.storage.UpdateRequest(lastRequest.ID, lastRequest)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		// sending message to bot
		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.ConfirmMessage], models.MenuMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.ConfirmMessage], models.MenuMarkupRu)
		}
	case 8:
		fRequest, err := h.storage.GetLastFilledRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		request, err := h.storage.GetLastRequestUser(user.ID)
		request.Address = fRequest.Address
		request.PhoneNumber = fRequest.PhoneNumber
		request.FullName = fRequest.FullName

		message := newRequestMessageToGroup(*request)
		_, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)
		if err != nil {
			h.logger.Error("error ", zap.Error(err))
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		err = h.storage.UpdateRequest(request.ID, request)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		err = h.storage.UpdatePhase(user.ID, 0)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.ConfirmMessage], models.MenuMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.ConfirmMessage], models.MenuMarkupRu)
		}
	case 12:
		fRequest, err := h.storage.GetLastFilledRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		request, err := h.storage.GetLastRequestUser(user.ID)
		request.Address = fRequest.Address
		request.PhoneNumber = fRequest.PhoneNumber
		request.FullName = fRequest.FullName

		message := newRequestMessageToGroup(*request)
		_, err = h.bot.Send(&tele.Chat{ID: constants.TelegramGroupId}, message)
		if err != nil {
			h.logger.Error("error ", zap.Error(err))
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		err = h.storage.UpdateRequest(request.ID, request)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		err = h.storage.UpdatePhase(user.ID, 0)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			return c.Send(constants.ConstMessages[constants.Tajik][constants.ConfirmMessage], models.MenuMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.ConfirmMessage], models.MenuMarkupRu)
		}
	}
	return nil
}

func (h BotHandler) IgnoreRequest(c tele.Context) error {
	// sending message to bot
	user, err := h.storage.GetUserByTgId(c.Sender().ID)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}
	h.logger.Info("onText started", zap.Uint("userId", user.ID), zap.String("userLanguage", user.Language))
	err = h.storage.UpdatePhase(user.ID, 0)
	if err != nil {
		return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
	}
	switch user.Language {
	case constants.Tajik:
		return c.Send(constants.ConstMessages[constants.Tajik][constants.IgnoreMessage], models.MenuMarkupTg)
	default:
		return c.Send(constants.ConstMessages[constants.Russian][constants.IgnoreMessage], models.MenuMarkupRu)
	}
}

func (h BotHandler) createRequest(userId int, provider, service string) error {
	var userPhase int

	err := h.storage.CreateRequest(models.Request{
		Service:  service,
		Provider: provider,
		UserId:   userId,
	})
	if err != nil {
		return err
	}
	userPhase = utils.UserPhaseToService(service)

	err = h.storage.UpdatePhase(uint(userId), userPhase)
	if err != nil {
		return err
	}

	return nil
}
