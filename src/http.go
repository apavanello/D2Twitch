package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getJsonFromWebAPI(currentServerSteamID string) {

	apiDomain := cfg.Dataset.SteamWebApi
	apiKey := "key=" + cfg.Steam.ApiKey
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
			if err := json.Unmarshal(body, &matchJson); err != nil {
				println(string(body))
				processError(err)
			}
			break
		}
	}
}
