package s2ink

import (
	"encoding/json"
	"github.com/fetus-hina/salmon-notify/conf"
	"github.com/fetus-hina/salmon-notify/types"
	"io/ioutil"
	"net/http"
)

const coopScheduleUrl = "https://splatoon2.ink/data/coop-schedules.json"

func FetchCoopSchedules() (*types.CoopSchedules, error) {
	body, err := fetchJson(coopScheduleUrl)
	if err != nil {
		return nil, err
	}

	retObj := types.CoopSchedules{}
	err = json.Unmarshal(body, &retObj)
	if err != nil {
		return nil, err
	}

	return &retObj, nil
}

func fetchJson(url string) ([]byte, error) {
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", conf.GetUserAgentString())

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
