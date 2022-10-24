package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/cfg"
	"gitnub.com/chernyakoff/telegram-auth-cli/pkg/prompt"
)

type sendResponse struct {
	Message string `json:"message"`
}

func Send(account prompt.Account) (string, error) {
	endpoint, _ := cfg.GetEndpoint()

	var message string
	url := strings.TrimSuffix(endpoint, "/") + "/account/upsert"

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	requestBody, err := json.Marshal(account)
	if err != nil {
		return message, err
	}

	reader := bytes.NewReader(requestBody)
	request, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return message, err
	}
	response, err := client.Do(request)
	if err != nil {
		return message, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return message, err
	}
	sendResponse := &sendResponse{}
	err = json.Unmarshal(body, sendResponse)
	if err != nil {
		return message, err
	}

	return sendResponse.Message, err

}
