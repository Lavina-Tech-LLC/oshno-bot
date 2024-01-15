package constants

import "fmt"

// map[messagesKey]messageValue
type Messages map[string]string

// mape[language]Messages
type Languages map[string]Messages

var ConstMessages = Languages{
	Tajik: Messages{
		TestALert:              "Ҳол дар барнома рӯйхаткунӣ карда мешавад!\n",
		ChooseLanguageStart:    "Заборо интихоб кунед хизматрасонӣ",
		ChooseLanguageSettings: "Заборо интихоб кунед хизматрасонӣ",
		LanguageSet:            "Забон бомуваффақият интихоб карда шуд. Шумо метавонед давом диҳед!",
		LanguageChanged:        "Барои пайваст шудан ба провайдер, тарофаро тағир диҳедан ё нуқтаи иловагии интернетро пайваст кадан, ба 'Провайдери дархост' -ро интихоб кунед.\nБарои гирифтани дастгирии техникӣ, 'Пуштибонии фаннӣ' -ро интихоб кунед.",
		ErrorReport:            "Мутасир карда шуд. Дармандаи бот!, Кӯшиш кунед ботро дубора пайваст карда маълумотро аз нав овардан.",
		Settings:               "Аз насбҳо интихоб кунед:",
		MenuButtonText:         "Меню",
		EnterPhone:             "Барои дархости, намбури телефони худро баҳам медиҳед ё дар формати 901234567 ворид кунед:",
		CodeSentAlready:        "Код аллакай фиристода шуд, баъд аз 15 сония оид барқарор кунед",
		Restart:                "/рестарт",
		Location:               "Сурати достоваро ворид кунед:",
		YourPhone:              "Намбурои телефони шумо: <b>%s</b>",
		YourLocation:           "Маҳолати шумо:\n <b>%s</b>",
		RequestProvider:        "Провайдерро интихоб кунед",
		ChooseService:          "Хидмати заруриро интихоб кунед",
		EnterFirstName:         "Ном ва насабатонро ворид кунед",
		EnterLastName:          "Фамилияи худро ворид кунед: ",
		EnterMiddleName:        "Номи падари худро ворид кунед",
		EnterDateBirth:         "Санаи таваллуди худро ворид кунед: dd-mm-yyyy",
		EnterPhoneNumber:       "Рақами телефони худро мубодила кунед ё дар формати +992ХХХХХХХХ ворид кунед",
		EnterEmail:             "Адреси почтаи электронии худро ворид кунед",
		EnterAddress:           "Суроғаеро, ки шумо бояд ба Интернет пайваст шавед, ворид кунед",
		EnterPersonalAccount:   "Хисоби шахсии худро ворид кунед",
		EnterNewPlan:           "Барномаи навро интихоб кунед",
		EnterPassportData:      "Маълумоти паспорти худро ворид кунед",
		EnterPassportBack:      "Сурати паспорти худро бо тарафи пойонӣ фиристед",
		EnterPassportFront:     "Сурати паспорти худро бо тарафи оварид фиристед",
		EnterPhotoWith:         "Сурати худро бо паспорти худ фиристед",

		ReEnterFirstName:                 "Номи шумо бояд танҳо извози барои як ҳарфи китоб аст: ",
		ReEnterLastName:                  "Насаби шумо бояд танҳо извози барои ҳарфҳо бошад: ",
		ReEnterMiddleName:                "Номи падари шумо бояд танҳо извози барои ҳарфҳо бошад: ",
		ReEnterDateBirth:                 "Санаи таваллуди шумо бояд дар формати dd-mm-yyyy бошад",
		ReEnterPhoneNumber:               "Маълумоти нодуруст. Рақами телефони худро мубодила кунед ё дар формати +992ХХХХХХХХ ворид кунед",
		EnterTariffSelectionOshnoService: fmt.Sprintf("Тарифро интихоб кунед\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n7. %s\n8. %s\n9. %s\n10. %s\n", OshnoPlans[PlanOne], OshnoPlans[PlanTwo], OshnoPlans[PlanThree], OshnoPlans[PlanFour], OshnoPlans[PlanFive], OshnoPlans[PlanSix], OshnoPlans[PlanSeven], OshnoPlans[PlanEight], OshnoPlans[PlanNine], OshnoPlans[PlanTen]),
		EnterTariffSelectionTojNet:       fmt.Sprintf("Тарифро интихоб кунед\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n7. %s\n8. %s\n9. %s\n10. %s\n11. %s\n12. %s\n13. %s\n", TojNetPlans[PlanOne], TojNetPlans[PlanTwo], TojNetPlans[PlanThree], TojNetPlans[PlanFour], TojNetPlans[PlanFive], TojNetPlans[PlanSix], TojNetPlans[PlanSeven], TojNetPlans[PlanEight], TojNetPlans[PlanNine], TojNetPlans[PlanTen], TojNetPlans[PlanEleventh], TojNetPlans[PlanTwelfth], TojNetPlans[PlanThirteenth]),
		PlanTariffLimitOshno:             fmt.Sprintf("Тарифро интихоб кунед\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n", OshnoPlans[PlanOne], OshnoPlans[PlanTwo], OshnoPlans[PlanThree], OshnoPlans[PlanFour], OshnoPlans[PlanFive], OshnoPlans[PlanSix]),
		PlanTariffLimitTojNet:            fmt.Sprintf("Тарифро интихоб кунед\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n", TojNetPlans[PlanOne], TojNetPlans[PlanTwo], TojNetPlans[PlanThree], TojNetPlans[PlanFour], TojNetPlans[PlanFive], TojNetPlans[PlanSix]),
		PlanTariffUnlimitOshno:           fmt.Sprintf("Тарифро интихоб кунед\n1. %s\n2. %s\n3. %s\n4. %s\n", OshnoPlans[PlanSeven], OshnoPlans[PlanEight], OshnoPlans[PlanNine], OshnoPlans[PlanTen]),
		PlanTariffUnlimitTojNet:          fmt.Sprintf("Тарифро интихоб кунед\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n7. %s\n", TojNetPlans[PlanSeven], TojNetPlans[PlanEight], TojNetPlans[PlanNine], TojNetPlans[PlanTen], TojNetPlans[PlanEleventh], TojNetPlans[PlanTwelfth], TojNetPlans[PlanThirteenth]),
		EnterPreferPlanTariff:            "Нақшаеро, ки ба шумо манфиатдор аст, интихоб кунед",
		ConfirmRequest:                   "Дархости фиристодан",
		ConfirmMessage:                   "Мо дархости шуморо қабул хоҳем кард.\nКормандони мо ба зудӣ бо шумо тамос хоҳанд гирифт ва дар посух ба дархостатон ба шумо кумак хоҳанд кард.\nМо аз ҳамкорӣ бо шумо хурсандем!",
		IgnoreMessage:                    "Дархости шумо бекор карда шуд!",
		NotSelectebleTarif:               "Шумо наметавонед тарифро тағир диҳед, шумо метавонед фавран тарифро тағир диҳед Тарифро васл кунед",
		NotSelectebleAddTarif:            "Шумо наметавонед тарифи иловагиро васл кунед, барои фаъол кардани тарифи иловагӣ аввал бояд тарифро васл кунед",
	},

	Russian: Messages{
		TestALert:              "Пока бот работает в тестовом режиме!\n",
		ChooseLanguageStart:    "Выберите язык обслуживание",
		ChooseLanguageSettings: "Выберите язык обслуживание",
		LanguageSet:            "Язык успешно установлен. Вы можете продолжить заказ!",
		LanguageChanged:        "Для подключения к провайдеру, смены тарифа или подключения дополнительную точку интернета выберите 'Оставить запрос на провайдера'\nДля получения техническую поддержку выберите 'Техническая поддержка'",
		ErrorReport:            "Прошу прощения. Ошибка бота! Попробуйте перезапустить бота.",
		Settings:               "Выберите в настройках:",
		MenuButtonText:         "Меню",
		EnterPhone:             "Для заказа поделитесь номером телефона или введите в формате 901234567:",
		CodeSentAlready:        "Код уже отправлен, попробуйте через 15 секунд",
		Restart:                "/restart",
		Location:               "Oтправьте адрес доставки:",
		YourPhone:              "Ваш номер телефона: <b>%s</b>",
		YourLocation:           "Ваш адрес:\n <b>%s</b>",
		RequestProvider:        "Выберите провайдера",
		ChooseService:          "Выберите необходимую услугу",
		EnterFirstName:         "Введите ваши имя и фамилию",
		EnterLastName:          "Укажите вашу фаимлию: ",
		EnterMiddleName:        "Укажите вашу отчество: ",
		EnterDateBirth:         "Укажите вашу дату рождение: дд-мм-гггг",
		EnterPhoneNumber:       "Поделитесь вашим номером телефона или введите в формате +992ХХХХХХХХХ",
		EnterEmail:             "Укажите вашу электроную почту:",
		EnterAddress:           "Введите адрес где необходимо подключить интернет",
		EnterPassportData:      "Укажите ваш пасспортные данные:",
		EnterPersonalAccount:   "Укажите ваш лицевой счёт:",
		EnterNewPlan:           "Выберите новый план",
		EnterPassportFront:     "Отправте фото пасспорт с передней стороны",
		EnterPassportBack:      "Отправте фото пасспорт с задней стороны",
		EnterPhotoWith:         "Отправте фото селфи с пасспортом",

		ReEnterFirstName:                 "Ваше имя должен состоять только из букв: ",
		ReEnterLastName:                  "Ваше фаимлия состоять только из букв: ",
		ReEnterMiddleName:                "Ваше отчество состоять только из букв: ",
		ReEnterDateBirth:                 "Ваше дата рождение дожен быть в формате: дд-мм-гггг",
		ReEnterPhoneNumber:               "Неправильные данные. Поделитесь вашим номером телефона или введите в формате +992ХХХХХХХХХ",
		ReEnterEmail:                     "Ваша электроная почта дожен быть в формате: xxxxx@xxx.xxx",
		ReEnterAddress:                   "Ваш адрес должен содежать:",
		ReEnterPassportData:              "Ваш пасспортные данные должен содержать: ",
		ReEnterPassportFront:             "Отправте фото пасспорт с передней стороны",
		ReEnterPassportBack:              "Отправте фото пасспорт с задней стороны",
		ReEnterPhotoWith:                 "Отправте фото селфи с пасспортом",
		EnterTariffSelectionOshnoService: fmt.Sprintf("Выберите тариф\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n7. %s\n8. %s\n9. %s\n10. %s\n", OshnoPlans[PlanOne], OshnoPlans[PlanTwo], OshnoPlans[PlanThree], OshnoPlans[PlanFour], OshnoPlans[PlanFive], OshnoPlans[PlanSix], OshnoPlans[PlanSeven], OshnoPlans[PlanEight], OshnoPlans[PlanNine], OshnoPlans[PlanTen]),
		EnterTariffSelectionTojNet:       fmt.Sprintf("Выберите тариф\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n7. %s\n8. %s\n9. %s\n10. %s\n11. %s\n12. %s\n13. %s\n", TojNetPlans[PlanOne], TojNetPlans[PlanTwo], TojNetPlans[PlanThree], TojNetPlans[PlanFour], TojNetPlans[PlanFive], TojNetPlans[PlanSix], TojNetPlans[PlanSeven], TojNetPlans[PlanEight], TojNetPlans[PlanNine], TojNetPlans[PlanTen], TojNetPlans[PlanEleventh], TojNetPlans[PlanTwelfth], TojNetPlans[PlanThirteenth]),
		PlanTariffLimitOshno:             fmt.Sprintf("Выберите тариф\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n", OshnoPlans[PlanOne], OshnoPlans[PlanTwo], OshnoPlans[PlanThree], OshnoPlans[PlanFour], OshnoPlans[PlanFive], OshnoPlans[PlanSix]),
		PlanTariffLimitTojNet:            fmt.Sprintf("Выберите тариф\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n", TojNetPlans[PlanOne], TojNetPlans[PlanTwo], TojNetPlans[PlanThree], TojNetPlans[PlanFour], TojNetPlans[PlanFive], TojNetPlans[PlanSix]),
		PlanTariffUnlimitOshno:           fmt.Sprintf("Выберите тариф\n1. %s\n2. %s\n3. %s\n4. %s\n", OshnoPlans[PlanSeven], OshnoPlans[PlanEight], OshnoPlans[PlanNine], OshnoPlans[PlanTen]),
		PlanTariffUnlimitTojNet:          fmt.Sprintf("Выберите тариф\n1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n6. %s\n7. %s\n", TojNetPlans[PlanSeven], TojNetPlans[PlanEight], TojNetPlans[PlanNine], TojNetPlans[PlanTen], TojNetPlans[PlanEleventh], TojNetPlans[PlanTwelfth], TojNetPlans[PlanThirteenth]),
		EnterPreferPlanTariff:            "Выберите план который вас интересует",
		ConfirmRequest:                   "Отправить заявку",
		ConfirmMessage:                   "Вас запрос принять.\nВ скором времени наши сотрудники свяжутся с вами и помогут вам с ответом вашей заявки.\nМы рады работать с вами!",
		IgnoreMessage:                    "Ваша заявка было отменена!",
		NotSelectebleTarif:               "Вы не можете сменить тариф, чтобы сменить тариф с начало нужно Подключить тариф",
		NotSelectebleAddTarif:            "Вы не можете подключить дополнительную точка, чтобы подключить дополнительную точку с начало нужно Подключить тариф",
	},
}
