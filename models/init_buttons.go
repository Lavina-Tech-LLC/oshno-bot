package models

import (
	"fmt"

	"github.com/enescakir/emoji"
	tele "gopkg.in/telebot.v3"
)

func InitAllButtons() {
	InitAdvertisement()
	InitConfirmButton()
	InitStart()
	InitPhone()
	InitLocation()
	InitLanguage()
	InitMenu()
	InitRequestProvider()
	InitOshnoProvider()
	InitTojNetProvider()
	InitUserRegistarion()
	InitConfirmRequest()
	InitPlanList()
	InitAIConfirm()
	InitDisplay()
	InitPlanStatus()
	InitPhoneFill()
}

// restart button
func InitStart() {
	StartMarkup = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	BtnStart = StartMarkup.Text("restart")
}

func InitPhone() {
	// phone buttons
	PhoneMarkup = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	PhoneMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	BtnSharePhone = PhoneMarkup.Contact("Поделиться контактом")
	BtnSharePhoneTg = PhoneMarkup.Contact("Телефонро мубодила кунед")

}

func InitLocation() {
	LocationMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	LocationMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnLocationRu = LocationMarkupRu.Location("Моя локация")
	BtnLocationTg = LocationMarkupTg.Location("My location")

}

func InitAdvertisement() {
	AdvertisementMarkup = &tele.ReplyMarkup{}
	BtnAdvertisement = AdvertisementMarkup.Data(fmt.Sprintf("%sДа", emoji.ThumbsUp), "da_ad")
}

func InitConfirmButton() {
	ConfirmAdMarkup = &tele.ReplyMarkup{}
	BtnConfirmAd = ConfirmAdMarkup.Data("Ҳа", "yes_confirm")
	BtnIgnoreAd = ConfirmAdMarkup.Data("Нест", "no_confirm")
}

// inline language buttons
func InitLanguage() {
	LanguageMarkup = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnTg = LanguageMarkup.Text(fmt.Sprintf("%s Таджийский", emoji.FlagForTajikistan))
	BtnRu = LanguageMarkup.Text(fmt.Sprintf("%s Русский", emoji.FlagForRussia))
}

func InitMenu() {
	MenuMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnRequestProviderTg = MenuMarkupTg.Text(fmt.Sprintf("%sПровайдери дархост", emoji.InboxTray))
	BtnTechnicalSupportTg = MenuMarkupTg.Text(fmt.Sprintf("%sПуштибонии фаннӣ", emoji.ManTechnologist))

	MenuMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnRequestProviderRu = MenuMarkupRu.Text(fmt.Sprintf("%s Оставить запрос на провайдера", emoji.InboxTray))
	BtnTechnicalSupportRu = MenuMarkupRu.Text(fmt.Sprintf("%s Tехническая подержка", emoji.ManTechnologist))

}

func InitRequestProvider() {
	RequestProviderMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	RequestProviderMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnOshnoServiceRu = RequestProviderMarkupRu.Text("Oshno service")
	BtnTojNetRu = RequestProviderMarkupRu.Text("TojNet")
	BtnBackRequestProviderRu = RequestProviderMarkupRu.Text(fmt.Sprintf("%sНɑзад", emoji.BackArrow))
	BtnMenuRequestProviderRu = RequestProviderMarkupRu.Text(fmt.Sprintf("%sМеню", emoji.House))

	BtnOshnoServiceTg = RequestProviderMarkupTg.Text("Oshno service")
	BtnTojNetTg = RequestProviderMarkupTg.Text("TojNet")
	BtnBackRequestProviderTg = RequestProviderMarkupTg.Text(fmt.Sprintf("%sБозгaшт", emoji.BackArrow)) // a => is russian letter
	BtnMenuRequestProviderTg = RequestProviderMarkupTg.Text(fmt.Sprintf("%sМеню", emoji.House))

}

func InitOshnoProvider() {
	OshnoProviderMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	OshnoProviderMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnOshSerConnectProviderRu = OshnoProviderMarkupRu.Text("Подключится к провайдеру")
	BtnOshSerChangeTariffRu = OshnoProviderMarkupRu.Text("Сменить тариф")
	BtnOshSerConnectAddittionalRu = OshnoProviderMarkupRu.Text("Подключить дополнительную точку")
	BtnBackProviderRu = OshnoProviderMarkupRu.Text(fmt.Sprintf("%sНазɑд", emoji.BackArrow))
	BtnMenuProviderRu = OshnoProviderMarkupRu.Text(fmt.Sprintf("%sМеню", emoji.House))

	BtnOshSerConnectProviderTg = OshnoProviderMarkupTg.Text("Провайдери пайваст кунед Ошно")
	BtnOshSerChangeTariffTg = OshnoProviderMarkupTg.Text("Тағйир додани қурби Ошно")
	BtnOshSerConnectAddittionalTg = OshnoProviderMarkupTg.Text("Пайвастшавӣ ба нархи иловагӣ Ошно")
	BtnBackProviderTg = OshnoProviderMarkupTg.Text(fmt.Sprintf("%sБозгɑшт", emoji.BackArrow))
	BtnMenuProviderTg = OshnoProviderMarkupTg.Text(fmt.Sprintf("%sМеню", emoji.House))

}

func InitTojNetProvider() {
	TojNetProviderMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	TojNetProviderMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnTojNetConnectProviderRu = TojNetProviderMarkupRu.Text("Пoдключится к провайдеру")           // first 'o' leter in enlish keyword
	BtnTojNetChangeTariffRu = TojNetProviderMarkupRu.Text("Смeнить тариф")                         // first 'e' leter in enlish keyword
	BtnTojNetConnectAddittionalRu = TojNetProviderMarkupRu.Text("Пoдключить дополнительную точку") // first 'o' leter in enlish keyword
	BtnBackTojRu = TojNetProviderMarkupRu.Text(fmt.Sprintf("%sНɑзɑд", emoji.BackArrow))
	BtnMenuTojRu = TojNetProviderMarkupRu.Text(fmt.Sprintf("%sМеню", emoji.House))

	BtnTojNetConnectProviderTg = TojNetProviderMarkupTg.Text("Провайдери пайваст кунед TojNet")
	BtnTojNetChangeTariffTg = TojNetProviderMarkupTg.Text("Тағйир додани қурби TojNet")
	BtnTojNetConnectAddittionalTg = TojNetProviderMarkupTg.Text("Пайвастшавӣ ба нархи иловагӣ TojNet")
	BtnBackTojTg = TojNetProviderMarkupTg.Text(fmt.Sprintf("%sБoзгaшт", emoji.BackArrow)) // ac -> is russian letters
	BtnMenuTojTg = TojNetProviderMarkupTg.Text(fmt.Sprintf("%sМеню", emoji.House))

}

func InitConfirmRequest() {
	ConfirmRequestMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	ConfirmRequestMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnConfirmRequestRu = ConfirmRequestMarkupRu.Text("Да")
	BtnIgnoreRequestRu = ConfirmRequestMarkupRu.Text("Нет")

	BtnConfirmRequestTg = ConfirmRequestMarkupTg.Text("Ха")
	BtnIgnoreRequestTg = ConfirmRequestMarkupTg.Text("Нест")
}

func InitUserRegistarion() {
	FirstNameMarkup = &tele.ReplyMarkup{}
}

func InitPlanList() {
	PlanListMarkupToj = &tele.ReplyMarkup{}
	PlanListMarkupOsh = &tele.ReplyMarkup{}

	PlanUnlimitListMarkupToj = &tele.ReplyMarkup{}
	PlanUnlimitListMarkupOsh = &tele.ReplyMarkup{}

	BtnPlanOne = PlanListMarkupToj.Data("1", "one")
	BtnPlanTwo = PlanListMarkupToj.Data("2", "two")
	BtnPlanThree = PlanListMarkupToj.Data("3", "three")
	BtnPlanFour = PlanListMarkupToj.Data("4", "four")
	BtnPlanFive = PlanListMarkupToj.Data("5", "five")
	BtnPlanSix = PlanListMarkupToj.Data("6", "six")

	BtnPlanUnlimitOne = PlanUnlimitListMarkupToj.Data("1", "seven")
	BtnPlanUnlimitTwo = PlanUnlimitListMarkupToj.Data("2", "eight")
	BtnPlanUnlimitThree = PlanUnlimitListMarkupToj.Data("3", "nine")
	BtnPlanUnlimitFour = PlanUnlimitListMarkupToj.Data("4", "ten")
	BtnPlanUnlimitFive = PlanUnlimitListMarkupToj.Data("5", "eleven")
	BtnPlanUnlimitSix = PlanUnlimitListMarkupToj.Data("6", "twelfth")
	BtnPlanUnlimitSeven = PlanUnlimitListMarkupToj.Data("7", "Thirteenth")

	BtnPlanOneOsh = PlanListMarkupOsh.Data("1", "one_osh")
	BtnPlanTwoOsh = PlanListMarkupOsh.Data("2", "two_osh")
	BtnPlanThreeOsh = PlanListMarkupOsh.Data("3", "three_osh")
	BtnPlanFourOsh = PlanListMarkupOsh.Data("4", "four_osh")
	BtnPlanFiveOsh = PlanListMarkupOsh.Data("5", "five_osh")
	BtnPlanSixOsh = PlanListMarkupOsh.Data("6", "six_osh")

	BtnPlanUnlimitOneOsh = PlanUnlimitListMarkupOsh.Data("1", "seven_osh")
	BtnPlanUnlimitTwoOsh = PlanUnlimitListMarkupOsh.Data("2", "eight_osh")
	BtnPlanUnlimitThreeOsh = PlanUnlimitListMarkupOsh.Data("3", "nine_osh")
	BtnPlanUnlimitFourOsh = PlanUnlimitListMarkupOsh.Data("4", "ten_osh")

}

func InitAIConfirm() {
	AIConfirmMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	AIConfirmMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnAIConfirmNoRu = AIConfirmMarkupRu.Text("нет")
	BtnAIConfirmYesRu = AIConfirmMarkupRu.Text("да")

	BtnAIConfirmNoTg = AIConfirmMarkupTg.Text("нест")
	BtnAIConfirmYesTg = AIConfirmMarkupTg.Text("ҳа")

}

func InitDisplay() {
	DisplayMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	DisplayMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnBackRu = DisplayMarkupRu.Text(fmt.Sprintf("%sНазад", emoji.BackArrow))
	BtnHomeRu = DisplayMarkupRu.Text(fmt.Sprintf("%sМеню", emoji.House))

	BtnBackTg = DisplayMarkupTg.Text(fmt.Sprintf("%sБозгашт", emoji.BackArrow))
	BtnHomeTg = DisplayMarkupTg.Text(fmt.Sprintf("%sМеню", emoji.House))
}

func InitPlanStatus() {
	PlanStatusMarkupRu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	PlanStatusMarkupToj = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnPlanStatusLimitRu = PlanStatusMarkupRu.Data("Лимит", "limit_ru")
	BtnPlanStatusUnlimitRu = PlanStatusMarkupRu.Data("Безлимит", "unlimit_ru")

	BtnPlanStatusLimitToj = PlanStatusMarkupToj.Data("Маҳдудият", "limit_toj")
	BtnPlanStatusUnlimitToj = PlanStatusMarkupToj.Data("Hомаҳдуд", "unlimit_toj")

}

func InitPhoneFill() {
	// phone buttons
	PhoneFillMarkup = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}
	PhoneFillMarkupTg = &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	BtnShareFillPhone = PhoneMarkup.Contact("Поделиться с контактом")
	BtnShareFillPhoneTg = PhoneMarkup.Contact("Телефонро мубодила кунед")
}
