package constants

type Plans map[int]string

type PlanLimit []int

var OshnoPlans = Plans{
	PlanOne:   "Скорость: 4мб/с, Трафик: 70Гбайт, Цена: 80сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanTwo:   "Скорость: 8мб/с, Трафик: 100Гбайт, Цена: 100сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanThree: "Скорость: 12мб/с, Трафик: 160Гбайт, Цена: 160сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanFour:  "Скорость: 15мб/с, Трафик: 200Гбайт, Цена: 200сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanFive:  "Скорость: 20мб/с, Трафик: 320Гбайт, Цена: 320сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanSix:   "Скорость: 30мб/с, Трафик: 500Гбайт, Цена: 500сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanSeven: "Скорость: 4мб/с, Трафик: Безлимит, Цена: 100сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanEight: "Скорость: 9мб/с, Трафик: Безлимит, Цена: 200сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanNine:  "Скорость: 15мб/с, Трафик: Безлимит, Цена: 300сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanTen:   "Скорость: 25мб/с, Трафик: Безлимит, Цена: 500сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
}

var TojNetPlans = Plans{
	PlanOne:        "Скорость: 4мб/с, Трафик: 65Гбайт, Цена: 80сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanTwo:        "Скорость: 8мб/с, Трафик: 100Гбайт, Цена: 105сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanThree:      "Скорость: 20мб/с, Трафик: 200Гбайт, Цена: 205сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanFour:       "Скорость: 30мб/с, Трафик: 320Гбайт, Цена: 330сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanFive:       "Скорость: 40мб/с, Трафик: 420Гбайт, Цена: 430сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanSix:        "Скорость: 50мб/с, Трафик: 500Гбайт, Цена: 520сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanSeven:      "Скорость: 5мб/с, Трафик: Безлимит, Цена: 145сом, Бонус: (+20 мбит/с YouTube, Instagram, Facebook)",
	PlanEight:      "Скорость: 10мб/с, Трафик: Безлимит, Цена: 285сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanNine:       "Скорость: 15мб/с, Трафик: Безлимит, Цена: 420сом, Бонус: (+30 мбит/с YouTube, Instagram, Facebook)",
	PlanTen:        "Скорость: 20мб/с, Трафик: Безлимит, Цена: 550сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanEleventh:   "Скорость: 25мб/с, Трафик: Безлимит, Цена: 740сом, Бонус: (+40 мбит/с YouTube, Instagram, Facebook)",
	PlanTwelfth:    "Скорость: 30мб/с, Трафик: Безлимит, Цена: 910сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
	PlanThirteenth: "Скорость: 40мб/с, Трафик: Безлимит, Цена: 1075сом, Бонус: (+50 мбит/с YouTube, Instagram, Facebook)",
}

var OshnoPlanLimit = PlanLimit{PlanOne, PlanTwo, PlanThree, PlanFour, PlanFive, PlanSix}
var OshnoPlanUnlimit = PlanLimit{PlanSeven, PlanEight, PlanNine, PlanTen}

var TojNetPlanLimit = PlanLimit{PlanOne, PlanTwo, PlanThree, PlanFour, PlanFive, PlanSix}
var TojNetPlanUnlimit = PlanLimit{PlanSeven, PlanEight, PlanNine, PlanTen, PlanEleventh, PlanTwelfth, PlanThirteenth}
