package handler

import (
	"oshno/models"
	"oshno/pkg/constants"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func (h BotHandler) BackButton(action string) func(c tele.Context) error {
	return func(c tele.Context) error {
		var backPhase int
		// var respMessage string
		user, err := h.storage.GetUserByTgId(c.Sender().ID)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}
		h.logger.Info("display started", zap.Any("user", user))

		if action == "request_provider" {
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.MenuButtonText], models.MenuMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.MenuButtonText], models.MenuMarkupRu)
			}
		} else if action == "oshno_provider" || action == "toj_net" {
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.RequestProvider], models.RequestProviderMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.RequestProvider], models.RequestProviderMarkupRu)
			}
		}

		if user.UserPhase != 0 {
			backPhase = user.UserPhase - 1
		}

		err = h.storage.UpdatePhase(user.ID, backPhase)
		if err != nil {
			return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
		}

		switch backPhase {
		case 0:
			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			if request.Provider == constants.RequestOshoProvider {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.OshnoProviderMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.OshnoProviderMarkupRu)
				}
			} else {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.TojNetProviderMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.TojNetProviderMarkupRu)
				}
			}
		case 1:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterFirstName], models.DisplayMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterFirstName], models.DisplayMarkupRu)
			}
		case 2:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPhoneNumber], models.PhoneFillMarkupTg)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPhoneNumber], models.PhoneFillMarkup)
			}
		case 3:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
			}
		case 4:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPreferPlanTariff], models.PlanStatusMarkupToj)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPreferPlanTariff], models.PlanStatusMarkupRu)
			}
		case 5:
			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			if request.Provider == constants.RequestOshoProvider {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.OshnoProviderMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.OshnoProviderMarkupRu)
				}
			} else {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.TojNetProviderMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.TojNetProviderMarkupRu)
				}
			}
		case 6:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPersonalAccount])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPersonalAccount])
			}
		case 8:
			request, err := h.storage.GetLastRequestUser(user.ID)
			if err != nil {
				return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
			}
			if request.Provider == constants.RequestOshoProvider {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.OshnoProviderMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.OshnoProviderMarkupRu)
				}
			} else {
				switch user.Language {
				case constants.Tajik:
					return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.TojNetProviderMarkupTg)
				default:
					return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.TojNetProviderMarkupRu)
				}
			}
		case 9:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPersonalAccount])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPersonalAccount])
			}
		case 10:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPreferPlanTariff], models.PlanStatusMarkupToj)
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPreferPlanTariff], models.PlanStatusMarkupRu)
			}
		case 11:
			switch user.Language {
			case constants.Tajik:
				return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
			default:
				return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
			}
			/*
				case 7:
					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
					}
				case 8:
					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPassportData])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPassportData])
					}


				case 11:
					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPhotoWith])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPhotoWith])
					}
				case 12:
					request, err := h.storage.GetLastRequestUser(user.ID)
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
				case 14:
					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPersonalAccount])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPersonalAccount])
					}
				case 15:
					request, err := h.storage.GetLastRequestUser(user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}
					switch user.Language {
					case constants.Tajik:
						if request.Provider == constants.RequestOshoProvider {
							return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterTariffSelectionOshnoService], models.PlanListMarkupOsh)
						} else {
							return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterTariffSelectionTojNet], models.PlanListMarkupToj)
						}
					default:
						if request.Provider == constants.RequestOshoProvider {
							return c.Send(constants.ConstMessages[constants.Russian][constants.EnterTariffSelectionOshnoService], models.PlanListMarkupOsh)
						} else {
							return c.Send(constants.ConstMessages[constants.Russian][constants.EnterTariffSelectionTojNet], models.PlanListMarkupToj)
						}
					}
				case 17:
					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterPersonalAccount], models.DisplayMarkupTg)
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterPersonalAccount], models.DisplayMarkupRu)
					}
				case 18:
					request, err := h.storage.GetLastRequestUser(user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}
					switch user.Language {
					case constants.Tajik:
						if request.Provider == constants.RequestOshoProvider {
							return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterTariffSelectionOshnoService], models.PlanListMarkupOsh)
						} else {
							return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterTariffSelectionTojNet], models.PlanListMarkupToj)
						}
					default:
						if request.Provider == constants.RequestOshoProvider {
							return c.Send(constants.ConstMessages[constants.Russian][constants.EnterTariffSelectionOshnoService], models.PlanListMarkupOsh)
						} else {
							return c.Send(constants.ConstMessages[constants.Russian][constants.EnterTariffSelectionTojNet], models.PlanListMarkupToj)
						}
					}
				case 19:
					switch user.Language {
					case constants.Tajik:
						return c.Send(constants.ConstMessages[constants.Tajik][constants.EnterAddress])
					default:
						return c.Send(constants.ConstMessages[constants.Russian][constants.EnterAddress])
					}

				case 13:
					request, err := h.storage.GetLastRequestUser(user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}
					if request.Provider == constants.RequestOshoProvider {
						switch user.Language {
						case constants.Tajik:
							return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.OshnoProviderMarkupTg)
						default:
							return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.OshnoProviderMarkupRu)
						}
					} else {
						switch user.Language {
						case constants.Tajik:
							return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.TojNetProviderMarkupTg)
						default:
							return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.TojNetProviderMarkupRu)
						}
					}

				case 16:
					request, err := h.storage.GetLastRequestUser(user.ID)
					if err != nil {
						return c.Send(constants.ConstMessages[constants.Russian][constants.ErrorReport], models.StartMarkup)
					}
					if request.Provider == constants.RequestOshoProvider {
						switch user.Language {
						case constants.Tajik:
							return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.OshnoProviderMarkupTg)
						default:
							return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.OshnoProviderMarkupRu)
						}
					} else {
						switch user.Language {
						case constants.Tajik:
							return c.Send(constants.ConstMessages[constants.Tajik][constants.ChooseService], models.TojNetProviderMarkupTg)
						default:
							return c.Send(constants.ConstMessages[constants.Russian][constants.ChooseService], models.TojNetProviderMarkupRu)
						}
					}
			*/
		}
		return nil
	}
}

func (h BotHandler) Menu() func(c tele.Context) error {
	return func(c tele.Context) error {
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
			return c.Send(constants.ConstMessages[constants.Tajik][constants.MenuButtonText], models.MenuMarkupTg)
		default:
			return c.Send(constants.ConstMessages[constants.Russian][constants.MenuButtonText], models.MenuMarkupRu)
		}
	}
}
