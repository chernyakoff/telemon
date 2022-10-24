package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type pingResponse struct {
	Message string `json:"message"`
}

func Ping(endpoint string) bool {
	url := strings.TrimSuffix(endpoint, "/") + "/ping"
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false
	}
	response, err := client.Do(request)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}
	pingResponse := &pingResponse{}
	err = json.Unmarshal(body, pingResponse)
	if err != nil {
		return false
	}
	if pingResponse.Message == "pong" {
		return true
	} else {
		return false
	}

}
