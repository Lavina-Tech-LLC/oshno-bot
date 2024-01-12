package constants

type Plans map[int]string

type PlanLimit []int

var OshnoPlans = Plans{
	PlanOne:   "Скорость: 4мб/с, Трафик: 70Гбайт, Сумма: 80сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanTwo:   "Скорость: 8мб/с, Трафик: 100Гбайт, Сумма: 100сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanThree: "Скорость: 12мб/с, Трафик: 160Гбайт, Сумма: 160сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanFour:  "Скорость: 15мб/с, Трафик: 200Гбайт, Сумма: 200сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanFive:  "Скорость: 20мб/с, Трафик: 320Гбайт, Сумма: 320сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanSix:   "Скорость: 30мб/с, Трафик: 500Гбайт, Сумма: 500сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanSeven: "Скорость: 4мб/с, Трафик: Безлимит, Сумма: 100сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanEight: "Скорость: 9мб/с, Трафик: Безлимит, Сумма: 200сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanNine:  "Скорость: 15мб/с, Трафик: Безлимит, Сумма: 300сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanTen:   "Скорость: 25мб/с, Трафик: Безлимит, Сумма: 500сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
}

var TojNetPlans = Plans{
	PlanOne:        "Скорость: 4мб/с, Трафик: 65Гбайт, Сумма: 80сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanTwo:        "Скорость: 8мб/с, Трафик: 100Гбайт, Сумма: 105сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanThree:      "Скорость: 20мб/с, Трафик: 200Гбайт, Сумма: 205сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanFour:       "Скорость: 30мб/с, Трафик: 320Гбайт, Сумма: 330сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanFive:       "Скорость: 40мб/с, Трафик: 420Гбайт, Сумма: 430сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanSix:        "Скорость: 50мб/с, Трафик: 500Гбайт, Сумма: 520сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanSeven:      "Скорость: 5мб/с, Трафик: Безлимит, Сумма: 145сом, Бонус: (+20 мбит/с YouTube, Instagram, Facebook)",
	PlanEight:      "Скорость: 10мб/с, Трафик: Безлимит, Сумма: 285сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanNine:       "Скорость: 15мб/с, Трафик: Безлимит, Сумма: 420сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanTen:        "Скорость: 20мб/с, Трафик: Безлимит, Сумма: 550сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanEleventh:   "Скорость: 25мб/с, Трафик: Безлимит, Сумма: 740сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanTwelfth:    "Скорость: 30мб/с, Трафик: Безлимит, Сумма: 910сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanThirteenth: "Скорость: 40мб/с, Трафик: Безлимит, Сумма: 1075сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
}

var OshnoPlanLimit = PlanLimit{PlanOne, PlanTwo, PlanThree, PlanFour, PlanFive, PlanSix}
var OshnoPlanUnlimit = PlanLimit{PlanSeven, PlanEight, PlanNine, PlanTen}

var TojNetPlanLimit = PlanLimit{PlanOne, PlanTwo, PlanThree, PlanFour, PlanFive, PlanSix}
var TojNetPlanUnlimit = PlanLimit{PlanSeven, PlanEight, PlanNine, PlanTen, PlanEleventh, PlanTwelfth, PlanThirteenth}
