package utils

import (
	"fmt"
	"oshno/pkg/constants"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func CreateMessage(msg *tele.Message) interface{} {
	var data interface{}

	if msg.Text != "" {
		data = msg.Text
		return data
	}

	switch msg.Media().MediaType() {
	case "photo":
		photo := msg.Photo
		photo.Caption = msg.Caption
		data = photo
	case "video":
		video := msg.Video
		video.Caption = msg.Caption
		data = video
	case "voice":
		voice := msg.Voice
		voice.Caption = msg.Caption
		data = voice
	case "audio":
		audio := msg.Audio
		audio.Caption = msg.Caption
		data = audio
	case "animation":
		animation := msg.Animation
		animation.Caption = msg.Caption
		data = animation
	case "sticker":
		data = msg.Sticker
	case "document":
		document := msg.Document
		document.Caption = msg.Caption
		data = document
	case "videoNote":
		data = msg.VideoNote
	}

	return data
}

func UserPhaseToService(service string) int {
	switch service {
	case constants.ServiceConnectProviderRu:
		return 1
	case constants.ServiceChangeTariffRu:
		return 6
	case constants.ServiceConnectAdditionalTariffRu:
		return 9
	default:
		return 0
	}
}

func AddWorker(w string) (uint, string) {
	if w == "" {
		return 1, "1"
	}
	wArray := strings.Split(w, "-")
	lastElement := wArray[len(wArray)-1]
	num, _ := strconv.Atoi(lastElement)
	return uint(num + 1), w + "-" + strconv.Itoa(num+1)
}

func DeleteWorker(w string, wNum uint) string {
	var wString string
	if wNum != 1 {
		wString = "-" + strconv.Itoa(int(wNum))
	} else {
		wString = "1"
	}

	index := strings.Index(w, wString)
	if index != -1 {
		// Подстрока найдена, удаляем её
		result := w[:index] + w[index+len(wString):]
		fmt.Println("Исходная строка:", w)
		fmt.Println("Строка после удаления подстроки:", result)
		return result
	} else {
		fmt.Println("Подстрока не найдена.")
		return w
	}

}
func IsLastWorker(w string, num uint) bool {
	myString := w
	substring := strconv.Itoa(int(num))

	index := strings.Index(myString, substring)
	if index != -1 && index+len(substring) < len(myString) {
		nextSubstring := myString[index+len(substring):]
		fmt.Println("Подстрока найдена, и следующая подстрока начинается с:", nextSubstring)
		return true
	} else if index != -1 {
		return false
	} else {
		fmt.Println("Строка не найдена")
		return true
	}
}
