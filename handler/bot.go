package handler

import (
	"fmt"
	"oshno/models"
	"oshno/pkg/constants"
	"oshno/storage"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	tgMiddleware "gopkg.in/telebot.v3/middleware"
)

type BotHandler struct {
	bot     *tele.Bot
	logger  *zap.Logger
	storage *storage.Storage
	workers map[uint]string
}

func NewBotHandler(bot *tele.Bot, logger *zap.Logger, storage *storage.Storage) BotHandler {
	workers := make(map[uint]string)
	return BotHandler{
		bot:     bot,
		logger:  logger,
		storage: storage,
		workers: workers,
	}
}

func Start(h BotHandler) {
	setButtons()
	me := h.bot.Me
	fmt.Println(me.Username)
	adminOnly := h.bot.Group()
	adminOnly.Use(tgMiddleware.Whitelist(getAdmins()...))
	h.bot.Handle("/start", h.Start)
	adminOnly.Handle("/admin", h.Admin)
	adminOnly.Handle(&models.BtnConfirmAd, h.SendAd)
	adminOnly.Handle(&models.BtnIgnoreAd, h.UnsendAd)
	adminOnly.Handle(tele.OnMedia, h.OnMedia)

	h.bot.Handle(&models.BtnStart, h.Start)
	h.bot.Handle(tele.OnContact, h.Contact)

	h.bot.Handle(&models.BtnRu, h.LanguageButton(constants.Russian))
	h.bot.Handle(&models.BtnTg, h.LanguageButton(constants.Tajik))

	h.bot.Handle(&models.BtnRequestProviderRu, h.RequestProvider)
	h.bot.Handle(&models.BtnTechnicalSupportRu, h.TechnicalSupport(constants.Russian))

	h.bot.Handle(&models.BtnRequestProviderTg, h.RequestProvider)
	h.bot.Handle(&models.BtnTechnicalSupportTg, h.TechnicalSupport(constants.Tajik))

	h.bot.Handle(&models.BtnBackRu, h.BackButton(""))
	h.bot.Handle(&models.BtnHomeRu, h.Menu())

	h.bot.Handle(&models.BtnBackTg, h.BackButton(""))
	h.bot.Handle(&models.BtnHomeTg, h.Menu())

	h.bot.Handle(&models.BtnOshnoServiceRu, h.OshnoService)
	h.bot.Handle(&models.BtnTojNetRu, h.TojNet)
	h.bot.Handle(&models.BtnBackRequestProviderRu, h.BackButton("request_provider"))
	h.bot.Handle(&models.BtnMenuRequestProviderRu, h.Menu())

	h.bot.Handle(&models.BtnOshnoServiceRu, h.OshnoService)
	h.bot.Handle(&models.BtnTojNetRu, h.TojNet)
	h.bot.Handle(&models.BtnBackRequestProviderTg, h.BackButton("request_provider"))
	h.bot.Handle(&models.BtnMenuRequestProviderTg, h.Menu())

	h.bot.Handle(&models.BtnOshSerConnectProviderRu, h.ConnectProvider(constants.RequestOshoProvider, constants.ServiceConnectProviderRu))
	h.bot.Handle(&models.BtnOshSerChangeTariffRu, h.ChangeTariff(constants.RequestOshoProvider, constants.ServiceChangeTariffRu))
	h.bot.Handle(&models.BtnOshSerConnectAddittionalRu, h.ConnectAdditional(constants.RequestOshoProvider, constants.ServiceConnectAdditionalTariffRu))
	h.bot.Handle(&models.BtnBackProviderRu, h.BackButton("oshno_provider"))
	h.bot.Handle(&models.BtnMenuProviderRu, h.Menu())

	h.bot.Handle(&models.BtnOshSerConnectProviderTg, h.ConnectProvider(constants.RequestOshoProvider, constants.ServiceConnectProviderRu))
	h.bot.Handle(&models.BtnOshSerChangeTariffTg, h.ChangeTariff(constants.RequestOshoProvider, constants.ServiceChangeTariffRu))
	h.bot.Handle(&models.BtnOshSerConnectAddittionalTg, h.ConnectAdditional(constants.RequestOshoProvider, constants.ServiceConnectAdditionalTariffRu))
	h.bot.Handle(&models.BtnBackProviderTg, h.BackButton("oshno_provider"))
	h.bot.Handle(&models.BtnMenuProviderTg, h.Menu())

	h.bot.Handle(&models.BtnTojNetConnectProviderRu, h.ConnectProvider(constants.RequestTojNet, constants.ServiceConnectProviderRu))
	h.bot.Handle(&models.BtnTojNetChangeTariffRu, h.ChangeTariff(constants.RequestTojNet, constants.ServiceChangeTariffRu))
	h.bot.Handle(&models.BtnTojNetConnectAddittionalRu, h.ConnectAdditional(constants.RequestTojNet, constants.ServiceConnectAdditionalTariffRu))
	h.bot.Handle(&models.BtnBackTojRu, h.BackButton("toj_net"))
	h.bot.Handle(&models.BtnMenuTojRu, h.Menu())

	h.bot.Handle(&models.BtnTojNetConnectProviderTg, h.ConnectProvider(constants.RequestTojNet, constants.ServiceConnectProviderRu))
	h.bot.Handle(&models.BtnTojNetChangeTariffTg, h.ChangeTariff(constants.RequestTojNet, constants.ServiceChangeTariffRu))
	h.bot.Handle(&models.BtnTojNetConnectAddittionalTg, h.ConnectAdditional(constants.RequestTojNet, constants.ServiceConnectAdditionalTariffRu))
	h.bot.Handle(&models.BtnBackTojTg, h.BackButton("toj_net"))
	h.bot.Handle(&models.BtnMenuTojTg, h.Menu())

	h.bot.Handle(&models.BtnConfirmRequestRu, h.ConfirmRequest)
	h.bot.Handle(&models.BtnIgnoreRequestRu, h.IgnoreRequest)

	h.bot.Handle(&models.BtnConfirmRequestTg, h.ConfirmRequest)
	h.bot.Handle(&models.BtnIgnoreRequestTg, h.IgnoreRequest)

	h.bot.Handle(&models.BtnPlanOne, h.PlanBtn(constants.PlanOne))
	h.bot.Handle(&models.BtnPlanTwo, h.PlanBtn(constants.PlanTwo))
	h.bot.Handle(&models.BtnPlanThree, h.PlanBtn(constants.PlanThree))
	h.bot.Handle(&models.BtnPlanFour, h.PlanBtn(constants.PlanFour))
	h.bot.Handle(&models.BtnPlanFive, h.PlanBtn(constants.PlanFive))
	h.bot.Handle(&models.BtnPlanSix, h.PlanBtn(constants.PlanSix))
	h.bot.Handle(&models.BtnPlanUnlimitOne, h.PlanBtn(constants.PlanSeven))
	h.bot.Handle(&models.BtnPlanUnlimitTwo, h.PlanBtn(constants.PlanEight))
	h.bot.Handle(&models.BtnPlanUnlimitThree, h.PlanBtn(constants.PlanNine))
	h.bot.Handle(&models.BtnPlanUnlimitFour, h.PlanBtn(constants.PlanTen))
	h.bot.Handle(&models.BtnPlanUnlimitFive, h.PlanBtn(constants.PlanEleventh))
	h.bot.Handle(&models.BtnPlanUnlimitSix, h.PlanBtn(constants.PlanTwelfth))
	h.bot.Handle(&models.BtnPlanUnlimitSeven, h.PlanBtn(constants.PlanThirteenth))

	h.bot.Handle(&models.BtnPlanOneOsh, h.PlanBtn(constants.PlanOne))
	h.bot.Handle(&models.BtnPlanTwoOsh, h.PlanBtn(constants.PlanTwo))
	h.bot.Handle(&models.BtnPlanThreeOsh, h.PlanBtn(constants.PlanThree))
	h.bot.Handle(&models.BtnPlanFourOsh, h.PlanBtn(constants.PlanFour))
	h.bot.Handle(&models.BtnPlanFiveOsh, h.PlanBtn(constants.PlanFive))
	h.bot.Handle(&models.BtnPlanSixOsh, h.PlanBtn(constants.PlanSix))
	h.bot.Handle(&models.BtnPlanUnlimitOneOsh, h.PlanBtn(constants.PlanSeven))
	h.bot.Handle(&models.BtnPlanUnlimitTwoOsh, h.PlanBtn(constants.PlanEight))
	h.bot.Handle(&models.BtnPlanUnlimitThreeOsh, h.PlanBtn(constants.PlanNine))
	h.bot.Handle(&models.BtnPlanUnlimitFourOsh, h.PlanBtn(constants.PlanTen))

	h.bot.Handle(&models.BtnAIConfirmNoRu, h.AIConfirm("no"))
	h.bot.Handle(&models.BtnAIConfirmYesRu, h.AIConfirm("yes"))
	h.bot.Handle(&models.BtnAIConfirmNoTg, h.AIConfirm("no"))
	h.bot.Handle(&models.BtnAIConfirmYesTg, h.AIConfirm("yes"))

	h.bot.Handle(&models.BtnPlanStatusLimitRu, h.PlanLimit())
	h.bot.Handle(&models.BtnPlanStatusUnlimitRu, h.PlanUnlimit())
	h.bot.Handle(&models.BtnPlanStatusLimitToj, h.PlanLimit())
	h.bot.Handle(&models.BtnPlanStatusUnlimitToj, h.PlanUnlimit())

	h.bot.Handle(tele.OnText, h.Text(constants.Russian))
	h.bot.Handle(tele.OnLocation, h.Location)

	h.bot.Start()
}

func setButtons() {
	models.InitAllButtons()

	models.StartMarkup.Reply(
		models.StartMarkup.Row(models.BtnStart),
	)

	models.ConfirmAdMarkup.Inline(
		models.ConfirmAdMarkup.Row(models.BtnConfirmAd, models.BtnIgnoreAd),
	)

	models.PhoneMarkup.Reply(
		models.PhoneMarkup.Row(models.BtnSharePhone),
	)

	models.PhoneMarkupTg.Reply(
		models.PhoneMarkupTg.Row(models.BtnSharePhoneTg),
	)

	models.LanguageMarkup.Reply(
		models.LanguageMarkup.Row(models.BtnTg, models.BtnRu),
	)

	models.MenuMarkupRu.Reply(
		models.MenuMarkupRu.Row(models.BtnRequestProviderRu, models.BtnTechnicalSupportRu),
	)

	models.MenuMarkupTg.Reply(
		models.MenuMarkupTg.Row(models.BtnRequestProviderTg, models.BtnTechnicalSupportTg),
	)

	models.RequestProviderMarkupRu.Reply(
		models.RequestProviderMarkupRu.Row(models.BtnOshnoServiceRu, models.BtnTojNetRu),
		models.RequestProviderMarkupRu.Row(models.BtnMenuRequestProviderRu),
	)

	models.RequestProviderMarkupTg.Reply(
		models.RequestProviderMarkupTg.Row(models.BtnOshnoServiceTg, models.BtnTojNetTg),
		models.RequestProviderMarkupTg.Row(models.BtnMenuRequestProviderTg),
	)

	models.OshnoProviderMarkupRu.Reply(
		models.OshnoProviderMarkupRu.Row(models.BtnOshSerConnectProviderRu, models.BtnOshSerChangeTariffRu, models.BtnOshSerConnectAddittionalRu),
		models.OshnoProviderMarkupRu.Row(models.BtnBackProviderRu, models.BtnMenuProviderRu),
	)

	models.OshnoProviderMarkupTg.Reply(
		models.OshnoProviderMarkupTg.Row(models.BtnOshSerConnectProviderTg, models.BtnOshSerChangeTariffTg, models.BtnOshSerConnectAddittionalTg),
		models.OshnoProviderMarkupTg.Row(models.BtnBackProviderTg, models.BtnMenuProviderTg),
	)

	models.TojNetProviderMarkupRu.Reply(
		models.TojNetProviderMarkupRu.Row(models.BtnTojNetConnectProviderRu, models.BtnTojNetChangeTariffRu, models.BtnTojNetConnectAddittionalRu),
		models.TojNetProviderMarkupRu.Row(models.BtnBackTojRu, models.BtnMenuTojRu),
	)

	models.TojNetProviderMarkupTg.Reply(
		models.TojNetProviderMarkupTg.Row(models.BtnTojNetConnectProviderTg, models.BtnTojNetChangeTariffTg, models.BtnTojNetConnectAddittionalTg),
		models.TojNetProviderMarkupRu.Row(models.BtnBackTojTg, models.BtnMenuTojTg),
	)

	models.LocationMarkupRu.Reply(
		models.LocationMarkupRu.Row(models.BtnLocationRu),
	)

	models.LocationMarkupTg.Reply(
		models.LocationMarkupTg.Row(models.BtnLocationTg),
	)

	models.ConfirmRequestMarkupRu.Reply(
		models.ConfirmRequestMarkupRu.Row(models.BtnConfirmRequestRu, models.BtnIgnoreRequestRu),
	)

	models.ConfirmRequestMarkupTg.Reply(
		models.ConfirmRequestMarkupTg.Row(models.BtnConfirmRequestTg, models.BtnIgnoreRequestTg),
	)

	models.PlanListMarkupToj.Inline(
		models.PlanListMarkupToj.Row(models.BtnPlanOne, models.BtnPlanTwo, models.BtnPlanThree),
		models.PlanListMarkupToj.Row(models.BtnPlanFour, models.BtnPlanFive, models.BtnPlanSix),
	)

	models.PlanListMarkupOsh.Inline(
		models.PlanListMarkupOsh.Row(models.BtnPlanOneOsh, models.BtnPlanTwoOsh, models.BtnPlanThreeOsh),
		models.PlanListMarkupOsh.Row(models.BtnPlanFourOsh, models.BtnPlanFiveOsh, models.BtnPlanSixOsh),
	)

	models.PlanUnlimitListMarkupToj.Inline(
		models.PlanUnlimitListMarkupToj.Row(models.BtnPlanUnlimitOne, models.BtnPlanUnlimitTwo, models.BtnPlanUnlimitThree),
		models.PlanUnlimitListMarkupToj.Row(models.BtnPlanUnlimitFour, models.BtnPlanUnlimitFive, models.BtnPlanUnlimitSix),
		models.PlanUnlimitListMarkupToj.Row(models.BtnPlanUnlimitSeven),
	)

	models.PlanUnlimitListMarkupOsh.Inline(
		models.PlanUnlimitListMarkupOsh.Row(models.BtnPlanUnlimitOneOsh, models.BtnPlanUnlimitTwoOsh, models.BtnPlanUnlimitThreeOsh),
		models.PlanUnlimitListMarkupOsh.Row(models.BtnPlanUnlimitFourOsh),
	)

	models.AIConfirmMarkupRu.Reply(
		models.AIConfirmMarkupRu.Row(models.BtnAIConfirmYesRu, models.BtnAIConfirmNoRu),
	)

	models.AIConfirmMarkupTg.Reply(
		models.AIConfirmMarkupTg.Row(models.BtnAIConfirmYesTg, models.BtnAIConfirmNoTg),
	)

	models.DisplayMarkupRu.Reply(
		models.DisplayMarkupRu.Row(models.BtnBackRu, models.BtnHomeRu),
	)

	models.DisplayMarkupTg.Reply(
		models.DisplayMarkupTg.Row(models.BtnBackTg, models.BtnHomeTg),
	)

	models.PlanStatusMarkupRu.Inline(
		models.PlanStatusMarkupRu.Row(models.BtnPlanStatusLimitRu, models.BtnPlanStatusUnlimitRu),
	)

	models.PlanStatusMarkupToj.Inline(
		models.PlanStatusMarkupToj.Row(models.BtnPlanStatusLimitToj, models.BtnPlanStatusUnlimitToj),
	)
	models.PhoneFillMarkup.Reply(
		models.PhoneMarkup.Row(models.BtnSharePhone),
		models.DisplayMarkupTg.Row(models.BtnBackRu, models.BtnHomeRu),
	)

	models.PhoneFillMarkupTg.Reply(
		models.PhoneMarkupTg.Row(models.BtnSharePhoneTg),
		models.DisplayMarkupTg.Row(models.BtnBackTg, models.BtnHomeTg),
	)

}

func getAdmins() []int64 {
	admins := make([]int64, 0)
	admins = append(admins, 639356141)
	admins = append(admins, 852501376)
	admins = append(admins, 65800827)
	admins = append(admins, 1677349501)
	return admins
}
