package handler

import (
	"oshno/models"
	"oshno/pkg/constants"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func (h BotHandler) PlanBtn(planNumber int) func(c tele.Context) error {
	return func(c tele.Context) error {
		c.Delete()
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		h.logger.Info("select plan started", zap.Any("user", user), zap.Int("plan number", planNumber))

		if user.UserPhase != 5 && user.UserPhase != 7 && user.UserPhase != 10 {
			return nil
		}

		request, err := h.storage.GetLastRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		request.PlanNumber = uint(planNumber)
		if request.Provider == constants.RequestOshoProvider {
			switch user.Language {
			case constants.Tajik:
				request.Plan = constants.OshnoPlansToj[planNumber]
			default:
				request.Plan = constants.OshnoPlans[planNumber]

			}
		} else {
			switch user.Language {
			case constants.Tajik:
				request.Plan = constants.TojNetPlansToj[planNumber]
			default:
				request.Plan = constants.TojNetPlansToj[planNumber]
			}
		}

		switch user.UserPhase {
		case 5:
			err = h.storage.UpdateRequest(request.ID, request)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			err = h.storage.UpdatePhase(user.ID, 5)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:

				c.Send(confirmMessageTg(*request))
				if err != nil {
					return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
				}
				return c.Send("шумо ин маълумотро тасдиқ мекунед?", models.ConfirmRequestMarkupTg)
			default:
				c.Send(confirmMessageRu(*request))
				return c.Send("Вы потдверждаете эти данные?", models.ConfirmRequestMarkupRu)
			}
		case 7:

			err = h.storage.UpdateRequest(request.ID, request)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			err = h.storage.UpdatePhase(user.ID, 8)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}

			switch user.Language {
			case constants.Tajik:
				err = c.Send("Шумо тасдиқ мекунед, ки шумо тарифро ба иваз кардан мехоҳед" + request.Plan)
				if err != nil {
					return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
				}
				return c.Send("шумо ин маълумотро тасдиқ мекунед?", models.ConfirmRequestMarkupTg)
			default:
				err = c.Send("Вы подтверждаете, что хотите сменить тариф на " + request.Plan)
				if err != nil {
					return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
				}
				return c.Send("Вы потдверждаете эти данные?", models.ConfirmRequestMarkupRu)
			}
		case 10:
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
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
			}

		}

		return nil
	}
}

func (h BotHandler) PlanLimit() func(c tele.Context) error {
	return func(c tele.Context) error {
		c.Delete()
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		h.logger.Info("select plan limited started", zap.Any("user", user))

		request, err := h.storage.GetLastRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		request.IsUnlimit = false

		err = h.storage.UpdateRequest(request.ID, request)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		switch user.Language {
		case constants.Tajik:
			if request.Provider == constants.RequestTojNet {
				return c.Send(constants.ConstMessages[constants.Tajik][constants.PlanTariffLimitTojNet], models.PlanListMarkupToj)
			} else {
				return c.Send(constants.ConstMessages[constants.Tajik][constants.PlanTariffLimitOshno], models.PlanListMarkupOsh)
			}
		default:
			if request.Provider == constants.RequestTojNet {
				return c.Send(constants.ConstMessages[constants.Russian][constants.PlanTariffLimitTojNet], models.PlanListMarkupToj)
			} else {
				return c.Send(constants.ConstMessages[constants.Russian][constants.PlanTariffLimitOshno], models.PlanListMarkupOsh)
			}
		}
	}
}

func (h BotHandler) PlanUnlimit() func(c tele.Context) error {
	return func(c tele.Context) error {
		c.Delete()
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		h.logger.Info("select plan unlimited started", zap.Any("user", user))

		request, err := h.storage.GetLastRequestUser(user.ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		request.IsUnlimit = true

		err = h.storage.UpdateRequest(request.ID, request)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch user.Language {
		case constants.Tajik:
			if request.Provider == constants.RequestTojNet {
				return c.Send(constants.ConstMessages[constants.Tajik][constants.PlanTariffUnlimitTojNet], models.PlanUnlimitListMarkupToj)
			} else {
				return c.Send(constants.ConstMessages[constants.Tajik][constants.PlanTariffUnlimitOshno], models.PlanUnlimitListMarkupOsh)
			}
		default:
			if request.Provider == constants.RequestTojNet {
				return c.Send(constants.ConstMessages[constants.Russian][constants.PlanTariffUnlimitTojNet], models.PlanUnlimitListMarkupToj)
			} else {
				return c.Send(constants.ConstMessages[constants.Russian][constants.PlanTariffUnlimitOshno], models.PlanUnlimitListMarkupOsh)
			}
		}
	}
}
