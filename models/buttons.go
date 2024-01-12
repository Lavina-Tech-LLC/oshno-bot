package models

import tele "gopkg.in/telebot.v3"

// start
var (
	StartMarkup *tele.ReplyMarkup
	BtnStart    tele.Btn
)

var (
	// share phone russian
	PhoneMarkup   *tele.ReplyMarkup
	PhoneMarkupTg *tele.ReplyMarkup

	BtnSharePhone   tele.Btn
	BtnSharePhoneTg tele.Btn
)

var (
	AdvertisementMarkup *tele.ReplyMarkup
	BtnAdvertisement    tele.Btn
)

var (
	ConfirmAdMarkup *tele.ReplyMarkup
	BtnConfirmAd    tele.Btn
	BtnIgnoreAd     tele.Btn
)

var (
	LanguageMarkup *tele.ReplyMarkup

	BtnTg tele.Btn
	BtnRu tele.Btn
)

var (
	MenuMarkupRu *tele.ReplyMarkup
	MenuMarkupTg *tele.ReplyMarkup

	BtnRequestProviderTg  tele.Btn
	BtnTechnicalSupportTg tele.Btn

	BtnRequestProviderRu  tele.Btn
	BtnTechnicalSupportRu tele.Btn
)

var (
	DisplayMarkupTg *tele.ReplyMarkup
	DisplayMarkupRu *tele.ReplyMarkup

	BtnBackRu tele.Btn
	BtnHomeRu tele.Btn
	BtnBackTg tele.Btn
	BtnHomeTg tele.Btn
)

var (
	RequestProviderMarkupRu *tele.ReplyMarkup
	RequestProviderMarkupTg *tele.ReplyMarkup

	BtnBackRequestProviderRu tele.Btn
	BtnBackRequestProviderTg tele.Btn

	BtnMenuRequestProviderRu tele.Btn
	BtnMenuRequestProviderTg tele.Btn

	BtnOshnoServiceRu tele.Btn
	BtnTojNetRu       tele.Btn

	BtnOshnoServiceTg tele.Btn
	BtnTojNetTg       tele.Btn
)

var (
	OshnoProviderMarkupRu *tele.ReplyMarkup
	OshnoProviderMarkupTg *tele.ReplyMarkup

	BtnOshSerConnectProviderRu    tele.Btn
	BtnOshSerChangeTariffRu       tele.Btn
	BtnOshSerConnectAddittionalRu tele.Btn

	BtnBackProviderRu tele.Btn
	BtnBackProviderTg tele.Btn

	BtnMenuProviderRu tele.Btn
	BtnMenuProviderTg tele.Btn

	BtnOshSerConnectProviderTg    tele.Btn
	BtnOshSerChangeTariffTg       tele.Btn
	BtnOshSerConnectAddittionalTg tele.Btn
)

var (
	TojNetProviderMarkupRu *tele.ReplyMarkup
	TojNetProviderMarkupTg *tele.ReplyMarkup

	BtnTojNetConnectProviderRu    tele.Btn
	BtnTojNetChangeTariffRu       tele.Btn
	BtnTojNetConnectAddittionalRu tele.Btn

	BtnBackTojRu tele.Btn
	BtnBackTojTg tele.Btn

	BtnMenuTojRu tele.Btn
	BtnMenuTojTg tele.Btn

	BtnTojNetConnectProviderTg    tele.Btn
	BtnTojNetChangeTariffTg       tele.Btn
	BtnTojNetConnectAddittionalTg tele.Btn
)

var (
	LocationMarkupRu *tele.ReplyMarkup
	LocationMarkupTg *tele.ReplyMarkup

	BtnLocationRu tele.Btn
	BtnLocationTg tele.Btn
)

var (
	ConfirmRequestMarkupRu *tele.ReplyMarkup
	ConfirmRequestMarkupTg *tele.ReplyMarkup

	BtnConfirmRequestRu tele.Btn
	BtnIgnoreRequestRu  tele.Btn

	BtnConfirmRequestTg tele.Btn
	BtnIgnoreRequestTg  tele.Btn
)

var (
	PlanStatusMarkupToj *tele.ReplyMarkup
	PlanStatusMarkupRu  *tele.ReplyMarkup

	BtnPlanStatusLimitToj   tele.Btn
	BtnPlanStatusUnlimitToj tele.Btn
	BtnPlanStatusLimitRu    tele.Btn
	BtnPlanStatusUnlimitRu  tele.Btn
)

var (
	PlanListMarkupToj *tele.ReplyMarkup
	PlanListMarkupOsh *tele.ReplyMarkup

	PlanUnlimitListMarkupToj *tele.ReplyMarkup
	PlanUnlimitListMarkupOsh *tele.ReplyMarkup

	BtnPlanOne          tele.Btn
	BtnPlanTwo          tele.Btn
	BtnPlanThree        tele.Btn
	BtnPlanFour         tele.Btn
	BtnPlanFive         tele.Btn
	BtnPlanSix          tele.Btn
	BtnPlanUnlimitOne   tele.Btn
	BtnPlanUnlimitTwo   tele.Btn
	BtnPlanUnlimitThree tele.Btn
	BtnPlanUnlimitFour  tele.Btn
	BtnPlanUnlimitFive  tele.Btn
	BtnPlanUnlimitSix   tele.Btn
	BtnPlanUnlimitSeven tele.Btn

	BtnPlanOneOsh          tele.Btn
	BtnPlanTwoOsh          tele.Btn
	BtnPlanThreeOsh        tele.Btn
	BtnPlanFourOsh         tele.Btn
	BtnPlanFiveOsh         tele.Btn
	BtnPlanSixOsh          tele.Btn
	BtnPlanUnlimitOneOsh   tele.Btn
	BtnPlanUnlimitTwoOsh   tele.Btn
	BtnPlanUnlimitThreeOsh tele.Btn
	BtnPlanUnlimitFourOsh  tele.Btn
)

var (
	AIConfirmMarkupRu *tele.ReplyMarkup
	AIConfirmMarkupTg *tele.ReplyMarkup

	BtnAIConfirmYesRu tele.Btn
	BtnAIConfirmNoRu  tele.Btn

	BtnAIConfirmYesTg tele.Btn
	BtnAIConfirmNoTg  tele.Btn
)

var (

	// share phone russian
	PhoneFillMarkup   *tele.ReplyMarkup
	PhoneFillMarkupTg *tele.ReplyMarkup

	BtnShareFillPhone   tele.Btn
	BtnShareFillPhoneTg tele.Btn
)
