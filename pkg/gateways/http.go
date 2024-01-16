package gateways

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oshno/models"
)

func GetHistoryChat(chatId string) (*models.AIRespBody, error) {
	url := "https://chatly-back.lavina.tech/chats/history/" + fmt.Sprintf("%s", chatId)
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
