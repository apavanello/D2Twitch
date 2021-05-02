package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getJSONFromWebAPI(currentServerSteamID string) {

	apiDomain := cfg.Dataset.SteamWebAPI
	apiKey := "key=" + cfg.Steam.APIKey
	matchPatch := "/IDOTA2MatchStats_570/GetRealtimeStats/v1/?"
	serverSteamID := "server_steam_id=" + currentServerSteamID
	URI := apiDomain + matchPatch + apiKey + "&" + serverSteamID

	if cfg.Mock.Enable == true {
		URI = cfg.Mock.Endpoints.MatchStats
	}

	for i := 0; i < 3; i++ {
		resp, err := http.Get(URI)
		if err != nil {
			processError(err)
		}

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				processError(err)
			}
			if err := json.Unmarshal(body, &matchJSON); err != nil {
				println(string(body))
				processError(err)
			}
			break
		}
	}
}
