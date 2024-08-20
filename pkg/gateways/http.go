package gateways

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"oshno/config"
	"oshno/models"
)

func GetHistoryChat(chatId string) (*models.AIRespBody, error) {
	cfg := config.Config()
	url := "https://chatly-back.lavina.tech/account/" + cfg.Chatly.Key + "/chats/history/" + chatId
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Origin", "https://oshno.lavina.tech")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: tr,
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}

	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	// Print the response body as a string
	var history models.AIRespBody
	fmt.Println(string(body))
	err = json.Unmarshal(body, &history)

	if err != nil {
		fmt.Println("error in parse json", err)
		return nil, err
	}

	return &history, nil
}

type SendMessageReq struct {
	Message string `json:"message"`
}

type SendMessageRes struct {
	Data    string `json:"data"`
	IsOk    bool   `json:"isOk"`
	Message string `json:"message"`
}

func SendMessage(chatId, message string) (SendMessageRes, error) {
	cfg := config.Config()
	url := "https://chatly-back.lavina.tech/services/" + cfg.Chatly.Key + "/chats/" + chatId + "/completion/"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	reqBody, err := json.Marshal(SendMessageReq{Message: message})
	if err != nil {

		return SendMessageRes{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {

		return SendMessageRes{}, err
	}

	req.Header.Add("Origin", "https://oshno.lavina.tech")
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Transport: tr,
	}

	response, err := client.Do(req)
	if err != nil {

		return SendMessageRes{}, err
	}

	defer response.Body.Close()

	fmt.Println("Status: ", response.StatusCode, response.Status)

	if response.StatusCode != http.StatusOK {
		return SendMessageRes{}, errors.New("error bad request: ")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return SendMessageRes{}, err
	}

	var responseBody SendMessageRes
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return SendMessageRes{}, err
	}

	return responseBody, nil
}
