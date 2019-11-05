package discord

import (
	"bytes"
	"encoding/json"
	"github.com/fetus-hina/salmon-notify/conf"
	"net/http"
)

type Webhook struct {
	UserName string `json:"username"`
	Content  string `json:"content"`
}

func (t Webhook) Post(endpointUrl string) error {
	client := http.Client{}

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpointUrl, bytes.NewBuffer(json))
	req.Header.Add("User-Agent", conf.GetUserAgentString())
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
